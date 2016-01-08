package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

// Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each.  The service goroutines read requests and
// then call srv.Handler to reply to them.

/*

File of server.go

func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		rw, e := l.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c, err := srv.newConn(rw)
		if err != nil {
			continue
		}
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve()
	}
}

...

// Serve a new connection.
func (c *conn) serve() {
	origConn := c.rwc // copy it before it's set nil on Close or Hijack
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
		}
		if !c.hijacked() {
			c.close()
			c.setState(origConn, StateClosed)
		}
	}()

	if tlsConn, ok := c.rwc.(*tls.Conn); ok {
		if d := c.server.ReadTimeout; d != 0 {
			c.rwc.SetReadDeadline(time.Now().Add(d))
		}
		if d := c.server.WriteTimeout; d != 0 {
			c.rwc.SetWriteDeadline(time.Now().Add(d))
		}
		if err := tlsConn.Handshake(); err != nil {
			c.server.logf("http: TLS handshake error from %s: %v", c.rwc.RemoteAddr(), err)
			return
		}
		c.tlsState = new(tls.ConnectionState)
		*c.tlsState = tlsConn.ConnectionState()
		if proto := c.tlsState.NegotiatedProtocol; validNPN(proto) {
			if fn := c.server.TLSNextProto[proto]; fn != nil {
				h := initNPNRequest{tlsConn, serverHandler{c.server}}
				fn(c.server, tlsConn, h)
			}
			return
		}
	}

	for {
		w, err := c.readRequest()
		if c.lr.N != c.server.initialLimitedReaderSize() {
			// If we read any bytes off the wire, we're active.
			c.setState(c.rwc, StateActive)
		}
		if err != nil {
			if err == errTooLarge {
				// Their HTTP client may or may not be
				// able to read this if we're
				// responding to them and hanging up
				// while they're still writing their
				// request.  Undefined behavior.
				io.WriteString(c.rwc, "HTTP/1.1 413 Request Entity Too Large\r\n\r\n")
				c.closeWriteAndWait()
				break
			} else if err == io.EOF {
				break // Don't reply
			} else if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				break // Don't reply
			}
			io.WriteString(c.rwc, "HTTP/1.1 400 Bad Request\r\n\r\n")
			break
		}

		// Expect 100 Continue support
		req := w.req
		if req.expectsContinue() {
			if req.ProtoAtLeast(1, 1) && req.ContentLength != 0 {
				// Wrap the Body reader with one that replies on the connection
				req.Body = &expectContinueReader{readCloser: req.Body, resp: w}
			}
			req.Header.Del("Expect")
		} else if req.Header.get("Expect") != "" {
			w.sendExpectationFailed()
			break
		}

		// HTTP cannot have multiple simultaneous active requests.[*]
		// Until the server replies to this request, it can't read another,
		// so we might as well run the handler in this goroutine.
		// [*] Not strictly true: HTTP pipelining.  We could let them all process
		// in parallel even if their responses need to be serialized.
		serverHandler{c.server}.ServeHTTP(w, w.req)
		if c.hijacked() {
			return
		}
		w.finishRequest()
		if !w.shouldReuseConnection() {
			if w.requestBodyLimitHit || w.closedRequestBodyEarly() {
				c.closeWriteAndWait()
			}
			break
		}
		c.setState(c.rwc, StateIdle)
	}
}

....

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux { return &ServeMux{m: make(map[string]muxEntry)} }

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = NewServeMux()

...

// Find a handler on a handler map given a path string
// Most-specific (longest) pattern wins
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}

...

// Handle registers the handler for the given pattern.
// If a handler already exists for pattern, Handle panics.
func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern " + pattern)
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if mux.m[pattern].explicit {
		panic("http: multiple registrations for " + pattern)
	}

	mux.m[pattern] = muxEntry{explicit: true, h: handler, pattern: pattern}

	if pattern[0] != '/' {
		mux.hosts = true
	}

	// Helpful behavior:
	// If pattern is /tree/, insert an implicit permanent redirect for /tree.
	// It can be overridden by an explicit registration.
	n := len(pattern)
	if n > 0 && pattern[n-1] == '/' && !mux.m[pattern[0:n-1]].explicit {
		// If pattern contains a host name, strip it and use remaining
		// path for redirect.
		path := pattern
		if pattern[0] != '/' {
			// In pattern, at least the last character is a '/', so
			// strings.Index can't be -1.
			path = pattern[strings.Index(pattern, "/"):]
		}
		url := &url.URL{Path: path}
		mux.m[pattern[0:n-1]] = muxEntry{h: RedirectHandler(url.String(), StatusMovedPermanently), pattern: pattern}
	}
}

*/

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

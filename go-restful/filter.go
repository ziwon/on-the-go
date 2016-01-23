package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id, Name string
}

type UserList struct {
	Users []User
}

// Filters (Container, WebService and Route)

func main() {
	// install a global(DefaultContainer)
	restful.Filter(globalLogging)

	restful.Add(NewUserService())
	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewUserService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	// install a webservice filter
	ws.Filter(webserviceLogging).Filter(measureTime)

	// install a counter filter
	ws.Route(ws.GET("").Filter(NewCountFilter().routeCounter).To(getAllUsers))

	// install 2 chained route filters (processed before calling findUser)
	ws.Route(ws.GET("/{user-id}").Filter(routeLogging).Filter(NewCountFilter().routeCounter).To(findUser))
	return ws
}

// Global Filter
func globalLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[global-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

// WebService Filter
func webserviceLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[webservice-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

// WebService (post-process) Filter - FilterFunction
func measureTime(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	log.Printf("[webservice-filter (timer)] %v \n", time.Now().Sub(now))
}

// Route Filter - FilterFunction
func routeLogging(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[route-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)
}

type CountFilter struct {
	count   int
	counter chan int
}

func NewCountFilter() *CountFilter {
	c := new(CountFilter)
	c.counter = make(chan int)
	go func() {
		for {
			c.count += <-c.counter
		}
	}()
	return c
}

func (c *CountFilter) routeCounter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	c.counter <- 1
	log.Printf("[route-filter (counter)] count:%d", c.count)
	chain.ProcessFilter(req, resp)
}

func getAllUsers(request *restful.Request, response *restful.Response) {
	log.Printf("getAllUsers")
	response.WriteEntity(UserList{[]User{User{"42", "Gandalf"}, User{"3.14", "Pi"}}})
}

func findUser(request *restful.Request, response *restful.Response) {
	log.Printf("findUser")
	response.WriteEntity(User{"42", "Gandalf"})
}

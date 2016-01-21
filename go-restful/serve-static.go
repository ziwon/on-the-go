package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/emicklei/go-restful"
)

var rootdir = "/tmp"

func main() {
	restful.DefaultContainer.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.Route(ws.GET("/static/{subpath:*}").To(staticFromPathParam))
	ws.Route(ws.GET("/static").To(staticFromQueryParam))
	restful.Add(ws)

	println("[go-restful] serving files on http://localhost:8080/static from local /tmp")
	http.ListenAndServe(":8080", nil)
}

func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	actual := path.Join(rootdir, req.PathParameter("subpath"))
	fmt.Println("serving %s ...(from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(resp.ResponseWriter, req.Request, actual)
}

func staticFromQueryParam(req *restful.Request, resp *restful.Response) {
	http.ServeFile(resp.ResponseWriter, req.Request, path.Join(rootdir, req.QueryParameter("resource")))
}

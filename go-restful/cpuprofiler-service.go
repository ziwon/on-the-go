package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

type ProfilingService struct {
	rootPath   string
	cpuprofile string
	cpufile    *os.File
}

func NewProfileService(rootPath string, outputFilename string) *ProfilingService {
	ps := new(ProfilingService)
	ps.rootPath = rootPath
	ps.cpuprofile = outputFilename
	return ps
}

func (p ProfilingService) AddWebServiceTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path(p.rootPath).Consumes("*/*").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/start").To(p.startProfiler))
	ws.Route(ws.GET("/stop").To(p.stopProfiler))
	container.Add(ws)
}

func (p *ProfilingService) startProfiler(req *restful.Request, resp *restful.Response) {
	if p.cpufile != nil {
		io.WriteString(resp.ResponseWriter, "[restful] CPU profiling already running")
		return
	}

	cpufile, err := os.Create(p.cpuprofile)
	if err != nil {
		log.Fatal(err)
	}

	p.cpufile = cpufile
	pprof.StartCPUProfile(cpufile)
	io.WriteString(resp.ResponseWriter, "[restful] CPU profiling started, writing on:"+p.cpuprofile)
}

func (p *ProfilingService) stopProfiler(req *restful.Request, resp *restful.Response) {
	if p.cpufile == nil {
		io.WriteString(resp.ResponseWriter, "[restful] CPU profiling not active")
		return
	}

	pprof.StopCPUProfile()
	p.cpufile.Close()
	p.cpufile = nil
	io.WriteString(resp.ResponseWriter, "[restful] CPU profiling stopped, closing:"+p.cpuprofile)
}

func main() {}

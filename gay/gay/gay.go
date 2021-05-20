package gay

import (
	"fmt"

	"net/http"
)

type HandleFunc func(ctx *Context)


func (router *Router) GET(adr string,handleFunc HandleFunc)  {
	router.AddRouter("GET",adr,handleFunc)
}

func (router *Router) POST(adr string,handleFunc HandleFunc)  {
	router.AddRouter("POST",adr,handleFunc)
}

func (router *Router)AddRouter(method,adr string,handleFunc HandleFunc)  {
	key := method+" "+adr
	router.handler[key]=handleFunc
}


func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + " " + req.URL.Path
	if handler, ok := router.handler[key]; ok {
		context := newContext(w, req)
		handler(context)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
package router

import (
	"excercise2/internal/handler"
	"net/http"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(txHandler handler.HandlerTransaction, eventHandler handler.HandlerEvent) *router {

	router := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	router.addRoute("POST", "/create-tx", txHandler.CreateTxHandler)
	router.addRoute("GET", "/findAll", txHandler.FindAllHandler)

	router.addRoute("POST", "/event/create", eventHandler.CreateEventHandler)
	router.addRoute("GET", "/event/findAll", eventHandler.FindAllEventHandler)
	router.addRoute("GET", "/event/findById", eventHandler.FindByIdEventHandler)

	return router
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.URL.Path]; ok {
		if handler, methodExists := handlers[req.Method]; methodExists {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

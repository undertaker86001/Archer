package server

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

type ServerHandler struct {
	GoRestfulContainer string
}

type director struct {
	name               string
	goRestfulContainer *restful.Container
}

func (d director) ServerHTTP(w http.ResponseWriter, req *http.Request) {

}

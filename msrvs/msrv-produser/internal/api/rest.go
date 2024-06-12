package api

import (
	"net/http"
	"projectX/msrvs/msrv-produser/config"
	"projectX/msrvs/msrv-produser/internal/service"
)

type IRestApi interface {
	Start(addr string)

	Set(w http.ResponseWriter, r *http.Request)
}

func InitRestApi(cnf *config.Config, srv *service.IService) IRestApi {
	return &handler{}
}

type handler struct {
}

func (h *handler) Start(addr string) {
	http.HandleFunc("/", h.Set)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return
	}
}

func (h *handler) Set(w http.ResponseWriter, r *http.Request) {
	return
}

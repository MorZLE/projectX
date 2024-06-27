package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"projectX/msrvs/msrv-produser/internal/service"
	"projectX/msrvs/pkg/cerrors"
	"projectX/msrvs/pkg/model"
	"time"
)

type IRestApi interface {
	Start(addr string)

	Default(w http.ResponseWriter, r *http.Request)
	Set(w http.ResponseWriter, r *http.Request)
}

func InitRestApi(srv service.IService) IRestApi {
	return &handler{
		srv: srv,
	}
}

type handler struct {
	srv service.IService
}

func (h *handler) Start(addr string) {
	http.HandleFunc("/", h.Default)
	http.HandleFunc("/send", h.Set)
	slog.Info("start rest api success addr:", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		slog.Error("rest api start error:", err)
		return
	}

}

func (h *handler) Default(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	return
}

func (h *handler) Set(w http.ResponseWriter, r *http.Request) {
	var err error

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("rest api set error:", err)
			return
		}

		var requestBody model.UserReq
		err = json.Unmarshal(body, &requestBody)
		response(w, err)

		err = h.srv.Set(&ctx, &requestBody)
		response(w, err)
	}
	if err == nil {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
	return
}

func response(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	if errors.Is(err, cerrors.ErrUserNil) || errors.Is(err, cerrors.ErrBodyNil) {
		w.WriteHeader(405)
		w.Write([]byte(err.Error()))
		return
	}

	if errors.Is(err, cerrors.ErrUnmarshalData) {
		w.WriteHeader(405)
		w.Write([]byte("unreachable data"))
		return
	}

	w.WriteHeader(505)
	w.Write([]byte(err.Error()))
	return
}

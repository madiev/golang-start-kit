package handler

import (
    "start_kit/src/application/config"
    "start_kit/src/application/route"
    "net/http"
)


type Handler struct {
    Route     *route.Route
    Config    *config.Config
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    switch path := req.URL.Path; path {
    case "/":
    case "/test":
    default:
    }
}

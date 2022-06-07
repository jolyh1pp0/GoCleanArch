package request

import (
	"GoClearArch/internal/domain/request"
	"context"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	requestService request.Service
}

func NewHandler(service request.Service) *Handler {
	return &Handler{requestService: service}
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc("/request", h.GetRequest)
	router.HandleFunc("/admin/requests", h.GetAdminRequests)
}

func (h *Handler) GetRequest(w http.ResponseWriter, r *http.Request) {
	s, err := h.requestService.GetRequest(context.Background())
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, "ERROR in GetRequest")
	}
	fmt.Fprintln(w, s)
}

func (h *Handler) GetAdminRequests(w http.ResponseWriter, r *http.Request) {
	s, err := h.requestService.GetAdminRequests(context.Background())
	if err != nil {
		log.Fatal(err)
		fmt.Fprintln(w, "ERROR in GetAdminRequest")
	}
	fmt.Fprintln(w, s)
}

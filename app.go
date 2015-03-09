package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct{}
type internalServerErrorHandler struct{}

type MyResponse struct {
	name string `json:"myname"`
}

func (Mux) ServeHTTP(w *http.ResponseWriter, r *http.Request) {
}

func (h *internalServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "error")
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	if err := enc.Encode(&MyResponse{}); nil != err {
        fmt.Fprintf(w, format, ...)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &handler{})
	http.HandleFunc("/", h.ServeHTTP)
	http.ListenAndServe(":8081", mux)
	fmt.Println("Listening on 8081...")
}

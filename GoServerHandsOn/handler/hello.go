package handler

import (
	"net/http"
	"fmt"
	"io"
	"log"
)
type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello{
	return &hello{l}
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	h.l.Println("Hello World")

	d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)

	
}
package handler

import (
	"io"
	"log"
	"net/http"
)

type goodbye struct {
	l *log.Logger
}

func NewGoodby(l *log.Logger) *goodbye {
	return &goodbye{l}
}

func (g *goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Good bye")
	_, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Ooop", http.StatusBadRequest)
		return
	}

	rw.Write([]byte("bye"))

}

package handlers

import (
	"net/http"
)

func (p *Courses) Options(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("[PRE-FLIGHT REQUEST]")
	rw.Header().Add("Content-Type", "application/json")
	rw.Header().Add("Access-Control-Allow-Origin", "*")
	enableCors(&rw)
	return
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

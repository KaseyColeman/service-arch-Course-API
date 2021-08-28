package handlers

import (
	"net/http"
)

// swagger:route PUT /courses courses updateCourse
// Update a courses details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update courses
func (p *Courses) Options(rw http.ResponseWriter, r *http.Request) {

	// fetch the course from the context
	// prod := r.Context().Value(KeyCourse{}).(data.Course)
	p.l.Println("[PRE-FLIGHT REQUEST]")
	// write the no content success header
	// rw.WriteHeader
	// (*rw).Header().Set("Access-Control-Allow-Origin", "*")
	enableCors(&rw)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

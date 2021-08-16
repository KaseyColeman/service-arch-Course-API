package handlers

import (
	"net/http"
	"module/data"
)

// swagger:route POST /courses courses createCourse
// Create a new course
//
// responses:
//	200: courseResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new courses
func (p *Courses) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the course from the context
	prod := r.Context().Value(KeyCourse{}).(*data.Course)

	p.l.Printf("[DEBUG] Inserting course: %#v\n", prod)
	data.AddCourse(prod)
}

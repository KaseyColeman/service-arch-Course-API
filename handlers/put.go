package handlers

import (
	"net/http"
	"module/data"
)

// swagger:route PUT /courses courses updateCourse
// Update a courses details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update courses
func (p *Courses) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the course from the context
	prod := r.Context().Value(KeyCourse{}).(data.Course)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateCourse(prod)
	if err == data.ErrCourseNotFound {
		p.l.Println("[ERROR] course not found", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Course not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}

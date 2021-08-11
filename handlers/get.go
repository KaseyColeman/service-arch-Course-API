package handlers

import (
	"net/http"
	"module/data"
)

// swagger:route GET /courses courses listCourses
// Return a list of Courses from the database
// responses:
//	200: courseResponse

// ListAll handles GET requests and returns all current courses
func (p *Courses) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	rw.Header().Add("Content-Type", "application/json")

	prods := data.GetCourses()

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing course", err)
	}
}

// swagger:route GET /courses/{id} courses listSingleCourse
// Return a list of course from the database
// responses:
//	200: courseResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Courses) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getCourseID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetCourseByID(id)

	switch err {
	case nil:

	case data.ErrCourseNotFound:
		p.l.Println("[ERROR] fetching course", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching course", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing course", err)
	}
}

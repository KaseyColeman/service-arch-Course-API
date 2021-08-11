package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"module/data"
)

// KeyCourses is a key used for the Course object in the context
type KeyCourse struct{}

// Courses handler for getting and updating courses
type Courses struct {
	l *log.Logger
	v *data.Validation
}

// NewCourses returns a new courses handler with the given logger
func NewCourse(l *log.Logger, v *data.Validation) *Courses {
	return &Courses{l, v}
}

// ErrInvalidCoursePath is an error message when the course path is not valid
var ErrInvalidCoursePath = fmt.Errorf("Invalid Path, path should be /courses/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getCourseID returns the course ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getCourseID(r *http.Request) int {
	// parse the course id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}

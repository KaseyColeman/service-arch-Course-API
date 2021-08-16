package handlers

import (
	"context"
	"net/http"
	"module/data"
)

// MiddlewareValidateCourse validates the course in the request and calls next if ok
func (p *Courses) MiddlewareValidateCourse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Course{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing course", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the course
		
		// add thecourse to the context
		ctx := context.WithValue(r.Context(), KeyCourse{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

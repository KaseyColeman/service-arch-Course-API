// Package classification of Course API
//
// Documentation for Course API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "module/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	//Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	//Body ValidationError
}

// A list of courses
// swagger:response coursesResponse
type coursesResponseWrapper struct {
	// All current products
	// in: body
	Body []*data.Course
}

// Data structure representing a single course
// swagger:response courseResponse
type courseResponseWrapper struct {
	// Newly created product
	// in: body
	Body *data.Course
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateCourse createCourse
type courseParamsWrapper struct {
	// Course data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body *data.Course
}

// swagger:parameters listSingleCourse deleteCourse
type courseIDParamsWrapper struct {
	// The id of the course for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

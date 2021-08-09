package data

import (
	"fmt"
)

// ErrCourseNotFound is an error raised when a course can not be found in the database
var ErrCourseNotFound = fmt.Errorf("Course not found")

// Product defines the structure for an API product
// swagger:model
type Course struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the Code for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	Code string `json:"code" validate :"sku"`

	// the name for the instructor
	//
	// required: true
	// max length: 255
	InstructorName string `json:"instructor" validate:"required"`

	// date the course starts
	//
	// required: true
	// max length: 255
	CourseTime string `json:"time" validate:"required"`

	// date the course starts
	//
	// required: true
	// max length: 255
	StartDate string `json:"startDate" validate:"required"`

	// date the course starts
	//
	// required: true
	// max length: 255
	EndDate string `json:"endDate" validate:"required"`

	
}

// Products defines a slice of Product
type Products []*Course

// GetProducts returns all products from the database
func GetProducts() Products {
	return courseList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Course, error) {
	i := findIndexByCourseID(id)
	if id == -1 {
		return nil, ErrCourseNotFound
	}

	return courseList[i], nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateCourse(p Course) error {
	i := findIndexByCourseID(p.ID)
	if i == -1 {
		return ErrCourseNotFound
	}

	// update the product in the DB
	courseList[i] = &p

	return nil
}

// AddProduct adds a new product to the database
func AddCourse(p Course) {
	// get the next id in sequence
	maxID := courseList[len(courseList)-1].ID
	p.ID = maxID + 1
	courseList = append(courseList, &p)
}

// DeleteProduct deletes a product from the database
func DeleteCourse(id int) error {
	i := findIndexByCourseID(id)
	if i == -1 {
		return ErrCourseNotFound
	}

	courseList = append(courseList[:i], courseList[i+1])

	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByCourseID(id int) int {
	for i, p := range courseList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var courseList = []*Course{
	&Course{
		ID:          1,
		Name:        "english",
		Code: "adc-adb-adg",
		InstructorName:   "Instructors name ",
		CourseTime:         "10-11",
		StartDate: "8/2/22",
		EndDate: "4/5/24",
	},
	&Course{
		ID:          1,
		Name:        "math",
		Code: "adc-adb-adg",
		InstructorName:   "Instructors name ",
		CourseTime:         "10-11",
		StartDate: "8/2/22",
		EndDate: "4/5/24",
	},
}

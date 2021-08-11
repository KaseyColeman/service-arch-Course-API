package data

import (
	"fmt"
)

// ErrCourseNotFound is an error raised when a course can not be found in the database
var ErrCourseNotFound = fmt.Errorf("Course not found")

// Course defines the structure for an API course
// swagger:model
type Course struct {
	// the id for the course
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the course

	// the name for this course
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the Code for the course
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	Code string `json:"code"`

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

// Course defines a slice of Course
type Courses []*Course

// GetCourses returns all courses from the database
func GetCourses() Courses {
	return courseList
}

// GetCourseByID returns a single course which matches the id from the
// database.
// If a course is not found this function returns a CourseNotFound error
func GetCourseByID(id int) (*Course, error) {
	i := findIndexByCourseID(id)
	if id == -1 {
		return nil, ErrCourseNotFound
	}

	return courseList[i], nil
}

// UpdateCourse replaces a course in the database with the given
// item.
// If a course with the given id does not exist in the database
// this function returns a CourseNotFound error
func UpdateCourse(p Course) error {
	i := findIndexByCourseID(p.ID)
	if i == -1 {
		return ErrCourseNotFound
	}

	// update the course in the DB
	courseList[i] = &p

	return nil
}

// AddCourse adds a new course to the database
func AddCourse(p Course) {
	// get the next id in sequence
	maxID := courseList[len(courseList)-1].ID
	p.ID = maxID + 1
	courseList = append(courseList, &p)
}

// DeleteCourse deletes a course from the database
func DeleteCourse(id int) error {
	i := findIndexByCourseID(id)
	if i == -1 {
		return ErrCourseNotFound
	}

	courseList = append(courseList[:i], courseList[i+1])

	return nil
}

// findIndex finds the index of a course in the database
// returns -1 when no course can be found
func findIndexByCourseID(id int) int {
	for i, p := range courseList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var courseList = []*Course{
	{
		ID:          1,
		Name:        "english",
		Code: "adc-adb-adg",
		InstructorName:   "Instructors name ",
		CourseTime:         "10-11",
		StartDate: "8/2/22",
		EndDate: "4/5/24",
	},
	{
		ID:          1,
		Name:        "math",
		Code: "adc-adb-adg",
		InstructorName:   "Instructors name ",
		CourseTime:         "10-11",
		StartDate: "8/2/22",
		EndDate: "4/5/24",
	},
}

package data

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	ID int `json:"id"`

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
	InstructorName string `json:"instructorName"`

	// date the course starts
	//
	// required: true
	// max length: 255
	CourseTime string `json:"courseTime"`

	// date the course starts
	//
	// required: true
	// max length: 255
	StartDate string `json:"startDate"`

	// date the course starts
	//
	// required: true
	// max length: 255
	EndDate string `json:"endDate"`	
}

// Course defines a slice of Course
type Courses []*Course

var courseList = []*Course{ }

// GetCourses returns all courses from the database
func GetCourses() Courses {
	courseList= courseList[:0]
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://connection_user:testpassword@cluster0.wql1e.mongodb.net/course?retryWrites=true&w=majority"))
	contex, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client.Connect(contex)
	courseDatabase := client.Database("course")
	courseCatalogCollection := courseDatabase.Collection("coursecatalog")
	cursor, err := courseCatalogCollection.Find(contex, bson.M{})
	var courses []bson.M
	if err = cursor.All(contex, &courses); err != nil { log.Fatal(err) }
	
	
	for _, courseLoop := range courses {
		courseL := &Course{ID:1,Name:"",Code:"",InstructorName:"",CourseTime:"",StartDate: "",EndDate: "", }
		str := fmt.Sprint(courseLoop["CourseTime"])
		courseL.CourseTime= str
		str=fmt.Sprint(courseLoop["Name"])
		courseL.Name = str
		str=fmt.Sprint(courseLoop["Code"])
		courseL.Code = str
		str=fmt.Sprint(courseLoop["InstructorName"])
		courseL.InstructorName=str
		str=fmt.Sprint(courseLoop["StartDate"])
		courseL.StartDate=str
		str=fmt.Sprint(courseLoop["EndDate"])
		courseL.EndDate=str
		str=fmt.Sprint(courseLoop["ID"])
		intVar, _ := strconv.Atoi(str)
		courseL.ID=intVar
		fmt.Print(courseL)
		courseList = append(courseList, courseL)
		fmt.Print("F")
	}

	defer client.Disconnect(contex)
	client.Ping(contex, readpref.Primary())
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
func AddCourse(p *Course) {
	GetCourses()
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://connection_user:testpassword@cluster0.wql1e.mongodb.net/course?retryWrites=true&w=majority"))
	contex, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client.Connect(contex)
	courseDatabase := client.Database("course")
	courseCatalogCollection := courseDatabase.Collection("coursecatalog")
	
	// get the next id in sequence
	maxID := courseList[len(courseList)-1].ID
	p.ID = maxID + 1

	courseResult, err := courseCatalogCollection.InsertOne(contex, bson.D{
		{Key: "ID", Value: p.ID},
		{Key: "Name",Value: p.Name},
		{Key: "Code", Value: p.Code},
		{Key: "CourseTime", Value: p.CourseTime},
		{Key: "InstructorName", Value: p.InstructorName},
		{Key: "StartDate",Value: p.StartDate},
		{Key: "EndDate", Value:p.EndDate},
	})

	if err != nil{log.Fatal(err)}

	fmt.Println(courseResult.InsertedID)
	courseList = append(courseList, p)
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

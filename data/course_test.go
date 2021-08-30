package data

import "testing"

func TestCourseGet(t *testing.T) {
	arry := GetCourses()
	if arry ==nil{
 		t.Errorf("Time to panic now")
	}
}

func TestSingleCourseGet(t *testing.T) {
	_, err := GetCourseByID(4)
	if err == nil{
 		t.Errorf("Time to panic now")
	}
}







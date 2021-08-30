package data

import "testing"

func TestCourseGet(t *testing.T) {
	arry := GetCourses()
	if arry ==nil{
 		t.Errorf("Time to panic now")
	}
}

func TestSingleCourseGet(t *testing.T) {
	arry:= findIndexByCourseID(1)
	if arry != -1{
 		t.Errorf("Time to panic now")
	}
}








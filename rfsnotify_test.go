package rfsnotify

import (
	"testing"
)

func TestDeletePath_IndexSmallerThanLen_DeletePathAtIndex(t *testing.T) {
	var index = 1
	var paths = getSamplePaths()
	var result = deletePath(paths, index)

	if len(result) != 2 {
		t.Fatal("Len(result) should not be bigger than len(paths)")
	}
}

func TestDeletePath_IndexGreaterThanLen_Panics(t *testing.T) {
	var index = 3
	var paths = getSamplePaths()

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("test must have paniced!")
		}
	}()
	_ = deletePath(paths, index)
}

func TestDeletePath_IndexIsLenMinusOne_DeletesLastElement(t *testing.T) {
	var index = 2
	var paths = getSamplePaths()

	result := deletePath(paths, index)

	if len(result) != 2 {
		t.Log("len(result) should have been two")
		t.Fail()
	}
	if result[1] != "world" {
		t.Log("result[1] must have been world")
		t.Fail()
	}
}

func getSamplePaths() []string{
	return []string{"hello", "world", "mars"}
}

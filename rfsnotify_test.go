package rfsnotify

import (
	"io/ioutil"
	"os"
	"path"
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

func TestInclude_AddingPaths_WorksProperly(t *testing.T) {
	var watcher = &Watcher{}
	watcher.Include("test1", "test2")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}
	if watcher.filePaths[0] != "test1" && watcher.filePaths[1] != "test2" {
		t.Error("watcher.filesPaths does not have correct items.")
	}
}

func TestInclude_AddingNothing_ReturnsNil(t *testing.T) {
	var watcher = new(Watcher)
	if watcher.filePaths != nil {
		t.Error(`watcher.filePaths must be nil`)
	}
}

func TestExclude_RemovingExistingItems_ItemsRemoved(t *testing.T){
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file2.txt", "file3.txt")
	watcher.Exclude("file2.txt")

	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}

	if watcher.filePaths[1] != "file3.txt"{
		t.Error("watcher.filePaths[1] must be file3.txt")
	}
}

func TestExclude_RemovingNonExistingItem_SliceRemainedTheSame(t *testing.T) {
	var watcher Watcher
	watcher.Include("file1.txt", "file2.txt")
	watcher.Exclude("file3.txt")

	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}
}

func  TestInclude_AddingDuplicateItem_DuplicateItemsNotAdded( t *testing.T) {
	var watcher Watcher
	watcher.Include("file1.txt")
	watcher.Include("file1.txt")
	watcher.Include("file2.txt")

	if len(watcher.filePaths) != 2{
		t.Error("len(watcher.filePaths) must be 2")
	}
}

func  TestInclude_AddingDuplicateItemAtTheSameTime_DuplicateItemsNotAdded( t *testing.T) {
	var watcher Watcher
	watcher.Include("file1.txt", "file1.txt", "file2.txt")
	if len(watcher.filePaths) != 2{
		t.Error("len(watcher.filePaths) must be 2")
	}
}

func TestNewWatcher_GivenDirectory_ReturnsAllFiles(t *testing.T) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil{
		t.Fatal("cannot create a temp directory")
	}

	defer os.RemoveAll(dir) //clean up

	err = os.MkdirAll(path.Join(dir, "dir1","dir2", "dir3"), os.ModePerm)

	if err != nil{
		t.Fatal("cannot create the temp directories")
	}

	var tempFiles  = []string {path.Join(dir, "dir1", "file1"),
		path.Join(dir, "dir1", "file2"),
		path.Join(dir, "dir1", "dir2", "file3"),
		path.Join(dir, "dir1", "dir2", "file4"),
		path.Join(dir, "dir1", "dir2", "dir3", "file5"),
		path.Join(dir, "dir1", "dir2", "dir3", "file6"),
	}

	for _, fileName := range tempFiles{
		err := ioutil.WriteFile(fileName, []byte("hello world"), os.ModePerm)
		if err != nil{
			t.Error("cannot create file " + fileName)
		}
	}

	watcher := NewWatcher(dir, true, nil)
	t.Log(watcher.filePaths)
	if len(watcher.filePaths) != 6{
		t.Error("watcher didn't find all the files.")
	}
}

func getSamplePaths() []string {
	return []string{"hello", "world", "mars"}
}

package rfsnotify

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestInclude_AddingPaths_WorksProperly(t *testing.T) {
	var watcher = &Watcher{}
	watcher.Include("test1", "test2")
	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}
	if !watcher.filePaths["test1"] || !watcher.filePaths["test2"] {
		t.Error("watcher.filesPaths does not have correct items.")
	}
}

func TestInclude_AddingNothing_ReturnsNil(t *testing.T) {
	var watcher = new(Watcher)
	if watcher.filePaths != nil {
		t.Error(`watcher.filePaths must be nil`)
	}
}

func TestExclude_RemovingExistingItems_ItemsRemoved(t *testing.T) {
	var watcher = new(Watcher)
	watcher.Include("file1.txt", "file2.txt", "file3.txt")
	watcher.Exclude("file2.txt")

	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}

	if !watcher.filePaths["file3.txt"] {
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

func TestNewWatcher_GivenDirectory_ReturnsAllFiles(t *testing.T) {
	dir := setupWatchedDirectory(t)
	defer os.RemoveAll(dir)
	watcher := NewWatcher(dir, true, nil)
	if len(watcher.filePaths) != 6 {
		t.Error("watcher didn't find all the files.")
	}
}

func TestNewWatcher_GivenDirectoryAndInclude_ReturnsAllFiles(t *testing.T) {
	dir := setupWatchedDirectory(t)
	watcher := NewWatcher(dir, true, nil)
	watcher.Include("includedFile1.txt", "includedFile1.txt")
	defer os.RemoveAll(dir)

	if len(watcher.filePaths) != 7 {
		t.Error("len(watcher.filePaths) must be 7")
	}

	watcher.Include("includedFile1.txt")

	if len(watcher.filePaths) != 7 {
		t.Error("Duplicated files should be allowed.")
	}
}

func setupWatchedDirectory(t *testing.T) string {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		t.Fatal("cannot create a temp directory")
	}

	//clean up
	err = os.MkdirAll(path.Join(dir, "dir1", "dir2", "dir3"), os.ModePerm)
	if err != nil {
		t.Fatal("cannot create the temp directories")
	}
	var tempFiles = []string{path.Join(dir, "dir1", "file1"),
		path.Join(dir, "dir1", "file2"),
		path.Join(dir, "dir1", "dir2", "file3"),
		path.Join(dir, "dir1", "dir2", "file4"),
		path.Join(dir, "dir1", "dir2", "dir3", "file5"),
		path.Join(dir, "dir1", "dir2", "dir3", "file6"),
	}
	for _, fileName := range tempFiles {
		err := ioutil.WriteFile(fileName, []byte("hello world"), os.ModePerm)
		if err != nil {
			t.Error("cannot create file " + fileName)
		}
	}
	return dir
}

func TestInclude_AddingDuplicateItem_DuplicateItemsNotAdded(t *testing.T) {
	var watcher Watcher
	watcher.Include("file1.txt")
	watcher.Include("file1.txt")
	watcher.Include("file2.txt")

	if len(watcher.filePaths) != 2 {
		t.Error("len(watcher.filePaths) must be 2")
	}
}

func TestRefresh_AddingNewFile_GetAllFile(t *testing.T) {
	dir := setupWatchedDirectory(t)
	watcher := NewWatcher(dir, true, nil)
	defer os.RemoveAll(dir)

	if len(watcher.filePaths) != 6 {
		t.Error("len(watcher.filePaths) must be 7")
	}

	err := ioutil.WriteFile(path.Join(dir, "file_new"), []byte("hello again"), os.ModePerm)
	if err != nil{
		t.Error(err)
	}

	watcher.Refresh()
	if len(watcher.filePaths) != 7{
		t.Error("Refresh method is not discovering the new files")
	}
}

package rfsnotify

import (
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

func TestInclude_AddingDuplicateItem_DuplicateItemsNotAdded(t *testing.T) {
	var watcher Watcher
	watcher.Include("file1.txt")
	watcher.Include("file1.txt")

	if len(watcher.filePaths) != 1 {
		t.Error("len(watcher.filePaths) must be 1")
	}
}

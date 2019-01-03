package rfsnotify

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

type Event int

const (
	Delete Event = iota
	Create
	Rename
	Write
)

type Watcher struct {
	Path      string
	Recursive bool
	Events    []Event
	// todo think about the usefulness of this backing slice. Because we can use fsnotify.Watcher internal backing slice.
	filePaths map[string]bool
	internalWatcher *fsnotify.Watcher
}

// Adds new files to the internal watch list to track.
func (w *Watcher) Include(paths ...string) {
	if w.filePaths == nil {
		w.filePaths = make(map[string]bool)
	}

	if w.internalWatcher == nil{
		w.internalWatcher, _ = fsnotify.NewWatcher()
	}

	for _, newPath := range paths {
		if !w.filePaths[newPath] {
			w.filePaths[newPath] = true
			// todo handle error
			w.internalWatcher.Add(newPath)
		}
	}
}

// Excludes paths from internal watch list.
func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
		// todo  handle the error
		w.internalWatcher.Remove(path)
	}
}

// Finds newly added files in given path.
func (w *Watcher) Refresh() {
	initFilePaths(w)
}

// Creates a new Watcher object and initializes the internal watch list
// based on the given path.
func NewWatcher(path string, events []Event) *Watcher {
	var watcher = &Watcher{
		Path:      path,
		Events:    events,
	}

	initFilePaths(watcher)

	var fsWatcher, err = fsnotify.NewWatcher()

	if err != nil{
		fmt.Print(err)
	}

	watcher.internalWatcher = fsWatcher

	for path := range watcher.filePaths {
		// todo handle error.
		fsWatcher.Add(path)
	}

	return watcher
}

func initFilePaths(w *Watcher) {
	givenFileInfo, err := os.Stat(w.Path)
	if err != nil {
		panic(err)
	}

	var allFilePaths []string

	switch mode := givenFileInfo.Mode(); {
	case mode.IsDir():
		allFilePaths = getAllFiles(w.Path)
		w.Include(allFilePaths...)
	case mode.IsRegular():
		w.Include(w.Path)
	}
}

func getAllFiles(dirPath string) []string {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return files
}

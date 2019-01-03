package rfsnotify

import (
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
	filePaths map[string]bool
}

// Adds new files to the internal watch list to track.
func (w *Watcher) Include(paths ...string) {
	if w.filePaths == nil {
		w.filePaths = make(map[string]bool)
	}
	for _, newPath := range paths {
		if !w.filePaths[newPath] {
			w.filePaths[newPath] = true
		}
	}
}

// Excludes paths from internal watch list.
func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
	}
}

// Finds newly added files in given path.
func (w *Watcher) Refresh() {
	initFilePaths(w)
}

// Creates a new Watcher object and initializes the internal watch list
// based on the given path.
func NewWatcher(path string, recursive bool, events []Event) *Watcher {
	var watcher = &Watcher{
		Path:      path,
		Recursive: recursive,
		Events:    events,
	}

	initFilePaths(watcher)

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
		// todo check this logic later.
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

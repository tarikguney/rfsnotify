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

func (w *Watcher) Exclude(paths ...string) {
	for _, path := range paths {
		delete(w.filePaths, path)
	}
}

func NewWatcher(path string, recursive bool, events []Event) *Watcher {
	var watcher = &Watcher{
		Path:      path,
		Recursive: recursive,
		Events:    events,
	}

	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	var allFilePaths []string

	switch mode := fi.Mode(); {
	case mode.IsDir():
		allFilePaths = walkDir(path)
		watcher.Include(allFilePaths...)
	case mode.IsRegular():
		watcher.Include(path)
	}

	return watcher
}

func walkDir(dirPath string) []string {
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
package rfsnotify

import (
	"fmt"
)

type Event int

const(
	Delete Event = iota
	Create
	Rename
	Write
)

type Watcher struct {
	Path string
	Recursive bool
	Events []Event
	filePaths []string
}

func (w *Watcher) Include(path ...string){
	for _,existingPath := range w.filePaths {
		for _, newPath := range path{
			if existingPath == newPath{
				return
			}
		}
	}
	w.filePaths = append(w.filePaths, path...)
}

func (w *Watcher) Exclude(path ...string){
	for _, value := range path{
		for i,v := range w.filePaths{
			if  value == v {
				w.filePaths = deletePath(w.filePaths, i)
			}
		}
	}
}

func deletePath(paths []string, index int) []string {
	if index > len(paths) -1 {
		panic(fmt.Sprintf("index %v is bigger than the size of the paths slice!", index))
	}

	if index < len(paths) -1 {
		return append(paths[:index], paths[index+1:]...)
	}

	return paths[:index]
}
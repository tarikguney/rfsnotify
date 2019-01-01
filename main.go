package rfsnotify

import (
	"fmt"
)

type Event int

const(
	Delete Event = 0
	Create Event = 1
	Rename Event = 2
	Write Event = 3
)

type Watcher struct {
	Path string
	Recursive bool
	Events []Event
	filePaths []string
}

func (w *Watcher) Include(path ...string){
	w.filePaths = append(w.filePaths, path...)
}

func (w *Watcher) Exclude(path ...string){
	var indices = make([]int, 5)
	for _, value := range path{
		for i,v := range w.filePaths{
			if  value == v {
				indices = append(indices, i)
			}
		}
	}

	for i := range indices{
		deletePath(w.filePaths, indices[i])
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
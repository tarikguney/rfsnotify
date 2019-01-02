package rfsnotify

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
	filePaths map[string]bool
}

func (w *Watcher) Include(paths ...string){
	if w.filePaths == nil{
		w.filePaths = make(map[string]bool)
	}
	for _,  newPath := range paths {
		if !w.filePaths[newPath]{
			w.filePaths[newPath] = true
		}
	}
}

func (w *Watcher) Exclude(paths ...string){
	for _, path := range paths {
		delete(w.filePaths, path)
	}
}
package rfsnotify

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
	for value := range path{
		for v,i := range w.filePaths{
			/*if  value == v {
				w.filePaths[]
			}*/
		}
	}

	// todo remove paths from w.filePaths
}
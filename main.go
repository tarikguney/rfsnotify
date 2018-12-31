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
	var indices = []int{}
	for _, value := range path{
		for i,v := range w.filePaths{
			if  value == v {
				indices = append(indices, i)
			}
		}
	}

	for i := range indices{
		delete(w.filePaths, indices[i])
	}
}

func delete(paths []string, index int){
	paths = append(paths[:index], paths[index+1:]...)
}
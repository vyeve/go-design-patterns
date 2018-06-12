package workers_pool

import "time"

type Dispatcher interface {
	LaunchWorker(WorkerLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	inCh chan Request
}

func (d *dispatcher) LaunchWorker(w WorkerLauncher) {
	w.LaunchWorker(d.inCh)
}

func (d *dispatcher) MakeRequest(r Request) {
	select {
	case d.inCh <- r:
	case <-time.After(time.Second * 5):
		return
	}
}

func (d *dispatcher) Stop() {
	close(d.inCh)
}

func NewDispatcher(bufSize int) Dispatcher {
	return &dispatcher{
		inCh: make(chan Request, bufSize),
	}
}

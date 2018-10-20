package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func newWorker(interval time.Duration,r *Runner) *Worker{
	return &Worker{
		ticker:time.NewTicker(interval * time.Second),
		runner:r,
	}
}

func(w *Worker) startWorker(){
	for  {
		select {
			case <-w.ticker.C:
				go 	w.runner.startAll()
		}
	}
}

func Start(){
	r := newRunner(30, false, clearVidRecDispatcher, clearCidRecExecutor)
	w := newWorker(30, r)
	go w.startWorker()
}

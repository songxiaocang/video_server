package taskrunner

type Runner struct {
	controller controlChann
	err        controlChann
	data       dataChan
	dataSize   int
	longLived  bool
	dispatcher fn
	executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		controller: make(chan string, 1),
		err:        make(chan string, 1),
		data:       make(chan interface{}, size),
		dataSize:   size,
		longLived:  longlived,
		dispatcher: d,
		executor:   e,
	}
}

func (r *Runner) startDispatcher() {
	defer func() {
		if !r.longLived {
			close(r.controller)
			close(r.data)
			close(r.err)
		}
	}()

	for {
		select {
		case c := <-r.controller:
			if c == READY_TO_DISPATCHER {
				err := r.dispatcher(r.data)
				if err != nil {
					r.err <- CLOSE
				} else {
					r.controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.executor(r.data)
				if err != nil {
					r.err <- CLOSE
				} else {
					r.controller <- READY_TO_DISPATCHER
				}

			}
		case err := <-r.err:
			if err == CLOSE {
				return
			}
		default:
		}
	}

}

func (r *Runner) StartAll() {
	r.controller <- READY_TO_DISPATCHER
	r.startDispatcher()
}

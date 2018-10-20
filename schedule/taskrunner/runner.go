package taskrunner

type Runner struct {
	controller controlChann
	err controlChann
	data dataChan
	dataSize int
	longLived bool
	dispatcher fn
	executor fn
}

func newRunner(size int,longLived bool,d fn,e fn) *Runner{
	return &Runner{
		controller:make(chan string,1),
		err:make(chan string,1),
		data:make(chan  interface{},1),
		dataSize:size,
		longLived:longLived,
		dispatcher:d,
		executor:e,
	}
}

func (r *Runner) startDispatcher(){
	defer func() {
		if !r.longLived {
			close(r.controller)
			close(r.err)
			close(r.data)
		}
	}()



	for {
		select {
			case c:=<-r.controller:
				if c==READY_TO_DISPATCHER {
					if err := r.dispatcher(r.data);err!=nil{
						r.err<-CLOSE
					}else {
						r.controller<-READY_TO_EXECUTE
					}
				}

				if c==READY_TO_EXECUTE {
					if err := r.executor(r.data);err!=nil{
						r.err<-CLOSE
					}else{
						r.controller<-READY_TO_DISPATCHER
					}

				}
			case err:=<-r.err:
				if err==CLOSE {
					return
				}
			default:
		}
	}


}


func (r *Runner) startAll(){
	r.controller<- READY_TO_DISPATCHER
	r.startDispatcher()
}
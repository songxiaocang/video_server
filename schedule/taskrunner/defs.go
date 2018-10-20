package taskrunner

const (
	READY_TO_DISPATCHER="d"
	READY_TO_EXECUTE="e"
	CLOSE="c"

	VIDEO_PATH = "E:\\videos\\"
)

type controlChann chan string

type dataChan chan interface{}

type fn func(dc dataChan) error


package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("data sent: %v", i)
		}

		return nil
	}

	e := func(dc dataChan) error {
	forLoop:
		for {
			select {
			case i := <-dc:
				log.Printf("data execute: %v", i)
			default:
				break forLoop
			}
		}

		return errors.New("executor")
	}

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()

	time.Sleep(3 * time.Second)

}

package taskrunner

import (
	"errors"
	"log"
	"os"
	"sync"
	"video_server/schedule/dbops"
)

func DelVideoById(id string) error {
	var videoPath = VIDEO_PATH + id
	err := os.Remove(videoPath)
	if err != nil {
		log.Printf("del video from disk err: %v", err)
		return err
	}

	return nil
}

func clearVidRecDispatcher(dc dataChan) error {

	ids, err := dbops.ReadVidRes(3)

	if err != nil {
		log.Printf("clear video dispatcher error: %v", err)
		return err
	}

	if len(ids) == 0 {
		return errors.New("all task has finished: %v")
	}

	for _, id := range ids {
		dc <- id
	}
	return nil
}

func clearVidRecExecutor(dc dataChan) error {

	errMap := &sync.Map{}

	var err error

forLoop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := dbops.DelVidRecs(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}

				if err := DelVideoById(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)

		default:
			break forLoop
		}
	}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		} else {
			return true
		}
	})

	return err

}

package taskrunner

import (
	"log"
	"os"
	"sync"
)

func delVideoById(id string) error{
	var videoPath =VIDEO_PATH+id
	err := os.Remove(videoPath)
	if err !=nil{
		log.Printf("del video from disk err: %v",err)
		return err
	}

	return nil
}

func clearVidRecDispatcher(dc dataChan) error{


	ids,err := readVidRes(3)

	if err!=nil || len(ids) == 0 {
		log.Printf("internal error: %v",err)
		return err
	}

	for _,id:= range ids {
		dc <- id
	}
	return nil
}


func clearCidRecExecutor(dc dataChan) error{

	errMap := &sync.Map{}

	var err error

	forLoop:for {
				select {
					case dc<-dc:
						go func(id interface{}) {
							if err := delVidRecs(id.(string)); err!=nil{
								errMap.Store(id,err)
							}

							if err := delVideoById(id.(string)); err!=nil{
								errMap.Store(id,err)
							}
						}(dc)

					default:
						break forLoop
				}
			}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err!=nil {
			return false
		}else {
			return true
		}
	})


	return err


}
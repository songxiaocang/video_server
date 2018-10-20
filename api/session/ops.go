package session

import (
	"../dbops"
	"../defs"
	"log"
	"sync"
	"time"
)
var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func timeInMills() int64{
	return time.Now().UnixNano()/1000000
}

func LoadSessionFromDB(){
	sessions, e := dbops.RetrieveAllSessions()

	if e!=nil {
		log.Printf("loadSessionFromDB error: %v",e)
		return
	}

	sessions.Range(func(k,v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})
}

func GenerateNewSessionId(uname string) string{
	uuid,err := utils.NewUUID()
	if err!=nil {
		log.Printf("generate uuid error: %v",err)
		return ""
	}
	//session valid time：30 min
	ttl := timeInMills()+30*60*1000
	ss := &defs.SimpleSession{UserName:uname, TTL:ttl}
	sessionMap.Store(uuid,ss)

	err2 := dbops.InsertSession(uuid, ttl, uname)
	if err2!=nil {
		log.Printf("insertSession error: %v",err2)
		return ""
	}
	return uuid
}

func IsSessionExpired(sid string) (string,bool){
	value, ok := sessionMap.Load(sid)
	if ok {
		ttl := value.(*defs.SimpleSession).TTL
		if ttl<= timeInMills() {
			return "",false
		}
		DeleteSession(sid)

		return value.(*defs.SimpleSession).UserName,true
	}

	return "",true
}

func DeleteSession(sid string){
	sessionMap.Delete(sid)
	dbops.DelSession(sid)
}

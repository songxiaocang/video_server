package session

import (
	"log"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func timeInMills() int64 {
	return time.Now().UnixNano() / 1000000
}

func LoadSessionFromDB() {
	sessions, e := dbops.RetrieveAllSessions()

	if e != nil {
		log.Printf("loadSessionFromDB error: %v", e)
		return
	}

	sessions.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	uuid, err := utils.NewUUID()
	if err != nil {
		log.Printf("generate uuid error: %v", err)
		return ""
	}
	//session valid time：30 min
	ttl := timeInMills() + 30*60*1000
	ss := &defs.SimpleSession{UserName: uname, TTL: ttl}
	sessionMap.Store(uuid, ss)

	err2 := dbops.InsertSession(uuid, ttl, uname)
	if err2 != nil {
		log.Printf("insertSession error: %v", err2)
		return ""
	}
	return uuid
}

func IsSessionExpired(sid string) (string, bool) {
	value, ok := sessionMap.Load(sid)
	ct := timeInMills()
	if ok {
		ttl := value.(*defs.SimpleSession).TTL
		if ttl < ct {
			DeleteSession(sid)
			return "", true
		}
		return value.(*defs.SimpleSession).UserName, false
	} else {
		ss, e := dbops.RetrieveSession(sid)
		if e != nil || ss == nil {
			return "", true
		}
		if ss.TTL < ct {
			DeleteSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.UserName, false

	}

	return "", true

}

func DeleteSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DelSession(sid)
}

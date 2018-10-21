package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, userName string) error {

	ttlStr := strconv.FormatInt(ttl, 10)

	stmtIns, e := dbConn.Prepare("INSERT INTO sessions(session_id,TTL,login_name) VALUES(?,?,?)")

	if e != nil {
		log.Printf("insert session error: %v", e)
		return e
	}

	_, err := stmtIns.Exec(sid, ttlStr, userName)

	if err != nil {
		log.Printf("insert session error: %v", err)
		return err
	}

	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	stmtOut, e := dbConn.Prepare("SELECT * FROM sessions WHERE session_id=?")
	if e != nil {
		log.Printf("retriveSession error: %v", e)
		return nil, e
	}

	var (
		sessionId string
		userName  string
		ttlStr    string
	)

	err := stmtOut.QueryRow(sid).Scan(&sessionId, &userName, &ttlStr)

	if err != nil && err != sql.ErrNoRows {
		log.Print("retrieveSession error")
		return nil, err
	}

	ttl, err2 := strconv.ParseInt(ttlStr, 10, 64)
	if err2 != nil {
		log.Print("retrieveSession error")
		return nil, err
	}

	simpleSession := &defs.SimpleSession{userName: userName, TTL: ttl}

	return simpleSession, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	stmtOuts, e := dbConn.Prepare("SELECT * FROM sessions")
	if e != nil {
		log.Printf("retrieveAllSessions error: %v", e)
		return nil, e
	}

	rows, err := stmtOuts.Query()
	if err != nil {
		log.Printf("retrieveAllSessions error: %v", e)
		return nil, e
	}

	var (
		id       string
		userName string
		ttlStr   string
	)

	var sessionMap *sync.Map
	if rows.Next() {
		err := rows.Scan(&id, &userName, &ttlStr)
		if err != nil {
			log.Printf("retrieveAllSessions error: %v", err)
			return nil, err
		}

		ttl, err2 := strconv.ParseInt(ttlStr, 10, 64)
		if err2 != nil {
			log.Printf("retrieveAllSessions error: %v", err2)
			return nil, err2
		}

		ss := &defs.SimpleSession{userName: userName, TTL: ttl}

		sessionMap.Store(id, ss)
	}

	return sessionMap, nil

}

func DelSession(sid string) error {
	stmtDel, e := dbConn.Prepare("DELETE FROM sessions WHEREE session_id=?")
	if e != nil {
		log.Printf("delSession error: %v", e)
		return e
	}

	_, err := stmtDel.Exec(sid)

	if err != nil {
		log.Printf("delSession error: %v", e)
		return e
	}

	return nil
}

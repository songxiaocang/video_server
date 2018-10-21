package dbops

import "log"

func AddVideoRec(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec(video_id) VALUES(?)")
	if err != nil {
		log.Printf("insert error: %v", err)
		return err
	}

	_, err2 := stmtIns.Exec(vid)
	if err2 != nil {
		log.Printf("insert error: %v", err2)
		return err2
	}

	return nil
}

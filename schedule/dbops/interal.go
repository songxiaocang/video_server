package dbops

import "log"

func readVidRes(count int) ([]string,error){
	stmtIns, e := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")

	var ids []string
	if e!=nil {
		log.Printf("select video_del_rec error: %v",e)
		return ids,e
	}
	rows, err2 := stmtIns.Query(count)
	if  err2!=nil{
		log.Printf("select video_del_rec error: %v",err2)
		return ids,err2
	}


	for rows.Next()  {
		var id string
		if err := rows.Scan(&id);err!=nil{
			log.Printf("select video_del_rec error: %v",err)
			return ids,err
		}
		ids = append(ids, id)
	}

	defer stmtIns.Close()

	return  ids,nil

}


func delVidRecs(vid string) error{
	stmtDel, e := dbConn.Prepare("DEL FROM video_del_rec WHERE video_id=?")
	if e!=nil {
		log.Printf("del video_del_rec error: %v",e)
		return e
	}

	_, err := stmtDel.Exec(vid)
	if err!=nil {
		log.Printf("del video_del_rec error: %v",err)
		return err
	}

	return nil
}
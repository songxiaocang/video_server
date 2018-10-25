package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func openConn() *sql.DB {
	return nil
}

func AddUserCredential(loginName string, pwd string) error {
	querySql := "INSERT INTO users(login_name,pwd) VALUES (?,?)"
	stmtIns, err := dbConn.Prepare(querySql)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	_, e := dbConn.Exec(querySql, loginName, pwd)
	if e != nil {
		log.Printf("%v", err)
		return e
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginUsername string) (string, error) {
	selectSql := "SELECT pwd FROM users where login_name=?"
	stmtOut, e := dbConn.Prepare(selectSql)
	if e != nil {
		log.Printf("%v", e)
		return "", e
	}
	var pwd string
	err := dbConn.QueryRow(selectSql, loginUsername).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v", err)
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

func GetUser(loginUsername string) (*defs.User, error) {
	selectSql := "SELECT * FROM users where login_name=?"
	stmtOut, e := dbConn.Prepare(selectSql)
	if e != nil {
		log.Printf("%v", e)
		return nil, e
	}
	var (
		id        int
		loginName string
		pwd       string
	)
	err := dbConn.QueryRow(selectSql, loginUsername).Scan(&id, &loginName, &pwd)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("%v", err)
		return nil, err
	}
	user := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}

	defer stmtOut.Close()
	return user, nil
}

func DeleteUser(LoginUsername string, pwd string) error {
	deleteSql := "DELETE FROM users where login_name=?"
	stmtDel, e := dbConn.Prepare(deleteSql)
	if e != nil {
		log.Printf("%v", e)
		return e
	}
	_, err := dbConn.Exec(deleteSql, LoginUsername)
	if err != nil {
		log.Printf("%v", e)
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddVideoInfo(authorId int, name string) (*defs.VideoInfo, error) {
	uuid, err := utils.NewUUID()
	if err != nil {
		log.Printf("generate uuid error: %v", err)
	}

	now := time.Now()

	formatTime := now.Format("Jan 02 2006, 15:04:05")

	var addSql string = "INSERT INTO video_info(id,author_id,name,display_ctime) VALUES(?,?,?,?)"
	stmtIns, e := dbConn.Prepare(addSql)

	if e != nil {
		log.Printf("insert error: %v", e)
		return nil, e
	}

	_, err = dbConn.Exec(addSql, uuid, authorId, name, formatTime)

	if err != nil {
		log.Printf("insert video_info error: %s", err)
		return nil, err
	}

	defer stmtIns.Close()
	res := &defs.VideoInfo{Id: uuid, AuthorId: authorId, Name: name, DisplayCtime: formatTime}
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	selectSql := "SELECT author_id,name,display_ctime FROM video_info WHERE id=?"
	stmtOut, e := dbConn.Prepare(selectSql)
	if e != nil {
		log.Printf("select video_info error: %v", e)
		return nil, e
	}

	var (
		authorId     int
		name         string
		displayCtime string
	)
	err := dbConn.QueryRow(selectSql, vid).Scan(&authorId, &name, &displayCtime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("select video_info error: %v", err)
		return nil, err
	}

	if err == sql.ErrNoRows {
		log.Printf("none video_info error : %v", err)
		return nil, nil
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: authorId, Name: name, DisplayCtime: displayCtime}

	defer stmtOut.Close()

	return res, nil

}

func ListAllVideos(vid string, from, to int) ([]*defs.VideoInfo, error) {
	selectAllSql := `SELECT * FROM video_info WHERE create_time>FROM_UNIXTIME(?) AND create_time<FROM_UNIXTIME(?) ORDER BY create_time DESC `

	stmtOut, e := dbConn.Prepare(selectAllSql)
	if e != nil {
		log.Printf("selectAll videos error: %s", e)
		return nil, e
	}

	rows, err := dbConn.Query(selectAllSql, vid, from, to)
	if err != nil {
		log.Printf("selectAll videos error: %s", e)
		return nil, e
	}

	var (
		id          string
		author      int
		name        string
		displayTime string
	)

	var videos []*defs.VideoInfo

	if rows.Next() {
		err := rows.Scan(&id, &author, &name, &displayTime)
		if err != nil {
			log.Printf("selectAll videos error: %s", e)
			return nil, err
		}
		video := &defs.VideoInfo{Id: id, AuthorId: author, Name: name, DisplayCtime: displayTime}
		videos = append(videos, video)

	}

	defer stmtOut.Close()
	return videos, nil
}

func DelVideoInfo(vid string) error {
	delSql := "DELETE FROM video_info WHERE id=?"
	stmtDel, e := dbConn.Prepare(delSql)
	if e != nil {
		log.Printf("del video_info error: %v", e)
		return e
	}

	_, err := dbConn.Exec(delSql, vid)
	if err != nil {
		log.Printf("del video_info error: %v", err)
		return err
	}

	defer stmtDel.Close()

	return nil
}

//comments

func AddComment(vid string, authorId string, content string) error {
	addSql := "INSERT INTO comments(id,video_id,author_id,content) VALUES(?,?,?,?)"
	stmtIns, e := dbConn.Prepare(addSql)
	if e != nil {
		log.Printf("insert comments error: %s", e)
		return e
	}
	uuids, err := uuid.NewUUID()
	if err != nil {
		log.Printf("insert comments error: %s", e)
		return e
	}
	_, err = dbConn.Exec(addSql, uuids, vid, authorId, content)
	if err != nil {
		log.Printf("insert comments error: %s", e)
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	selectAllSql := `SELECT comments.id,users.login_name,comments.content FROM comments INNER JOIN 
users ON comments.author_id=users.id WHERE comments.video_id=? and comments.time>FROM_UNIXTIME(?) and comments.time<=FROM_UNIXTIME(?)
ORDER  BY comments.time DESC `

	stmtOut, e := dbConn.Prepare(selectAllSql)
	if e != nil {
		log.Printf("selectAll comments error: %s", e)
		return nil, e
	}

	rows, err := dbConn.Query(selectAllSql, vid, from, to)
	if err != nil {
		log.Printf("selectAll comments error: %s", e)
		return nil, e
	}

	var (
		id      string
		author  string
		content string
	)

	var comments []*defs.Comment

	if rows.Next() {
		err := rows.Scan(&id, &author, &content)
		if err != nil {
			log.Printf("selectAll comments error: %s", e)
			return nil, err
		}
		comment := &defs.Comment{Id: id, VideoId: vid, Author: author, Content: content}
		comments = append(comments, comment)

	}

	defer stmtOut.Close()
	return comments, nil
}

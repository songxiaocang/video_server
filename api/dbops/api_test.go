package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M){
	fmt.Println("testMain for users exec")
	clearTables()
	m.Run()
	clearTables()
}

//func TestUserWorkflow(t *testing.T){
//	t.Run("addUser",testAddUserCredential)
//	t.Run("getUser",testGetUserCredential)
//	t.Run("delUser",testDeleteUser)
//	t.Run("regetUser",testRegetUser)
//}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser：%v",err)
	}
}

func testGetUserCredential(t *testing.T){
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser fail")
	}
}


func testDeleteUser(t *testing.T){
	err := DeleteUser("avenssi", "123")
	if err!=nil {
		t.Errorf("Error of DeleteUser: %v",err)
	}
}

func testRegetUser(t *testing.T){
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v",err)
	}

	if pwd != ""{
		t.Errorf("delete user fail")
	}
}

var videoInfoId string

//func TestVideoInfoWorkflow(t *testing.T){
//	clearTables()
//	t.Run("prepareUser",testAddUserCredential)
//	t.Run("addVideoInfo",testAddVideoInfo)
//	t.Run("getVideoInfo",testGetVideoInfo)
//	t.Run("delVideoInfo",testDelVideoInfo)
//	t.Run("regetVideoInfo",testRegetVideoInfo)
//}

func testAddVideoInfo(t *testing.T) {
	videoInfo, e := AddVideoInfo(1, "view")
	if e!=nil {
		t.Errorf("add video_info error: %v",e)
	}

	videoInfoId = videoInfo.Id
}

func testGetVideoInfo(t *testing.T){
	_, e := GetVideoInfo(videoInfoId)
	if e !=nil {
		t.Errorf("get video_info error: %v",e)
	}
}

func testDelVideoInfo(t *testing.T) {
	e := DelVideoInfo(videoInfoId)
	if e !=nil {
		t.Errorf("del video_info error: %v",e)
	}
}


func testRegetVideoInfo(t *testing.T){
	info, e := GetVideoInfo(videoInfoId)
	if e!=nil || info != nil{
		t.Errorf("reget video_info error: %v",e)
	}
}


func TestComments(t *testing.T){
	clearTables()
	t.Run("prepareUser",testAddUserCredential)
	t.Run("addComment",testAddComment)
	t.Run("listComments",testListComments)
}

func testAddComment(t *testing.T){
	err := AddComment("12345", "1", "风景秀丽")

	if err!=nil {
		t.Errorf("addComment error: %v",err)
	}
}

func testListComments(t *testing.T){

	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	comments, e := ListComments("12345", 1538911829, to)
	if e!=nil {
		t.Errorf("listComments error: %v",err)
	}

	for i,v := range comments {
		fmt.Printf("comment: %d, %v\n",i,v)
	}
}
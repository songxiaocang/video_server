package defs

//request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string  `json:"pwd"`
}

//response
type SignUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}


//data model
type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct{
	Id string `json:"id"`
	VideoId string `json:"video_id"`
	Author string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct{
	UserName string `json:"user_name"`
	TTL int64 `json:"ttl"`
}
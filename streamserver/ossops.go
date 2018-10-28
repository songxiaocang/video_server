package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"sxc/config"
)

var EP string
var AK string
var SK string

func init() {
	EP = config.GetOssAddr()
	AK = ""
	SK = ""
}

func UploadToOss(filename string, path string, bn string) bool {
	client, e := oss.New(EP, AK, SK)
	if e != nil {
		log.Printf("init oss service error:%v", e)
		return false
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		log.Printf("getting bucket error: %v", err)
		return false
	}

	err1 := bucket.UploadFile(filename, path, 500*1024, oss.Routines(3))
	if err1 != nil {
		log.Printf("upload error: %v", err1)
		return false
	}

	return true
}

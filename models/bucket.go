package models

import (
	"crypto/md5"
	"encoding/hex"
)

type Bucket struct {
	BucketName    string `json:"bucketName" bson:"bucket_name"`
	BucketRegion  string `json:"bucketRegion" bson:"bucket_region"`
	ObjectLocking bool   `json:"objectLocking"`
	HashName      string `bson:"_id"`
}

type File struct {
	BucketName string   `json:"bucketName"`
	FileName   []string `json:"fileName"`
}
type Link struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
}

func GetMD5hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

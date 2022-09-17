package models

type Bucket struct {
	BucketName    string `json:"bucketName"`
	BucketRegion  string `json:"bucketRegion"`
	ObjectLocking bool   `json:"objectLocking"`
}

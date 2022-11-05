package models

type BucketDescription struct {
	BucketName  string      `json:"bucketName"`
	Description interface{} `json:"description"`
	ID          string      `bson:"_id"`
}

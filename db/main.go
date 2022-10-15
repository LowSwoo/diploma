package db

import (
	"context"
	"errors"
	"log"
	server "lowswoo/minio-server"
	"lowswoo/models"

	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var url = "mongodb://localhost:27017"
var ctx = context.TODO()

func CreateBucket(b *models.Bucket) error {
	b.HashName = models.GetMD5hash(b.BucketName)
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Client().Disconnect(ctx)
	collection := db.Collection(b.BucketName)
	result, err := collection.InsertOne(ctx, b)
	if err != nil {
		return err
	}
	err = server.CreateBucket(b)
	log.Default().Println("Result of creating new bucket: ", result)
	return err

}

func bucketExists(bucketName string) bool {
	buckets := GetBucketListNames()
	for _, bucket := range buckets {
		if bucket == bucketName {
			return true
		}
	}
	return false
}

func UploadFile(file *models.File) ([]models.Link, error) {
	links := []models.Link{}
	var err error
	file.BucketName, err = GetBucketHashName(file.BucketName)
	if err != nil {
		return nil, err
	}
	for _, fn := range file.FileName {
		url := server.GetFileLink(file.BucketName, fn)
		url_string := url.Scheme + "://" + url.Host + url.Path + "?" + url.RawQuery
		links = append(links, models.Link{FileName: fn, Url: url_string})
	}
	log.Default().Println(links)
	return links, nil
}

func RemoveFile(bucketName string, fileName string) {
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		log.Default().Println(err)
	}
	server.RemoveFile(hashName, fileName)
}

func GetBucketHashName(bucketName string) (string, error) {
	if !bucketExists(bucketName) {
		return "", errors.New("Bucket doesn't exist | Get bucket hash name")
	}
	var bucket = models.Bucket{}
	db, err := connect()
	if err != nil {
		return "", err
	}
	err = db.Collection(bucketName).FindOne(ctx, bson.M{"bucket_name": bucketName}).Decode(&bucket)
	return bucket.HashName, nil
}

func GetFileList(bucketName string) []minio.ObjectInfo {
	if len(bucketName) == 0 {
		return nil
	}
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		log.Default().Println(err)
		return nil
	}
	return server.GetFileList(hashName)
}

func RemoveBucket(bucketName string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		return err
	}
	err = server.RemoveBucket(hashName)
	if err != nil {
		return err
	}
	return db.Collection(bucketName).Drop(ctx)

}

func GetBucketListNames() []string {
	db, err := connect()
	if err != nil {
		// log.Fatal(err)
		log.Default().Println("Тут")
		log.Default().Println(err)
	}
	// defer db.Client().Disconnect(ctx)
	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return collections
}

func connect() (*mongo.Database, error) {
	session, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	err = session.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return session.Database("buckets"), nil
}

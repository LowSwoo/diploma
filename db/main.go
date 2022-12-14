package db

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"lowswoo/bucket"
	server "lowswoo/minio-server"
	"lowswoo/models"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var uri = "mongodb://localhost:27017"
var ctx = context.TODO()

func CreateBucket(b *models.BucketDB) error {
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

func UploadFile(file *models.FileUpload, hostURL string) ([]models.Link, error) {
	links := []models.Link{}
	var err error
	file.BucketName, err = GetBucketHashName(file.BucketName)
	if err != nil {
		return nil, err
	}
	for _, fn := range file.FileName {
		url := server.GetFileLink(file.BucketName, file.Folder+"/"+fn)
		host := strings.Split(hostURL, ":")[0] + ":" + url.Port()
		url_string := url.Scheme + "://" + host + url.Path + "?" + url.RawQuery
		log.Default().Println(url.Path)
		links = append(links, models.Link{FileName: fn, Url: url_string})
	}
	log.Default().Println(links)
	return links, nil
}

func DownloadFile(bucketName string, fileName string, hostURL string) string {
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		log.Default().Println(err)
	}
	url := server.GetFileLinkDownload(hashName, fileName)
	// log.Default().Printf("Default link: %v\n", url)
	// log.Default().Printf("hostURL: %v\n", hostURL)
	host := strings.Split(hostURL, ":")[0] + ":" + url.Port()
	url_string := url.Scheme + "://" + host + url.Path + "?" + url.RawQuery
	return url_string
}

func RemoveFile(bucketName string, fileName string) {
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		log.Default().Println(err)
	}
	server.RemoveFile(hashName, fileName)
}

func GetBucketHashName(bucketName string) (string, error) {
	// start := time.Now()
	if !bucketExists(bucketName) {
		return "", errors.New("Bucket doesn't exist | Get bucket hash name")
	}
	// log.Default().Println(time.Since(start))
	var bucket = models.BucketDB{}
	db, err := connect()
	if err != nil {
		return "", err
	}
	err = db.Collection(bucketName).FindOne(ctx, bson.M{"bucket_name": bucketName}).Decode(&bucket)
	return bucket.HashName, nil
}

func GetFileList(bucketName string, hostURL string) map[string]bucket.Folder {
	if len(bucketName) == 0 {
		return nil
	}
	b := bucket.NewBucket(bucketName)
	hashName, err := GetBucketHashName(bucketName)
	if err != nil {
		log.Default().Println(err)
		return nil
	}
	files := server.GetFileList(hashName)
	for idx, file := range files {
		files[idx].UserTags = map[string]string{
			"link": DownloadFile(bucketName, file.Key, hostURL),
		}
		b.AppendFile(file.Key, DownloadFile(bucketName, file.Key, hostURL), nil)
	}
	b.SetFoldersList()
	WriteToFile(b)
	return b.Folders
}

func WriteToFile(a interface{}) {
	data, err := json.Marshal(a)
	if err != nil {
		log.Default().Println(err)
	}
	err = os.WriteFile("bucket_data.json", data, 0644)
	if err != nil {
		log.Default().Println(err)
	}
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
		log.Default().Println("??????")
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
	session, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = session.Connect(ctx)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	// session.Ping(ctx, nil)
	// log.Default().Println(time.Since(start))
	return session.Database("buckets"), nil
}

func SetDescription(d *models.BucketDescription) error {
	db, err := connect()
	if err != nil {
		return err
	}
	if bucketExists(d.BucketName) {
		_, err := db.Collection(d.BucketName).InsertOne(ctx, d)
		if err != nil {
			return err
		}
		// log.Default().Println(result)
	} else {
		err = errors.New("no such bucket")
	}
	return err
}

// func GetDescription(bucketName string) (interface{}, error) {
// 	db, err := connect()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if bucketExists(bucketName) {
// 		// db.Collection(bucketName).FindOne(ctx, )
// 	}
// }

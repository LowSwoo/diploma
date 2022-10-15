package server

import (
	"context"
	"errors"
	"log"
	"lowswoo/models"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var d models.Data

func Login() (*minio.Client, context.Context) {
	ctx := context.Background()
	data := d.GetData()
	minioClient, err := minio.New(data.Host+":"+data.Port, &minio.Options{
		Creds:  credentials.NewStaticV4(data.AccessKeyID, data.SecretAccessKey, ""),
		Secure: data.UseSSL,
	})
	if err != nil {
		log.Fatalln("КАВО БЛЯТЬ")
	}
	return minioClient, ctx
}

func GetBucketListNames() []string {
	client, ctx := Login()
	l, err := client.ListBuckets(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	var elements []string
	for _, value := range l {
		elements = append(elements, value.Name)
	}
	return elements
}

func CreateBucket(bucket *models.Bucket) error {
	client, ctx := Login()
	log.Default().Println("Creating bucket with name", bucket.HashName)
	err := client.MakeBucket(ctx, bucket.HashName, minio.MakeBucketOptions{
		Region:        bucket.BucketRegion,
		ObjectLocking: bucket.ObjectLocking,
	})
	return err
}

func RemoveBucket(bucketName string) error {
	client, ctx := Login()
	ex, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if ex == false {
		return errors.New("Bucket doesn't exist | RemoveBucket")
	}
	return client.RemoveBucket(ctx, bucketName)
}

func GetFileLink(bucketName string, fileName string) *url.URL {
	client, ctx := Login()
	url, err := client.PresignedPutObject(ctx, bucketName, fileName, time.Second*2400)
	if err != nil {
		log.Default().Println(err)
	}
	// log.Default().Println(url)
	return url
}

func GetFileList(bucketName string) []minio.ObjectInfo {
	var objects []minio.ObjectInfo
	client, ctx := Login()
	for object := range client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{}) {
		objects = append(objects, object)
	}
	return objects
}

func RemoveFile(bucketName string, fileName string) {
	client, ctx := Login()
	client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}

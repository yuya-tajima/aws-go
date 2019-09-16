package s3

import (
	_"fmt"
	_aws "github.com/aws/aws-sdk-go/aws"
	_session "github.com/aws/aws-sdk-go/aws/session"
	_s3 "github.com/aws/aws-sdk-go/service/s3"
	_ "regexp"
	"time"
)

type S3 struct {
	svc *_s3.S3
}

type Object struct {
	*_s3.ListObjectsV2Output
}

type ObjectItem struct {
	Key   string    `json:"key"`
	Etag  string    `json:"etag"`
	Mdate time.Time `json:"last_modified"`
}

type ObjectDir struct {
	Name string `json:"name"`
}

type ObjectData struct {
	Items []*ObjectItem `json:"items"`
	Dirs  []*ObjectDir  `json:"dirs"`
}

type Bucket struct {
	*_s3.ListBucketsOutput
}

type BucketItem struct {
	Name  string    `json:"name"`
	Cdate time.Time `json:"creation_date"`
}

type BucketData struct {
	Items []*BucketItem `json:"items"`
}

func New(session *_session.Session, cfgs *_aws.Config) *S3 {
	return &S3{
		svc: _s3.New(session, cfgs),
	}
}

func (a *S3) GetBuckets(path string) (data *BucketData, err error) {

	result, err := a.ListsBuckets()

	if err == nil {
		c := &Bucket{result}
		data = c.getItems()
	}

	return
}

func (a *S3) ListsBuckets() (result *_s3.ListBucketsOutput, err error) {

	input := &_s3.ListBucketsInput{}
	result, err = a.svc.ListBuckets(input)

	return
}

func (b *Bucket) getItems() (data *BucketData) {

	data = &BucketData{}
	for _, i := range b.Buckets {
		item := b.createItem(i)
		data.Items = append(data.Items, item)
	}

	return
}

func (b *Bucket) createItem(i *_s3.Bucket) (item *BucketItem) {
	item = &BucketItem{
		Name:  _aws.StringValue(i.Name),
		Cdate: _aws.TimeValue(i.CreationDate),
	}
	return
}

func (a *S3) GetObjects(bucket, path string) (data *ObjectData, err error) {

	result, err := a.ListsObjects(bucket, path)

	if err == nil {
		c := &Object{result}
		items := c.getItems()
		dirs := c.getDirs()
		data = &ObjectData{
			Items: items,
			Dirs:  dirs,
		}
	}

	return
}

func (a *S3) ListsObjects(bucket, path string) (result *_s3.ListObjectsV2Output, err error) {

	input := &_s3.ListObjectsV2Input{
		Bucket:    _aws.String(bucket),
		Prefix:    _aws.String(path),
		Delimiter: _aws.String("/"),
		MaxKeys:   _aws.Int64(1000),
	}

	result, err = a.svc.ListObjectsV2(input)

	return
}

func (b *Object) getItems() (items []*ObjectItem) {
	for _, i := range b.Contents {
		item := b.createItem(i)
		if item != nil {
			items = append(items, item)
		}
	}
	return
}

func (b *Object) getDirs() (dirs []*ObjectDir) {
	for _, d := range b.CommonPrefixes {
		dir := b.createDir(d)
		if dir != nil {
			dirs = append(dirs, dir)
		}
	}
	return
}

func (b *Object) createItem(i *_s3.Object) (item *ObjectItem) {

	item = &ObjectItem{
		Key:   _aws.StringValue(i.Key),
		Etag:  _aws.StringValue(i.ETag),
		Mdate: _aws.TimeValue(i.LastModified),
	}
	return
}

func (b *Object) createDir(d *_s3.CommonPrefix) (dir *ObjectDir) {
	dir = &ObjectDir{
		Name: _aws.StringValue(d.Prefix),
	}
	return
}

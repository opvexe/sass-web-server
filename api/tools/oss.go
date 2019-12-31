package tools

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"pea-web/cmd"
	"time"
)

//上传图片
func UploadImage(data []byte) (string, error) {
	md5 := MD5Bytes(data)
	key := "images/" + TimeFormat(time.Now(), "2006/01/02/") + md5 + ".jpg"
	return PutObject(key, data)
}

// 上传
func PutObject(key string, data []byte) (string, error) {
	client, err := oss.New(cmd.Conf.AliyunOss.Endpoint, cmd.Conf.AliyunOss.AccessId, cmd.Conf.AliyunOss.AccessSecret)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(cmd.Conf.AliyunOss.Bucket)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(key, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return cmd.Conf.AliyunOss.Host + key, nil
}

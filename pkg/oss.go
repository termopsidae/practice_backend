package pkg

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"strings"
)

const (
	imgUrl          = "https://xuanwu-nft.oss-cn-beijing.aliyuncs.com"
	endpoint        = "http://oss-cn-beijing.aliyuncs.com"
	accessKeyID     = ""
	accessKeySecret = ""
	bucketName      = "xuanwu-nft"
	//fileFolder      = "img"
)

type FileType int

const (
	Image FileType = 0
)

func (f FileType) GetString() string {
	switch f {
	case Image:
		return "png"
	default: //d其他文件方Img
		return "Img"
	}
}

// oss添加文件
func SetFileOss(reader io.Reader, filename string, fileFolder FileType) (string, error) {
	var (
		fileName = filename
		client   *oss.Client
		bucket   *oss.Bucket
		err      error
	)
	// 创建OSSClient实例。
	if client, err = oss.New(endpoint, accessKeyID, accessKeySecret); err != nil {
		return "", err
	}
	// 获取存储空间。
	if bucket, err = client.Bucket(bucketName); err != nil {
		return "", err
	}

	urlRoute := fileFolder.GetString() + "/" + fileName

	if err = bucket.PutObject(urlRoute, reader); err != nil {
		return "", err
	}

	fmt.Println(imgUrl + "/" + urlRoute)

	return imgUrl + "/" + urlRoute, nil
}

// oss获取文件
func GetFileOss(url string) ([]byte, error) {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// yourBucketName填写存储空间名称。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	urlRoute := strings.Split(url, "/")
	object := urlRoute[4]
	format := strings.Split(object, ".")[1]
	if format == "jpeg" || format == "jpg" {
		format = "png"
	}
	// 下载文件到流。
	body, err := bucket.GetObject(format + "/" + object)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return data, nil
}

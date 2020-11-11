package aliUtil

import (
	"bytes"
	"os"
	"net/http"
	"io"
	"github.com/astaxie/beego"
	. "github.com/gyf841010/infra/errorUtil"
	"io/ioutil"
	. "github.com/gyf841010/infra/logging"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type BaseUrlAndEndPoint struct {
	BaseUrl  string
	EndPoint string
}

const (
	DEFAULT_BUCKET = "pingthe-private"
	PUBLIC_BUCKET  = "pingthe-public-img"

	MIN_EXPIRED_TIME = 60 * 5
	MAX_EXPIRED_TIME = 60 * 60 * 24 * 7
)

type CorpBucketNames struct {
	//阿里云的私有桶名
	PrivateBucket string
	//阿里云的公有桶名
	PublicBucket string
	//阿里云的临时文件桶名
	TmpBucket string
}

var corpIDAndBucketNameMap = map[int]CorpBucketNames{
	1: CorpBucketNames{PrivateBucket: "pingthe-private", PublicBucket: "pingthe-public-img", TmpBucket: "pingthe-tmp"},
	3: CorpBucketNames{PrivateBucket: "pingthe-private", PublicBucket: "pingthe-public-img", TmpBucket: "pingthe-tmp"},
}

var bucketNameAndBaseUrlAndEndPintMap = map[string]BaseUrlAndEndPoint{
	"pingthe-private":    BaseUrlAndEndPoint{BaseUrl: "http://pingthe-private.oss-cn-hangzhou.aliyuncs.com", EndPoint: "oss-cn-hangzhou.aliyuncs.com"},
	"pingthe-public-img": BaseUrlAndEndPoint{BaseUrl: "http://pingthe-public-img.oss-cn-hangzhou.aliyuncs.com", EndPoint: "oss-cn-hangzhou.aliyuncs.com"},
	"pingthe-tmp":        BaseUrlAndEndPoint{BaseUrl: "http://pingthe-tmp.oss-cn-hangzhou.aliyuncs.com", EndPoint: "oss-cn-hangzhou.aliyuncs.com"},
}

type UploadFileContentInput struct {
	//上传到阿里云的桶名
	Bucket string
	//上传到七牛云的文件路径 例: a/b/c/1.txt
	UploadPath string
	//文件内容
	FileContent []byte
	//返回Url结果的过期时间，缺省为5分钟
	ExpiredInSec int64
}

func getExpiredTime(expiredInSec int64) int64 {
	if expiredInSec < MIN_EXPIRED_TIME {
		expiredInSec = MIN_EXPIRED_TIME
	} else if expiredInSec > MAX_EXPIRED_TIME {
		expiredInSec = MAX_EXPIRED_TIME
	}

	return expiredInSec
}

func UploadFileContent(input *UploadFileContentInput) (string, error) {
	if input.FileContent == nil {
		err := NewHErrorCustom(ERROR_CODE_UPLOAD_FILE_CONTENT_IS_EMPTY)
		Log.Error("Failed to Transmit Get File Request ", WithError(err))
		return "", err
	}
	aliAccessKeyId := beego.AppConfig.String("aliAccessKeyId")
	aliSecretAccessKey := beego.AppConfig.String("aliSecretAccessKey")
	if aliAccessKeyId == "" || aliSecretAccessKey == "" {
		err := NewHErrorCustom(ERROR_CODE_ALIOSS_CONFIG_IS_EMPTY)
		Log.Error("Failed to Transmit Get File Request ", WithError(err))
		return "", err
	}

	_, endpoint, err := getBaseUrlAndEndpointByBucket(input.Bucket)
	if err != nil {
		Log.Error("Failed to Get Base Url", With("UploadPath", input.UploadPath), WithError(err))
		return "", err
	}

	// New client
	client, err := oss.New(endpoint, aliAccessKeyId, aliSecretAccessKey)
	if err != nil {
		Log.Error("Failed to New Oss Client", With("endpoint", endpoint), WithError(err))
		return "", err
	}

	// Get bucket
	bucket, err := client.Bucket(input.Bucket)
	if err != nil {
		Log.Error("Failed to Get Bucket", With("input.Bucket", input.Bucket), WithError(err))
		return "", err
	}

	err = bucket.PutObject(input.UploadPath, bytes.NewReader(input.FileContent))
	if err != nil {
		Log.Error("Failed to Put Object", WithError(err))
		return "", err
	}

	strUrl, err := bucket.SignURL(input.UploadPath, oss.HTTPGet, getExpiredTime(input.ExpiredInSec))
	if err != nil {
		Log.Error("Failed to Put Object", WithError(err))
		return "", err
	}

	//return aliUrl + "/" + input.UploadPath, nil
	return strUrl, nil
}

type UploadFileInput struct {
	//上传到七牛云的桶名
	Bucket string
	//上传到aws的文件路径 例: a/b/c/1.txt
	UploadPath string
	//文件的本地路径
	LocalFilePath string
	//返回Url结果的过期时间，缺省为5分钟
	ExpiredInSec int64
}

func UploadFile(input *UploadFileInput) (string, error) {
	fileContent, err := ioutil.ReadFile(input.LocalFilePath)
	if err != nil {
		return "", err
	}
	return UploadFileContent(&UploadFileContentInput{
		Bucket:       input.Bucket,
		FileContent:  fileContent,
		UploadPath:   input.UploadPath,
		ExpiredInSec: input.ExpiredInSec,
	})
}

func DownloadFile(url, localPath string) error {
	out, err := os.Create(localPath)
	if err != nil {
		Log.Error("Failed to Create File", With("localPath", localPath), WithError(err))
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		Log.Error("Failed to Download From", With("url", url), WithError(err))
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		Log.Error("Failed to Copy Body For File", WithError(err))
		return err
	}
	return nil
}

func GetFileUrlFromOssWithExpired(bucketName, fileKey string, expiredInSec int64) (string, error) {
	aliAccessKeyId := beego.AppConfig.String("aliAccessKeyId")
	aliSecretAccessKey := beego.AppConfig.String("aliSecretAccessKey")
	if aliAccessKeyId == "" || aliSecretAccessKey == "" {
		return "", NewHErrorCustom(ERROR_CODE_ALIOSS_CONFIG_IS_EMPTY)
	}

	_, endpoint, err := getBaseUrlAndEndpointByBucket(bucketName)
	if err != nil {
		Log.Error("Failed to Get Base Url", With("bucketName", bucketName), WithError(err))
		return "", err
	}

	// New client
	client, err := oss.New(endpoint, aliAccessKeyId, aliSecretAccessKey)
	if err != nil {
		Log.Error("Failed to New Oss Client", With("endpoint", endpoint), WithError(err))
		return "", err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		Log.Error("Failed to Get Bucket", With("bucketName", bucketName), WithError(err))
		return "", err
	}

	strUrl, err := bucket.SignURL(fileKey, oss.HTTPGet, getExpiredTime(expiredInSec))
	if err != nil {
		Log.Error("Failed to Put Object", WithError(err))
		return "", err
	}

	return strUrl, nil
}

func DeleteFile(bucketName, fileKey string) error {
	aliAccessKeyId := beego.AppConfig.String("aliAccessKeyId")
	aliSecretAccessKey := beego.AppConfig.String("aliSecretAccessKey")
	if aliAccessKeyId == "" || aliSecretAccessKey == "" {
		return NewHErrorCustom(ERROR_CODE_ALIOSS_CONFIG_IS_EMPTY)
	}

	_, endpoint, err := getBaseUrlAndEndpointByBucket(bucketName)
	if err != nil {
		Log.Error("Failed to Get Base Url", With("bucketName", bucketName), WithError(err))
		return err
	}

	// New client
	client, err := oss.New(endpoint, aliAccessKeyId, aliSecretAccessKey)
	if err != nil {
		Log.Error("Failed to New Oss Client", With("endpoint", endpoint), WithError(err))
		return err
	}

	// Get bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		Log.Error("Failed to Get Bucket", With("bucketName", bucketName), WithError(err))
		return err
	}

	err = bucket.DeleteObject(fileKey)
	if err != nil {
		Log.Error("Failed to Delete Object", WithError(err))
		return err
	}

	return nil
}

func getBaseUrlAndEndpointByBucket(bucket string) (string, string, error) {
	if baseUrlAndEndPoint, ok := bucketNameAndBaseUrlAndEndPintMap[bucket]; ok {
		return baseUrlAndEndPoint.BaseUrl, baseUrlAndEndPoint.EndPoint, nil
	} else {
		return "", "", NewHErrorCustom(ERROR_CODE_GET_BASE_URL_ERROR)
	}
}

func GetBucketByCorpID(corpID int, privateOrPublic int, tmpFlag int) string {
	if corpBucketNames, ok := corpIDAndBucketNameMap[corpID]; ok {
		if tmpFlag == 1 {
			return corpBucketNames.TmpBucket
		}
		if privateOrPublic == 1 {
			return corpBucketNames.PublicBucket
		} else {
			return corpBucketNames.PrivateBucket
		}
	} else {
		return DEFAULT_BUCKET
	}
}

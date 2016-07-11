package controllers

import (
	"fmt"
	goio "io"
	"os"
	"path"
	"time"

	cf "github.com/qiniu/api.v6/conf"
	io "github.com/qiniu/api.v6/io"
	rs "github.com/qiniu/api.v6/rs"
	"github.com/revel/revel"
)

var (
	QiniuEnable bool
	QiniuAccess string
	QiniuSecret string
	QiniuBucket string
	QiniuDomain string
)

func QiniuToken(bucket string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucket,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}

func Upload(file_name string, rd goio.Reader) (string, error) {
	var url string
	if !QiniuEnable {
		sub := fmt.Sprintf("%d", time.Now().Unix()%31)
		dir := path.Join(revel.BasePath, "upload", sub)
		os.MkdirAll(dir, os.ModePerm)

		fname := fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file_name)
		wr, err := os.Create(path.Join(dir, fname))
		if err != nil {
			return "", err
		}
		defer wr.Close()

		if _, err := goio.Copy(wr, rd); err != nil {
			return "", err
		}

		url = fmt.Sprintf("/upload/%s/%s", sub, fname)
	} else {
		var ret io.PutRet
		var extra = &io.PutExtra{}

		uptoken := QiniuToken(QiniuBucket)
		if file_name != "" {
			if err := io.Put(nil, &ret, uptoken, file_name, rd, extra); err != nil {
				return "", err
			}
		} else {
			if err := io.PutWithoutKey(nil, &ret, uptoken, rd, extra); err != nil {
				return "", err
			}
		}
		url = fmt.Sprintf("http://%s/%s", QiniuDomain, ret.Key)
	}
	return url, nil
}

func InitWidget() {
	QiniuEnable = revel.Config.BoolDefault("qiniu.enable", false)
	QiniuAccess = revel.Config.StringDefault("qiniu.access", "")
	QiniuSecret = revel.Config.StringDefault("qiniu.secret", "")
	QiniuBucket = revel.Config.StringDefault("qiniu.bucket", "")
	QiniuDomain = revel.Config.StringDefault("qiniu.domain", "")

	cf.ACCESS_KEY = QiniuAccess
	cf.SECRET_KEY = QiniuSecret
}

func init() {
	revel.OnAppStart(InitWidget)
}

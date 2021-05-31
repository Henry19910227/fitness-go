package tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type uploader struct {
	setting setting.Uploader
}

func NewUploader(setting setting.Uploader) Uploader {
	return &uploader{setting}
}

func (u *uploader) UploadFile(file io.Reader, filename string, filepath string) error {
	// 創建資料夾
	dst := u.setting.GetUploadSavePath() + filepath
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}
	// 創建空白檔案
	out, err := os.Create(dst + "/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// 將檔案複製到空白檔案上
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

func (u *uploader) RemoveFile(fileNamed string, filepath string) error {
	dst := u.setting.GetUploadSavePath() + filepath
	return os.Remove(dst + "/" + fileNamed)
}

func (u *uploader) CheckUploadImageAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range u.setting.GetUploadImageAllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

func (u *uploader) CheckUploadImageMaxSize(file io.Reader) bool {
	content, _ := ioutil.ReadAll(file)
	size := len(content)
	return size <= u.setting.GetUploadImageMaxSize()*1024*1024
}

func (u *uploader) createUploadSavePath() (string, error) {
	err := os.MkdirAll(u.setting.GetUploadSavePath(), os.ModePerm)
	if err != nil {
		return "", err
	}
	return u.setting.GetUploadSavePath(), nil
}

func getFileName(ext string) string {
	timeStr := time.Now().Format("20060102150405")
	m := md5.New()
	m.Write([]byte(timeStr))
	return hex.EncodeToString(m.Sum(nil)) + ext
}


package uploader

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/setting/uploader"
	"io"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type tool struct {
	setting uploader.Setting
}

func New(setting uploader.Setting) Tool {
	return &tool{setting: setting}
}

func (t *tool) Save(file io.Reader, fileNamed string) (string, error) {
	if !t.checkMaxSize(file) {
		return "", errors.New("上傳檔案大小超過限制")
	}
	if !t.checkAllowExt(path.Ext(fileNamed)) {
		return "", errors.New("上傳檔案不符合規範")
	}
	newFileNamed := generateFileName(path.Ext(fileNamed))
	// 創建資料夾
	dst := t.setting.FilePath()
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return "", err
	}
	// 創建空白檔案
	out, err := os.Create(dst + "/" + newFileNamed)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 將檔案複製到空白檔案上
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return newFileNamed, nil
}

func (t *tool) checkMaxSize(file io.Reader) bool {
	if sizeValue, ok := file.(Size); ok {
		size := int(sizeValue.Size())
		return size < t.setting.MaxSize()*1024*1024
	}
	return false
}

func (t *tool) checkAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range t.setting.AllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

func generateFileName(ext string) string {
	seed := randRange(1, 1000)
	timeStr := time.Now().Format("20060102150405.000") + strconv.Itoa(int(seed))
	m := md5.New()
	m.Write([]byte(timeStr))
	return hex.EncodeToString(m.Sum(nil)) + ext
}

func randRange(min int64, max int64) int64 {
	if min > max || min < 0 {
		return 0
	}
	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()
}

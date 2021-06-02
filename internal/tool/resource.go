package tool

import (
	"github.com/Henry19910227/fitness-go/internal/setting"
	"io"
	"os"
)

type resource struct {
	setting setting.Resource
}

func NewFile(setting setting.Resource) Resource {
	return &resource{setting}
}

func (r *resource) SaveFile(file io.Reader, filename string, filepath string) error {
	// 創建資料夾
	dst := r.setting.RootPath() + filepath
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

func (r *resource) RemoveFile(filepath string, fileNamed string) error {
	dst := r.setting.RootPath() + filepath
	return os.Remove(dst + "/" + fileNamed)
}


package handler

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/Henry19910227/fitness-go/internal/setting"
	"github.com/Henry19910227/fitness-go/internal/tool"
	"io"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

type Size interface {
	Size() int64
}

type uploader struct {
	resTool       tool.Resource
	uploadSetting setting.Upload
}

func NewUploader(resTool tool.Resource, uploadSetting setting.Upload) Uploader {
	return &uploader{resTool: resTool, uploadSetting: uploadSetting}
}

func (u *uploader) UploadCourseCover(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/course/cover"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadActionCover(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/action/cover"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadTrainerAvatar(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/trainer/avatar"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadUserAvatar(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/user/avatar"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadWorkoutStartAudio(file io.Reader, audioNamed string) (string, error) {
	if !u.checkUploadAudioAllowExt(path.Ext(audioNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newAudioNamed := generateFileName(path.Ext(audioNamed))
	if err := u.resTool.SaveFile(file, newAudioNamed, "/workout/start_audio"); err != nil {
		return "", err
	}
	return newAudioNamed, nil
}

func (u *uploader) UploadWorkoutEndAudio(file io.Reader, audioNamed string) (string, error) {
	if !u.checkUploadAudioAllowExt(path.Ext(audioNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newAudioNamed := generateFileName(path.Ext(audioNamed))
	if err := u.resTool.SaveFile(file, newAudioNamed, "/workout/end_audio"); err != nil {
		return "", err
	}
	return newAudioNamed, nil
}

func (u *uploader) UploadWorkoutSetStartAudio(file io.Reader, audioNamed string) (string, error) {
	if !u.checkUploadAudioAllowExt(path.Ext(audioNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newAudioNamed := generateFileName(path.Ext(audioNamed))
	if err := u.resTool.SaveFile(file, newAudioNamed, "/workout_set/start_audio"); err != nil {
		return "", err
	}
	return newAudioNamed, nil
}

func (u *uploader) UploadWorkoutSetProgressAudio(file io.Reader, audioNamed string) (string, error) {
	if !u.checkUploadAudioAllowExt(path.Ext(audioNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newAudioNamed := generateFileName(path.Ext(audioNamed))
	if err := u.resTool.SaveFile(file, newAudioNamed, "/workout_set/progress_audio"); err != nil {
		return "", err
	}
	return newAudioNamed, nil
}


func (u *uploader) UploadActionVideo(file io.Reader, videoNamed string) (string, error) {
	if !u.checkUploadVideoAllowExt(path.Ext(videoNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkVideoMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newVideoNamed := generateFileName(path.Ext(videoNamed))
	if err := u.resTool.SaveFile(file, newVideoNamed, "/action/video"); err != nil {
		return "", err
	}
	return newVideoNamed, nil
}

func (u *uploader) UploadCardFrontImage(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/trainer/card_front_image"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadCardBackImage(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/trainer/card_back_image"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) UploadTrainerAlbumPhoto(file io.Reader, imageNamed string) (string, error) {
	if !u.checkUploadImageAllowExt(path.Ext(imageNamed)) {
		return "", errors.New("9007-上傳檔案不符合規範")
	}
	if !u.checkImageMaxSize(file) {
		return "", errors.New("9008-上傳檔案大小超過限制")
	}
	newImageNamed := generateFileName(path.Ext(imageNamed))
	if err := u.resTool.SaveFile(file, newImageNamed, "/trainer/album"); err != nil {
		return "", err
	}
	return newImageNamed, nil
}

func (u *uploader) checkUploadImageAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range u.uploadSetting.ImageAllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

func (u *uploader) checkUploadAudioAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range u.uploadSetting.AudioAllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

func (u *uploader) checkUploadVideoAllowExt(ext string) bool {
	ext = strings.ToUpper(ext)
	for _, v := range u.uploadSetting.VideoAllowExts() {
		if ext == strings.ToUpper(v) {
			return true
		}
	}
	return false
}

func (u *uploader) checkImageMaxSize(file io.Reader) bool {
	if sizeValue, ok := file.(Size); ok {
		 size := int(sizeValue.Size())
		 return size < u.uploadSetting.ImageMaxSize() * 1024 * 1024
	}
	return false
}

func (u *uploader) checkAudioMaxSize(file io.Reader) bool {
	if sizeValue, ok := file.(Size); ok {
		size := int(sizeValue.Size())
		return size < u.uploadSetting.AudioMaxSize() * 1024 * 1024
	}
	return false
}

func (u *uploader) checkVideoMaxSize(file io.Reader) bool {
	if sizeValue, ok := file.(Size); ok {
		size := int(sizeValue.Size())
		return size < u.uploadSetting.VideoMaxSize() * 1024 * 1024
	}
	return false
}

// 舊的上傳判斷法
func (u *uploader) checkUploadImageMaxSize(file io.Reader) (io.Reader, bool) {
	content, _ := ioutil.ReadAll(file)
	//因ReadAll讀取完後第二次會讀取不到，必須使用NopCloser將資料寫回
	data := ioutil.NopCloser(bytes.NewBuffer(content))
	return data, len(content) <= u.uploadSetting.ImageMaxSize() * 1024 * 1024
}

func generateFileName(ext string) string {
	timeStr := time.Now().Format("20060102150405.000")
	m := md5.New()
	m.Write([]byte(timeStr))
	return hex.EncodeToString(m.Sum(nil)) + ext
}

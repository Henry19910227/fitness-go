package handler

import (
	"github.com/Henry19910227/fitness-go/internal/tool"
)

type resource struct {
	resTool tool.Resource
}

func NewResource(resTool tool.Resource) Resource {
	return &resource{resTool: resTool}
}

func (r *resource) DeleteCourseCover(imageNamed string) error {
	return r.resTool.RemoveFile("/course/cover", imageNamed)
}

func (r *resource) DeleteTrainerAvatar(imageNamed string) error {
	return r.resTool.RemoveFile("/trainer/avatar", imageNamed)
}

func (r *resource) DeleteUserAvatar(imageNamed string) error {
	return r.resTool.RemoveFile("/user/avatar", imageNamed)
}

func (r *resource) DeleteWorkoutSetStartAudio(audioNamed string) error {
	return r.resTool.RemoveFile("/workout_set/start_audio", audioNamed)
}

func (r *resource) DeleteWorkoutSetProgressAudio(audioNamed string) error {
	return r.resTool.RemoveFile("/workout_set/progress_audio", audioNamed)
}

func (r *resource) DeleteWorkoutStartAudio(audioNamed string) error {
	return r.resTool.RemoveFile("/workout/start_audio", audioNamed)
}

func (r *resource) DeleteWorkoutEndAudio(audioNamed string) error {
	return r.resTool.RemoveFile("/workout/end_audio", audioNamed)
}

func (r *resource) DeleteCardFrontImage(imageNamed string) error {
	return r.resTool.RemoveFile("/trainer/card_front_image", imageNamed)
}

func (r *resource) DeleteCardBackImage(imageNamed string) error {
	return r.resTool.RemoveFile("/trainer/card_back_image", imageNamed)
}

func (r *resource) DeleteTrainerAlbumPhoto(imageNamed string) error {
	return r.resTool.RemoveFile("/trainer/album", imageNamed)
}

func (r *resource) DeleteActionCover(coverNamed string) error {
	return r.resTool.RemoveFile("/action/cover", coverNamed)
}

func (r *resource) DeleteActionVideo(videoNamed string) error {
	return r.resTool.RemoveFile("/action/video", videoNamed)
}

func (r *resource) DeleteCertificateImage(imageNamed string) error {
	return r.resTool.RemoveFile("/trainer/certificate", imageNamed)
}


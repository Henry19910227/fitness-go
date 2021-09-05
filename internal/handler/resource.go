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

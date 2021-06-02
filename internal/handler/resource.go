package handler

import "github.com/Henry19910227/fitness-go/internal/tool"

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

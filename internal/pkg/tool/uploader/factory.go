package uploader

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/uploader"

func NewCourseCoverTool() Tool {
	return New(setting.NewCourseCover())
}

func NewTrainerAvatarTool() Tool {
	return New(setting.NewTrainerAvatar())
}

func NewActionCoverTool() Tool {
	return New(setting.NewActionCover())
}

func NewActionVideoTool() Tool {
	return New(setting.NewActionVideo())
}

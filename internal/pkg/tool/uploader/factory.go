package uploader

import setting "github.com/Henry19910227/fitness-go/internal/pkg/setting/uploader"

func NewCourseCoverTool() Tool {
	return New(setting.NewCourseCover())
}

func NewTrainerAvatarTool() Tool {
	return New(setting.NewTrainerAvatar())
}

func NewUserAvatarTool() Tool {
	return New(setting.NewUserAvatar())
}

func NewActionCoverTool() Tool {
	return New(setting.NewActionCover())
}

func NewActionVideoTool() Tool {
	return New(setting.NewActionVideo())
}

func NewBodyImageTool() Tool {
	return New(setting.NewBodyImage())
}

func NewFeedbackImageTool() Tool {
	return New(setting.NewFeedbackImage())
}

func NewReviewImageTool() Tool {
	return New(setting.NewReviewImage())
}

func NewBannerImageTool() Tool {
	return New(setting.NewBannerImage())
}

func NewAccountImageTool() Tool {
	return New(setting.NewAccountImage())
}

func NewWorkoutStartAudioTool() Tool {
	return New(setting.NewWorkoutStartAudio())
}

func NewWorkoutEndAudioTool() Tool {
	return New(setting.NewWorkoutEndAudio())
}

package dto

type TrainerAlbumPhotoResult struct {
	Photo string `json:"photo" example:"dkf2se51fsdds.png"` // 教練相簿照片
}

type TrainerAlbumPhoto struct {
	ID int64 `json:"id" example:"1"` // 教練相簿照片id
	Photo string `json:"photo" example:"dkf2se51fsdds.png"` // 教練相簿照片
}
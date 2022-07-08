package review_image

type FindInput struct {
	IDOptional
}

type DeleteInput struct {
	IDOptional
}

// APIDeleteCMSReviewImageInput /v2/cms/review_image/{review_image_id} [DELETE]
type APIDeleteCMSReviewImageInput struct {
	Uri APIDeleteCMSReviewImageUri
}
type APIDeleteCMSReviewImageUri struct {
	IDRequired
}

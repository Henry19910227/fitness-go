package validator

type CreateCertificateQuery struct {
	Name string `form:"name" binding:"required,min=1,max=50" example:"A級教練證照"` //證照名稱(1~50字元)
}

type CertificateIDUri struct {
	CerID int64 `uri:"certificate_id" binding:"required" example:"1"` // 證照ID
}

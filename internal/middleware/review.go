package middleware

import (
	"errors"
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type review struct {
	Base
	errHandler errcode.Handler
}

func NewReview(errHandler errcode.Handler) Review {
	return &review{errHandler: errHandler}
}

func (r *review) ReviewCreatorVerify(reviewOwner func(c *gin.Context, reviewID int64) (int64, errcode.Error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, isExists := c.Get("uid")
		if !isExists {
			r.JSONErrorResponse(c, r.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.DataNotFound))))
			c.Abort()
			return
		}
		var uri validator.ReviewIDUri
		if err := c.ShouldBindUri(&uri); err != nil {
			r.JSONValidatorErrorResponse(c, err)
			return
		}
		ownerID, err := reviewOwner(c, uri.ReviewID)
		if err != nil {
			r.JSONErrorResponse(c, err)
			c.Abort()
			return
		}
		if uid != ownerID {
			r.JSONErrorResponse(c, r.errHandler.Set(c, "course repo", errors.New(strconv.Itoa(errcode.PermissionDenied))))
			c.Abort()
			return
		}
	}
}

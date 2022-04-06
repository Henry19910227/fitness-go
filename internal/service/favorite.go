package service

import (
	"github.com/Henry19910227/fitness-go/errcode"
	"github.com/Henry19910227/fitness-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type favorite struct {
	favoriteRepo repository.Favorite
	errHandler   errcode.Handler
}

func NewFavorite(favoriteRepo repository.Favorite,
	errHandler errcode.Handler) Favorite {
	return &favorite{favoriteRepo: favoriteRepo, errHandler: errHandler}
}

func (f *favorite) CreateFavoriteCourse(c *gin.Context, userID int64, courseID int64) errcode.Error {
	if err := f.favoriteRepo.CreateFavoriteCourse(userID, courseID); err != nil {
		return f.errHandler.Set(c, "favorite repo", err)
	}
	return nil
}

func (f *favorite) DeleteFavoriteCourse(c *gin.Context, userID int64, courseID int64) errcode.Error {
	if err := f.favoriteRepo.DeleteFavoriteCourse(userID, courseID); err != nil {
		return f.errHandler.Set(c, "favorite repo", err)
	}
	return nil
}

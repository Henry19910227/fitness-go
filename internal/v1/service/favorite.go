package service

import (
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
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

func (f *favorite) CreateFavoriteTrainer(c *gin.Context, userID int64, trainerID int64) errcode.Error {
	if err := f.favoriteRepo.CreateFavoriteTrainer(userID, trainerID); err != nil {
		return f.errHandler.Set(c, "favorite repo", err)
	}
	return nil
}

func (f *favorite) CreateFavoriteAction(c *gin.Context, userID int64, actionID int64) errcode.Error {
	if err := f.favoriteRepo.CreateFavoriteAction(userID, actionID); err != nil {
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

func (f *favorite) DeleteFavoriteTrainer(c *gin.Context, userID int64, trainerID int64) errcode.Error {
	if err := f.favoriteRepo.DeleteFavoriteTrainer(userID, trainerID); err != nil {
		return f.errHandler.Set(c, "favorite repo", err)
	}
	return nil
}

func (f *favorite) DeleteFavoriteAction(c *gin.Context, userID int64, actionID int64) errcode.Error {
	if err := f.favoriteRepo.DeleteFavoriteAction(userID, actionID); err != nil {
		return f.errHandler.Set(c, "favorite repo", err)
	}
	return nil
}

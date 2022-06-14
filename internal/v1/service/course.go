package service

import (
	"errors"
	"github.com/Henry19910227/fitness-go/internal/pkg/errcode"
	"github.com/Henry19910227/fitness-go/internal/pkg/global"
	"github.com/Henry19910227/fitness-go/internal/pkg/tool"
	"github.com/Henry19910227/fitness-go/internal/v1/dto"
	"github.com/Henry19910227/fitness-go/internal/v1/handler"
	"github.com/Henry19910227/fitness-go/internal/v1/model"
	"github.com/Henry19910227/fitness-go/internal/v1/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type course struct {
	Base
	courseRepo               repository.Course
	userCourseAssetRepo      repository.UserCourseAsset
	trainerRepo              repository.Trainer
	trainerStatRepo          repository.TrainerStatistic
	planRepo                 repository.Plan
	workoutRepo              repository.Workout
	workoutSetRepo           repository.WorkoutSet
	saleRepo                 repository.Sale
	subscribeInfoRepo        repository.UserSubscribeInfo
	userCourseStatisticRepo  repository.UserCourseStatistic
	favoriteRepo             repository.Favorite
	reviewStatisticRepo      repository.ReviewStatistic
	courseUsageStatisticRepo repository.CourseUsageStatistic
	transactionRepo repository.Transaction
	uploader        handler.Uploader
	resHandler      handler.Resource
	logger     handler.Logger
	jwtTool    tool.JWT
	errHandler errcode.Handler
}

func NewCourse(courseRepo repository.Course,
	userCourseAssetRepo repository.UserCourseAsset,
	trainerRepo repository.Trainer,
	trainerStatRepo repository.TrainerStatistic,
	planRepo repository.Plan,
	workoutRepo repository.Workout,
	workoutSetRepo repository.WorkoutSet,
	saleRepo repository.Sale,
	subscribeInfoRepo repository.UserSubscribeInfo,
	userCourseStatisticRepo repository.UserCourseStatistic,
	favoriteRepo repository.Favorite,
	reviewStatisticRepo repository.ReviewStatistic,
	courseUsageStatisticRepo repository.CourseUsageStatistic,
	transactionRepo repository.Transaction,
	uploader handler.Uploader, resHandler handler.Resource, logger handler.Logger,
	jwtTool tool.JWT,
	errHandler errcode.Handler) Course {
	return &course{courseRepo: courseRepo, userCourseAssetRepo: userCourseAssetRepo,
		trainerRepo: trainerRepo, trainerStatRepo: trainerStatRepo, planRepo: planRepo, workoutRepo: workoutRepo, workoutSetRepo: workoutSetRepo,
		saleRepo: saleRepo, subscribeInfoRepo: subscribeInfoRepo, userCourseStatisticRepo: userCourseStatisticRepo,
		favoriteRepo: favoriteRepo, reviewStatisticRepo: reviewStatisticRepo, courseUsageStatisticRepo: courseUsageStatisticRepo, transactionRepo: transactionRepo,
		uploader: uploader, resHandler: resHandler, logger: logger, jwtTool: jwtTool, errHandler: errHandler}
}

func (cs *course) CreateCourseByToken(c *gin.Context, token string, param *dto.CreateCourseParam) (*dto.Course, errcode.Error) {
	uid, err := cs.jwtTool.GetIDByToken(token)
	if err != nil {
		return nil, cs.errHandler.InvalidToken()
	}
	return cs.CreateCourse(c, uid, param)
}

func (cs *course) CreateCourse(c *gin.Context, uid int64, param *dto.CreateCourseParam) (*dto.Course, errcode.Error) {
	var courseID int64
	var err error
	if param.ScheduleType == 1 {
		courseID, err = cs.courseRepo.CreateSingleWorkoutCourse(uid, &model.CreateCourseParam{
			Name:     param.Name,
			Level:    param.Level,
			Category: param.Category,
		})
	} else {
		courseID, err = cs.courseRepo.CreateCourse(uid, &model.CreateCourseParam{
			Name:     param.Name,
			Level:    param.Level,
			Category: param.Category,
		})
	}
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) UpdateCourse(c *gin.Context, courseID int64, param *dto.UpdateCourseParam) (*dto.Course, errcode.Error) {
	if err := cs.courseRepo.UpdateCourseByID(nil, courseID, &model.UpdateCourseParam{
		Category:    param.Category,
		SaleType:    param.SaleType,
		SaleID:      param.SaleID,
		Name:        param.Name,
		Intro:       param.Intro,
		Food:        param.Food,
		Level:       param.Level,
		Suit:        param.Suit,
		Equipment:   param.Equipment,
		Place:       param.Place,
		TrainTarget: param.TrainTarget,
		BodyTarget:  param.BodyTarget,
		Notice:      param.Notice,
	}); err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) UpdateCourseSaleType(c *gin.Context, courseID int64, saleType int, saleID *int64) (*dto.Course, errcode.Error) {
	//檢查付費課表需一起帶入銷售id
	if saleType == int(global.SaleTypeCharge) {
		//檢查付費類型課表需帶入saleID
		if saleID == nil {
			return nil, cs.errHandler.Custom(8999, errors.New("付費類型課表需帶入saleID"))
		}
		//檢查saleID是否為付費課表類型
		item, err := cs.saleRepo.FindSaleItemByID(*saleID)
		if err != nil {
			return nil, cs.errHandler.Set(c, "sale repo", err)
		}
		if item.Type != int(global.SaleTypeCharge) {
			return nil, cs.errHandler.Custom(8999, errors.New("銷售類型不符"))
		}
	}
	//檢查免費課表或訂閱課表銷售id需為空
	if saleType != int(global.SaleTypeCharge) {
		saleID = nil
	}
	//執行更新d
	if err := cs.courseRepo.UpdateCourseByID(nil, courseID, &model.UpdateCourseParam{
		SaleType: &saleType,
		SaleID:   saleID,
	}); err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	return cs.GetCourseDetailByCourseID(c, courseID)
}

func (cs *course) DeleteCourse(c *gin.Context, courseID int64) (*dto.CourseID, errcode.Error) {
	if err := cs.courseRepo.DeleteCourseByID(courseID); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	return &dto.CourseID{ID: courseID}, nil
}

func (cs *course) GetCourseSummariesByUID(c *gin.Context, uid int64, status []int, orderByParam *dto.OrderByParam, pagingParam *dto.PagingParam) ([]*dto.CourseSummary, *dto.Paging, errcode.Error) {
	//設置排序
	var orderBy *model.OrderBy
	if orderByParam != nil {
		orderBy = &model.OrderBy{
			OrderType: global.DESC,
			Field:     "update_at",
		}
		if orderByParam.OrderType != nil {
			orderBy.OrderType = global.OrderType(*orderByParam.OrderType)
		}
		if orderByParam.OrderField != nil {
			orderBy.Field = *orderByParam.OrderField
		}
	}
	//設置分頁
	var paging *model.PagingParam
	if pagingParam != nil {
		offset, limit := cs.GetPagingIndex(pagingParam.Page, pagingParam.Size)
		paging = &model.PagingParam{
			Offset: offset,
			Limit:  limit,
		}
	}
	//查詢數據
	var totalCount int64
	datas, err := cs.courseRepo.FindCourseSummaries(&totalCount, &model.FindCourseSummariesParam{
		UID:    &uid,
		Status: status,
	}, orderBy, paging)
	if err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, nil, cs.errHandler.SystemError()
	}
	courses := make([]*dto.CourseSummary, 0)
	for _, data := range datas {
		courses = append(courses, dto.NewCourseSummary(data))
	}
	var pagingResult *dto.Paging
	if pagingParam != nil {
		pagingResult = &dto.Paging{
			TotalCount: int(totalCount),
			TotalPage:  cs.GetTotalPage(int(totalCount), pagingParam.Size),
			Page:       pagingParam.Page,
			Size:       pagingParam.Size,
		}
	}
	return courses, pagingResult, nil
}

func (cs *course) GetCourseDetailByCourseID(c *gin.Context, courseID int64) (*dto.Course, errcode.Error) {
	entity, err := cs.courseRepo.FindCourseByCourseID(courseID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, cs.errHandler.DataNotFound()
		}
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	course := dto.Course{
		ID:           entity.ID,
		SaleType:     entity.SaleType,
		CourseStatus: entity.CourseStatus,
		Category:     entity.Category,
		ScheduleType: entity.ScheduleType,
		Name:         entity.Name,
		Cover:        entity.Cover,
		Intro:        entity.Intro,
		Food:         entity.Food,
		Level:        entity.Level,
		Suit:         entity.Suit,
		Equipment:    entity.Equipment,
		Place:        entity.Place,
		TrainTarget:  entity.TrainTarget,
		BodyTarget:   entity.BodyTarget,
		Notice:       entity.Notice,
		PlanCount:    entity.PlanCount,
		WorkoutCount: entity.WorkoutCount,
		CreateAt:     entity.CreateAt,
		UpdateAt:     entity.UpdateAt,
	}
	trainer := &dto.TrainerSummary{
		UserID:   entity.Trainer.UserID,
		Nickname: entity.Trainer.Nickname,
		Avatar:   entity.Trainer.Avatar,
		Skill:    entity.Trainer.Skill,
	}
	course.Trainer = trainer
	if entity.Sale != nil {
		sale := &dto.SaleItem{
			ID:   entity.Sale.ID,
			Type: entity.Sale.Type,
		}
		course.Sale = sale
		if entity.Sale.ProductLabel != nil {
			course.Sale.Name = entity.Sale.Name
			course.Sale.Twd = entity.Sale.ProductLabel.Twd
			course.Sale.ProductID = entity.Sale.ProductLabel.ProductID
		}
	}
	return &course, nil
}

func (cs *course) GetCourseProductByCourseID(c *gin.Context, userID int64, courseID int64) (*dto.CourseProduct, errcode.Error) {
	//查詢課表詳情
	data, err := cs.courseRepo.FindCourseProduct(courseID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	courseDto := dto.NewCourseProduct(data)
	courseDto.AllowAccess = 0
	if global.SaleType(courseDto.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := cs.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo != nil {
			courseDto.AllowAccess = subscribeInfo.Status
		}
	}
	if global.SaleType(courseDto.SaleType) == global.SaleTypeFree || global.SaleType(courseDto.SaleType) == global.SaleTypeCharge {
		asset, err := cs.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: courseID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "course repo", err)
		}
		if asset != nil {
			if asset.Available == 1 {
				courseDto.AllowAccess = 1
			}
		}
	}
	//查詢課表收藏狀態
	courseDto.Favorite = 0
	favoriteCourse, err := cs.favoriteRepo.FindFavoriteCourse(userID, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "favorite repo", err)
	}
	if favoriteCourse.CourseID > 0 {
		courseDto.Favorite = 1
	}
	return &courseDto, nil
}

func (cs *course) GetCourseOverviewByCourseID(c *gin.Context, courseID int64) (*dto.CourseProduct, errcode.Error) {
	//查詢課表詳情
	data, err := cs.courseRepo.FindCourseProduct(courseID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	courseDto := dto.NewCourseProduct(data)
	courseDto.AllowAccess = 1
	return &courseDto, nil
}

func (cs *course) GetCourseProductSummaries(c *gin.Context, param *dto.GetCourseProductSummariesParam, page, size int) ([]*dto.CourseProductSummary, *dto.Paging, errcode.Error) {
	var field = string(global.UpdateAt)
	offset, limit := cs.GetPagingIndex(page, size)
	datas, err := cs.courseRepo.FindCourseProductSummaries(model.FindCourseProductSummariesParam{
		UserID:       param.UserID,
		Name:         param.Name,
		Score:        param.Score,
		Level:        param.Level,
		Category:     param.Category,
		Suit:         param.Suit,
		Equipment:    param.Equipment,
		Place:        param.Place,
		TrainTarget:  param.TrainTarget,
		BodyTarget:   param.BodyTarget,
		SaleType:     param.SaleType,
		TrainerSex:   param.TrainerSex,
		TrainerSkill: param.TrainerSkill,
	}, &model.OrderBy{
		Field:     field,
		OrderType: global.DESC,
	}, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	totalCount, err := cs.courseRepo.FindCourseProductCount(model.FindCourseProductCountParam{
		UserID:       param.UserID,
		Name:         param.Name,
		Score:        param.Score,
		Level:        param.Level,
		Category:     param.Category,
		Suit:         param.Suit,
		Equipment:    param.Equipment,
		Place:        param.Place,
		TrainTarget:  param.TrainTarget,
		BodyTarget:   param.BodyTarget,
		SaleType:     param.SaleType,
		TrainerSex:   param.TrainerSex,
		TrainerSkill: param.TrainerSkill,
	})
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	paging := dto.Paging{
		TotalCount: totalCount,
		TotalPage:  cs.GetTotalPage(totalCount, size),
		Page:       page,
		Size:       size,
	}
	return parserCourseProductSummaries(datas), &paging, nil
}

func (cs *course) GetCourseStatistic(c *gin.Context, courseID int64) (*dto.CourseStatistic, errcode.Error) {
	var course dto.CourseStatistic
	if err := cs.courseRepo.FindCourseOutput(courseID, &course); err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	//查找 ReviewStatistic
	var reviewStatistic dto.ReviewStatisticSummary
	if err := cs.reviewStatisticRepo.FindReviewStatisticOutput(courseID, &reviewStatistic); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "review_statistic repo", err)
	}
	course.Review = &reviewStatistic
	//查找 CourseUsageStatistic
	var courseUsageStatistic dto.CourseUsageStatistic
	if err := cs.courseUsageStatisticRepo.FindCourseUsageStatisticOutput(courseID, &courseUsageStatistic); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "course_usage_statistic repo", err)
	}
	course.CourseUsageStatistic = &courseUsageStatistic
	return &course, nil
}

func (cs *course) GetCourseStatisticSummaries(c *gin.Context, userID int64, pagingParam *dto.PagingParam) ([]*dto.CourseStatisticSummary, *dto.Paging, errcode.Error) {
	//設置分頁
	var paging *model.PagingParam
	if pagingParam != nil {
		offset, limit := cs.GetPagingIndex(pagingParam.Page, pagingParam.Size)
		paging = &model.PagingParam{
			Offset: offset,
			Limit:  limit,
		}
	}
	datas, amount, err := cs.courseRepo.FindCourseStatisticSummaries(userID, &model.OrderBy{
		OrderType: global.ASC,
		Field:     "create_at",
	}, paging)
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	courses := make([]*dto.CourseStatisticSummary, 0)
	for _, data := range datas {
		courses = append(courses, dto.NewCourseStatisticSummary(data))
	}
	pagingResult := dto.Paging{
		TotalCount: amount,
		TotalPage:  cs.GetTotalPage(amount, pagingParam.Size),
		Page:       pagingParam.Page,
		Size:       pagingParam.Size,
	}
	return courses, &pagingResult, nil
}

func (cs *course) GetProgressCourseAssetSummaries(c *gin.Context, userID int64, page int, size int) ([]*dto.CourseAssetSummary, *dto.Paging, errcode.Error) {
	offset, limit := cs.GetPagingIndex(page, size)
	datas, err := cs.courseRepo.FindProgressCourseAssetSummaries(userID, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	totalCount, err := cs.courseRepo.FindProgressCourseAssetCount(userID)
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	paging := dto.Paging{
		TotalCount: totalCount,
		TotalPage:  cs.GetTotalPage(totalCount, size),
		Page:       page,
		Size:       size,
	}
	courses := make([]*dto.CourseAssetSummary, 0)
	for _, data := range datas {
		course := dto.NewCourseAssetSummary(data)
		courses = append(courses, &course)
	}
	return courses, &paging, nil
}

func (cs *course) GetChargeCourseAssetSummaries(c *gin.Context, userID int64, page int, size int) ([]*dto.CourseAssetSummary, *dto.Paging, errcode.Error) {
	offset, limit := cs.GetPagingIndex(page, size)
	datas, err := cs.courseRepo.FindChargeCourseAssetSummaries(userID, &model.PagingParam{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	totalCount, err := cs.courseRepo.FindChargeCourseAssetCount(userID)
	if err != nil {
		return nil, nil, cs.errHandler.Set(c, "course repo", err)
	}
	paging := dto.Paging{
		TotalCount: totalCount,
		TotalPage:  cs.GetTotalPage(totalCount, size),
		Page:       page,
		Size:       size,
	}
	courses := make([]*dto.CourseAssetSummary, 0)
	for _, data := range datas {
		course := dto.NewCourseAssetSummary(data)
		courses = append(courses, &course)
	}
	return courses, &paging, nil
}

func (cs *course) GetCourseAsset(c *gin.Context, userID int64, courseID int64) (*dto.CourseAsset, errcode.Error) {
	//查詢課表詳情
	courseData, err := cs.courseRepo.FindCourseAsset(courseID, userID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	courseDto := dto.NewCourseAsset(courseData)
	//查詢課表購買狀態
	courseDto.AllowAccess = 0
	if global.SaleType(courseDto.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := cs.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo != nil {
			courseDto.AllowAccess = subscribeInfo.Status
		}
	}
	if global.SaleType(courseDto.SaleType) == global.SaleTypeFree || global.SaleType(courseDto.SaleType) == global.SaleTypeCharge {
		asset, err := cs.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: courseID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "course repo", err)
		}
		if asset != nil {
			if asset.Available == 1 {
				courseDto.AllowAccess = 1
			}
		}
	}
	//查詢課表收藏狀態
	courseDto.Favorite = 0
	favoriteCourse, err := cs.favoriteRepo.FindFavoriteCourse(userID, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "favorite repo", err)
	}
	if favoriteCourse.CourseID > 0 {
		courseDto.Favorite = 1
	}
	return &courseDto, nil
}

func (cs *course) GetCourseAssetStructure(c *gin.Context, userID int64, courseID int64) (*dto.CourseAssetStructure, errcode.Error) {
	courseData, err := cs.courseRepo.FindCourseAsset(courseID, userID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	if global.ScheduleType(courseData.ScheduleType) != global.SingleScheduleType {
		return nil, cs.errHandler.Custom(8999, errors.New("此查詢只支援單一訓練課表"))
	}
	courseDto := dto.NewCourseAssetStructure(courseData)
	planDatas, err := cs.planRepo.FindPlanAssets(userID, courseID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "plan repo", err)
	}
	for _, planData := range planDatas {
		plan := dto.NewPlanAssetStructure(planData)
		workoutDatas, err := cs.workoutRepo.FindWorkoutAssets(userID, planData.ID)
		if err != nil {
			return nil, cs.errHandler.Set(c, "workout repo", err)
		}
		for _, workoutData := range workoutDatas {
			workout := dto.NewWorkoutAssetStructure(workoutData)
			workoutSetDatas, err := cs.workoutSetRepo.FindWorkoutSetsByWorkoutID(workout.ID, &userID)
			if err != nil {
				return nil, cs.errHandler.Set(c, "workout set repo", err)
			}
			for _, workoutSetData := range workoutSetDatas {
				workoutSet := dto.NewWorkoutSet(workoutSetData)
				workout.WorkoutSets = append(workout.WorkoutSets, &workoutSet)
			}
			plan.Workouts = append(plan.Workouts, &workout)
		}
		courseDto.Plans = append(courseDto.Plans, &plan)
	}
	//查詢課表購買狀態
	courseDto.AllowAccess = 0
	if global.SaleType(courseDto.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := cs.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo != nil {
			courseDto.AllowAccess = subscribeInfo.Status
		}
	}
	if global.SaleType(courseDto.SaleType) == global.SaleTypeFree || global.SaleType(courseDto.SaleType) == global.SaleTypeCharge {
		asset, err := cs.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: courseID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "course repo", err)
		}
		if asset != nil {
			if asset.Available == 1 {
				courseDto.AllowAccess = 1
			}
		}
	}
	//查詢課表收藏狀態
	courseDto.Favorite = 0
	favoriteCourse, err := cs.favoriteRepo.FindFavoriteCourse(userID, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "favorite repo", err)
	}
	if favoriteCourse.CourseID > 0 {
		courseDto.Favorite = 1
	}
	return &courseDto, nil
}

func (cs *course) GetCourseProductStructure(c *gin.Context, userID int64, courseID int64) (*dto.CourseProductStructure, errcode.Error) {
	courseData, err := cs.courseRepo.FindCourseProduct(courseID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "course repo", err)
	}
	if global.ScheduleType(courseData.ScheduleType) != global.SingleScheduleType {
		return nil, cs.errHandler.Custom(8999, errors.New("此查詢只支援單一訓練課表"))
	}
	courseDto := dto.NewCourseProductStructure(courseData)
	planDatas, err := cs.planRepo.FindPlansByCourseID(courseID)
	if err != nil {
		return nil, cs.errHandler.Set(c, "plan repo", err)
	}
	for _, planData := range planDatas {
		plan := dto.NewPlanStructure(planData)
		workoutDatas, err := cs.workoutRepo.FindWorkoutsByPlanID(planData.ID)
		if err != nil {
			return nil, cs.errHandler.Set(c, "workout repo", err)
		}
		for _, workoutData := range workoutDatas {
			workout := dto.NewWorkoutStructure(workoutData)
			workoutSetDatas, err := cs.workoutSetRepo.FindWorkoutSetsByWorkoutID(workout.ID, &userID)
			if err != nil {
				return nil, cs.errHandler.Set(c, "workout set repo", err)
			}
			for _, workoutSetData := range workoutSetDatas {
				workoutSet := dto.NewWorkoutSet(workoutSetData)
				workout.WorkoutSets = append(workout.WorkoutSets, &workoutSet)
			}
			plan.Workouts = append(plan.Workouts, &workout)
		}
		courseDto.Plans = append(courseDto.Plans, &plan)
	}
	//查詢課表購買狀態
	courseDto.AllowAccess = 0
	if global.SaleType(courseDto.SaleType) == global.SaleTypeSubscribe {
		subscribeInfo, err := cs.subscribeInfoRepo.FindSubscribeInfo(userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "subscribe info repo", err)
		}
		if subscribeInfo != nil {
			courseDto.AllowAccess = subscribeInfo.Status
		}
	}
	if global.SaleType(courseDto.SaleType) == global.SaleTypeFree || global.SaleType(courseDto.SaleType) == global.SaleTypeCharge {
		asset, err := cs.userCourseAssetRepo.FindUserCourseAsset(&model.FindUserCourseAssetParam{
			UserID:   userID,
			CourseID: courseID,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cs.errHandler.Set(c, "course repo", err)
		}
		if asset != nil {
			if asset.Available == 1 {
				courseDto.AllowAccess = 1
			}
		}
	}
	//查詢課表收藏狀態
	courseDto.Favorite = 0
	favoriteCourse, err := cs.favoriteRepo.FindFavoriteCourse(userID, courseID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, cs.errHandler.Set(c, "favorite repo", err)
	}
	if favoriteCourse.CourseID > 0 {
		courseDto.Favorite = 1
	}
	return &courseDto, nil
}

func (cs *course) UploadCourseCoverByID(c *gin.Context, courseID int64, param *dto.UploadCourseCoverParam) (*dto.CourseCover, errcode.Error) {
	//上傳照片
	newImageNamed, err := cs.uploader.UploadCourseCover(param.File, param.CoverNamed)
	if err != nil {
		if strings.Contains(err.Error(), "9007") {
			return nil, cs.errHandler.FileTypeError()
		}
		if strings.Contains(err.Error(), "9008") {
			return nil, cs.errHandler.FileSizeError()
		}
		cs.logger.Set(c, handler.Error, "Resource Handler", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//查詢課表封面
	var course struct {
		Cover string `gorm:"column:cover"`
	}
	if err := cs.courseRepo.FindCourseByID(nil, courseID, &course); err != nil {
		cs.logger.Set(c, handler.Error, "Course Repo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//修改課表資訊
	if err := cs.courseRepo.UpdateCourseByID(nil, courseID, &model.UpdateCourseParam{
		Cover: &newImageNamed,
	}); err != nil {
		cs.logger.Set(c, handler.Error, "CourseRepo", cs.errHandler.SystemError().Code(), err.Error())
		return nil, cs.errHandler.SystemError()
	}
	//刪除舊照片
	if len(course.Cover) > 0 {
		if err := cs.resHandler.DeleteCourseCover(course.Cover); err != nil {
			cs.logger.Set(c, handler.Error, "ResHandler", cs.errHandler.SystemError().Code(), err.Error())
		}
	}
	return &dto.CourseCover{Cover: newImageNamed}, nil
}

func (cs *course) CourseSubmit(c *gin.Context, courseID int64) errcode.Error {
	//驗證課表填寫完整性
	entity, err := cs.courseRepo.FindCourseByCourseID(courseID)
	if err != nil {
		return cs.errHandler.Set(c, "course repo", err)
	}
	if err := cs.VerifyCourse(entity); err != nil {
		return cs.errHandler.Set(c, "verify course", err)
	}
	tx := cs.transactionRepo.CreateTransaction()
	//送審課表(測試暫時將課表狀態改為"銷售中")
	var courseStatus = int(global.Sale)
	if err := cs.courseRepo.UpdateCourseByID(tx, courseID, &model.UpdateCourseParam{
		CourseStatus: &courseStatus,
	}); err != nil {
		tx.Rollback()
		return cs.errHandler.Set(c, "course repo", err)
	}
	courseCount, err := cs.trainerStatRepo.CalculateTrainerCourseCount(tx, entity.UserID)
	if err != nil {
		tx.Rollback()
		return cs.errHandler.Set(c, "trainer stat repo", err)
	}
	if err := cs.trainerStatRepo.SaveTrainerStatistic(tx, entity.UserID, &model.SaveTrainerStatisticParam{
		CourseCount: &courseCount,
	}); err != nil {
		tx.Rollback()
		return cs.errHandler.Set(c, "trainer stat repo", err)
	}
	cs.transactionRepo.FinishTransaction(tx)
	return nil
}

func (cs *course) GetCourseStatus(c *gin.Context, courseID int64) (global.CourseStatus, errcode.Error) {
	course := struct {
		CourseStatus int `json:"course_status"`
	}{}
	if err := cs.courseRepo.FindCourseByID(nil, courseID, &course); err != nil {
		return 0, cs.errHandler.Set(c, "course repo", err)
	}
	return global.CourseStatus(course.CourseStatus), nil
}

func (cs *course) VerifyCourse(course *model.Course) error {
	if global.SaleType(course.SaleType) == global.SaleTypeNone {
		return errors.New("需指定 sale type")
	}
	if global.SaleType(course.SaleType) == global.SaleTypeCharge && course.Sale == nil {
		return errors.New("付費課表需指定 sale item")
	}
	if course.Trainer == nil {
		return errors.New(strconv.Itoa(errcode.UpdateError))
	}
	return nil
}

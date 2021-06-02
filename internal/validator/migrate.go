package validator

type MigrateStepUri struct {
	Step int `uri:"step" binding:"required" example:"1"`
}

type MigrateVersionUri struct {
	Version int `uri:"version" binding:"required" example:"20210317124555"`
}


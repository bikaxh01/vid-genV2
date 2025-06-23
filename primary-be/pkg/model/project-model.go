package model

import (
	
	

	"gorm.io/gorm"
)

type Status string

const (
	GENERATING_PLAN   Status = "generating_plan"
	GENERATING_SCENES Status = "generating_scenes"
	VIDEO_MERGING     Status = "video_merging"
)

type Project struct {
	gorm.Model
	ID          string `gorm:"type:uuid;default:uuid_generate_v4():primaryKey`
	UserId      string `gorm:"type:uuid"`
	Prompt      string
	Title       string
	Description string
	Plan        string
	Status      Status
	Scenes      []Scene  `gorm:"foreignKey:ProjectId"`

}

type Scene struct {
	gorm.Model
	ID          string `gorm:"type:uuid;default:uuid_generate_v4():primaryKey`
	ProjectId   string  `gorm:"type:uuid"`
	Title       string
	Description string
	Code        string
	FileKey     string
	VideoUrl    string
	
}







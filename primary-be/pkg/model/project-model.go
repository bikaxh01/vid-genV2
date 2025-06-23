package model

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	PLAN_GENERATING   Status = "generating_plan"
	GENERATING_SCENES Status = "generating_scenes"
	VIDEO_MERGING     Status = "video_merging"
)

type Project struct {
	gorm.Model
	ID          string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserId      string         `gorm:"type:uuid" json:"userId"`
	Prompt      string         `json:"prompt"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Plan        sql.NullString `json:"plan"`
	Status      Status         `json:"status"`
	Scenes      []Scene        `gorm:"foreignKey:ProjectId" json:"scenes"`
}

type Scene struct {
	gorm.Model
	ID          string `gorm:"type:uuid;default:uuid_generate_v4():primaryKey" json:"id"`
	ProjectId   string `gorm:"type:uuid" json:"projectId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	FileKey     string `json:"fileKey"`
	VideoUrl    string `json:"videoUrl"`
}

func (p *Project) CreateProject() {
	p.ID = uuid.New().String()
	p.Status = Status(PLAN_GENERATING)
// save to db
}

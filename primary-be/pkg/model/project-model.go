package model

import (
	

	"github.com/bikaxh/vid-gen/primary-be/pkg/db"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"time"
)

type Status string

const (
	PLAN_GENERATING   Status = "generating_plan"
	GENERATING_SCENES Status = "generating_scenes"
	VIDEO_MERGING     Status = "video_merging"
)

type Project struct {
	ID          string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserId      string         `gorm:"type:uuid" json:"userId"`
	Prompt      string         `json:"prompt"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Plan        datatypes.JSON `json:"plan"`
	Status      Status         `json:"status"`
	Scenes      []Scene        `gorm:"foreignKey:ProjectId" json:"scenes"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Scene struct {
	ID          string `gorm:"type:uuid;primaryKey" json:"id"`
	ProjectId   string `gorm:"type:uuid" json:"projectId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	FileKey     string `json:"fileKey"`
	VideoUrl    string `json:"videoUrl"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *Project) CreateProject(plan map[string]interface{}) (*Project, error) {
	p.ID = uuid.New().String()
	p.Status = Status(PLAN_GENERATING)

	// Extract metadata safely

	if metadata, ok := plan["metaData"].(map[string]interface{}); ok {
		if title, ok := metadata["title"].(string); ok {
			p.Title = title
		}
		if description, ok := metadata["description"].(string); ok {
			p.Description = description
		}
	}

	result := db.Db.Create(p)

	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil

}

func (p *Project) SavePlan() (*Project, error) {

	result := db.Db.Model(p).Updates(Project{Plan: p.Plan, Title: p.Title, Description: p.Description})

	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

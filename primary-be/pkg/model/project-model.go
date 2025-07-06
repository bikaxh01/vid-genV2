package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bikaxh/vid-gen/primary-be/pkg/db"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Status string

const (
	PLAN_GENERATING   Status = "generating_plan"
	GENERATING_SCENES Status = "generating_scenes"
	VIDEO_MERGING     Status = "video_merging"
	COMPLETED         Status = "completed"
)

type Project struct {
	ID          string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserId      string         `gorm:"type:uuid" json:"userId"`
	Prompt      string         `json:"prompt"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Plan        datatypes.JSON `json:"plan"`

	Status    Status  `json:"status"`
	VideoUrl  *string `json:"videoUrl,omitempty"`
	Scenes    []Scene `gorm:"foreignKey:ProjectId" json:"scenes"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Scene struct {
	ID          string  `gorm:"type:uuid;primaryKey" json:"id"`
	ProjectId   string  `gorm:"type:uuid" json:"projectId"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Code        string  `json:"code"`
	FileKey     *string `json:"fileKey,omitempty"`
	VideoUrl    *string `json:"videoUrl,omitempty"`
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

func (p *Project) SaveVideoUrl() error {

	// result := db.Db.Model(p).Where("id = ?", p.ID).UpdateColumn("video_url", p.VideoUrl)
	result := db.Db.Model(p).Updates(Project{VideoUrl: p.VideoUrl, Status: COMPLETED})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Scene) SaveScene(projectId string) (*Scene, error) {
	s.ID = uuid.New().String()
	s.ProjectId = projectId

	result := db.Db.Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil

}

func PushToQueue(data Project) {
	var ctx = context.Background()
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	jsonString := string(jsonBytes)

	db.RedisClient.LPush(ctx, "generate_scene", string(jsonString))
}

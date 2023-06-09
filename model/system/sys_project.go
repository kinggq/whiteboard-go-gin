package system

import "time"

type ProjProjects struct {
	ProjectId   int        `json:"project_id"`
	ProjectName string     `json:"project_name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	IsDeleted   bool       `json:"is_deleted"`
}

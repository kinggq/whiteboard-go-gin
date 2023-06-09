package system

import "time"

type ProjTaskLists struct {
	ListId    int        `json:"list_id"`
	ProjectId int        `json:"project_id"`
	ListTitle string     `json:"list_title"`
	SortOrder int        `json:"sort_order"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	IsDeleted bool       `json:"is_deleted"`
}

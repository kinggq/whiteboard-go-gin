package system

import "time"

type PriorityID string

const (
	PriorityID3 PriorityID = "3"
	PriorityID2 PriorityID = "2"
	PriorityID1 PriorityID = "1"
	PriorityID0 PriorityID = "0"
)

type ProjTasks struct {
	TaskId      int         `json:"task_id"`
	ListId      int         `json:"list_id"`
	TaskTitle   string      `json:"task_title"`
	PriorityId  *PriorityID `json:"priority_id" binding:"omitempty,oneof=3 2 1 0"`
	CreatorId   int         `json:"creator_id"`
	Description string      `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
	IsDeleted   bool        `json:"is_deleted"`
}

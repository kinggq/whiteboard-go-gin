package request

type Task struct {
	TaskTitle string `json:"task_title"`
	CreatorId string `json:"creator_id"`
}

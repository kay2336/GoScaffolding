package form

type CreateTask struct {
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
}

type UpdateTaskReq struct {
	Id         uint   `json:"id" form:"id"`
	Title      string `gorm:"index;not null"`
	Status     int    `gorm:"default:0"`
	Content    string `gorm:"type:longtext"`
	UpdateTime int64  `gorm:"default:0"`
}

type ShowTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type DeleteTaskReq struct {
	Id uint `json:"id" form:"id"`
}

type SearchTaskReq struct {
	Info string `form:"info" json:"info"`
}

type ListTasksReq struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

// swagger:response Resp
type TaskResp struct {
	UserId     uint
	Id         uint   `json:"id" example:"1"`       // 任务ID
	Title      string `json:"title" example:"吃饭"`   // 题目
	Content    string `json:"content" example:"睡觉"` // 内容
	View       uint64 `json:"view" example:"32"`    // 浏览量
	Status     int    `json:"status" example:"0"`   // 状态(0未完成，1已完成)
	CreatedAt  int64  `json:"created_at"`
	UpdateTime int64  `json:"update_time"`
}

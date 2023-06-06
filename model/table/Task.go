package table

// Task
// 任务模型
type Task struct {
	Id         uint `gorm:"AUTO_INCREMENT"`
	UserId     uint
	Title      string
	Status     int    `gorm:"default:0"`
	Content    string `gorm:"type:longtext"`
	CreatedAt  int64
	UpdateTime int64
}

//func (Task *Task) View() uint64 {
//	// 增加点击数
//	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.UserId)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}
//
//// AddView
//func (Task *Task) AddView() {
//	cache.RedisClient.Incr(cache.TaskViewKey(Task.UserId))                      // 增加视频点击数
//	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.UserId))) // 增加排行点击数
//}

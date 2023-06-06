package dao

import (
	"awesomeProject/model/table"
	"awesomeProject/sql"
	"gorm.io/gorm"
)

type TaskDao struct {
	db *gorm.DB
}

// NewTaskDao
// 实例化dao层
func NewTaskDao() *TaskDao {
	return &TaskDao{db: sql.GetMysqlDB()}
}

// Create
// 创建
func (dao *TaskDao) Create(task *table.Task) (err error) {
	return dao.db.Model(&table.Task{}).Create(&task).Error
}

// ListTask
// 列表，通过上下文id查询表（task id与user id关联）
// 返回备忘录内容和 总和
func (dao *TaskDao) ListTask(start, limit int, userId uint) (t []*table.Task, total int64, err error) {
	err = dao.db.Model(&table.Task{}).Preload("UserId").Where("user_id = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&t).Error

	return
}

// FindById
// 通过task_id查找
func (that *TaskDao) FindById(id uint) (task *table.Task, err error) {
	err = that.db.Model(&table.Task{}).Where("id=?", id).First(&task).Error

	return
}

// Delete
// 删除
func (that *TaskDao) Delete(id uint) (err error) {
	return that.db.Model(&table.Task{}).Delete(&table.Task{}, id).Error
}

// Update
// 更新
func (dao *TaskDao) Update(task *table.Task) (err error) {
	return dao.db.Model(&table.Task{}).Where("id = ?", task.Id).
		Updates(&table.Task{Title: task.Title, Content: task.Content,
			Status: task.Status, UpdateTime: task.UpdateTime}).Error
}

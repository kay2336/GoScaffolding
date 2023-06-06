package service

import (
	"awesomeProject/dao"
	"awesomeProject/model/form"
	"awesomeProject/model/response"
	"awesomeProject/model/table"
	"awesomeProject/pkg/utils"
	"context"
	"log"
	"time"
)

type TaskSrv struct {
}

// NewTaskSrv
// 实例化TaskSrv
func NewTaskSrv() *TaskSrv {
	return &TaskSrv{}
}

// Create
// 创建一个Task
func (that *TaskSrv) Create(ctx context.Context, cTask *form.CreateTask) (response interface{}, err error) {
	// dao层实例化
	taskDao := dao.NewTaskDao()
	//userDao := dao.NewUserDao()

	// 从上下文中获取用户信息（此处只需要为id）
	userInfo, err := utils.GetUserInfoFromCtx(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	// 实体转换
	task := &table.Task{
		UserId:     userInfo.Id,
		Title:      cTask.Title,
		Content:    cTask.Content,
		Status:     cTask.Status,
		CreatedAt:  time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	// 创建task
	if err = taskDao.Create(task); err != nil {
		//utils.LogrusObj.Info(err)
		log.Println(err)
		return
	}

	//返回responseSuccess
	return resp.RespSuccess(), nil
}

// ListTask
// 列表
func (that *TaskSrv) ListTask(ctx context.Context, lTask *form.ListTasksReq) (response interface{}, err error) {
	// dao层实例化
	taskDao := dao.NewTaskDao()

	// 从上下文中获取用户信息（此处只需要为id）
	userInfo, err := utils.GetUserInfoFromCtx(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	// 列表
	tasks, total, err := taskDao.ListTask(lTask.Start, lTask.Limit, userInfo.Id)
	if err != nil {
		log.Println(err)
		//utils.LogrusObj.Info(err)
		return
	}

	// 遍历备忘录，存入taskRespList
	taskRespList := make([]*form.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &form.TaskResp{
			Id:      v.UserId,
			Title:   v.Title,
			Content: v.Content,
			//View:       v.View(),
			Status:     v.Status,
			CreatedAt:  v.CreatedAt,
			UpdateTime: v.UpdateTime,
		})
	}

	return resp.RespList(taskRespList, total), nil
}

// ShowTask
// 展示一个已有的Task
func (s *TaskSrv) ShowTask(ctx context.Context, sTask *form.ShowTaskReq) (response interface{}, err error) {
	// dao层实例化
	taskDao := dao.NewTaskDao()

	// 从上下文中获取用户信息
	userInfo, err := utils.GetUserInfoFromCtx(ctx)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	// 通过id获取task
	task, err := taskDao.FindById(sTask.Id)
	if err != nil {
		log.Println(err)
		//utils.LogrusObj.Info(err)
		return
	}

	// 实体转换
	respTask := &form.TaskResp{
		Id:      sTask.Id,
		UserId:  userInfo.Id,
		Title:   task.Title,
		Content: task.Content,
		//View:       task.View(),
		Status:     task.Status,
		CreatedAt:  task.CreatedAt,
		UpdateTime: task.UpdateTime,
	}

	//task.AddView() // 增加点击数
	return resp.RespSuccessWithData(respTask), nil
}

// Delete
// 删除一个Task
func (s *TaskSrv) Delete(ctx context.Context, dTask *form.DeleteTaskReq) (response interface{}, err error) {
	// dao层实例化
	taskDao := dao.NewTaskDao()

	//通过task id删除
	err = taskDao.Delete(dTask.Id)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}

	return resp.RespSuccess(), nil
}

// Update
// 更新一个Task
func (s *TaskSrv) Update(ctx context.Context, uTask *form.UpdateTaskReq) (response interface{}, err error) {
	// dao层实例化
	taskDao := dao.NewTaskDao()

	// 实体转换
	task := &table.Task{
		Id:         uTask.Id,
		Title:      uTask.Title,
		Content:    uTask.Content,
		Status:     uTask.Status,
		CreatedAt:  time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	//log.Println(task)
	//更新task
	err = taskDao.Update(task)
	if err != nil {
		utils.LogrusObj.Info(err)
		return
	}
	return resp.RespSuccess(), nil
}

//func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
//	u, err := ctl.GetUserInfo(ctx)
//	if err != nil {
//		util.LogrusObj.Info(err)
//		return
//	}
//	tasks, err := dao.NewTaskDao(ctx).SearchTask(u.Id, req.Info)
//	if err != nil {
//		util.LogrusObj.Info(err)
//		return
//	}
//	taskRespList := make([]*types.TaskResp, 0)
//	for _, v := range tasks {
//		taskRespList = append(taskRespList, &types.TaskResp{
//			ID:        v.ID,
//			Title:     v.Title,
//			Content:   v.Content,
//			Status:    v.Status,
//			View:      v.View(),
//			CreatedAt: v.CreatedAt.Unix(),
//			StartTime: v.StartTime,
//			EndTime:   v.EndTime,
//		})
//	}
//	return ctl.RespList(taskRespList, int64(len(taskRespList))), nil
//}

package goflow

import (
	"strings"
	"time"
)

//任务实体类
type Task struct {
	Id           string        `gorm:"size:36;primary_key"` //主键ID
	Version      int           //版本
	OrderId      string        `gorm:"size:36;not null;index"`  //流程实例ID
	TaskName     string        `gorm:"size:100;not null;index"` //任务名称
	DisplayName  string        `gorm:"size:200;not null"`       //任务显示名称
	PerformType  PERFORM_ORDER //任务参与方式
	TaskType     TASK_ORDER    //任务类型
	Operator     string        `gorm:"size:36"` //任务处理者ID
	CreateTime   time.Time     //任务创建时间
	FinishTime   *time.Time     //任务完成时间
	ExpireTime   *time.Time     //期望任务完成时间
	RemindTime   *time.Time     //提醒时间
	Action       string        `gorm:"size:200"`      //任务关联的Action(WEB为表单URL)
	ParentTaskId string        `gorm:"size:36;index"` //父任务ID
	Variable     string        `gorm:"size:2000"`     //任务附属变量(json存储)
	Model        *TaskModel    `gorm:"-"`             //Model对象
}

//根据ID得到任务
func (p *Task) GetTaskById(id string) bool {
	p.Id = id
	err := orm.Where("id = ?", id).Find(p).Error()
	PanicIf(err, "fail to GetTaskById")
	return err == nil
}

//得到活动任务
//func (p *Task) GetActiveTasks() []*Task {
//	tasks := make([]*Task, 0)
//	err := orm.Find(&tasks, p)
//	PanicIf(err, "fail to GetActiveTasks")
//	return tasks
//}

//根据OrderID得到活动任务
func (p *Task) GetActiveTasksByOrderId(orderId string) []*Task {
	tasks := make([]*Task, 0)
	err := orm.Where("order_id = ?", orderId).Find(&tasks, p).Error()
	PanicIf(err, "fail to GetActiveTasksByOrderId")
	return tasks
}

//得到任务角色
func (p *Task) GetTaskActors() []*TaskActor {
	taskActors := make([]*TaskActor, 0)

	err := orm.Where("task_id = ?", p.Id).Find(&taskActors).Error()
	PanicIf(err, "fail to GetTaskActors")
	return taskActors
}

//得到下一个ANY类型的任务
func GetNextAnyActiveTasks(parentTaskId string) []*Task {
	tasks := make([]*Task, 0)
	err := orm.Where("parent_task_id = ?", parentTaskId).Find(&tasks).Error()
	PanicIf(err, "fail to GetNextAnyActiveTasks")
	return tasks
}

//得到下一个ALL类型的任务
func GetNextAllActiveTasks(orderId string, taskName string, parentTaskId string) []*Task {
	historyTask := &HistoryTask{
		OrderId:      orderId,
		TaskName:     taskName,
		ParentTaskId: parentTaskId,
	}
	historyTasks := make([]*HistoryTask, 0)
	err := orm.Where("order_id = ? and task_name = ? and parent_task_id = ?", parentTaskId).Find(&historyTasks).Error()
	PanicIf(err, "fail to GetNextAllActiveTasks One")

	ids := make([]string, 0)
	for _, historyTask = range historyTasks {
		ids = append(ids, historyTask.Id)
	}
	tasks := make([]*Task, 0)
	err = orm.Where(`"ParentTaskId" in (?)`, strings.Join(ids, ",")).Find(&tasks).Error()
	PanicIf(err, "fail to GetNextAllActiveTasks Two")

	return tasks
}

//得到活动的任务（通过SQL）
func GetActiveTasksSQL(querystring string, args ...interface{}) ([]*Task, error) {
	tasks := make([]*Task, 0)
	err := orm.Where(querystring, args).Find(&tasks).Error()
	return tasks, err
}

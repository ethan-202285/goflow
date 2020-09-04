package goflow

import "time"

//任务实体类
type HistoryTask struct {
	Id           string        `gorm:"size:36;primary_key"`     //主键ID
	OrderId      string        `gorm:"size:36;not null;index"`  //流程实例ID
	TaskName     string        `gorm:"size:100;not null;index"` //任务名称
	DisplayName  string        `gorm:"size:200;not null"`       //任务显示名称
	PerformType  PERFORM_ORDER //任务参与方式
	TaskType     TASK_ORDER    `gorm:"not null"` //任务类型
	Operator     string        `gorm:"size:36"`  //任务处理者ID
	CreateTime   time.Time     //任务创建时间
	FinishTime   *time.Time     //任务完成时间
	ExpireTime   *time.Time     //期望任务完成时间
	Action       string        `gorm:"size:200"`      //任务关联的Action(WEB为表单URL)
	ParentTaskId string        `gorm:"size:36;index"` //父任务ID
	Variable     string        `gorm:"size:2000"`     //任务附属变量(json存储)
	TaskState    FLOW_STATUS   //任务状态
}

//根据ID得到HistoryTask
func (p *HistoryTask) GetHistoryTaskById(id string) bool {
	p.Id = id
	err := orm.Find(p, "id = ？", id).Error()
	PanicIf(err, "fail to GetHistoryTaskById")
	return err == nil
}

//通过HistoryTask生成Task
func (p *HistoryTask) Undo() *Task {
	task := &Task{
		Id:           p.Id,
		TaskName:     p.TaskName,
		DisplayName:  p.DisplayName,
		TaskType:     p.TaskType,
		ExpireTime:   p.ExpireTime,
		Action:       p.Action,
		ParentTaskId: p.ParentTaskId,
		Variable:     p.Variable,
		PerformType:  p.PerformType,
		Operator:     p.Operator,
	}
	return task
}

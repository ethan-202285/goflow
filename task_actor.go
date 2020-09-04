package goflow

//任务参与者
type TaskActor struct {
	Id      string `gorm:"size:36;primary_key;not null"`                        //主键ID
	TaskId  string `gorm:"size:36;not null;index"` //任务ID
	ActorId string `gorm:"size:36;not null"`                           //参与者ID
}

//通过任务ID，得到任务角色
func GetTaskActorsByTaskId(taskId string) []*TaskActor {
	taskActors := make([]*TaskActor, 0)
	err := orm.Where("task_id = ?", taskId).Find(&taskActors).Error()
	PanicIf(err, "fail to GetTaskActorsByTaskId")
	return taskActors
}

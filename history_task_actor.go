package goflow

//历史任务参与者
type HistoryTaskActor struct {
	Id      string `gorm:"size:36;primary_key"`    //主键ID
	TaskId  string `gorm:"size:36;not null;index"` //任务ID
	ActorId string `gorm:"size:36;not null"`       //参与者ID
}

package goflow

import (
	"strings"
	"time"
)

//抄送实例表
type CCOrder struct {
	Id         string      `gorm:"size:36;primary_key"` //主键ID
	OrderId    string      `gorm:"size:36;index"`       //流程实例ID
	ActorId    string      `gorm:"size:36"`             //操作者ID
	Creator    string      `gorm:"size:36"`             //流程实例创建者ID
	CreateTime time.Time   //流程实例创建时间
	FinishTime *time.Time   //流程实例完成时间
	State      FLOW_STATUS `gorm:"type:tinyint"` //流程实例状态
}

func GetCCOrder(orderId string, actorIds ...string) []*CCOrder {
	ccorders := make([]*CCOrder, 0)
	err := orm.Where(`"OrderId" = ? and "ActorId" in (?)`, orderId, strings.Join(actorIds, ",")).Find(&ccorders).Error()
	PanicIf(err, "fail to GetCCOrder")
	return ccorders
}

package goflow

import "time"

//历史流程实例实体类
type HistoryOrder struct {
	Id         string      `gorm:"size:36;primary_key"`    //主键ID
	ProcessId  string      `gorm:"size:36;index;not null"` //流程定义ID
	Creator    string      `gorm:"size:36"`                //流程实例创建者ID
	CreateTime time.Time   //流程实例创建时间
	ParentId   string      `gorm:"size:36;index"` //流程实例为子流程时，该字段标识父流程实例ID
	ExpireTime *time.Time   //流程实例期望完成时间
	Priority   int         //流程实例优先级
	OrderNo    string      `gorm:"size:100;index"` //流程实例编号
	Variable   string      `gorm:"size:2000"`      //流程实例附属变量
	OrderState FLOW_STATUS `gorm:"not null"`       //流程实例状态
	FinishTime *time.Time   //完成时间
}

//从Order对象获取数据构件HistoryOrder
func (p *HistoryOrder) DataByOrder(order *Order) {
	p.Id = order.Id
	p.ProcessId = order.ProcessId
	p.CreateTime = order.CreateTime
	p.ExpireTime = order.ExpireTime
	p.Creator = order.Creator
	p.ParentId = order.ParentId
	p.Priority = order.Priority
	p.OrderNo = order.OrderNo
	p.Variable = order.Variable
}

//根据ID得到HistoryOrder
func (p *HistoryOrder) GetHistoryOrderById(id string) bool {
	p.Id = id
	err := orm.Find(p, "id = ？", id).Error()
	PanicIf(err, "fail to GetHistoryOrderById")
	return err == nil
}

//通过HistoryOrder生成Order
func (p *HistoryOrder) Undo() *Order {
	order := &Order{
		Id:             p.Id,
		ProcessId:      p.ProcessId,
		CreateTime:     p.CreateTime,
		ExpireTime:     p.ExpireTime,
		Creator:        p.Creator,
		LastUpdator:    p.Creator,
		LastUpdateTime: p.FinishTime,
		ParentId:       p.ParentId,
		Priority:       p.Priority,
		OrderNo:        p.OrderNo,
		Variable:       p.Variable,
		Version:        0,
	}
	return order
}

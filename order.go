package goflow

import "time"

//流程工作单实体类（一般称为流程实例）
type Order struct {
	Id             string    `gorm:"size:36;primary_key"` //主键ID
	Version        int       //版本
	ProcessId      string    `gorm:"size:36;not null;index"` //流程定义ID
	Creator        string    `gorm:"size:36"`                //流程实例创建者ID
	CreateTime     time.Time //流程实例创建时间
	ParentId       string    `gorm:"size:36;index"` //流程实例为子流程时，该字段标识父流程实例ID
	ParentNodeName string    `gorm:"size:100"`      //流程实例为子流程时，该字段标识父流程哪个节点模型启动的子流程
	ExpireTime     *time.Time //流程实例期望完成时间
	LastUpdateTime *time.Time //流程实例上一次更新时间
	LastUpdator    string    `gorm:"size:36"` //流程实例上一次更新人员ID
	Priority       int       //流程实例优先级
	OrderNo        string    `gorm:"size:100;index"` //流程实例编号
	Variable       string    `gorm:"size:3000"`      //流程实例附属变量
}

//根据ID得到Order
func (p *Order) GetOrderById(id string) bool {
	p.Id = id
	err := orm.Where("id = ?", id).Find(p).Error()
	PanicIf(err, "fail to GetOrderById")
	return err == nil
}

//得到活动的Order（通过SQL）
func GetActiveOrdersSQL(querystring string, args ...interface{}) []*Order {
	orders := make([]*Order, 0)
	err := orm.Where(querystring, args).Find(&orders).Error()
	PanicIf(err, "fail to GetActiveOrdersSQL")
	return orders
}

package goflow

import (
	"reflect"
	"tingnide.pro/infra/tboot/pkg/tlog"
	"tingnide.pro/infra/tgorm"
)

var orm tgorm.Repository

//初始化数据库ORM引擎
func InitAccess(xorm tgorm.Repository) {
	orm = xorm

	xorm.AutoMigrate(new(HistoryOrder), new(HistoryTask), new(HistoryTaskActor),
		new(Order), new(Process), new(Surrogate), new(Task), new(TaskActor), new(CCOrder)).Error()
}

//保存实体对象
func Save(inf interface{}, id interface{}) {
	err := orm.Save(inf).Error()
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to insert %v %v", t, id)
	tlog.Infof("%v %v inserted", t, id)
}

//更新实体对象
func Update(inf interface{}, id interface{}) {
	err := orm.Where("id = ?", id).Save(inf).Error()

	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to update %v %v", t, id)
	tlog.Infof("%v %v updated", t, id)
}

//删除实体对象
func Delete(inf interface{}, id interface{}) {
	err := orm.Delete(inf, "id = ?", id).Error()
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to delete %v %v", t, id)
	tlog.Infof("%v %v deleted", t, id)
}

//删除实体对象
func DeleteObj(where string, inf interface{}) {
	err := orm.Where(where).Delete(inf).Error()
	t := reflect.TypeOf(inf)
	PanicIf(err, "fail to delete %v", t)
	tlog.Infof("%v deleted", t)
}

package goflow

import "time"

//流程定义实体类
type Process struct {
	Id             string        `gorm:"size:36;primary_key"` //主键ID
	Version        int           //版本
	Name           string        `gorm:"size:100;index"` //流程定义名称
	DisplayName    string        `gorm:"size:200"`       //流程定义显示名称
	InstanceAction string        `gorm:"size:200"`       //当前流程的实例Action,(Web为URL,一般为流程第一步的URL;APP需要自定义),该字段可以直接打开流程申请的表单
	State          FLOW_STATUS   //状态
	CreateTime     time.Time     //创建时间
	Creator        string        `gorm:"size:36"`    //创建人
	Content        string        `gorm:"type:text"` //流程定义XML
	Model          *ProcessModel `gorm:"-"`          //Model对象
}

//根据ID得到Process
func (p *Process) GetProcessById(id string) bool {
	p.Id = id
	err := orm.Where("id = ?", id).Find(p).Error()
	PanicIf(err, "fail to GetProcessById")
	return err == nil
}

//根据Process本身条件得到Process
func (p *Process) GetProcess() bool {
	err := orm.Find(p).Error()
	PanicIf(err, "fail to GetProcess")
	return err == nil
}

//设定Model对象
func (p *Process) SetModel(model *ProcessModel) {
	p.Model = model
	p.Name = model.Name
	p.DisplayName = model.DisplayName
	p.InstanceAction = model.InstanceAction
}

//得到最新的Process
func GetLatestProcess(name string) *Process {
	processes := make([]*Process, 0)
	err := orm.Where("name = ?", name).Order("version DESC").Limit(1).First(&processes).Error()
	PanicIf(err, "fail to GetLatestProcess")
	if len(processes) > 0 {
		return processes[0]
	} else {
		return nil
	}
}

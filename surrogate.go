package goflow

import (
	"time"
)

//委托代理
type Surrogate struct {
	Id          string           `gorm:"size:36;primary_key"`    //主键ID
	ProcessName string           `gorm:"size:36;not null"`       //流程名称
	Operator    string           `gorm:"size:36;not null;index"` //授权人
	Surrogate   string           `gorm:"size:36"`                //代理人
	OpTime      time.Time        //操作时间
	StartTime   time.Time        //开始时间
	EndTime     time.Time        //结束时间
	State       SURROGATE_STATUS //状态
}

//得到代理人（通过SQL）
func GetSurrogateSQL(querystring string, args ...interface{}) []*Surrogate {
	surrogates := make([]*Surrogate, 0)
	err := orm.Where(querystring, args).Find(&surrogates).Error()
	PanicIf(err, "fail to GetSurrogateSQL")
	return surrogates
}

package po

type Slink struct {
	Id         int64  `json:"idgen,omitempty" gorm:"idgen"`             // 主键
	SLink      string `json:"s_link,omitempty" gorm:"s_link"`           // 短链
	OLink      string `json:"o_link,omitempty" gorm:"o_link"`           // 原链接
	Deleted    int8   `json:"deleted,omitempty" gorm:"deleted"`         // 逻辑删除
	CreateTime int64  `json:"create_time,omitempty" gorm:"create_time"` // 创建时间
	UpdateTime int64  `json:"update_time,omitempty" gorm:"update_time"` // 更新时间
}

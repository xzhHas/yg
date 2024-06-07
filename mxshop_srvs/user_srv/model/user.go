package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`                       // 用户手机号，唯一索引，类型为 varchar(11)，不能为空
	Password string     `gorm:"type:varchar(100);not null"`                                              // 用户密码，类型为 varchar(100)，不能为空
	NickName string     `gorm:"type:varchar(20)"`                                                        // 用户昵称，类型为 varchar(20)
	Birthday *time.Time `gorm:"type:datetime"`                                                           // 用户生日，类型为 datetime，可为空
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'"` // 用户性别，类型为 varchar(6)，默认为男性，有相关注释
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员'"`                // 用户角色，类型为 int，默认为普通用户，有相关注释
}

package model

import (
	"gopkg.in/guregu/null.v4"
)

const LinkTable = "links"

type Link struct {
	ID        uint      `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Active    bool      `json:"active" gorm:"not null"`                           // 服务状态
	Comment   string    `json:"comment" gorm:"not null"`                          // 备注信息
	Origin    string    `json:"origin" gorm:"not null"`                           // 原始链接
	Short     string    `json:"short" gorm:"primaryKey;not null"`                 // 短链ID，全局唯一
	OwnerID   uint      `json:"owner_id" gorm:"foreignKey:OwnerID;references:ID"` // 创建短连接的userid
	StartTime null.Time `json:"start_time"`                                       // 起始时间，UTC
	EndTime   null.Time `json:"end_time"`                                         // 到期时间，UTC
}

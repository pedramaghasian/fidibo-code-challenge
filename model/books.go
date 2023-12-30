package model

type Books struct {
	ID    uint64 `gorm:"type:bigint;primary_key"`
	Title string `gorm:"type:varchar(191)"`
	Cover string `gorm:"type:varchar(191)"`
	Type  uint8  `gorm:"type:tinyint"`
}

package model

type Tag struct {
	TagID int `gorm:"primaryKey:autoIncrement:false"`
}

func (Tag) TableName() string {
	return "tags"
}

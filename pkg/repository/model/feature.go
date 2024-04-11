package model

type Feature struct {
	FeatureID int `gorm:"primaryKey:autoIncrement:false"`
}

func (Feature) TableName() string {
	return "features"
}

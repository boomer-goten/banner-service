package model

import (
	"banner-server/internal/api/model"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	BannerID  int  `gorm:"primaryKey"`
	Content   JSON `gorm:"type:json"`
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Banner) TableName() string {
	return "banners"
}

type JSON json.RawMessage
type ArrayInt []int32

func (a *ArrayInt) Scan(value interface{}) error {
	if value == nil {
		*a = ArrayInt{}
		return nil
	}
	data, ok := value.(string)
	if !ok {
		return errors.New("invalid data")
	}
	dataValue := []byte(data[1 : len(data)-1])
	var arr []int32
	for _, b := range dataValue {
		if b != 44 {
			value, err := strconv.Atoi(string(b))
			if err != nil {
				return errors.New("invalid data")
			}
			arr = append(arr, int32(value))
		}
	}
	*a = ArrayInt(arr)
	return nil
}

func (a ArrayInt) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Ошибка распаковки значения JSONB:", value))
	}
	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

func ConvertPostRequest(data *model.BannerPostRequest) (Banner, error) {
	var jsonData JSON
	dataByte, err := json.Marshal(data.Content)
	if err != nil {
		return Banner{}, err
	}
	err = jsonData.Scan(dataByte)
	if err != nil {
		return Banner{}, err
	}
	return Banner{
		Content:   jsonData,
		IsActive:  data.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// func (b *Banner) BeforeDelete(tx *gorm.DB) (err error) {
// 	tx.Model(&BannerTag{}).Where("banner_id = ?", b.BannerID).Delete(&BannerTag{}, b.BannerID)
// 	return nil
// }

func (b *Banner) AfterCreate(tx *gorm.DB) (err error) {
	data := tx.Statement.Context.Value("banner_tags").([]BannerTag)
	for i := range data {
		data[i].BannerID = b.BannerID
	}
	tx.Create(data)
	return nil
}

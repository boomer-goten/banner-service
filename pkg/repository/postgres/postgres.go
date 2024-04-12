package postgres

import (
	"banner-server/pkg/repository"
	"banner-server/pkg/repository/model"
	"database/sql"
	"encoding/json"
	"os"
	"strconv"
	"sync"

	request "banner-server/internal/api/model"

	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	once     = sync.Once{}
	database *Postgres
)

type Postgres struct {
	Db *gorm.DB
}

func New() *Postgres {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES")), &gorm.Config{})
		rawDB, _ := db.DB()
		rawDB.SetMaxOpenConns(52)
		rawDB.SetMaxIdleConns(52)
		rawDB.SetConnMaxLifetime(0)
		if err != nil {
			panic("coudn't connect to database")
		}
		database = &Postgres{db}
		database.MigrateDB()
	})
	return database
}

func (p *Postgres) MigrateDB() {
	if !p.Db.Migrator().HasTable(&model.Feature{}) {
		p.Db.AutoMigrate(&model.Feature{})
		insertFeature(p.Db)
	}
	if !p.Db.Migrator().HasTable(&model.Tag{}) {
		p.Db.AutoMigrate(&model.Tag{})
		insertTags(p.Db)
	}
	if !p.Db.Migrator().HasTable(&model.Banner{}) {
		p.Db.AutoMigrate(&model.Banner{})
	}
	if !p.Db.Migrator().HasTable(&model.BannerTag{}) {
		p.Db.AutoMigrate(&model.BannerTag{})
	}
}

func (p *Postgres) BannerGet(tagID, featureID, offset, limit int, role string) ([]request.BannerGet200ResponseInner, error) {
	tx := p.Db.Begin()
	var err error
	var rows *sql.Rows
	if tagID > 0 && featureID > 0 {
		tx = tx.Table("banners").Where("? = ANY(tag_ids) AND feature_id = ?", tagID, featureID).Select("banners.banner_id, tag_ids, feature_id, content, is_active, created_at, updated_at").Joins("join (SELECT banner_id, ARRAY_AGG(tag_id)::integer [] as tag_ids, feature_id from banner_tags GROUP BY banner_id, feature_id) as bt on banners.banner_id = bt.banner_id").Offset(offset).Order("banner_id")
	} else if tagID > 0 {
		tx = tx.Table("banners").Where("? = ANY(tag_ids)", tagID).Select("banners.banner_id, tag_ids, feature_id, content, is_active, created_at, updated_at").Joins("join (SELECT banner_id, ARRAY_AGG(tag_id)::integer [] as tag_ids, feature_id from banner_tags GROUP BY banner_id, feature_id) as bt on banners.banner_id = bt.banner_id").Offset(offset).Order("banner_id")
	} else if featureID > 0 {
		tx = tx.Table("banners").Where("feature_id = ?", featureID).Select("banners.banner_id, tag_ids, feature_id, content, is_active, created_at, updated_at").Joins("join (SELECT banner_id, ARRAY_AGG(tag_id)::integer [] as tag_ids, feature_id from banner_tags GROUP BY banner_id, feature_id) as bt on banners.banner_id = bt.banner_id").Offset(offset).Order("banner_id")
	} else {
		tx = tx.Table("banners").Select("banners.banner_id, tag_ids, feature_id, content, is_active, created_at, updated_at").Joins("join (SELECT banner_id, ARRAY_AGG(tag_id)::integer [] as tag_ids, feature_id from banner_tags GROUP BY banner_id, feature_id) as bt on banners.banner_id = bt.banner_id").Offset(offset).Order("banner_id")
	}
	if limit > 0 {
		rows, err = tx.Limit(limit).Rows()
	} else {
		rows, err = tx.Rows()
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		return nil, repository.ErrSelectItem
	}
	if err != nil {
		tx.Rollback()
		return nil, repository.ErrDb
	}
	result := make([]request.BannerGet200ResponseInner, 0, limit)
	for rows.Next() {
		var banner request.BannerGet200ResponseInner
		var value model.JSON
		var tag_ids model.ArrayInt
		rows.Scan(&banner.BannerId, &tag_ids, &banner.FeatureId, &value, &banner.IsActive, &banner.CreatedAt, &banner.UpdatedAt)
		if role == "user" && !banner.IsActive {
			continue
		}
		json.Unmarshal(value, &banner.Content)
		banner.TagIds = tag_ids
		result = append(result, banner)
	}
	tx.Commit()
	return result, nil
}

func (p *Postgres) BannerIdDelete(bannerID int) error {
	result := p.Db.Model(&model.BannerTag{}).Where("banner_id = ?", bannerID).Delete(&model.BannerTag{BannerID: bannerID})
	if result.RowsAffected == 0 {
		return repository.ErrDeleteFind
	}
	if result.Error != nil {
		return repository.ErrDb
	}
	return nil
}

func (p *Postgres) BannerIdPatch(id int, data *request.BannerIdPatchRequest) error {
	dataBannerTags, errConvert := model.ConvertPatchRequestTags(id, data)
	var updateMap map[string]interface{}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return repository.ErrDb
	}
	json.Unmarshal(dataByte, &updateMap)
	tx := p.Db.Begin()
	var banner_tags model.BannerTag
	var result *gorm.DB
	if len(updateMap) != 0 {
		result = tx.Model(model.Banner{}).Where("banner_id = ?", id).Updates(updateMap)
		if result.RowsAffected == 0 {
			tx.Rollback()
			return repository.ErrFoundItem
		}
		if result.Error != nil {
			tx.Rollback()
			return repository.ErrDb
		}
	}
	switch errConvert {
	case repository.ErrPatchTags:
		result = tx.Model(model.BannerTag{}).Where("banner_id = ?", id).Updates(model.BannerTag{FeatureID: dataBannerTags[0].FeatureID})
		if result.RowsAffected == 0 {
			tx.Rollback()
			return repository.ErrFoundItem
		}
		if result.Error != nil {
			tx.Rollback()
			return repository.ErrDb
		}
		tx.Commit()
		return nil
	case repository.ErrPatchTags:
		if err = tx.Model(model.BannerTag{}).Clauses(clause.Returning{}).Where("banner_id = ?", id).Delete(&banner_tags).Error; err != nil {
			tx.Rollback()
			return repository.ErrDb
		}
		for i := range dataBannerTags {
			dataBannerTags[i].FeatureID = banner_tags.FeatureID
		}
	case nil:
		if err = tx.Model(model.BannerTag{}).Where("banner_id = ?", id).Delete(&banner_tags).Error; err != nil {
			tx.Rollback()
			return repository.ErrDb
		}
	}
	result = tx.Model(model.BannerTag{}).CreateInBatches(dataBannerTags, len(dataBannerTags))
	if result.RowsAffected == 0 {
		tx.Rollback()
		return repository.ErrFoundItem
	}
	if result.Error != nil {
		tx.Rollback()
		return repository.ErrDb
	}
	tx.Commit()
	return nil
}

func (p *Postgres) BannerPost(data *request.BannerPostRequest) (int, error) {
	banner, err := model.ConvertPostRequest(data)
	if err != nil {
		return 0, err
	}
	tags := model.ConvertPostRequestTags(data)
	ctx := context.WithValue(context.Background(), "banner_tags", tags)
	if err := p.Db.WithContext(ctx).Create(&banner).Error; err != nil {
		return 0, repository.ErrCreateItem
	}
	return banner.BannerID, nil
}

func (p *Postgres) UserBannerGet(tag, feature int) ([]byte, bool, error) {
	tx := p.Db.Begin()
	var banner model.Banner
	var content []byte
	if err := tx.Table("banners").Where("tag_id = ? AND feature_id = ?", tag, feature).Select("content, is_active").Joins("join banner_tags on banners.banner_id = banner_tags.banner_id").First(&banner).Error; err != nil {
		tx.Rollback()
		if err != gorm.ErrRecordNotFound {
			return content, false, repository.ErrDb
		}
		return content, true, repository.ErrSelectItem
	}
	tx.Commit()
	value, err := banner.Content.Value()
	if err != nil {
		return content, false, repository.ErrDb
	}
	content = value.([]byte)
	return content, banner.IsActive, nil
}

func insertFeature(Db *gorm.DB) error {
	countTagsFeatures, err := strconv.ParseInt(os.Getenv("TAGS_FEATURES"), 10, 32)
	if err != nil {
		return err
	}
	features := make([]model.Feature, countTagsFeatures)
	Db.CreateInBatches(features, 1)
	return nil

}

func insertTags(Db *gorm.DB) error {
	countTagsFeatures, err := strconv.ParseInt(os.Getenv("TAGS_FEATURES"), 10, 32)
	if err != nil {
		return err
	}
	tags := make([]model.Tag, countTagsFeatures)
	Db.CreateInBatches(tags, 1)
	return nil
}

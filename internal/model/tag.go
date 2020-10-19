package model

import (
	"github.com/jinzhu/gorm"
	"go-blog/pkg/app"
)

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// 获得指定名称/状态的tag数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// 获得玩家指定页数的tag列表
func (t Tag) List(db *gorm.DB, pageOffset int, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Get(db *gorm.DB, tagID uint32, state uint8) (Tag, error) {
	var tag Tag
	var err error
	db = db.Where("id = ? AND state = ? AND is_del = ?", tagID, state, 0)
	err = db.First(&tag).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}

	return tag, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}
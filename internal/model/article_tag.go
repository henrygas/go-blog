package model

import "github.com/jinzhu/gorm"

type ArticleTag struct {
	*Model
	TagID uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

// 根据article_id查询ArticleTag记录
func (a ArticleTag) GetByArticleID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).
		First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}

	return articleTag, nil
}

// 根据tag_id查询ArticleTag记录
func (a ArticleTag) ListByTagID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND is_del = ?", a.TagID, 0).
		Find(&articleTags).Error; err != nil {
			return nil, err
	}

	return articleTags, nil
}

// 根据一批tag_ids查询对应的ArticleTag记录列表
func (a ArticleTag) ListByArticleIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND is_del = ?", articleIDs, 0).
		Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articleTags, nil
}

// 创建ArticleTag记录
func (a ArticleTag) Create(db *gorm.DB) error {
	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

// 根据article_id更新一条ArticleTag记录
func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Where("article_id = ? AND is_del = ?", a.ArticleID, 0).
		Limit(1).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

// 根据ID删除指定的ArticleTag记录
func (a ArticleTag) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

// 根据ID删除指定的一条ArticleTag记录
func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	if err := db.Where("article_id = ? AND is_del = ?", a.ArticleID, 0).Delete(&a).Limit(1).Error; err != nil {
		return err
	}

	return nil
}
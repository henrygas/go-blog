package model

import (
	"github.com/jinzhu/gorm"
	"go-blog/pkg/app"
)

type Article struct {
	*Model
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(&a).Updates(values).Where("id = ? AND is_del = ?", a.ID).Error; err != nil {
		return err
	}

	return nil
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.ID, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, nil
	}

	return article, nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error; err != nil {
		return err
	}

	return nil
}

// 关联查询的结果行对象
type ArticleRow struct {
	ArticleID uint32
	TagID uint32
	TagName string
	ArticleTitle string
	ArticleDesc string
	CoverImageUrl string
	Content string
}

// 	根据blog_article_tag表记录为驱动表，来关联查询article和tag大表，得到结果后收集到ArticleRow对象中
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset int, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title",
		"ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName() + " AS at").
		Joins("LEFT JOIN `" + Tag{}.TableName() + "` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `" + Article{}.TableName() + "` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}
	return articles, nil
}

// 根据指定tagID查询记录数
func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName() + " AS at").
		Joins("LEFT JOIN `" + Tag{}.TableName() + "` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `" + Article{}.TableName() + "` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}
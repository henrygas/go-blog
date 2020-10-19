package dao

import (
	"go-blog/internal/model"
	"go-blog/pkg/app"
)

type Article struct {
	ID uint32 `json:"id"`
	TagID uint32 `json:"tag_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State uint8 `json:"state"`
}

// 创建Article对象
func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title: param.Title,
		Desc: param.Desc,
		Content: param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State: param.State,
		Model: &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

// 更新Article记录
func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{
		Model: &model.Model{ID: param.ID},
	}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state": param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)
}

// 根据id和state查询Article记录
func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

// 删除指定ID的article记录
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

// 根据id和state查询Article数量
func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

// 根据id和state查询指定页的Article列表
func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page int, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

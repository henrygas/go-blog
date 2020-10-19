package dao

import "go-blog/internal/model"

func (d *Dao) GetArticleTagByArticleID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByArticleID(d.engine)
}

func (d *Dao) GetArticleTagListByTagID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{TagID: tagID}
	return articleTag.ListByTagID(d.engine)
}

func (d *Dao) GetArticleTagListByArticleIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByArticleIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID uint32, tagID uint32, createdBy string) error {
	articleTag := model.ArticleTag{
		Model: &model.Model{CreatedBy: createdBy},
		ArticleID: articleID,
		TagID: tagID,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID uint32, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	values := map[string]interface{}{
		"article_id": articleID,
		"tag_id": tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.UpdateOne(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.DeleteOne(d.engine)
}



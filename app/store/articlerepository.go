package store

import "github.com/Mishon-pon-pon/Blog/app/model"

// ArticleRepository ...
type ArticleRepository struct {
	store *Store
}

// Create ...
func (r *ArticleRepository) Create(a *model.Article) (*model.Article, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO Articles(Title, TextArticle) VALUES($1, $2);",
		a.Title,
		a.TextArticle,
	).Scan(&a.ArticleID); err != nil {
		return nil, err
	}
	return nil, nil
}

// GetArticles ...
func (r *ArticleRepository) GetArticles() ([]model.Article, error) {
	result, err := r.store.db.Query(`SELECT 
										ArticleId, 
										Title, 
										substr(TextArticle, 0, 219) || ltrim(substr(TextArticle, 219, 220), ' ') || '...' as TextArticle   
										FROM Articles;`)
	if err != nil {
		return nil, err
	}
	var articles []model.Article
	var a model.Article
	for result.Next() {
		err := result.Scan(&a.ArticleID, &a.Title, &a.TextArticle)
		if err != nil {
			panic(err)
		}
		articles = append(articles, a)
	}
	return articles, nil
}

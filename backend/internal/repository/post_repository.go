package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/model"
)

type PostRepository struct {
	db *gorm.DB
}

type PostSearchQuery struct {
	Title    string
	UserName string
	Category string
	Limit    int
	Offset   int
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Search(q PostSearchQuery) ([]model.Post, error) {
	var posts []model.Post

	query := r.db.
		Model(&model.Post{}).
		Joins("JOIN users ON users.id = posts.user_id").
		Joins("JOIN categories ON categories.id = posts.category_id")

	if q.Title != "" {
		query = query.Where("posts.title LIKE ?", "%"+q.Title+"%")
	}

	if q.UserName != "" {
		query = query.Where("users.name LIKE ?", "%"+q.UserName+"%")
	}

	if q.Category != "" {
		query = query.Where("categories.name LIKE ?", "%"+q.Category+"%")
	}

	if q.Limit == 0 {
		q.Limit = 10
	}

	query = query.
		Limit(q.Limit).
		Offset(q.Offset).
		Order("posts.created_at DESC")

	err := query.Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) Get(id uint) (model.Post, error) {
	var post model.Post
	err := r.db.Preload("User").Preload("Category").First(&post, id).Error
	return post, err
}

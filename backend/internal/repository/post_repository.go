package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/model"
	"hobby-blog/internal/domain/post"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Search(q post.SearchQuery) ([]model.Post, error) {
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

func (r *PostRepository) Create(param post.CreateInput) error {
	post := model.Post{
		UserID: param.UserID,
		CategoryID: param.CategoryID,
		Title: param.Title,
		Content: param.Content,
		Status: "draft",
	}

	return r.db.Create(&post).Error
}

package repository

import (
	"gorm.io/gorm"
	"hobby-blog/internal/domain/post"
	"hobby-blog/internal/model"
)

type postRepository struct {
	db *gorm.DB
}

type PostRepository interface {
	Search(post.SearchQuery) ([]model.Post, error)
	Get(uint) (model.Post, error)
	Create(post.CreateInput) error
	Update(post.UpdateInput) (model.Post, error)
	Delete(uint, uint) error
	GetMyPostsByUserID(uint) ([]model.Post, error)
	FindByID(uint) (model.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Search(q post.SearchQuery) ([]model.Post, error) {
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

	limit := q.Limit
	if limit == 0 {
		limit = 10
	}

	query = query.
		Limit(limit).
		Offset(q.Offset).
		Order("posts.created_at DESC")

	err := query.Preload("User").Preload("Category").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepository) Get(id uint) (model.Post, error) {
	var post model.Post
	err := r.db.Preload("User").Preload("Category").First(&post, id).Error
	return post, err
}

func (r *postRepository) Create(param post.CreateInput) error {
	p := model.Post{
		UserID:     param.UserID,
		CategoryID: param.CategoryID,
		Title:      param.Title,
		Content:    param.Content,
		Status:     post.StatusDraft,
	}

	return r.db.Create(&p).Error
}

func (r *postRepository) Update(param post.UpdateInput) (model.Post, error) {
	var post model.Post
	result := r.db.Model(&model.Post{}).
		Where("id = ? AND user_id = ?", param.ID, param.UserID).
		Updates(map[string]interface{}{
			"title":       param.Title,
			"content":     param.Content,
			"category_id": param.CategoryID,
			"status":      param.Status,
		})

	if result.Error != nil {
		return post, result.Error
	}

	if result.RowsAffected == 0 {
		return post, gorm.ErrRecordNotFound
	}
	err := r.db.Preload("User").Preload("Category").First(&post, param.ID).Error

	return post, err
}

func (r *postRepository) Delete(id uint, userID uint) error {

	result := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&model.Post{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *postRepository) GetMyPostsByUserID(userID uint) ([]model.Post, error) {
	var posts []model.Post
	result := r.db.Preload("Category").Where("user_id = ?", userID).Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (r *postRepository) FindByID(postID uint) (model.Post, error) {
	var post model.Post
	err := r.db.First(&post, postID).Error

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

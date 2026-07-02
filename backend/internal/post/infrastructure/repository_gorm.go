package infrastructure

import (
    "context"
    "gorm.io/gorm"

    postDomain "hobby-blog/internal/post/domain"
    userDomain "hobby-blog/internal/user/domain"
)

type postRepository struct {
    db *gorm.DB
}

func NewPostRepository(db *gorm.DB) postDomain.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Search(ctx context.Context, title, userName, category string, limit, offset int) ([]postDomain.Post, error) {
	var posts []Post

	query := r.db.WithContext(ctx).
		Model(&Post{}).
		Joins("JOIN users ON users.id = posts.user_id").
		Joins("JOIN categories ON categories.id = posts.category_id")

	if title != "" {
		query = query.Where("posts.title LIKE ?", "%"+title+"%")
	}

	if userName != "" {
		query = query.Where("users.name LIKE ?", "%"+userName+"%")
	}

	if category != "" {
		query = query.Where("categories.name LIKE ?", "%"+category+"%")
	}

	if limit == 0 {
		limit = 10
	}

	query = query.
		Limit(limit).
		Offset(offset).
		Order("posts.created_at DESC")

	err := query.Preload("User").Preload("Category").Preload("MediaFiles").Find(&posts).Error

	if err != nil {
		return nil, err
	}

	result := make([]postDomain.Post, 0, len(posts))

	for _, p := range posts {
		result = append(result, postDomain.Post{
			ID:         uint(p.ID),
			UserID:     userDomain.ID(p.UserID),
			CategoryID: p.CategoryID,
			Title:      p.Title,
			Content:    p.Content,
			Status:     postDomain.Status(p.Status),
		})
	}

	return result, nil
}

func (r *postRepository) GetByID(ctx context.Context, id uint) (*postDomain.Post, error) {
	var p Post
	err := r.db.WithContext(ctx).Preload("User").Preload("Category").Preload("MediaFiles").First(&p, id).Error
	if err != nil {
		return nil, err
	}

	result := postDomain.Post{
		ID:         uint(p.ID),
		UserID:     userDomain.ID(p.UserID),
		CategoryID: p.CategoryID,
		Title:      p.Title,
		Content:    p.Content,
		Status:     postDomain.Status(p.Status),
	}

	return &result, nil
}

func (r *postRepository) Create(ctx context.Context, post postDomain.Post) (*postDomain.Post, error) {
	p := Post{
		UserID:     uint(post.UserID),
		CategoryID: post.CategoryID,
		Title:      post.Title,
		Content:    post.Content,
		Status:     string(post.Status),
	}

	if err := r.db.WithContext(ctx).Create(&p).Error; err != nil {
		return nil, err
	}

	result := postDomain.Post{
		ID:         uint(p.ID),
		UserID:     userDomain.ID(p.UserID),
		CategoryID: p.CategoryID,
		Title:      p.Title,
		Content:    p.Content,
		Status:     postDomain.Status(p.Status),
	}
	return &result, nil
}

func (r *postRepository) Update(ctx context.Context, post postDomain.Post) (*postDomain.Post, error) {
	var existing Post
	result := r.db.WithContext(ctx).Model(&Post{}).
		Where("id = ?", post.ID).
		Updates(map[string]interface{}{
			"title":       post.Title,
			"content":     post.Content,
			"category_id": post.CategoryID,
			"status":      post.Status,
			"user_id":     post.UserID,
		})

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	err := r.db.WithContext(ctx).Preload("User").Preload("Category").First(&existing, post.ID).Error
	if err != nil {
		return nil, err
	}
	updated := postDomain.Post{
		ID:         uint(existing.ID),
		UserID:     userDomain.ID(existing.UserID),
		CategoryID: existing.CategoryID,
		Title:      existing.Title,
		Content:    existing.Content,
		Status:     postDomain.Status(existing.Status),
	}
	return &updated, nil
}

func (r *postRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *postRepository) GetMyPostsByUserID(ctx context.Context, userID userDomain.ID) ([]postDomain.Post, error) {

    var posts []Post

    err := r.db.
        WithContext(ctx).
        Where("user_id = ?", uint(userID)).
        Find(&posts).Error

    if err != nil {
        return nil, err
    }

    result := make([]postDomain.Post, 0, len(posts))

    for _, p := range posts {
        result = append(result, postDomain.Post{
            ID:         uint(p.ID),
            UserID:     userDomain.ID(p.UserID),
            CategoryID: p.CategoryID,
            Title:      p.Title,
            Content:    p.Content,
            Status:     postDomain.Status(p.Status),
        })
    }

    return result, nil
}

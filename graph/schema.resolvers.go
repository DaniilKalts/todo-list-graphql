package graph

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DaniilKalts/todo-list-graphql/graph/model"
	"github.com/DaniilKalts/todo-list-graphql/internal/domain"
)

// CreateCategory creates a new Category
func (r *mutationResolver) CreateCategory(
	ctx context.Context, name string, description string,
) (*model.Category, error) {
	cat := &domain.Category{
		Name:        name,
		Description: description,
	}
	if err := r.DB.Create(cat).Error; err != nil {
		return nil, err
	}
	return &model.Category{
		ID:          fmt.Sprint(cat.ID),
		Name:        cat.Name,
		Description: cat.Description,
		CreatedAt:   cat.CreatedAt,
		UpdatedAt:   cat.UpdatedAt,
	}, nil
}

// CreateTodo creates a new Todo under a Category
func (r *mutationResolver) CreateTodo(
	ctx context.Context, name string, description string, categoryID string,
) (*model.Todo, error) {
	cid, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		return nil, err
	}
	t := &domain.Todo{
		Name:        name,
		Description: description,
		CategoryID:  uint(cid),
	}
	if err := r.DB.Create(t).Error; err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:          fmt.Sprint(t.ID),
		Name:        t.Name,
		Description: t.Description,
		Done:        t.Done,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}, nil
}

// AddImage associates an Image to a Todo
func (r *mutationResolver) AddImage(
	ctx context.Context, todoID string, url string,
) (*model.Image, error) {
	tid, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		return nil, err
	}
	img := &domain.Image{
		URL:    url,
		TodoID: ptrUint(uint(tid)),
	}
	if err := r.DB.Create(img).Error; err != nil {
		return nil, err
	}
	return &model.Image{
		ID:        fmt.Sprint(img.ID),
		URL:       img.URL,
		CreatedAt: img.CreatedAt,
		UpdatedAt: img.UpdatedAt,
	}, nil
}

// ToggleTodoDone flips the done flag on a Todo
func (r *mutationResolver) ToggleTodoDone(
	ctx context.Context, id string,
) (*model.Todo, error) {
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var t domain.Todo
	if err := r.DB.First(&t, uint(tid)).Error; err != nil {
		return nil, err
	}
	t.Done = !t.Done
	if err := r.DB.Save(&t).Error; err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:          fmt.Sprint(t.ID),
		Name:        t.Name,
		Description: t.Description,
		Done:        t.Done,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}, nil
}

// DeleteTodo removes a Todo (soft-delete)
func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (
	bool, error,
) {
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return false, err
	}
	if err := r.DB.Delete(&domain.Todo{}, uint(tid)).Error; err != nil {
		return false, err
	}
	return true, nil
}

// Todos returns all Todos with Category and Images preloaded
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var list []domain.Todo
	if err := r.DB.Preload("Category").Preload("Images").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*model.Todo, len(list))
	for i, t := range list {
		out[i] = &model.Todo{
			ID:          fmt.Sprint(t.ID),
			Name:        t.Name,
			Description: t.Description,
			Done:        t.Done,
			Category: &model.Category{
				ID:   fmt.Sprint(t.Category.ID),
				Name: t.Category.Name,
			},
			Images:    mapImages(t.Images),
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		}
	}
	return out, nil
}

// Todo returns a single Todo by ID
func (r *queryResolver) Todo(ctx context.Context, id string) (
	*model.Todo, error,
) {
	tid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var t domain.Todo
	if err := r.DB.Preload("Category").Preload("Images").First(
		&t, uint(tid),
	).Error; err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:          fmt.Sprint(t.ID),
		Name:        t.Name,
		Description: t.Description,
		Done:        t.Done,
		Category: &model.Category{
			ID:   fmt.Sprint(t.Category.ID),
			Name: t.Category.Name,
		},
		Images:    mapImages(t.Images),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

// Categories returns all Categories
func (r *queryResolver) Categories(ctx context.Context) (
	[]*model.Category, error,
) {
	var cats []domain.Category
	if err := r.DB.Find(&cats).Error; err != nil {
		return nil, err
	}
	out := make([]*model.Category, len(cats))
	for i, c := range cats {
		out[i] = &model.Category{
			ID:          fmt.Sprint(c.ID),
			Name:        c.Name,
			Description: c.Description,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		}
	}
	return out, nil
}

// Category returns a single Category by ID with its Todos
func (r *queryResolver) Category(
	ctx context.Context, id string,
) (*model.Category, error) {
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var c domain.Category
	if err := r.DB.Preload("Todos").First(&c, uint(cid)).Error; err != nil {
		return nil, err
	}
	todos := make([]*model.Todo, len(c.Todos))
	for i, t := range c.Todos {
		todos[i] = &model.Todo{ID: fmt.Sprint(t.ID), Name: t.Name}
	}
	return &model.Category{
		ID:          fmt.Sprint(c.ID),
		Name:        c.Name,
		Description: c.Description,
		Todos:       todos,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}, nil
}

// mapImages converts domain.Images to []*model.Image
func mapImages(imgs []domain.Image) []*model.Image {
	out := make([]*model.Image, len(imgs))
	for i, im := range imgs {
		out[i] = &model.Image{
			ID:        fmt.Sprint(im.ID),
			URL:       im.URL,
			CreatedAt: im.CreatedAt,
			UpdatedAt: im.UpdatedAt,
		}
	}
	return out
}

// helper ptrUint returns *uint
func ptrUint(u uint) *uint { return &u }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/larien/plus/fundamentos/graphql/graph/generated"
	"github.com/larien/plus/fundamentos/graphql/graph/model"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category := model.Category{
		ID:          "categoriazita",
		Name:        input.Name,
		Description: &input.Description,
	}
	r.Categories = append(r.Categories, &category)
	return &category, nil
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	category := r.Categories[0]

	course := model.Course{
		ID:          "cursozito",
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}
	r.Courses = append(r.Courses, &course)

	return &course, nil
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	category := r.Categories[0]

	chapter := model.Chapter{
		ID:       "capitulozito",
		Name:     input.Name,
		Category: category,
	}
	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

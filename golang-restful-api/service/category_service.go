package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, category web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, category web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)

	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}

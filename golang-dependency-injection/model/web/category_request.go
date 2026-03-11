package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,max=200,min=3"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required,numeric"`
	Name string `json:"name" validate:"required,max=200,min=3"`
}

package repository

import "4-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

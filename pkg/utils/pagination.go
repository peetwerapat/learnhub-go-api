package utils

import (
	"fmt"
	"strings"

	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller/dto"
	"gorm.io/gorm"
)

func Paginate[T any](
	db *gorm.DB,
	out *[]T,
	q dto.PaginationQuery,
	searchableFields []string,
	sortableFields []string,
) (totalItems int64, totalPages int, err error) {

	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit <= 0 {
		q.Limit = 10
	}
	offset := (q.Page - 1) * q.Limit

	countQuery := db.Session(&gorm.Session{})
	countQuery = countQuery.Model(new(T))

	if q.Search != "" && len(searchableFields) > 0 {
		searchPattern := "%" + q.Search + "%"
		for i, field := range searchableFields {
			if i == 0 {
				countQuery = countQuery.Where(fmt.Sprintf("%s ILIKE ?", field), searchPattern)
			} else {
				countQuery = countQuery.Or(fmt.Sprintf("%s ILIKE ?", field), searchPattern)
			}
		}
	}

	err = countQuery.Count(&totalItems).Error
	if err != nil {
		return 0, 0, err
	}

	dataQuery := db.Session(&gorm.Session{})
	dataQuery = dataQuery.Model(new(T))

	if q.Search != "" && len(searchableFields) > 0 {
		searchPattern := "%" + q.Search + "%"
		for i, field := range searchableFields {
			if i == 0 {
				dataQuery = dataQuery.Where(fmt.Sprintf("%s ILIKE ?", field), searchPattern)
			} else {
				dataQuery = dataQuery.Or(fmt.Sprintf("%s ILIKE ?", field), searchPattern)
			}
		}
	}

	sortField := "id"
	for _, allowed := range sortableFields {
		if q.Sort == allowed {
			sortField = q.Sort
			break
		}
	}

	order := "asc"
	if strings.ToLower(q.Order) == "desc" {
		order = "desc"
	}

	err = dataQuery.Order(fmt.Sprintf("%s %s", sortField, order)).
		Limit(q.Limit).
		Offset(offset).
		Find(out).Error
	if err != nil {
		return 0, 0, err
	}

	totalPages = int((totalItems + int64(q.Limit) - 1) / int64(q.Limit))
	return totalItems, totalPages, nil
}

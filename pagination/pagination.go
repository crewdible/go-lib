package pagination

import (
	"math"

	"gorm.io/gorm"
)

type (
	Pagination struct {
		Limit      int         `json:"limit,omitempty" query:"limit"`
		Page       int         `json:"page,omitempty" query:"page"`
		Sort       string      `json:"sort,omitempty" query:"sort"`
		TotalRows  int64       `json:"total_rows"`
		TotalPages int         `json:"total_pages"`
		Rows       interface{} `json:"rows"`
	}
)

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// Validate the string WARNING!! SQL Injection
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

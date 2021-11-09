package storage

import (
	"github.com/jinzhu/gorm"
)

//
//type Pagination struct {
//	Limit      int         `json:"limit,omitempty;query:limit"`
//	Page       int         `json:"page,omitempty;query:page"`
//	Sort       string      `json:"sort,omitempty;query:sort"`
//	TotalRows  int64       `json:"total_rows"`
//	TotalPages int         `json:"total_pages"`
//	Rows       interface{} `json:"rows"`
//}

//Paginate for pagination by page
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

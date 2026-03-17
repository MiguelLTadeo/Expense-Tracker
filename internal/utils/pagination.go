package utils

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type PageInfo struct {
	Page     int
	PageSize int
}

func Paginate(r *http.Request, info *PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		info.Page = page
		info.PageSize = pageSize

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

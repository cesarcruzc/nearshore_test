package meta

import (
	"os"
	"strconv"
)

type (
	Meta struct {
		TotalCount int64 `json:"total_count"`
		Page       int64 `json:"page"`
		PerPage    int64 `json:"per_page"`
		PageCount  int64 `json:"page_count"`
	}
)

func New(page, perPage, totalCount int64) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.ParseInt(os.Getenv("PAGINATOR_LIMIT_PAGE"), 10, 64)

		if err != nil {
			return nil, err
		}
	}

	var pageCount int64 = 0
	if totalCount >= 0 {
		pageCount = (totalCount + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page <= 0 {
		page = 1
	}

	return &Meta{
		TotalCount: totalCount,
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
	}, nil
}

func (meta *Meta) Offset() int64 {
	return (meta.Page - 1) * meta.PerPage
}

func (meta *Meta) Limit() int64 {
	return meta.PerPage
}

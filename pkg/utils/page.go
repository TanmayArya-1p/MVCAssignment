package utils

import (
	"inorder/pkg/types"
	"net/http"
	"strconv"
)

func Paginate(r *http.Request) (types.Page, error) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	var pg types.Page
	var err error

	if limit == "" {
		pg.Limit = 10
	} else {
		pg.Limit, err = strconv.Atoi(limit)
		if err != nil || pg.Limit <= 0 {
			return pg, ErrInvalidLimit
		}
	}

	if offset == "" {
		pg.Offset = types.DefaultOffset
	} else {
		pg.Offset, err = strconv.Atoi(offset)
		if err != nil || pg.Offset < 0 {
			return pg, ErrInvalidOffset
		}
	}
	return pg, nil
}

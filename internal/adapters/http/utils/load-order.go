package httpUtils

import (
	"net/http"
	"strconv"
	"tom/internal/templates"
)

func GetLoadOrder(r *http.Request) templates.LoadOrder {
    skipString := r.URL.Query().Get("skip")
    skip, err := strconv.Atoi(skipString)
    if err != nil {
        skip = 0
    }
    limitString := r.URL.Query().Get("limit")
    limit, err := strconv.Atoi(limitString)
    if err != nil {
        limit = 10
    }

    return templates.LoadOrder{
        Skip: skip,
        Limit: limit,
    }
}

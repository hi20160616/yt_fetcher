package data

import (
	"github.com/hi20160616/yt_fetcher/internal/biz"
)

var _ biz.FetcherRepo = new(fetcherRepo)

type fetcherRepo struct {
}

func NewFetcherRepo() biz.FetcherRepo {
	return &fetcherRepo{}
}

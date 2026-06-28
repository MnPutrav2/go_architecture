package pagination

import (
	"fmt"
	"net/http"
	"strconv"
)

type PaginationParam struct {
	Page    int
	Size    int
	Keyword *string
	Sort    *string
	Order   *string
	Custom  []string
	Offset  int
}

type Param struct {
	param   *http.Request
	keyword *string
	sort    *string
	order   *string
	custom  []string
}

func BuildPagination(r *http.Request) *Param {
	return &Param{
		param: r,
	}
}

func (p *Param) Keyword() *Param {

	url := p.param.URL.Query().Get("keyword")

	if url == "" {
		return p
	}

	p.keyword = &url
	return p
}

func (p *Param) Sort() *Param {

	url := p.param.URL.Query().Get("sort")

	if url == "" {
		return p
	}

	p.sort = &url
	return p
}

func (p *Param) Order() *Param {

	url := p.param.URL.Query().Get("order")

	if url == "" {
		return p
	}

	p.order = &url
	return p
}

func (p *Param) Custom(pr string) *Param {
	url := p.param.URL.Query().Get(pr)

	if url == "" {
		return p
	}

	p.custom = append(p.custom, url)
	return p
}

func (p *Param) Result() (PaginationParam, error) {
	var Pagination PaginationParam

	page, err := strconv.Atoi(p.param.URL.Query().Get("page"))
	if err != nil {
		return PaginationParam{}, fmt.Errorf("Invalid parameter")
	}

	size, err := strconv.Atoi(p.param.URL.Query().Get("size"))
	if err != nil {
		return PaginationParam{}, fmt.Errorf("Invalid parameter")
	}

	Pagination.Page = page
	Pagination.Size = size
	Pagination.Keyword = p.keyword
	Pagination.Sort = p.sort
	Pagination.Order = p.order
	Pagination.Custom = p.custom
	Pagination.Offset = offset(page, size)

	return Pagination, nil
}

func offset(page, size int) int {
	offsite := page * size
	return offsite
}

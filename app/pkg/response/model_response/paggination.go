package modelresponse

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/MnPutrav2/go_architecture/app/model"
	logging "github.com/MnPutrav2/go_architecture/app/pkg/log"
	"github.com/MnPutrav2/go_architecture/app/pkg/pagination"
)

func Pagination(body any, page, size, total int, keyword string, ty, log string, w http.ResponseWriter, r *http.Request, more ...model.NextLink) {
	var next []string
	var nextString *string

	for _, item := range more {
		next = append(next, fmt.Sprintf("%s=%s", item.Param, item.Value))
	}

	if len(next) == 0 {
		nextString = nil
	} else {
		st := strings.Join(next, "&")
		nextString = &st
	}

	previousLink, nextLink := pagination.Link(page, size, total, keyword, nextString)
	pg := float64(total) / float64(size)

	res, _ := json.Marshal(model.PaginationResponse{
		Result: body,
		Meta: model.PaginationMeta{
			TotalData: total,
			TotalPage: int(math.Ceil(pg)),
			Page:      page,
			Size:      size,
			Previous:  previousLink,
			Next:      nextLink,
		},
	})
	logging.Log(log, ty, r)
	w.WriteHeader(200)
	w.Write(res)
}

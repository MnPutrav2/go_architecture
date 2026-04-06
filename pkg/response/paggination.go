package response

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/MnPutrav2/go_architecture/internal/model"
	logging "github.com/MnPutrav2/go_architecture/pkg/log"
	"github.com/MnPutrav2/go_architecture/pkg/pagination"
)

func Pagination(body any, page, size, total int, keyword string, ty, log string, w http.ResponseWriter, r *http.Request) {
	previousLink, nextLink := pagination.Link(page, size, total, keyword)
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

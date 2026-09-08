package migration

import (
	"reflect"
	"strings"

	"github.com/MnPutrav2/go_architecture/app/pkg/query"
	"github.com/MnPutrav2/go_architecture/app/util"
)

func Seed(keyword string) {
	for _, m := range util.Models {
		if matchModel(m, keyword) {
			query.CreateSeed(m)
		}
	}
}

func matchModel(model any, keyword string) bool {
	t := reflect.TypeOf(model)

	name := strings.ToLower(t.Name())
	keyword = strings.ToLower(keyword)

	name = strings.TrimSuffix(name, "s")

	return strings.Contains(name, keyword)
}

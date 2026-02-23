package param

import (
	"fmt"
	"net/http"
	"time"
)

func ParseToDate(param []string, r *http.Request) ([]time.Time, error) {
	var params []time.Time

	url := r.URL.Query()
	for _, p := range param {
		paramParse := url.Get(p)
		if paramParse == "" {
			return nil, fmt.Errorf("empty param %s", p)
		}

		date, err := time.Parse("2006-01-02", paramParse)
		if err != nil {
			return nil, fmt.Errorf("error parse to date : %s", p)
		}

		params = append(params, date)
	}

	return params, nil
}

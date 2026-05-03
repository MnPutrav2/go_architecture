package prefix

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func String(param string, r *http.Request) string {
	idStr := r.PathValue(param)
	return idStr
}

func UUID(param string, r *http.Request) (uuid.UUID, error) {
	idStr := r.PathValue(param)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Invalid parameter")
	}

	return id, nil
}

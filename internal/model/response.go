package model

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseBody struct {
	Status   string `json:"status"`
	Response any    `json:"result"`
}

type PaginationResponse struct {
	Result any            `json:"result"`
	Meta   PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	TotalData int    `json:"total_data"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Previous  string `json:"previous"`
	Next      string `json:"next"`
}

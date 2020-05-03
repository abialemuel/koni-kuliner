package response

type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
	Meta MetaInfo    `json:"meta"`
}

type SuccessResponseWithMeta struct {
	Data interface{}                `json:"data,omitempty"`
	Meta MetaInfoWithOffsetAndLimit `json:"meta"`
}

type ErrorResponse struct {
	Errors []ErrorInfo `json:"errors,omitempty"`
	Meta   MetaInfo    `json:"meta"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Field   string `json:"field,omitempty"`
}

type MetaInfo struct {
	HttpStatus int `json:"http_status"`
}

type MetaInfoWithOffsetAndLimit struct {
	HttpStatus int `json:"http_status"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
}

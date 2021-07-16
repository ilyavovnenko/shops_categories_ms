package response

type Meta struct {
	Page            int    `json:"page,omitempty"`
	PerPage         int    `json:"per_page,omitempty"`
	Self            string `json:"self"`
	Next            string `json:"next"`
	Previous        string `json:"previous"`
	TotalPages      int    `json:"total_pages,omitempty"`
	TotalItemsCount int64  `json:"total_items_count,omitempty"`
}

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

type AppError struct {
	Code    string
	Message string
}

type Response struct {
	Data       interface{}       `json:"data,omitempty"`
	Errors     []AppError        `json:"errors,omitempty"`
	Message    string            `json:"message,omitempty"`
	Meta       Meta              `json:"meta"`
	Validation []ValidationError `json:"validation_errors,omitempty"`
}

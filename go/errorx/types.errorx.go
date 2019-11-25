package errorx

type Error *CustomError

type CustomError struct {
	Code        string        `json:"id"`
	Status      string        `json:"status"`
	Title       string        `json:"title"`
	TitleParams []interface{} `json:"title_params"`
	Meta        interface{}   `json:"meta"`
}

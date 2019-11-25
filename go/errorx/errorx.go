package errorx

import (
	"encoding/json"
	"errors"
)

var title = map[string]string{}

// InitTitleMap for errorx title
func InitTitleMap(titleMap map[string]string) {
	title = titleMap
}

// New errorx
func New(code string) *CustomError {
	return &CustomError{
		Code:  code,
		Title: Title(code),
	}
}

// NewArr of errorx return multiple error message on json format
func NewArr(arrErr []*CustomError) error {
	newTitle, _ := json.Marshal(arrErr)
	return errors.New(string(newTitle))
}

// Title returns a text for the Error code.
func Title(code string) string {
	return title[code]
}

func (e *CustomError) Error() string {
	return "Errorx Code: " + e.Code
}

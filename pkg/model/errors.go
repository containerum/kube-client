package model

import (
	"fmt"
)

type ResourceError struct {
	StatusCode int    `json:"status_code"`
	ErrorMsg   string `json:"error"`
}

func (err *ResourceError) Error() string {
	return fmt.Sprintf("status %d: %s ", err.StatusCode, err.ErrorMsg)
}

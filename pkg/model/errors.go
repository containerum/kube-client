package model

type ResourceError struct {
	ErrorMsg string `json:"error"`
}

func (err *ResourceError) Error() string {
	return err.ErrorMsg
}

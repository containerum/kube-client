package cherry

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	HeadDelimiter    = ": "
	DetailsDelimiter = "; "
)

type Err struct {
	Message string   `json:"message"`
	ID      uint64   `json:"id"`
	Details []string `json:"details"`
}

func NewErr(msg string, ID uint64) *Err {
	return &Err{
		Message: msg,
		ID:      ID,
	}
}

func (err *Err) Error() string {
	buf := bytes.NewBufferString(err.Message +
		strconv.FormatUint(err.ID, 10) +
		HeadDelimiter)
	for _, msg := range err.Details {
		buf.WriteString(DetailsDelimiter + msg)
	}
	return buf.String()
}

func (err *Err) AddDetails(details ...string) *Err {
	err.Details = append(err.Details, details...)
	return err
}

func (err *Err) AddDetailF(formatS string, args ...interface{}) *Err {
	return err.AddDetails(fmt.Sprintf(formatS, args...))
}

func (err *Err) AddDetailsErr(details ...error) *Err {
	for _, detail := range details {
		err.AddDetails(detail.Error())
	}
	return err
}

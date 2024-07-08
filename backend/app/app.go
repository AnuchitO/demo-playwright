package app

import "time"

type status string

const (
	Success status = "success"
	Fail    status = "error"
)

type Paging struct {
	Page       uint `json:"page"`
	Total      uint `json:"total"`
	TotalPages uint `json:"totalPages"`
	TotalCount uint `json:"totalCount"`
	NextPage   uint `json:"nextPage"`
	PrevPage   uint `json:"prevPage"`
}

type Response struct {
	Status  status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Pagination struct {
	Response
	Paging Paging `json:"paging"`
}

type ErrorField struct {
	Field string `json:"field"`
	Value any    `json:"value"`
	Tag   string `json:"tag"`
}

type Error Response

func (err *Error) Error() string {
	return err.Message
}

type Clock interface {
	Now() time.Time
	After(d time.Duration) <-chan time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time                         { return time.Now() }
func (RealClock) After(d time.Duration) <-chan time.Time { return time.After(d) }

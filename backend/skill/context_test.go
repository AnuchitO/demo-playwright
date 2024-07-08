package skill

import (
	"encoding/json"
	"net/http"

	"demo/app"
)

type context struct {
	app.Context
	payload    []byte
	err        error
	ok         any
	status     int
	page       uint
	perPage    uint
	totalItems uint
	q          map[string]string
	p          map[string]string
}

func (c *context) OK(data interface{}) {
	c.status = http.StatusOK
	c.ok = data
}

func (c *context) CREATED(data interface{}) {
	c.status = http.StatusCreated
	c.ok = data
}

func (c *context) UPDATED() {
	c.status = http.StatusNoContent
}

func (c *context) Bind(v any) error {
	return json.Unmarshal(c.payload, v)
}

func (c *context) Query(key string) string {
	return c.q[key]
}

func (c *context) Param(key string) string {
	return c.p[key]
}

func (c *context) OkWithPagination(data interface{}, page, perPage, totalItems uint) {
	c.status = http.StatusOK
	c.ok = data
	c.page = page
	c.perPage = perPage
	c.totalItems = totalItems
}

func (c *context) InternalServerError(err error) {
	c.status = http.StatusInternalServerError
	c.err = err
}

func (c *context) BadRequest(err error) {
	c.status = http.StatusBadRequest
	c.err = err
}

const skillTable = `CREATE TABLE skill (
			key TEXT PRIMARY KEY,
			name TEXT NOT NULL DEFAULT '',
			description TEXT NOT NULL DEFAULT '',
			logo TEXT NOT NULL DEFAULT '',
			levels JSONB NOT NULL DEFAULT '[]',
			tags TEXT[] NOT NULL DEFAULT '{}'
		);`

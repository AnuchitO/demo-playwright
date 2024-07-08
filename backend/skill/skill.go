package skill

import (
	"encoding/json"

	"github.com/lib/pq"
)

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Levels      []Level  `json:"levels"`
	Tags        []string `json:"tags"`
}

type Level struct {
	Key          string   `json:"key"`
	Name         string   `json:"name"`
	Brief        string   `json:"brief"`
	Descriptions []string `json:"descriptions"`
	Level        int      `json:"level"`
}

func (s Skill) toRecord() record {
	return record{
		Key:         s.Key,
		Name:        s.Name,
		Description: s.Description,
		Logo:        s.Logo,
		Levels:      s.levels(),
		Tags:        s.tags(),
	}
}

func (s Skill) levels() []byte {
	return levels(s.Levels)
}

func levels(l []Level) []byte {
	if len(l) == 0 {
		return []byte(`[]`)
	}

	levels, err := json.Marshal(l)
	if err != nil {
		return []byte(`[]`)
	}

	return levels
}

func (s Skill) tags() pq.StringArray {
	return tags(s.Tags)
}

func tags(tg []string) pq.StringArray {
	if len(tg) == 0 {
		return pq.StringArray{}
	}

	return pq.StringArray(tg)
}

type Handler struct {
	skill storager
}

func NewHandler(rec storager) *Handler {
	return newHandler(rec)
}

func newHandler(rec storager) *Handler {
	return &Handler{
		skill: rec,
	}
}

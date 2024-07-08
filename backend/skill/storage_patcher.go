package skill

import "github.com/lib/pq"

type patcher interface {
	PatchName(key string, name string) error
	PatchDescription(key string, description string) error
	PatchLogo(key string, logo string) error
	PatchTags(key string, tags pq.StringArray) error
	PatchLevels(key string, levels []byte) error
}

// PatchName updates the name of a skill.
func (s storage) PatchName(key string, name string) error {
	_, err := s.conn.Exec("UPDATE skill SET name = $1 WHERE key = $2", name, key)
	return err
}

// PatchDescription updates the description of a skill.
func (s storage) PatchDescription(key string, description string) error {
	_, err := s.conn.Exec("UPDATE skill SET description = $1 WHERE key = $2", description, key)
	return err
}

// PatchLogo updates the logo of a skill.
func (s storage) PatchLogo(key string, logo string) error {
	_, err := s.conn.Exec("UPDATE skill SET logo = $1 WHERE key = $2", logo, key)
	return err
}

// PatchTags updates the tags of a skill.
func (s storage) PatchTags(key string, tags pq.StringArray) error {
	_, err := s.conn.Exec("UPDATE skill SET tags = $1 WHERE key = $2", tags, key)
	return err
}

// PatchLevels updates the levels of a skill.
func (s storage) PatchLevels(key string, levels []byte) error {
	_, err := s.conn.Exec("UPDATE skill SET levels = $1 WHERE key = $2", levels, key)
	return err
}

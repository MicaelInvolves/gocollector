package access

import "time"

type Access struct {
	Id       string
	ClientId string
	Path     string
	Date     time.Time
}

type Gateway interface {
	Save(access *Access) error
}

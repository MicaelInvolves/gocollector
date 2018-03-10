package access

type AccessGateway interface {
	Save(access *Access) error
}

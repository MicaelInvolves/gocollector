package subscriber

type Subscriber struct {
	ClientId string
	Name     string
	Email    string
}

type Gateway interface {
	Save(subscriber *Subscriber) error
}

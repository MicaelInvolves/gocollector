package subscriber

type Subscriber struct {
	ClientId string
	Name     string
	Email    string
}

type SubscriberGateway interface {
	Save(subscriber *Subscriber) error
}

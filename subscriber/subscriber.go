package subscriber

type Subscriber struct {
	ClientId string
	Name     string
	Email    string
}

type SubscribersAccessData struct {
	Subscriber  *Subscriber
	AccessCount int
	AccessPaths []string
}

type Gateway interface {
	Save(subscriber *Subscriber) error
	All() ([]*SubscribersAccessData, error)
}

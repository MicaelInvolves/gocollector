package subscriber

type ListSubscribersAccessDataUseCase struct {
	Gateway Gateway
}

func (this *ListSubscribersAccessDataUseCase) List() (*ListSubscribersAccessDataResponse, error) {
	subscribers, _ := this.Gateway.All()
	return &ListSubscribersAccessDataResponse{
		Subscribers: subscribers,
	}, nil
}

type ListSubscribersAccessDataResponse struct {
	Subscribers []*SubscribersAccessData
}

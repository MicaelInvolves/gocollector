package subscriber

type ListSubscribersAccessDataUseCase struct {
	Gateway Gateway
}

func (this *ListSubscribersAccessDataUseCase) List() (*ListSubscribersAccessDataResponse, error) {
	subscribers, err := this.Gateway.All()
	if err != nil {
		return nil, err
	}
	return &ListSubscribersAccessDataResponse{
		SubscribersAccessData: subscribers,
	}, nil
}

type ListSubscribersAccessDataResponse struct {
	SubscribersAccessData []*SubscribersAccessData
}

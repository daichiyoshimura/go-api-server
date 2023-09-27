package getOneService

type GreetingService struct {
	repo IGreetingRepository
}

func NewGreetingService(repo IGreetingRepository) *GreetingService {
	return &GreetingService{
		repo: repo,
	}
}

func (s *GreetingService) GetOne(in *Input) (*Output, error) {
	return nil, nil
}

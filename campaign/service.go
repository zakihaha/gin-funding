package campaign

type Service interface {
	GetAll() ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Campaign, error) {
	campaigns, err := s.repository.GetAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

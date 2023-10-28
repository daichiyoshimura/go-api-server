package account

import (
	"github.com/cockroachdb/errors"
)

const (
	errMsgAccountService string = "account service: %w"
)

type AccountService struct {
	repo IAccountRepository
}

func NewAccountService(repo IAccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) Get(id AccountID) (*AccountEntity, error) {
	ac, err := s.repo.Get(id)
	if err != nil {
		return nil, errors.Errorf(errMsgAccountService, err)
	}

	return NewAccountEntity(ac), nil
}

func (s *AccountService) Create(in *AccountDTO) (*AccountEntity, error) {
	ac, err := s.repo.Create(in)
	if err != nil {
		return nil, errors.Errorf(errMsgAccountService, err)
	}

	return NewAccountEntity(ac), nil
}

func (s *AccountService) Update(in *AccountDTO) (*AccountEntity, error) {
	ac, err := s.repo.Update(in)
	if err != nil {
		return nil, errors.Errorf(errMsgAccountService, err)
	}

	return NewAccountEntity(ac), nil
}

func (s *AccountService) Delete(id AccountID) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.Errorf(errMsgAccountService, err)
	}

	return nil
}

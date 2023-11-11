package domain

import (
	"github.com/cockroachdb/errors"
)

type AccountService struct {
	repo iAccountRepository
}

func NewAccountService(repo iAccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) Get(id int64) (*AccountEntity, error) {
	acDTO, err := s.repo.Get(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ac, err := NewAccountEntity(acDTO)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ac, nil
}

func (s *AccountService) Create(udto *AccountUnspecifiedDTO) (*AccountEntity, error) {
	acDTO, err := s.repo.Create(udto)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ac, err := NewAccountEntity(acDTO)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ac, nil
}

func (s *AccountService) Update(dto *AccountDTO) (*AccountEntity, error) {
	acDTO, err := s.repo.Update(dto)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ac, err := NewAccountEntity(acDTO)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ac, nil
}

func (s *AccountService) Delete(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

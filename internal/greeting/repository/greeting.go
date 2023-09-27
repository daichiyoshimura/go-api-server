package repository

import (
	"github.com/uptrace/bun"
)

type GreetingRepository struct {
	conn bun.IDB
}

// Find implements getOneService.IGreetingRepository.
func (*GreetingRepository) Find() {
	panic("unimplemented")
}

func NewGreetingRepository(conn bun.IDB) *GreetingRepository {
	return &GreetingRepository{
		conn: conn,
	}
}

type GreetingGetOneInput struct{}

func (r *GreetingRepository) GetOne(in GreetingGetOneInput) *GreetingModel {
	return &GreetingModel{}
}

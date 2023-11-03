package domain_test

import (
	"awsomeapp/internal/module/account/internal/domain"
	"awsomeapp/internal/module/account/internal/domain/mock"

	"reflect"
	"testing"

	"github.com/cockroachdb/errors"

	"go.uber.org/mock/gomock"
)

func TestAccountService_Get(t *testing.T) {
	id := domain.AccountID(10)

	entity, _ := domain.NewAccountEntity(&domain.AccountDTO{
		ID:   &id,
		Name: "John Smith",
	})
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repo := mock.NewMockIAccountRepository(ctrl)

	type fields struct {
		repo domain.IAccountRepository
	}

	type args struct {
		id domain.AccountID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountEntity
		wantErr bool
	}{
		{
			name: "get an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Get(id).Return(entity.DTO(), nil)

					return repo
				}(repo),
			},
			args: args{
				id: id,
			},
			want:    entity,
			wantErr: false,
		},
		{
			name: "failed to get an acount",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Get(id).Return(nil, errors.Newf("mock err"))

					return repo
				}(repo),
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewAccountService(repo)
			got, err := s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Get() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Create(t *testing.T) {
	name := domain.AccountName("John Smith")
	dto := &domain.AccountDTO{
		Name: name,
	}
	id := domain.AccountID(10)
	entity, _ := domain.NewAccountEntity(&domain.AccountDTO{
		ID:   &id,
		Name: name,
	})
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repo := mock.NewMockIAccountRepository(ctrl)

	type fields struct {
		repo domain.IAccountRepository
	}

	type args struct {
		in *domain.AccountDTO
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountEntity
		wantErr bool
	}{
		{
			name: "create an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Create(dto).Return(entity.DTO(), nil)

					return repo
				}(repo),
			},
			args: args{
				in: dto,
			},
			want:    entity,
			wantErr: false,
		},
		{
			name: "failed to create an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Create(dto).Return(nil, errors.Newf("mock err"))

					return repo
				}(repo),
			},
			args: args{
				in: dto,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewAccountService(repo)
			got, err := s.Create(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Create() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Update(t *testing.T) {
	name := domain.AccountName("John Smith")
	dto := &domain.AccountDTO{
		Name: name,
	}
	id := domain.AccountID(10)
	entity, _ := domain.NewAccountEntity(&domain.AccountDTO{
		ID:   &id,
		Name: name,
	})
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repo := mock.NewMockIAccountRepository(ctrl)

	type fields struct {
		repo domain.IAccountRepository
	}

	type args struct {
		in *domain.AccountDTO
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountEntity
		wantErr bool
	}{
		{
			name: "update an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Update(dto).Return(entity.DTO(), nil)

					return repo
				}(repo),
			},
			args: args{
				in: dto,
			},
			want:    entity,
			wantErr: false,
		},
		{
			name: "failed to update an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Update(dto).Return(nil, errors.Newf("mock err"))

					return repo
				}(repo),
			},
			args: args{
				in: dto,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewAccountService(repo)
			got, err := s.Update(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Update() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountService.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Delete(t *testing.T) {
	id := domain.AccountID(10)

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repo := mock.NewMockIAccountRepository(ctrl)

	type fields struct {
		repo domain.IAccountRepository
	}

	type args struct {
		id domain.AccountID
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Delete(id).Return(nil)

					return repo
				}(repo),
			},
			args: args{
				id: id,
			},
			wantErr: false,
		},
		{
			name: "failed to delete an account",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Delete(id).Return(errors.Newf("mock err"))

					return repo
				}(repo),
			},
			args: args{
				id: id,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewAccountService(repo)
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package domain

import (
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestNewAccountService(t *testing.T) {
	repo := NewMockiAccountRepository(gomock.NewController(t))

	type args struct {
		repo iAccountRepository
	}
	tests := []struct {
		name string
		args args
		want *AccountService
	}{
		{
			name: "define",
			args: args{
				repo: repo,
			},
			want: &AccountService{
				repo: repo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountService_Get(t *testing.T) {
	repo := NewMockiAccountRepository(gomock.NewController(t))
	var id int64 = 1
	name := "JohnSmith"
	dto := &AccountDTO{
		ID:   id,
		Name: name,
	}
	invalidDTO := &AccountDTO{
		ID:   id,
		Name: "",
	}
	account, _ := NewAccountEntity(dto)

	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountEntity
		wantErr bool
	}{
		{
			name: "DB error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Get(gomock.Eq(id)).Return(nil, errors.Newf("db error"))
					return repo
				}(),
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Entity error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Get(gomock.Eq(id)).Return(invalidDTO, nil)
					return repo
				}(),
			},
			args: args{
				id: id,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Get(gomock.Eq(id)).Return(dto, nil)
					return repo
				}(),
			},
			args: args{
				id: id,
			},
			want:    account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountService{
				repo: tt.fields.repo,
			}
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
	repo := NewMockiAccountRepository(gomock.NewController(t))
	var id int64 = 1
	name := "JohnSmith"
	udto := &AccountUnspecifiedDTO{
		Name: name,
	}
	dto := &AccountDTO{
		ID:   id,
		Name: name,
	}
	invalidDTO := &AccountDTO{
		ID:   id,
		Name: "",
	}
	account, _ := NewAccountEntity(dto)

	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		dto *AccountUnspecifiedDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountEntity
		wantErr bool
	}{
		{
			name: "DB error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Create(gomock.Eq(udto)).Return(nil, errors.Newf("DB error"))
					return repo
				}(),
			},
			args: args{
				dto: udto,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Entity error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Create(gomock.Eq(udto)).Return(invalidDTO, nil)
					return repo
				}(),
			},
			args: args{
				dto: udto,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Create(gomock.Eq(udto)).Return(dto, nil)
					return repo
				}(),
			},
			args: args{
				dto: udto,
			},
			want:    account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountService{
				repo: tt.fields.repo,
			}
			got, err := s.Create(tt.args.dto)
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
	repo := NewMockiAccountRepository(gomock.NewController(t))
	var id int64 = 1
	name := "JohnSmith"
	dto := &AccountDTO{
		ID:   id,
		Name: name,
	}
	invalidDTO := &AccountDTO{
		ID:   id,
		Name: "",
	}
	account, _ := NewAccountEntity(dto)

	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		dto *AccountDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountEntity
		wantErr bool
	}{
		{
			name: "DB error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Update(gomock.Eq(dto)).Return(nil, errors.Newf("DB error"))
					return repo
				}(),
			},
			args: args{
				dto: dto,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Entity error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Update(gomock.Eq(dto)).Return(invalidDTO, nil)
					return repo
				}(),
			},
			args: args{
				dto: dto,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Update(gomock.Eq(dto)).Return(dto, nil)
					return repo
				}(),
			},
			args: args{
				dto: dto,
			},
			want:    account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountService{
				repo: tt.fields.repo,
			}
			got, err := s.Update(tt.args.dto)
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
	repo := NewMockiAccountRepository(gomock.NewController(t))
	var id int64 = 1

	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "DB error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Delete(gomock.Eq(id)).Return(errors.Newf("DB error"))
					return repo
				}(),
			},
			args: args{
				id: id,
			},
			wantErr: true,
		},
		{
			name: "DB error",
			fields: fields{
				repo: func() *MockiAccountRepository {
					repo.EXPECT().Delete(gomock.Eq(id)).Return(nil)
					return repo
				}(),
			},
			args: args{
				id: id,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AccountService{
				repo: tt.fields.repo,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

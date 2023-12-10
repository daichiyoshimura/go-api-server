package account

import (
	"awsomeapp/internal/module/account/internal/domain"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

func TestAccountUsecase_Create(t *testing.T) {
	repo := NewMockiAccountRepository(gomock.NewController(t))
	id, _ := uuid.NewRandom()
	strID := id.String()
	name := "JohnSmith"
	input := &AccountCreateInput{
		Name: name,
	}
	udto := &domain.AccountUnspecifiedDTO{
		Name: name,
	}
	dto := &domain.AccountDTO{
		ID:   strID,
		Name: name,
	}
	account := &Account{
		ID:   strID,
		Name: name,
	}
	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		in *AccountCreateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				repo: func() iAccountRepository {
					repo.EXPECT().Create(gomock.Eq(udto)).Return(dto, nil)
					return repo
				}(),
			},
			args: args{
				in: input,
			},
			want:    account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccountUsecase(tt.fields.repo).Create(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUsecase.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Get(t *testing.T) {
	repo := NewMockiAccountRepository(gomock.NewController(t))
	id, _ := uuid.NewRandom()
	strID := id.String()
	name := "JohnSmith"
	input := &AccountGetInput{
		ID: strID,
	}
	dto := &domain.AccountDTO{
		ID:   strID,
		Name: name,
	}
	account := &Account{
		ID:   strID,
		Name: name,
	}
	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		in *AccountGetInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				repo: func() iAccountRepository {
					repo.EXPECT().Get(gomock.Eq(strID)).Return(dto, nil)
					return repo
				}(),
			},
			args: args{
				in: input,
			},
			want:    account,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AccountUsecase{
				repo: tt.fields.repo,
			}
			got, err := u.Get(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUsecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUsecase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Update(t *testing.T) {
	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		in *AccountUpdateInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		// TODO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AccountUsecase{
				repo: tt.fields.repo,
			}
			got, err := u.Update(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountUsecase.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUsecase.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Delete(t *testing.T) {
	type fields struct {
		repo iAccountRepository
	}
	type args struct {
		in *AccountDeleteInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AccountUsecase{
				repo: tt.fields.repo,
			}
			if err := u.Delete(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("AccountUsecase.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

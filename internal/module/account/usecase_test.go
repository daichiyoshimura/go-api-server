package account

import (
	"awsomeapp/internal/module/account/internal/domain"
	"reflect"
	"testing"
)

func TestNewAccountUsecase(t *testing.T) {
	type args struct {
		repo iAccountRepository
	}
	tests := []struct {
		name string
		args args
		want *AccountUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountUsecase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createAccountFromDTO(t *testing.T) {
	type args struct {
		dto *domain.AccountDTO
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createAccountFromDTO(tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createAccountFromDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountCreateInput_UnspecifiedDTO(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.AccountUnspecifiedDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccountCreateInput{
				Name: tt.fields.Name,
			}
			if got := i.UnspecifiedDTO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountCreateInput.UnspecifiedDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountUsecase_Create(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AccountUsecase{
				repo: tt.fields.repo,
			}
			got, err := u.Create(tt.args.in)
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
		// TODO: Add test cases.
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

func TestAccountUpdateInput_DTO(t *testing.T) {
	type fields struct {
		ID   int64
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.AccountDTO
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &AccountUpdateInput{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if got := i.DTO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountUpdateInput.DTO() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
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

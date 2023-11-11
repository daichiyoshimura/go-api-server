package repository

import (
	"awsomeapp/internal/db"
	"awsomeapp/internal/env"
	"awsomeapp/internal/module/account/internal/domain"
	"reflect"
	"testing"

	"github.com/uptrace/bun"
)

func TestNewAccountRepository(t *testing.T) {
	t.Setenv("STAGE", "TEST")
	_, dbEnv, _ := env.NewReader().Read()
	dbClient, _ := db.NewConnection().Establish(dbEnv)
	
	type args struct {
		db bun.IDB
	}
	tests := []struct {
		name string
		args args
		want *AccountRepository
	}{
		{
			name: "define",
			args: args{
				db: dbClient,
			},
			want: &AccountRepository{
				db: dbClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_Get(t *testing.T) {
	type fields struct {
		db bun.IDB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				db: tt.fields.db,
			}
			got, err := r.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_Create(t *testing.T) {
	type fields struct {
		db bun.IDB
	}
	type args struct {
		udto *domain.AccountUnspecifiedDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				db: tt.fields.db,
			}
			got, err := r.Create(tt.args.udto)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_Update(t *testing.T) {
	type fields struct {
		db bun.IDB
	}
	type args struct {
		dto *domain.AccountDTO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.AccountDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				db: tt.fields.db,
			}
			got, err := r.Update(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_Delete(t *testing.T) {
	type fields struct {
		db bun.IDB
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{
				db: tt.fields.db,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

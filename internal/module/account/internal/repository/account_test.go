package repository_test

import (
	"awsomeapp/internal/db"
	"awsomeapp/internal/env"
	"awsomeapp/internal/module/account/internal/domain"
	"awsomeapp/internal/module/account/internal/repository"
	"reflect"
	"testing"

	"github.com/uptrace/bun"
)

func TestAccountRepository_Get(t *testing.T) {
	t.Setenv("STAGE", "TEST")
	_, dbEnv, _, _ := env.NewReader().Read() //nolint
	dbClient, _ := db.NewPool().Establish(dbEnv)

	var idExists int64 = 1
	nameExists := "JohnSmith"
	var idNotExits int64 = 2

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
		{
			name: "exists",
			fields: fields{
				db: dbClient,
			},
			args: args{
				id: idExists,
			},
			want: &domain.AccountDTO{
				ID:   idExists,
				Name: nameExists,
			},
			wantErr: false,
		},
		{
			name: "not exists",
			fields: fields{
				db: dbClient,
			},
			args: args{
				id: idNotExits,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repository.NewAccountRepository(tt.fields.db).Get(tt.args.id)
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewAccountRepository(tt.fields.db)
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewAccountRepository(tt.fields.db)
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
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewAccountRepository(tt.fields.db)
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

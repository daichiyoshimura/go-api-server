package model

import (
	"awsomeapp/internal/module/account/internal/domain"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun"
)

func TestCreateAccountFromDTO(t *testing.T) {
	var id int64 = 1
	name := "JohnSmith"
	dto := &domain.AccountDTO{
		ID:   id,
		Name: name,
	}

	type args struct {
		dto *domain.AccountDTO
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "define",
			args: args{
				dto: dto,
			},
			want: &Account{
				ID:   id,
				Name: name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAccountFromDTO(tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountFromDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAccountFromUnspecifiedDTO(t *testing.T) {
	name := "JohnSmith"
	udto := &domain.AccountUnspecifiedDTO{
		Name: name,
	}
	type args struct {
		udto *domain.AccountUnspecifiedDTO
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "define",
			args: args{
				udto: udto,
			},
			want: &Account{
				Name: name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAccountFromUnspecifiedDTO(tt.args.udto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountFromUnspecifiedDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_DTO(t *testing.T) {
	var id int64 = 1
	name := "JohnSmith"

	type fields struct {
		BaseModel bun.BaseModel
		ID        int64
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.AccountDTO
	}{
		{
			name: "define",
			fields: fields{
				ID:   id,
				Name: name,
			},
			want: &domain.AccountDTO{
				ID:   id,
				Name: name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				BaseModel: tt.fields.BaseModel,
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
				DeletedAt: tt.fields.DeletedAt,
			}
			if got := a.DTO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.DTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

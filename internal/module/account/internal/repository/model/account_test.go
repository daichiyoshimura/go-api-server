package model

import (
	"awsomeapp/internal/module/account/internal/domain"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func TestCreateAccountFromDTO(t *testing.T) {
	id, _ := uuid.NewRandom()
	strid := id.String()
	binid, _ := id.MarshalBinary()
	name := "JohnSmith"
	dto := &domain.AccountDTO{
		ID:   strid,
		Name: name,
	}

	type args struct {
		dto *domain.AccountDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "define",
			args: args{
				dto: dto,
			},
			want: &Account{
				ID:   binid,
				Name: name,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAccountFromDTO(tt.args.dto)
			if err != nil && !tt.wantErr {
				t.Errorf("CreateAccountFromDTO() = %v, wantErr %v", err, tt.want)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountFromDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAccountFromUnspecifiedDTO(t *testing.T) {
	id, _ := uuid.NewRandom()
	binid, _ := id.MarshalBinary()
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
				ID:   binid,
				Name: name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAccountFromUnspecifiedDTO(tt.args.udto)
			if err != nil {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccountFromUnspecifiedDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_DTO(t *testing.T) {
	id, _ := uuid.NewRandom()
	binid, _ := id.MarshalBinary()
	strid := id.String()
	name := "JohnSmith"

	type fields struct {
		BaseModel bun.BaseModel
		ID        []byte
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
				ID:   binid,
				Name: name,
			},
			want: &domain.AccountDTO{
				ID:   strid,
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
			got, err := a.DTO()
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.DTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

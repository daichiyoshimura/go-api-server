package domain

import (
	"awsomeapp/internal/module/account/internal/domain/value"
	"reflect"
	"testing"
)

func TestNewAccountEntity(t *testing.T) {
	var id int64 = 1
	var name string = "JohnSmith"
	accountName, _ := value.NewAccountName(name)

	type args struct {
		dto *AccountDTO
	}
	tests := []struct {
		name    string
		args    args
		want    *AccountEntity
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				dto: &AccountDTO{
					ID:   id,
					Name: name,
				},
			},
			want: &AccountEntity{
				id:   id,
				name: accountName,
			},
			wantErr: false,
		},
		{
			name: "invalid name",
			args: args{
				dto: &AccountDTO{
					ID:   id,
					Name: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccountEntity(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccountEntity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountEntity_DTO(t *testing.T) {
	var id int64 = 1
	var name string = "JohnSmith"
	accountName, _ := value.NewAccountName(name)

	type fields struct {
		id   int64
		name *value.AccountName
	}
	tests := []struct {
		name   string
		fields fields
		want   *AccountDTO
	}{
		{
			name: "name",
			fields: fields{
				id:   id,
				name: accountName,
			},
			want: &AccountDTO{
				ID:   id,
				Name: name,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &AccountEntity{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := e.DTO(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountEntity.DTO() = %v, want %v", got, tt.want)
			}
		})
	}
}

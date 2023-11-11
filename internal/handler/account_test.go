package handler

import (
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNewAccountHandler(t *testing.T) {
	type args struct {
		usecase iAccountUsecase
	}
	tests := []struct {
		name string
		args args
		want *AccountHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountHandler(tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountHandler_GetAccount(t *testing.T) {
	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		ctx echo.Context
		id  int64
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
			h := &AccountHandler{
				usecase: tt.fields.usecase,
			}
			if err := h.GetAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountHandler_PostAccount(t *testing.T) {
	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		ctx echo.Context
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
			h := &AccountHandler{
				usecase: tt.fields.usecase,
			}
			if err := h.PostAccount(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.PostAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountHandler_PutAccount(t *testing.T) {
	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		ctx echo.Context
		id  int64
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
			h := &AccountHandler{
				usecase: tt.fields.usecase,
			}
			if err := h.PutAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.PutAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountHandler_DeleteAccount(t *testing.T) {
	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		ctx echo.Context
		id  int64
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
			h := &AccountHandler{
				usecase: tt.fields.usecase,
			}
			if err := h.DeleteAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

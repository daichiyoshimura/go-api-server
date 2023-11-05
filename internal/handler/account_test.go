package handler_test

import (
	"awsomeapp/internal/handler"
	"awsomeapp/internal/module/account"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAccountHandler_GetAccount(t *testing.T) {
	ctx := echo.New().NewContext(nil, nil)
	
	var id int64 = 10

	usecase, _ := account.WireMock(t)

	type fields struct {
		usecase handler.IAccountUsecase
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
		{
			name: "success",
			fields: fields{
				usecase: usecase,
			},
			args: args{
				ctx: ctx,
				id:  id,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewAccountHandler(usecase)
			if err := h.GetAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package handler_test

import (
	"awsomeapp/internal/handler"
	"awsomeapp/internal/module/account"
	"awsomeapp/internal/module/account/mock"
	"awsomeapp/internal/server"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestAccountHandler_GetAccount(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost/user/10", nil)
	res := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, res)

	var id int64 = 10

	name := "John Smith"

	mockErr := errors.Newf("mock err")

	ctrl := gomock.NewController(t)
	usecase := mock.NewMockIAccountUsecase(ctrl)

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
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				usecase: func(usecase *mock.MockIAccountUsecase) *mock.MockIAccountUsecase {
					usecase.EXPECT().Get(&account.AccountGetInput{
						ID: id,
					}).Return(&account.AccountGetOutput{
						ID:   id,
						Name: name,
					}, nil)

					return usecase
				}(usecase),
			},
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: response(&server.Account{
				Id:   id,
				Name: name,
			}),
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				usecase: func(usecase *mock.MockIAccountUsecase) *mock.MockIAccountUsecase {
					usecase.EXPECT().Get(&account.AccountGetInput{
						ID: id,
					}).Return(nil, mockErr)

					return usecase
				}(usecase),
			},
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: response(&server.Error{
				Code:    0,
				Message: mockErr.Error(),
			}),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler.NewAccountHandler(usecase)
			if err := h.GetAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AccountHandler.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
			} else if !reflect.DeepEqual(res.Body.String(), tt.want) {
				t.Errorf("AccountHandler.GetAccount() = %#v, want %#v", res.Body.String(), tt.want)
			}
			res.Body.Reset()
		})
	}
}

func response(v any) string {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	var sb strings.Builder

	sb.Write(res)

	sb.WriteString("\n")

	return sb.String()
}

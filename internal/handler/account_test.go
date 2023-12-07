package handler

import (
	"awsomeapp/internal/module/account"
	"awsomeapp/internal/server"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/mock/gomock"
)

func TestAccountHandler_GetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	u := NewMockiAccountUsecase(ctrl)

	id, _ := uuid.NewRandom()
	strid := id.String()
	name := "JohnSmith"
	in := &account.AccountGetInput{
		ID: strid,
	}
	out := &account.Account{
		ID:   strid,
		Name: name,
	}

	responseBody := &server.Account{
		Id:   strid,
		Name: name,
	}
	errMsg := "mock err"
	errResponseBody := &server.Error{
		Message: errMsg,
	}

	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		id string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		want       *server.Account
		wantErrRes *server.Error
	}{
		{
			name: "got one",
			fields: fields{
				usecase: func() *MockiAccountUsecase {
					u.EXPECT().Get(gomock.Eq(in)).Return(out, nil)
					return u
				}(),
			},
			args: args{
				id: strid,
			},
			wantErr:    false,
			want:       responseBody,
			wantErrRes: nil,
		},
		{
			name: "failed to get one",
			fields: fields{
				usecase: func() *MockiAccountUsecase {
					u.EXPECT().Get(gomock.Eq(in)).Return(nil, errors.Newf(errMsg))
					return u
				}(),
			},
			args: args{
				id: strid,
			},
			wantErr:    true,
			want:       nil,
			wantErrRes: errResponseBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/account/1", nil)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			h := NewAccountHandler(tt.fields.usecase)
			h.GetAccount(ctx, tt.args.id)
			if tt.wantErr {
				gotErrRes := &server.Error{}
				json.Unmarshal(rec.Body.Bytes(), gotErrRes)
				if !reflect.DeepEqual(gotErrRes, tt.wantErrRes) {
					t.Errorf("AccountHandler.GetAccount() got = %v, want %v", gotErrRes, tt.wantErrRes)
				}
			} else {
				got := &server.Account{}
				json.Unmarshal(rec.Body.Bytes(), got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AccountHandler.GetAccount() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAccountHandler_PostAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	u := NewMockiAccountUsecase(ctrl)

	id, _ := uuid.NewRandom()
	strid := id.String()
	name := "JohnSmith"
	in := &account.AccountCreateInput{
		Name: name,
	}
	out := &account.Account{
		ID:   strid,
		Name: name,
	}

	requestBody := &server.NewAccount{
		Name: name,
	}

	responseBody := &server.Account{
		Id:   strid,
		Name: name,
	}
	errMsg := "mock err"

	errResponseBody := &server.Error{
		Message: errMsg,
	}

	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		reqBody *server.NewAccount
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		want       *server.Account
		wantErrRes *server.Error
	}{
		{
			name: "create an account",
			fields: fields{
				usecase: func() *MockiAccountUsecase {
					u.EXPECT().Create(gomock.Eq(in)).Return(out, nil)
					return u
				}(),
			},
			args: args{
				reqBody: requestBody,
			},
			wantErr:    false,
			want:       responseBody,
			wantErrRes: nil,
		},
		{
			name: "failed to create an account",
			fields: fields{
				usecase: func() *MockiAccountUsecase {
					u.EXPECT().Create(gomock.Eq(in)).Return(nil, errors.Newf(errMsg))
					return u
				}(),
			},
			args: args{
				reqBody: requestBody,
			},
			wantErr:    true,
			want:       nil,
			wantErrRes: errResponseBody,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteBody, _ := json.Marshal(requestBody)
			bufBody := bytes.NewReader(byteBody)

			req := httptest.NewRequest(http.MethodPost, "/account", bufBody)
			rec := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, rec)

			h := NewAccountHandler(tt.fields.usecase)
			h.PostAccount(ctx)
			if tt.wantErr {
				gotErrRes := &server.Error{}
				json.Unmarshal(rec.Body.Bytes(), gotErrRes)
				if !reflect.DeepEqual(gotErrRes, tt.wantErrRes) {
					t.Errorf("AccountHandler.PostAccount() got = %v, want %v", gotErrRes, tt.wantErrRes)
				}
			} else {
				got := &server.Account{}
				json.Unmarshal(rec.Body.Bytes(), got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AccountHandler.PostAccount() got = %v, want %v", got, tt.want)
				}
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
		id  string
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
	id, _ := uuid.NewRandom()
	_ = id.String()
	type fields struct {
		usecase iAccountUsecase
	}
	type args struct {
		ctx echo.Context
		id  string
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

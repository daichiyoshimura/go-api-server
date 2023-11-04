package account_test

import (
	"awsomeapp/internal/module/account"
	"awsomeapp/internal/module/account/internal/domain"
	"awsomeapp/internal/module/account/internal/repository/mock"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestAccountUsecase_Create(t *testing.T) {
	name := domain.AccountName("John Smith")
	createDTO := &domain.AccountCreateDTO{
		Name: name,
	}
	id := domain.AccountID(10)

	DTO := &domain.AccountDTO{
		ID:   id,
		Name: name,
	}
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	repo := mock.NewMockIAccountRepository(ctrl)

	type fields struct {
		repo account.IAccountRepository
	}

	type args struct {
		in *account.AccountCreateInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *account.AccountCreateOutput
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Create(createDTO).Return(DTO, nil)

					return repo
				}(repo),
			},
			args: args{
				in: &account.AccountCreateInput{
					Name: string(name),
				},
			},
			want: &account.AccountCreateOutput{
				ID:   int64(id),
				Name: string(name),
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				repo: func(repo *mock.MockIAccountRepository) *mock.MockIAccountRepository {
					repo.EXPECT().Create(createDTO).Return(nil, errors.Newf("mock err"))

					return repo
				}(repo),
			},
			args: args{
				in: &account.AccountCreateInput{
					Name: string(name),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := account.NewAccountUsecase(tt.fields.repo)
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

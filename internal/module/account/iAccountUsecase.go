package account

type IAccountUsecase interface {
	Create(in *AccountCreateInput) (*AccountCreateOutput, error)
	Get(in *AccountGetInput) (*AccountGetOutput, error)
	Update(in *AccountUpdateInput) (*AccountUpdateOutput, error)
	Delete(in *AccountDeleteInput) error
}

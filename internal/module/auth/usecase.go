package auth

type AuthUsecase struct{}

func NewAuthUsecase() *AuthUsecase {
	return &AuthUsecase{}
}

type SigninInput struct {
	ID   string
	Pass string
}

func (u *AuthUsecase) Signin(in *SigninInput) (bool, error) {
	// TODO: account client
	return true, nil
}

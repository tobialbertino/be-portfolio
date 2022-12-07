package usecase

import (
	"context"
	"time"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"
	"tobialbertino/portfolio-be/pkg/tokenize"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthUseCaseImpl struct {
	UserUseCase    UserUseCase
	AuthRepository postgres.AuthRepository
	DB             *pgxpool.Pool
	Validate       *validator.Validate
}

func NewAuthUseCase(userUC UserUseCase, authRepo postgres.AuthRepository, DB *pgxpool.Pool, validate *validator.Validate) AuthUseCase {
	return &AuthUseCaseImpl{
		UserUseCase:    userUC,
		AuthRepository: authRepo,
		DB:             DB,
		Validate:       validate,
	}
}

// AddRefreshToken implements AuthUseCase
func (useCase *AuthUseCaseImpl) AddRefreshToken(req *domain.ReqLoginUser) (*domain.ResToken, error) {
	var (
		tokenRes *domain.ResToken = new(domain.ResToken)
		token    *entity.Token    = new(entity.Token)
		user     *entity.User     = new(entity.User)
	)

	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	user = &entity.User{
		Username:  req.Username,
		Passwword: req.Passwword,
	}
	// verifyUserCredential
	_, err = useCase.UserUseCase.VerifyUserCredential(user)
	if err != nil {
		return nil, err
	}

	// Create the Claims
	myClaims := tokenize.AccountClaims{
		Username:  req.Username,
		Passwword: req.Passwword,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken, err := tokenize.GenerateAccessToken(myClaims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := tokenize.GenerateRefreshToken(myClaims)
	if err != nil {
		return nil, err
	}

	token = &entity.Token{
		Token: refreshToken,
	}
	i, err := useCase.AuthRepository.AddRefreshToken(context.Background(), useCase.DB, token)
	if err != nil || i != 0 {
		return nil, err
	}

	tokenRes = &domain.ResToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenRes, nil
}

// DeleteRefreshToken implements AuthUseCase
func (useCase *AuthUseCaseImpl) DeleteRefreshToken(req *domain.ReqLoginUser) (*domain.ResToken, error) {
	panic("unimplemented")
}

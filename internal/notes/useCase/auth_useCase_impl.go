package usecase

import (
	"context"
	"time"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"
	"tobialbertino/portfolio-be/pkg/tokenize"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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
	userDetail, err := useCase.UserUseCase.VerifyUserCredential(user)
	if err != nil {
		return nil, err
	}

	// Create the Claims
	myClaims := tokenize.AccountClaims{
		ID:        userDetail.Id,
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	}

	myRefreshClaims := tokenize.AccountClaims{
		ID:        userDetail.Id,
		ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
	}

	accessToken, err := tokenize.GenerateAccessToken(myClaims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := tokenize.GenerateRefreshToken(myRefreshClaims)
	if err != nil {
		return nil, err
	}

	token = &entity.Token{
		Token: refreshToken,
	}
	i, err := useCase.AuthRepository.AddRefreshToken(context.Background(), useCase.DB, token)
	if err != nil && (i == 0) {
		return nil, err
	}

	tokenRes = &domain.ResToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokenRes, nil
}

func (useCase *AuthUseCaseImpl) VerifyRefreshToken(req *domain.ReqRefreshToken) (*domain.ResToken, error) {
	var (
		tokenRes *domain.ResToken = new(domain.ResToken)
		token    *entity.Token    = new(entity.Token)
	)

	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	token = &entity.Token{
		Token: req.RefreshToken,
	}

	// validasi dari database
	s, err := useCase.AuthRepository.VerifyRefreshToken(context.Background(), useCase.DB, token)
	if err != nil || s == "" {
		return nil, exception.NewClientError("Refresh token tidak valid", 400)
	}
	// validasi dari token signature
	tokenDetail, err := tokenize.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, exception.NewClientError("Refresh token tidak valid", 400)
	}

	// Cast data to map[string]interface{} and cast data["name"] to string
	claims := tokenDetail.Claims.(jwt.MapClaims)
	dataID := claims["ID"].(string)

	// Create the Claims
	myClaims := tokenize.AccountClaims{
		ID:        dataID,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}

	accessToken, err := tokenize.GenerateAccessToken(myClaims)
	if err != nil {
		return nil, err
	}

	tokenRes = &domain.ResToken{
		AccessToken: accessToken,
	}

	return tokenRes, nil
}

// DeleteRefreshToken implements AuthUseCase
func (useCase *AuthUseCaseImpl) DeleteRefreshToken(req *domain.ReqRefreshToken) (*domain.ResToken, error) {
	var (
		tokenRes *domain.ResToken = new(domain.ResToken)
		token    *entity.Token    = new(entity.Token)
	)

	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, err
	}

	token = &entity.Token{
		Token: req.RefreshToken,
	}

	s, err := useCase.AuthRepository.VerifyRefreshToken(context.Background(), useCase.DB, token)
	if err != nil || s == "" {
		return nil, exception.NewClientError("Refresh token tidak valid", 400)
	}

	i, err := useCase.AuthRepository.DeleteRefreshToken(context.Background(), useCase.DB, token)
	if err != nil || i == 0 {
		return nil, err
	}

	return tokenRes, nil
}

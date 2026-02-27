package auth

import (
	"context"
	"net/mail"

	"github.com/Syrix42/link-shortener/internal/services/repositories"
)

type LoginService struct {
	UserRepo repositories.UserRepository
	Comparer Comparer
}

func NewLoginService(Userrepository repositories.UserRepository, Comparer Comparer) *LoginService {
	return &LoginService{
		UserRepo: Userrepository,
		Comparer: Comparer,
	}
}

func (l *LoginService) Login(ctx context.Context, Email, Password string) (string, string, error) {
	ValidEmail, err := mail.ParseAddress(Email)
	if err != nil {
		return "", "", ErrInvalidEmailFormat
	}
	existance, err := l.UserRepo.GetByEmail(ctx, ValidEmail.Address)
	if existance == nil {
		return "", "", ErrUserNotFound
	}
	if err != nil {
		return "", "", err
	}
	err = l.Comparer.Compare(ctx, existance.HashedPassword, Password)

	if err != nil {
		return "", "", err
	}

}

package di

import "go/adv-dev/internal/user"

type IStatRepository interface {
	AddClick(linkId uint)
}

type IUserRepository interface {
	Create(user *user.User) (*user.User, error)
	GetByEmail(email string) (*user.User, error)
}

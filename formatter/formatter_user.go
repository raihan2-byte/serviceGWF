package formatter

import (
	"payment-gwf/entity"
	"time"
)

type UserFormatter struct {
	ID        int       `json:"ID"`
	Username  string    `json:"Name"`
	Email     string    `json:"Email"`
	Token     string    `json:"Token"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type LoginFormatter struct {
	Token string `json:"Token"`
}

func LoginFormatterUser(user *entity.User, Token string) LoginFormatter {
	formatter := LoginFormatter{
		Token: Token,
	}
	return formatter
}

func UpdatedFormatterUser(user *entity.User, Token string) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Token:     Token,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}

package main

import "fmt"

type AccountManager struct {
	accounts   []*AccountItem
	nextUserId int
}

func NewAccountManager() *AccountManager {
	return &AccountManager{
		nextUserId: 0,
		accounts:   make([]*AccountItem, 0),
	}
}

type AccessPermission uint64

const (
	AccessPermission_NormalUser = 1 << 0
	AccessPermission_Admin      = 1 << 63
)

type AccountItem struct {
	UserId       uint64
	Username     string
	PasswordHash string
	Permission   AccessPermission
}

// TODO: use DB (ex. PostgreSQL)
func (am *AccountManager) Add(username string, passwordHash string, permission AccessPermission) {
	am.accounts = append(
		am.accounts,
		&AccountItem{
			UserId:       uint64(am.nextUserId),
			Username:     username,
			PasswordHash: passwordHash,
			Permission:   permission,
		},
	)
	am.nextUserId++
}

// TODO: use DB (ex. PostgreSQL)
func (am *AccountManager) Remove(userId uint64) error {
	for _, account := range am.accounts {
		if userId == account.UserId {
			account = nil
			return nil
		}
	}

	// if account not found,
	return fmt.Errorf("account(id=%v) not found", userId)
}

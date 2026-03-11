package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("absdefghigklmnopqrstuvwhyzABSCIFGHIGKLMNOPQRSTUVWXYZ1234567890!.~?*")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *Account) OutputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, urlErr := url.ParseRequestURI(urlString)
	if urlErr != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		Account: Account{
			login:    login,
			url:      urlString,
			password: password,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var letterRunes = []rune("absdefghigklmnopqrstuvwhyzABSCIFGHIGKLMNOPQRSTUVWXYZ1234567890!.~?*")

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, urlErr := url.ParseRequestURI(urlString)
	if urlErr != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimeStamp{
		account: account{
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

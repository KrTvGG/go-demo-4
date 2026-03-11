package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"
	"github.com/fatih/color"
)

var letterRunes = []rune("absdefghigklmnopqrstuvwhyzABSCIFGHIGKLMNOPQRSTUVWXYZ1234567890!.~?*")

type Account struct {
	login    string `json:"login" xml:"test"`
	password string
	url      string
}

type AccountWithTimeStamp struct {
	Account
	createdAt time.Time
	updatedAt time.Time
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
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

	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
	fmt.Println(string(field.Tag))

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

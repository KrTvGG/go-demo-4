package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"
	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault() (*Vault) {
	db := files.NewJsonDB("data.json")
	file, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatcked := strings.Contains(account.Url, url)
		if isMatcked {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) DestroyAccountByUrl(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatcked := strings.Contains(account.Url, url)
		if !isMatcked {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *Vault) save() {
	db := files.NewJsonDB("data.json")
	vault.UpdatedAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}
	db.Write(data)
}

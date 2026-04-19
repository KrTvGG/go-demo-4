package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type data[T any] struct {
	el []T
}

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": destroyAccount,
}

func main() {
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewJsonDB("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		meneFunc := menu[variant]
		if meneFunc == nil {
			break Menu
		}
		meneFunc(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	destroyAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль (enter для генерации)"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	accounts := vault.FindAccounts(url, ckeckUrl)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}

func ckeckUrl(acc account.Account, str string) bool {
	return strings.Contains(acc.Url, str)
}

func destroyAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	isDeleted := vault.DestroyAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData[T any](prompts []T) string {
	for index, line := range prompts {
		if len(prompts) - 1 == index {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

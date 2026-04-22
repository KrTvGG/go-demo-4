package main

import (
	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type data[T any] struct {
	el []T
}

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": destroyAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func menuCounter() func() {
	i := 0
	return func () {
		i++
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("__Менеджер паролей__")

	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось найти .env файл")
	}

	vault := account.NewVault(files.NewJsonDB("data.vault"), *encrypter.NewEncrypter())
	counter := menuCounter()
	// vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		counter()
		variant := promptData(menuVariants...)
		meneFunc := menu[variant]
		if meneFunc == nil {
			break Menu
		}
		meneFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль (enter для генерации)")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	accounts := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range *accounts {
		account.Output()
	}
}

func destroyAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	isDeleted := vault.DestroyAccountByUrl(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData(prompts ...string) string {
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

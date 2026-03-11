package main

import (
	"fmt"
	"demo/password/account"
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль (enter для генерации)")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

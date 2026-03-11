package main

import (
	"fmt"
)

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль (enter для генерации)")
	url := promptData("Введите URL")

	myAccount, err := newAccountWithTimeStamp(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}

	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

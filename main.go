package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	files.ReadFile()
	files.WriteFile("Привет, я файл", "file.txt")
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

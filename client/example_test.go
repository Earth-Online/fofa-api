package client

import (
	"fmt"
	"os"
)

func ExampleNewUser() {
	fakeEmail := "fake"
	fakeToken := "fake"
	_ = NewUser(fakeEmail, fakeToken)
}

func ExampleUser_Me() {
	err := user.Me()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s have %d message %d Fcoin", user.UserName, user.MessageNum, user.Fcoin)
}

func ExampleUser_GetMessages() {
	msg, err := user.GetMessages()
	if err != nil {
		fmt.Print(err)
		return
	}
	if len(msg) == 0 {
		fmt.Print("not have message")
		return
	}
	for _, value := range msg {
		fmt.Printf("%s : <msg>", value.Time)
	}
}

func ExampleUser_GetPocCode() {
	code, err := user.GetPocCode("phpyun_v4_install_getshell.rb")
	if err != nil {
		fmt.Print(err)
		return
	}
	output, err := os.OpenFile("phpyun_v4_install_getshell.rb", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return
	}
	_, err = output.WriteString(code)
	if err != nil {
		fmt.Print(err)
		return
	}

}

package practice

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func AboutSetuid() {
	file, err := os.Open("my_shadow.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data,err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("my_shadow:%s \n", data)

	//打印Uid相关信息
	printUser("smbuser")
}

func printUser(username string) {
	fmt.Printf("current user is:")
	fmt.Println(user.Current())
	fmt.Printf("look up user:%s", username)
	fmt.Println( user.Lookup(username))
	fmt.Printf("look up userid:0")
	fmt.Println(user.LookupId("0"))
}
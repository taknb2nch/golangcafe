package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printf("Current User", u)

	// change user id to yours
	u, err = user.LookupId("S-1-5-21-..........")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printf("LookupId User", u)

	// change username to yours
	u, err = user.Lookup("hogeuser")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printf("Lookup User", u)

	fmt.Printf("Getuid: %v\n", os.Getuid())
	fmt.Printf("Getgid: %v\n", os.Getgid())
}

func printf(method string, u *user.User) {
	fmt.Printf("#%s#\nUid: %s\nGid: %s\nUsername: %s\nName: %s\nHomeDir: %s\n\n",
		method, u.Uid, u.Gid, u.Username, u.Name, u.HomeDir)
}

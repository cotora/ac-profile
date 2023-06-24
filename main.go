package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/cotora/ac-profile/acuser"
)


func main() {
	var (
		//u=flag.String("u","tourist","user name flag")
		h = flag.Bool("h", false, "information about heuristic is displayed")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: ac-profile [options] user_name\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("[Error] : no input user name")
		return
	}

	var user acuser.ACuser

	err := user.Init(flag.Arg(0), *h)
	if err != nil {
		fmt.Println(err)
		return
	}

	user.PrintInformation()
}

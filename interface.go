package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/legowerewolf/cryptowrapper/wrapper"
)

func main() {
	var file, key, mode string
	for index, arg := range os.Args {
		if index == 0 {
			continue
		} else if arg == "-file" && len(os.Args) >= 1+index {
			file = os.Args[index+1]
		} else if arg == "-key" && len(os.Args) >= 1+index {
			key = os.Args[index+1]
		} else if arg == "-mode" && len(os.Args) >= 1+index {
			mode = os.Args[index+1]
		}
	}

	switch mode {
	case "encrypt":
		raw, err := ioutil.ReadFile(file)
		checkErr(err, "file read")
		fmt.Println(wrapper.SymmetricEncrypt(string(raw), key))
	}
}

func checkErr(err error, key string) {
	if err != nil {
		log.Fatal("ERROR at \""+key+"\": ", err)
	}
}

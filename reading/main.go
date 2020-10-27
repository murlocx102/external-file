package reading

import "fmt"

// esc -o=./reading/read.go -pkg=reading ./readfile

func ReadFile() {
	result, err := FSString(false, "/readfile") // false use static go file. name is key
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

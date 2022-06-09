package main

import (
	"back-end/utils"
	"fmt"
	"os"
)

func main() {
	utils.LoadEnv(".env")

	fmt.Println(os.Environ())
}

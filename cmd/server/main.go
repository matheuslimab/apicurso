package main

import (
	"fmt"

	"github.com/matheuslimab/apicurso/configs"
)

func main() {
	config, _ := configs.LoadConfig("./configs")

	fmt.Println(config.DBDriver)
}

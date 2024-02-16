package main

import (
	"fmt"

	"github.com/GoSeoTaxi/olegMicroserviceAuth/cmd"
	"github.com/fatih/color"
)

func main() {

	fmt.Println(color.HiRedString("Starting..."))
	cmd.RunService()

}

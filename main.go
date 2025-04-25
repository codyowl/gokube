package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

// cobra root command
var cobrarootCmd = &cobra.Command{
	Use:     "gokube",
	Version: "1.0.0",
	Short:   "A swiss army knife cli tool for kubernetes",
}

func main(){
	fmt.Println("Things are working fine....")
}
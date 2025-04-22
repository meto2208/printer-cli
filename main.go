package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	printer := &cobra.Command{
		Use:   "printer",
		Short: "printer - simple CLI that prints the first argument",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("please provide an argument.")
				return
			}
			fmt.Println("Printer says smth: ", args[0])
		},
	}

	if err := printer.Execute(); err != nil {
		fmt.Println("something went wrong")
		os.Exit(1)
	}
}

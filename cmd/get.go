/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
    Use:   "get",
    Short: "This command will get the desired Gopher",
    Long:  `This get command will call GitHub respository in order to return the desired Gopher.`,
    Run: func(cmd *cobra.Command, args []string) {
        var gopherName = "dr-who.png"

        if len(args) >= 1 && args[0] != "" {
            gopherName = args[0]
        }

        URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

        fmt.Println("Try to get '" + gopherName + "' Gopher...")

        response, err := http.Get(URL)
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {

			out, err := os.Create(gopherName + ".png")
            if err != nil {
                fmt.Println(err)
            }
            defer out.Close()

            _, err = io.Copy(out, response.Body)
            if err != nil {
                fmt.Println(err)
            }

            fmt.Println("Perfect! Just saved in " + out.Name() + "!")
        } else {
            fmt.Println("Error: " + gopherName + " not exists! :-(")
        }
    },
}


func init() {
	rootCmd.AddCommand(getCmd)
}
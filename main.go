/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "marcocd/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

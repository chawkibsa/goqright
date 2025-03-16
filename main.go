/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
 "github.com/chawkibsa/goqright/cmd"
 "github.com/chawkibsa/goqright/data"
 )

func main() {
 data.OpenDatabase()
	cmd.Execute()
}

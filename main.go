/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"inscriptiontool/cmd"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	cmd.Execute()
}

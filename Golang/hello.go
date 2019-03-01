package main

import (
		"fmt"
		"os"
		"strconv"
	)

func ms() {

	// Get number of Players
	fmt.Printf("Input Number of Players: ")
	var pNum string
	fmt.Scanln(&pNum)
	nPlayers, err := strconv.Atoi(pNum)
	if err != nil {
			fmt.Println(err)
			os.Exit(2)
	}

	// Populate Player Array
	players := make([]string, nPlayers)
	for i := 0; i < nPlayers; i++ {
		fmt.Printf("Player: ")
		var player string
		fmt.Scanln(&player)
		players = append(players,player)
	}

	fmt.Println(players)

	// Get Questions
	questions := make([]string, 5*nPlayers)
	fmt.Printf("Enter your questions: ")
	for i := 0; i < 5; i++ {
		fmt.Printf(">> ")
		var question string
		fmt.Scanln(&question)
		questions = append(questions,question)
	}


}

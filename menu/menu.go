package menu

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"script_manager/scripts"
	"strconv"
	"strings"
)

func PrintMenu() {

	err := scripts.LoadScripts()
	if err != nil {
		log.Fatalln("❌ Error loading scripts:", err)
		return
	}

	fmt.Println("--------------------------------")
	fmt.Println("📋 What would you like to do?")
	fmt.Println("1. 📜 List scripts")
	fmt.Println("2. ➕ Add script")
	fmt.Println("3. 🗑️ Delete script")
	fmt.Println("4. ⚙️ Config")
	fmt.Println("5. 📥 Download and install scripts")
	fmt.Println("6. 🚪 Exit")
	fmt.Println("--------------------------------")
}

func GetUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("🔢 Enter your choice (1-6): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err == nil && choice >= 1 && choice <= 6 {
			fmt.Printf("✅ You chose option %d\n", choice)
			return choice
		}
		fmt.Println("❌ Invalid input. Please enter a number between 1 and 6.")
	}
}

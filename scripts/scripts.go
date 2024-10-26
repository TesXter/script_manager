package scripts

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Script struct {
	Name string
	Link string
}

var Scripts []Script

func ListScripts() {
	if len(Scripts) == 0 {
		fmt.Println("📜 No Scripts available.")
		return
	}
	fmt.Println("📜 Available Scripts:")
	for i, script := range Scripts {
		fmt.Printf("%d. %s - %s\n", i+1, script.Name, script.Link)
	}
}

func AddScript() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("🔗 Enter GitHub link (format: https://github.com/username/repository.git): ")
	link, _ := reader.ReadString('\n')
	link = strings.TrimSpace(link)

	// Validate the link format
	if !strings.HasPrefix(link, "https://github.com/") || !strings.HasSuffix(link, ".git") {
		fmt.Println("❌ Invalid GitHub link format. Please use the format: https://github.com/username/repository.git")
		return
	}

	// Extract the repository name
	parts := strings.Split(strings.TrimSuffix(link, ".git"), "/")
	if len(parts) != 5 {
		fmt.Println("❌ Invalid GitHub link format.")
		return
	}

	name := parts[4] // The repository name is the last part of the URL

	// Check if the script already exists
	for _, script := range Scripts {
		if script.Link == link {
			fmt.Println("❌ Script already exists.")
			return
		}
	}

	Scripts = append(Scripts, Script{Name: name, Link: link})
	fmt.Println("✅ Script added successfully.")

	saveScriptsToFile()
}

func DeleteScript() {
	if len(Scripts) == 0 {
		fmt.Println("📜 No Scripts available to delete.")
		return
	}

	ListScripts()
	fmt.Print("🔢 Enter the number of the script to delete: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)

	if err != nil || index < 1 || index > len(Scripts) {
		fmt.Println("❌ Invalid input. No script deleted.")
		return
	}

	Scripts = append(Scripts[:index-1], Scripts[index:]...)
	fmt.Println("🗑️ Script deleted successfully.")

	saveScriptsToFile()
}

func saveScriptsToFile() {
	file, err := os.Create("Scripts.json")
	if err != nil {
		fmt.Println("❌ Error creating Scripts.json:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(Scripts, "", "  ")
	if err != nil {
		fmt.Println("❌ Error marshalling Scripts:", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("❌ Error writing to Scripts.json:", err)
	}
}

func LoadScripts() error {
	file, err := os.Open("Scripts.json")
	if err != nil {
		return fmt.Errorf("error opening Scripts.json: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Scripts)
	if err != nil {
		return fmt.Errorf("error decoding Scripts.json: %w", err)
	}
	return nil
}

func EnsureScriptsFileExists() {
	// 🔍 Check if scripts.json exists, create if not and write empty JSON
	if _, err := os.Stat("scripts.json"); os.IsNotExist(err) {
		file, err := os.Create("scripts.json")
		if err != nil {
			fmt.Println("❌ Error creating scripts.json:", err)
			return
		}
		defer file.Close()

		emptyJSON := []byte("[]")
		_, err = file.Write(emptyJSON)
		if err != nil {
			fmt.Println("❌ Error writing to scripts.json:", err)
			return
		}
		fmt.Println("✅ Created scripts.json file with empty JSON")
	}
}

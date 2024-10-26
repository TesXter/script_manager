package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	ScriptsFolderPath string `json:"scripts_folder_path"`
}

func LoadOrCreateConfig() Config {
	if _, err := os.Stat("config.txt"); os.IsNotExist(err) {
		return createConfig()
	}

	file, err := os.Open("config.txt")
	if err != nil {
		fmt.Println("❌ Error opening config.txt:", err)
		return createConfig()
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("❌ Error reading config.txt:", err)
		return createConfig()
	}

	return config
}

func createConfig() Config {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("📁 Enter the path to your scripts folder: ")
	path, _ := reader.ReadString('\n')
	path = filepath.Clean(strings.TrimSpace(path))

	config := Config{ScriptsFolderPath: path}

	file, err := os.Create("config.txt")
	if err != nil {
		fmt.Println("❌ Error creating config.txt:", err)
		return config
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		fmt.Println("❌ Error writing to config.txt:", err)
	}

	fmt.Println("✅ Created config.txt file")
	return config
}

func ConfigureSettings() {
	if checkConfigFileExists() {
		deleteConfigFile()
	}
	LoadOrCreateConfig()
}

func deleteConfigFile() {
	os.Remove("config.txt")
}

func checkConfigFileExists() bool {
	if _, err := os.Stat("config.txt"); os.IsNotExist(err) {
		return false
	}
	return true
}

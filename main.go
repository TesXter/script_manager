package main

import (
	"fmt"

	"script_manager/config"
	"script_manager/menu"
	"script_manager/scripts"
)

func main() {
	cfg := config.LoadOrCreateConfig()
	fmt.Println("Scripts Folder Path:", cfg.ScriptsFolderPath)

	scripts.EnsureScriptsFileExists()

	// menu.PrintMenu()
	// scripts.DownloadAndInstallScripts()

	for {
		menu.PrintMenu()
		choice := menu.GetUserChoice()

		switch choice {
		case 1:
			scripts.ListScripts()
		case 2:
			scripts.AddScript()
		case 3:
			scripts.DeleteScript()
		case 4:
			config.ConfigureSettings()
		case 5:
			scripts.DownloadAndInstallScripts()
		case 6:
			fmt.Println("🚪 Exiting the application. Goodbye! 👋")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again. 🔄")
		}
	}
}

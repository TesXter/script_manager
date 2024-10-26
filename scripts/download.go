package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"script_manager/config"
)

func DownloadAndInstallScripts() {
	fmt.Println("📥 Downloading and installing scripts...")

	if len(Scripts) == 0 {
		fmt.Println("❌ No scripts available to download.")
		return
	}

	cfg := config.LoadOrCreateConfig()
	scriptsPath := cfg.ScriptsFolderPath

	printAvailableScripts()

	for _, script := range Scripts {
		scriptPath := filepath.Join(scriptsPath, script.Name)
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			cloneScript(script, scriptsPath)
		} else {
			updateScript(script, scriptPath)
		}
	}

	fmt.Println("🎉 Finished downloading and installing scripts.")
}

func printAvailableScripts() {
	fmt.Println("📜 Available scripts:")
	for i, script := range Scripts {
		fmt.Printf("%d. %s - %s\n", i+1, script.Name, script.Link)
	}
}

func cloneScript(script Script, scriptsPath string) {
	fmt.Printf("🔄 Cloning %s...\n", script.Name)
	cmd := getGitCommand(script, "clone")
	cmd.Dir = scriptsPath
	executeGitCommand(cmd, script.Name, "cloning")
}

func updateScript(script Script, scriptPath string) {
	fmt.Printf("🔄 Updating %s...\n", script.Name)
	if err := os.RemoveAll(scriptPath); err != nil {
		fmt.Printf("❌ Error deleting existing folder for %s: %v\n", script.Name, err)
		return
	}
	cmd := getGitCommand(script, "clone")
	cmd.Dir = filepath.Dir(scriptPath)
	executeGitCommand(cmd, script.Name, "updating")
}

func getGitCommand(script Script, action string) *exec.Cmd {
	if script.Name == "piteertest" {
		return exec.Command("git", action, "-b", "quest-id-bug-test", "https://github.com/Zewx1776/piteertest.git")
	}
	return exec.Command("git", action, script.Link)
}

func executeGitCommand(cmd *exec.Cmd, scriptName, action string) {
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Error %s %s: %v\n", action, scriptName, err)
	} else {
		fmt.Printf("✅ Successfully %s %s\n", action, scriptName)
	}
}

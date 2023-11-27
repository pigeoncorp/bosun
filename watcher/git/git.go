package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashibuto/oof"
	"github.com/pigeoncorp/bosun/watcher/config"
)

const StorageDir = "/var/lib/bosun/data"

func StoreCredentials() error {
	if config.Config.GithubUsername == "" {
		return nil
	}

	srcCredential := fmt.Sprintf("https://%s:%s@github.com\n", config.Config.GithubUsername, config.Config.GithubPassword)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return oof.Trace(err)
	}

	credFile := filepath.Join(homeDir, ".git-credentials")
	data, _ := os.ReadFile(credFile)
	destCredential := string(data)
	if srcCredential == destCredential {
		return nil
	}

	err = os.WriteFile(credFile, []byte(srcCredential), 0600)
	if err != nil {
		return oof.Trace(err)
	}

	return nil
}

func CloneRepoIfNotExists() error {
	url := config.Config.GithubRepositoryUrl
	repoDir := getRepoDir()
	_, err := os.Stat(repoDir)
	if err == nil {
		return nil
	}

	cmd := exec.Command("git", "clone", url)
	cmd.Dir = StorageDir
	err = cmd.Run()
	if err != nil {
		return oof.Trace(err)
	}

	return nil
}

func PullRepoChanges() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = getRepoDir()
	err := cmd.Run()
	if err != nil {
		return oof.Trace(err)
	}

	return nil
}

func getRepoDir() string {
	url := config.Config.GithubRepositoryUrl
	repoName := parseRepoName(url)
	return filepath.Join(StorageDir, repoName)
}

func parseRepoName(url string) string {
	parts := strings.Split(url, "/")
	finalPart := parts[len(parts)-1]
	return strings.TrimSuffix(finalPart, ".git")
}

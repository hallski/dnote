package ext

import (
	"os"
	"os/exec"
)

type GitStatus int

const (
	Clean GitStatus = iota
	Dirty
)

type GitClient struct {
	repoPath string
	gitPath  string
}

func getExecPath() (string, error) {
	path, err := exec.LookPath("git")
	if err != nil {
		return "", err
	}

	return path, nil
}

func NewGitClient(path string) (*GitClient, error) {
	gitPath, err := getExecPath()
	if err != nil {
		return nil, err
	}

	client := &GitClient{path, gitPath}

	return client, nil
}

func (client *GitClient) runCommand(args ...string) (string, error) {
	if err := os.Chdir(client.repoPath); err != nil {
		return "", err
	}

	cmd := exec.Command(client.gitPath, args...)

	res, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func (client *GitClient) Status() (GitStatus, error) {
	out, err := client.runCommand("status", "-s")
	if err != nil {
		return Dirty, err
	}

	if string(out) != "" {
		return Dirty, nil
	}

	return Clean, nil
}

func (client *GitClient) Pull(path string) error {
	_, err := client.runCommand("commit", "-m", "Commit from dNote")

	if err != nil {
		return err
	}

	return nil
}

func (client *GitClient) Commit(msg string) error {
	if _, err := client.runCommand("add", "."); err != nil {
		return err
	}

	if _, err := client.runCommand("commit", "-m", msg); err != nil {
		return err
	}

	return nil
}

func (client *GitClient) PullRebasePush() error {
	if _, err := client.runCommand("pull"); err != nil {
		return err
	}

	if _, err := client.runCommand("rebase", "origin/main"); err != nil {
		return err
	}

	if _, err := client.runCommand("push", "origin"); err != nil {
		return err
	}

	return nil
}

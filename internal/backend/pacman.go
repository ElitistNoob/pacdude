package backend

import "os/exec"

type PacmanBackend struct{}

func (p PacmanBackend) ListInstalled() ResultMsg {
	cmd := exec.Command("pacman", "-Q")
	return newResultMsg(*cmd)
}

func (p PacmanBackend) Search(query string) ResultMsg {
	cmd := exec.Command("pacman", "-Ss", query)
	return newResultMsg(*cmd)
}

func (p PacmanBackend) Install(pkg string) ResultMsg {
	cmd := exec.Command("sudo", "pacman", pkg, "--noconfirm", "--needed")
	return newResultMsg(*cmd)
}

func (p PacmanBackend) Remove(pkg string) ResultMsg {
	cmd := exec.Command("sudo", "pacman", "-Rns")
	return newResultMsg(*cmd)
}

func (p PacmanBackend) ListUpgradable() ResultMsg {
	cmd := exec.Command("pacman", "Qu")
	return newResultMsg(*cmd)
}

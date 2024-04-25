package repositer

import (
    "os/exec"
    "strings"
)

func GetUserName() (string, error) {
    cmd := exec.Command("git", "config", "--global", "user.name")
    output, err := cmd.CombinedOutput()
    return strings.TrimSpace(string(output)), err
}
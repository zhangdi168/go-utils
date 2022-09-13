// Package tools /** package  tools
package tools

import "os/exec"

func OpenBrowser(uri string) error {
	cmd := exec.Command("cmd", "/C", "start "+uri)
	return cmd.Run()
}

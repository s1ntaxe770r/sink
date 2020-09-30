package main

import (
	"fmt"
	"io"
	"os"
)

// Dofile struct
type DotFile struct {
	Name     string
	Location string
}

// sync individual dotfile
func syncf(d DotFile, repo string) error {
	in, err := os.Open(d.Location + string(os.PathSeparator) + d.Name)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(repo + string(os.PathSeparator) + d.Name)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		fmt.Println(err)
	}

	out.Close()

	return err
}

func main() {
	defaultdirs := map[string]string{
		"nvimconf": "~/.config/nvim/",
		"awesome":  "~/.config/awesome",
		"zshrc":    "~/",
	}

	configrepo := "~/dots"
	nvimconf := DotFile{"init.vim", defaultdirs["nvimconf"]}
	awesomeconf := DotFile{"rc.lua", defaultdirs["awesome"]}
	zshrc := DotFile{".zshrc", defaultdirs["zshrc"]}

}

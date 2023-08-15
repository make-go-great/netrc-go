package netrc

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	prefixMachine  = "machine"
	prefixLogin    = "login"
	prefixPassword = "password"

	homeSymbol = '~'
)

func ParseFile(filename string) (Data, error) {
	filename, err := trimHomeSymbol(filename)
	if err != nil {
		return Data{}, fmt.Errorf("failed to trim home symbol: %w", err)
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return Data{}, fmt.Errorf("os: failed to read file: %w", err)
	}

	return Parse(bytes), nil
}

// trimHomeSymbol replace ~ with full path
// https://stackoverflow.com/a/17609894
func trimHomeSymbol(path string) (string, error) {
	if path == "" || path[0] != homeSymbol {
		return path, nil
	}

	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	newPath := filepath.Join(currentUser.HomeDir, path[1:])
	return newPath, nil
}

func Parse(bytes []byte) Data {
	machines := make([]Machine, 0, 10)

	var curMachine *Machine

	for _, line := range strings.Split(string(bytes), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, prefixMachine) {
			// New machine
			// Add old machine first
			if curMachine != nil {
				machines = append(machines, *curMachine)
			}

			// Replace cur machine
			curMachine = &Machine{
				Name: strings.TrimSpace(strings.TrimPrefix(line, prefixMachine)),
			}
		}

		// There is no cur machine
		if curMachine == nil {
			continue
		}

		if strings.HasPrefix(line, prefixLogin) {
			curMachine.Login = strings.TrimSpace(strings.TrimPrefix(line, prefixLogin))
		} else if strings.HasPrefix(line, prefixPassword) {
			curMachine.Password = strings.TrimSpace(strings.TrimPrefix(line, prefixPassword))
		}
	}

	// Add last machine
	if curMachine != nil {
		machines = append(machines, *curMachine)
	}

	return Data{
		Machines: machines,
	}
}

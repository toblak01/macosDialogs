package macosDialogs

import (
	"os/exec"
	"strings"
)

// ChooseEntry allows the user to choose a specific entry from a list of entries
func ChooseEntry(title, text string, entries []string) string {
	// Ask the user to choose a repository to remove
	entry, success, err := list(title, text, entries)
	if err != nil || !success {
		return ""
	}

	return entry
}

// ChooseFile allows the user to choose a file or directory
func ChooseFile(title, filter string, directory bool) (string, bool, error) {
	path, success, err := file(title, filter, directory)
	// path = strings.Replace(path, " ", "\\ ", -1)

	return path, success, err
}

// YesNoButtons will create a dialog with a Yes and a No button
func YesNoButtons(title, text string) (bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return false, err
	}

	o, err := exec.Command(osa, "-e", `set T to button returned of (display dialog "`+text+`" with title "`+title+`" buttons {"No", "Yes"} default button 2)`).Output()
	if err != nil {
		return false, err
	}

	out := strings.TrimSpace(string(o))
	result := out == "Yes"

	return result, nil
}

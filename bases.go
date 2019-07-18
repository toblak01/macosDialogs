package macosDialogs

import (
	"os/exec"
	"strings"
	"syscall"
)

// Entry displays input dialog, returning the entered value and a bool for success.
func entry(title, text, defaultText string) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	o, err := exec.Command(osa, "-e", `set T to text returned of (display dialog "`+text+`" with title "`+title+`" default answer "`+defaultText+`")`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	return out, ret, err
}

// list displays a list dialog, returning the selected value and a bool for success.
func list(title, text string, items []string) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	list := ""
	for i, l := range items {
		list += `"` + l + `"`
		if i != len(items)-1 {
			list += ", "
		}
	}

	o, err := exec.Command(osa, "-e", `choose from list {`+list+`} with prompt "`+text+`" with title "`+title+`"`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" || out == "false" {
		ret = false
	}

	return out, ret, err
}

func osaDialogNoIcon(title, text string) (bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return false, err
	}

	out, err := exec.Command(osa, "-e", `display dialog "`+text+`" with title "`+title+`" buttons {"OK"} default button "OK"`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0, nil
		}
	}

	ret := false
	if strings.TrimSpace(string(out)) == "OK" {
		ret = true
	}

	return ret, err
}

// File displays a file dialog, returning the selected file/directory and a bool for success.
func file(title, filter string, directory bool) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	f := "file"
	if directory {
		f = "folder"
	}

	t := ""
	if filter != "" {
		t = ` of type {"` + filter + `"}`
	}

	o, err := exec.Command(osa, "-e", `choose `+f+t+` with prompt "`+title+`"`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	tmp := strings.Split(out, ":")
	tmp[0] = strings.Replace(tmp[0], "alias ", "Volumes/", -1)

	outPath := "/" + strings.Join(tmp[0:len(tmp)], "/")

	return outPath, ret, err
}

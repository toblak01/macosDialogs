package macosDialogs

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
	return path, success, err
}

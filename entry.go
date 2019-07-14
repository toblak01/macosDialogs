package macosDialogs

import (
	"fmt"
	"strconv"
)

// EnterInteger provides a dialog for the user to enter an integer
func EnterInteger(title, text string, minValue, defaultValue int) (int, error) {
	value, success, err := entry(title, text, strconv.Itoa(defaultValue))
	if err != nil {
		return 0, fmt.Errorf("error")
	}

	if !success {
		return 0, fmt.Errorf("error")
	}

	if len(value) == 0 {
		return 0, fmt.Errorf("error")
	}

	newValue, err := strconv.Atoi(value)
	if err != nil {
		UserDialogNotification("Invalid Value", "An integer value needs to be entered.")
		return 0, fmt.Errorf("error")
	}

	if newValue < minValue {
		msg := fmt.Sprintf("A value less than %d is not allowed.", minValue)
		UserDialogNotification("Invalid Value", msg)
		return 0, fmt.Errorf("error")
	}

	return newValue, nil
}

// EntryWithDefault will return a value and supply a default
func EntryWithDefault(title, text, defaultResult string) (string, error) {
	result, success, err := entry(title, text, defaultResult)
	if err != nil {
		return "", fmt.Errorf("error")
	}

	if !success {
		return "", fmt.Errorf("error")
	}

	return result, nil
}

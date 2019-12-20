package macosDialogs

import (
	"fmt"
	"strconv"
	"strings"
)

// EnterIntegerWithDefaultAndMinValue provides a dialog for the user to enter an integer
func EnterIntegerWithDefaultAndMinValue(title, text string, minValue, defaultValue int) (int, error) {
	text = text + fmt.Sprintf("\nMust be greater or equal to %d.", minValue)
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

// EnterTextWithDefault will return a value and supply a default
func EnterTextWithDefault(title, text, defaultResult string) (string, bool, error) {
	result, success, err := entry(title, text, defaultResult)
	if err != nil {
		return "", false, fmt.Errorf("error")
	}

	if !success {
		return "", false, fmt.Errorf("error")
	}

	result = strings.Trim(result, " ")
	if len(result) == 0 {
		success = false
	}

	return result, success, nil
}

// EnterText will return a value
func EnterText(title, text string) (string, bool, error) {
	result, success, err := entry(title, text, "")
	if err != nil {
		return "", false, fmt.Errorf("error")
	}

	if !success {
		return "", false, fmt.Errorf("error")
	}

	result = strings.Trim(result, " ")
	if len(result) == 0 {
		success = false
	}

	return result, success, nil
}

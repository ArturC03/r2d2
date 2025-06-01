package visitor

import (
	"fmt"
	"github.com/ArturC03/r2d2Styles"
)

// Creates a standardized error message with line number
func formatErrorMessage(message string, line int) string {
	return fmt.Sprintf("%s on line %s", message, r2d2Styles.Bold(fmt.Sprintf("%d", line)))
}

// Creates a standardized warning message with line number
func formatWarningMessage(message string, line int) string {
	return fmt.Sprintf("%s on line %s", message, r2d2Styles.Bold(fmt.Sprintf("%d", line)))
}

// Creates a standardized error message without a line number
func formatErrorMessageNoLine(message string) string {
	return message
}

// Creates a standardized warning message without line number
func formatWarningMessageNoLine(message string) string {
	return message
}

// Creates a standardized error message for operations that envolve files
func formatFileErrorMessage(message string, filePath string) string {
	return fmt.Sprintf("%s: %s", message, r2d2Styles.Bold(filePath))
}

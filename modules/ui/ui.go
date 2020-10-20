package ui

import (
	"github.com/fatih/color"
)

// Heading : formatting for headings
func Heading(heading string) {
	headingStyling := color.New(color.FgHiYellow).Add(color.Underline).Add(color.Bold)
	headingStyling.Println(heading)
}

// Description : formatting for description lines
func Description(description string) {
	descriptionStyling := color.New(color.FgMagenta).Add(color.Italic)
	descriptionStyling.Println(description)
}

// Error : formatting for error messages
func Error(errorMessage string) {
	errorMessageStyling := color.New(color.FgHiRed)
	errorMessageStyling.Println("[error] " + errorMessage)
}

// Info : formatting for info messages
func Info(infoMessage string) {
	infoMessageStyling := color.New(color.FgHiCyan)
	infoMessageStyling.Println("[info] " + infoMessage)
}

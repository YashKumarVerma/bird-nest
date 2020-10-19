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

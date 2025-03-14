package form

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

// public types --------------------------------------------------------------------------------------

type Institution struct {
	Name string
	File string
}

// public methods ------------------------------------------------------------------------------------

func DisplayInstitutionForm() Institution {
	var institution Institution

	// Should we run in accessible mode?
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	// File picker current directory
	userHomeDir, _ := os.UserHomeDir()

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Terminal Budget").
			Description("Welcome to the terminal budget.\n\n").
			Next(true).
			NextLabel("Next"),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("American Express", "Capital One", "Chase", "Discover")...).
				Title("Choose your financial institution.").
				Description("Your financial institution to import your transactions from.").
				Value(&institution.Name),

			huh.NewFilePicker().
				Title("CSV File").
				Description("Select your csv file.").
				Picking(true).
				AllowedTypes([]string{".csv", ".CSV"}).
				CurrentDirectory(userHomeDir).
				Value(&institution.File),
		),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	{
		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		/*
			There is a bug/lack of feature in lipgloss where styling is lost on string wrap...

			The wordwrap package wrapping is currently naive and only happens at white-space.
			So in this case where we have a really long file name we need to do something really hacky??
			1. Insert whitespace at the point we want to wrap
			2. Wrap the string
			3. Split by new line and style each line
			4. Join the lines back together
		*/
		preWrapFileName := insertWhitespace(institution.File, 56) // accounts for padding in 60 width print
		wrappedFileNameLines := strings.Split(wordwrap.WrapString(preWrapFileName, 56), "\n")
		var styledFileNameLines []string
		for _, line := range wrappedFileNameLines {
			styledFileNameLines = append(styledFileNameLines, keyword(line))
		}
		styledFileName := strings.Join(styledFileNameLines, "\n")

		fmt.Fprintf(&sb,
			"%s\n\nFinancial Institution\n%s \n\nCSV File\n%s",
			lipgloss.NewStyle().Bold(true).Render("TRANSACTIONS"),
			keyword(institution.Name),
			styledFileName,
		)

		fmt.Println(
			lipgloss.NewStyle().
				Width(60).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)
	}

	return institution
}

// private methods ------------------------------------------------------------------------------------

func insertWhitespace(s string, index int) string {
	return s[:index] + " " + s[index:]
}

package form

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
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
				Title("CSV Transaction File").
				Description("Select your csv transaction file.").
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
			fmt.Println(s)
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\nFinancial Institution\n%s \n\nTransaction File\n%s",
			lipgloss.NewStyle().Bold(true).Render("TRANSACTIONS"),
			keyword(institution.Name),
			keyword(institution.File),
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

package ui

import (
	"github.com/charmbracelet/lipgloss"
	c "worker/cfg"
)

const (
	header = "<|!|> AFK Worker v0.1 <|!|>"
)
const (
	hotPink     = lipgloss.Color("#FF06B7")
	black       = lipgloss.Color("0")
	white       = lipgloss.Color("#FFFFFF")
	darkGray    = lipgloss.Color("#767676")
	purple      = lipgloss.Color("99")
	brightGreen = lipgloss.Color("#00FF00")
	bloodRed    = lipgloss.Color("#FF0000")
	someG       = lipgloss.Color("#00FFa0")
	someR       = lipgloss.Color("#FFa000")
	orange      = lipgloss.Color("#faa805")
	sep         = "|>"
	cyanny      = lipgloss.Color("#000FFF")
)

var (
	//////////////
	/// LEFT /////
	// Panel ////
	////////////

	// Title
	titleStl = lipgloss.NewStyle().
			Bold(true).
			Foreground(white).
			Align(lipgloss.Right)

	// list Title Bar
	titbarStl = lipgloss.NewStyle().
			MarginBottom(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(hotPink)

	menulistStl = lipgloss.NewStyle().
			ColorWhitespace(true).
			Align(lipgloss.Bottom, lipgloss.Left).
			Margin(0, 0, 0, 0).
			Width(50)

	///////////////
	/// RIGHT ////
	// Panel ////
	////////////
	statusStl = lipgloss.NewStyle().
			MarginLeft(1).
			Border(lipgloss.RoundedBorder()).
			Bold(true).
			Width(52).
			PaddingLeft(3).
			PaddingRight(5).
			Align(lipgloss.Top, lipgloss.Right).
			BorderForeground(bloodRed)

	spinnerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("69"))

	runnunTaskStyle = statusStl.Copy().
			MarginTop(2).
			UnsetPaddingLeft().
			Width(55).
			UnsetBorderStyle().
			Align(lipgloss.Bottom, lipgloss.Left)

	emuStatus = statusStl.Copy().
			BorderForeground(cyanny).
			Foreground(brightGreen).
			Align(lipgloss.Center)

	helpStyle = lipgloss.NewStyle().
			MarginLeft(1).
			Align(lipgloss.Bottom).
			Foreground(darkGray)

	quitStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")).
			MarginBackground(lipgloss.Color("#00FF00")).
			Margin(10)

	taskName = lipgloss.NewStyle().Foreground(orange)
)

// ////////////////
// / settings ////
// // input /////
// /////////////
var (
	// MultiText Input Form
	topInputStyle = lipgloss.NewStyle().
			PaddingLeft(30).
			PaddingTop(8)

	focusedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))

	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	cursorStyle = focusedStyle.Copy()

	noStyle = lipgloss.NewStyle()

	cursorModeHelpStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().
			Render("[ Submit ]")

	blurredButton = c.F("[ %s ]", blurredStyle.Render("Submit"))
)

var (
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			MarginLeft(1).
			MarginBottom(3)
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("170"))
)

/////////////////////
/////////host input/////
////////////////
const ()

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

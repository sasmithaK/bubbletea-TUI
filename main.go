package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	menuScreen screen = iota
	progressScreen
	searchScreen
	helpScreen
)

type tickMsg time.Time

type model struct {
	screen    screen
	cursor    int
	width     int
	height    int
	progress  progress.Model
	percent   float64
	textInput textinput.Model
	items     []string
	filtered  []string
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type to filter features..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	p := progress.New(progress.WithDefaultGradient())

	items := []string{
		"Multi-screen Navigation",
		"Progress Bar with Ticks",
		"Search & Filter",
		"Responsive Layout",
		"Styled Components (Lipgloss)",
		"Keyboard Shortcuts",
		"State Management",
	}

	return model{
		screen:    menuScreen,
		progress:  p,
		textInput: ti,
		items:     items,
		filtered:  items,
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second/20, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tickMsg:
		if m.screen == progressScreen {
			m.percent += 0.005
			if m.percent > 1 {
				m.percent = 0
			}
			return m, tick()
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "1":
			m.screen = menuScreen

		case "2":
			m.screen = progressScreen

		case "3":
			m.screen = searchScreen

		case "4":
			m.screen = helpScreen

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.filtered)-1 {
				m.cursor++
			}
		}
	}

	if m.screen == searchScreen {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)

		query := strings.ToLower(m.textInput.Value())
		m.filtered = nil
		for _, item := range m.items {
			if strings.Contains(strings.ToLower(item), query) {
				m.filtered = append(m.filtered, item)
			}
		}
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FFCC")).
		Padding(0, 1)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		BorderForeground(lipgloss.Color("#00FFFF"))

	footer := "\n[1] Menu  [2] Progress  [3] Search  [4] Help  [q] Quit"

	switch m.screen {

	case menuScreen:
		content := "🚀 Enhanced Bubble Tea TUI\n\n"
		content += "This interface demonstrates:\n\n"
		content += "• Multi-screen architecture\n"
		content += "• Real-time tick updates\n"
		content += "• Search filtering\n"
		content += "• Responsive layout handling\n"
		content += "• Styled components with Lipgloss\n"
		content += "• Component-based state management\n\n"
		content += "Press numbers to navigate."

		return boxStyle.Render(headerStyle.Render("MAIN MENU")+"\n\n"+content) + footer

	case progressScreen:
		bar := m.progress.ViewAs(m.percent)
		content := fmt.Sprintf(
			"📊 Live Progress Demo\n\nAnimated using tea.Tick()\n\n%s\n\n%.0f%%",
			bar,
			m.percent*100,
		)
		return boxStyle.Render(headerStyle.Render("PROGRESS DEMO")+"\n\n"+content) + footer

	case searchScreen:
		content := "🔍 Search & Filter Demo\n\n"
		content += m.textInput.View() + "\n\n"

		for i, item := range m.filtered {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			content += fmt.Sprintf("%s %s\n", cursor, item)
		}

		return boxStyle.Render(headerStyle.Render("SEARCH DEMO")+"\n\n"+content) + footer

	case helpScreen:
		content := "📘 Help & Keybindings\n\n"
		content += "Navigation:\n"
		content += "  1 → Menu screen\n"
		content += "  2 → Progress demo\n"
		content += "  3 → Search demo\n"
		content += "  4 → Help screen\n\n"
		content += "Movement:\n"
		content += "  ↑ / k → Move up\n"
		content += "  ↓ / j → Move down\n\n"
		content += "Exit:\n"
		content += "  q or Ctrl+C\n\n"
		content += "This TUI is structured using the Elm Architecture\n"
		content += "(Model → Update → View pattern)."

		return boxStyle.Render(headerStyle.Render("HELP")+"\n\n"+content) + footer
	}

	return ""
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

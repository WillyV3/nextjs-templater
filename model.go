
package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// getBorderedTitleStyle returns a responsive bordered title style
func (m model) getBorderedTitleStyle() lipgloss.Style {
	// Use minimal padding if terminal height is small
	padding := 1
	margin := 1
	if m.height < 20 {
		padding = 0
		margin = 0
	}

	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#006666")).
		Width(m.width - 6).
		Align(lipgloss.Center).
		Padding(padding).
		MarginBottom(margin)
}

func (m model) View() string {
	switch m.step {
	case stepAppName:
		// Create styled components for the app name page
		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86")).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#006666")).
			Width(m.width - 6).
			Align(lipgloss.Center).
			Padding(1).
			MarginBottom(1)

		questionStyle := lipgloss.NewStyle().
			Background(lipgloss.Color("234")).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#006666")).
			Padding(1).
			Width(m.width - 6).
			MarginBottom(1)

		controlsStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Background(lipgloss.Color("236")).
			Padding(0, 1)

		// Combine question and input in one bordered container
		combinedContent := fmt.Sprintf("What would you like to name your app?\n\n%s", m.appName.View())

		return fmt.Sprintf(
			"\n%s\n%s\n%s",
			headerStyle.Render(getAsciiArt()),
			questionStyle.Render(combinedContent),
			controlsStyle.Render("Enter: continue • Ctrl+C: quit"),
		)

	case stepDirectory:
		var b strings.Builder
		b.WriteString(m.getBorderedTitleStyle().Render("Choose Directory"))
		b.WriteString(fmt.Sprintf("\nCurrent: %s\n", m.directory))

		// Show appropriate input based on mode
		if m.creatingNewDir {
			b.WriteString(fmt.Sprintf("Creating new directory in: %s\n", m.directory))
			b.WriteString("Name: " + m.newDirInput.View() + "\n")
		} else if m.searching {
			b.WriteString("Search: " + m.searchInput.View() + "\n")
		}

		// Create separator that fits terminal width
		separator := strings.Repeat("━", m.width-2)
		if len(separator) < 20 {
			separator = "━━━━━━━━━━━━━━━━━━━━"
		}
		b.WriteString(separator + "\n")

		// Show files and directories (only viewport)
		for i := m.viewportStart; i < m.viewportEnd; i++ {
			if i >= len(m.filteredFiles) {
				break
			}
			entry := m.filteredFiles[i]
			line := entry.Name
			if entry.IsDir {
				line = folderStyle.Render(line + "/")
			} else {
				line = fileStyle.Render(line)
			}

			if i == m.cursor && !m.creatingNewDir && !m.searching {
				line = selectedStyle.Render(line)
			} else if i == m.cursor && m.searching {
				line = selectedStyle.Render(line)
			}
			b.WriteString(line + "\n")
		}

		b.WriteString(separator + "\n")

		// Show help based on mode with file count
		if m.creatingNewDir {
			b.WriteString("Type folder name • enter: create & select • esc: cancel")
		} else if m.searching {
			totalFiles := len(m.filteredFiles)
			if totalFiles == 0 {
				b.WriteString("No matches • esc: exit search")
			} else {
				info := fmt.Sprintf("(%d/%d matches) ↑↓: navigate • enter: select • esc: exit search",
					m.cursor+1, totalFiles)

				// Wrap to terminal width, keeping key:action pairs together
				if len(info) > m.width-4 {
					// Split by bullet points to keep key:action pairs together
					parts := strings.Split(info, " • ")
					var lines []string
					var currentLine string

					for i, part := range parts {
						// Add bullet back except for first part
						if i > 0 {
							part = "• " + part
						}

						testLine := currentLine
						if testLine != "" {
							testLine += " "
						}
						testLine += part

						if len(testLine) <= m.width-4 {
							currentLine = testLine
						} else {
							if currentLine != "" {
								lines = append(lines, currentLine)
							}
							currentLine = part
						}
					}
					if currentLine != "" {
						lines = append(lines, currentLine)
					}

					b.WriteString(strings.Join(lines, "\n"))
				} else {
					b.WriteString(info)
				}
			}
		} else {
			info := fmt.Sprintf("(%d/%d) ↑↓/jk: navigate • →: open • ←: up dir • enter: select • s: search • n: new folder • esc: back",
				m.cursor+1, len(m.filteredFiles))

			// Wrap to terminal width, keeping key:action pairs together
			if len(info) > m.width-4 {
				// Split by bullet points to keep key:action pairs together
				parts := strings.Split(info, " • ")
				var lines []string
				var currentLine string

				for i, part := range parts {
					// Add bullet back except for first part
					if i > 0 {
						part = "• " + part
					}

					testLine := currentLine
					if testLine != "" {
						testLine += " "
					}
					testLine += part

					if len(testLine) <= m.width-4 {
						currentLine = testLine
					} else {
						if currentLine != "" {
							lines = append(lines, currentLine)
						}
						currentLine = part
					}
				}
				if currentLine != "" {
					lines = append(lines, currentLine)
				}

				b.WriteString(strings.Join(lines, "\n"))
			} else {
				b.WriteString(info)
			}
		}

		return b.String()

	case stepTheme:
		return fmt.Sprintf("\n%s\n%s\n%s\n\n%s\n\n%s",
			m.getBorderedTitleStyle().Render("Choose Theme"),
			fmt.Sprintf("Name of your Next.js App: %s", m.appName.Value()),
			fmt.Sprintf("Parent Directory: %s", m.directory),
			m.theme.View(),
			"Enter: continue • Esc: back • Ctrl+C: quit",
		)

	case stepAuthChoice:
		return fmt.Sprintf(
			"\n%s\n\n%s\n\n%s",
			m.getBorderedTitleStyle().Render("Choose Authentication"),
			m.authChoice.View(),
			"Enter: continue • Esc: back • Ctrl+C: quit",
		)


	case stepProgress:
		return fmt.Sprintf(
			"\n%s\n\n%s\n\n%s\n\n%s\n%s\n%s\n\n%s\n\n%s",
			m.getBorderedTitleStyle().Render("Creating Your Project"),
			fmt.Sprintf("Name of your Next.js App: %s", m.appName.Value()),
			fmt.Sprintf("Parent Directory: %s", m.directory),
			m.progress.View(),
			m.progress2.View(),
			m.progress3.View(),
			m.outputViewport.View(),
			lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Installation output • Ctrl+C to cancel"),
		)

	case stepComplete:
		status := "Project created successfully!"
		if m.err != nil {
			status = "Error: " + m.err.Error()
		}

		// Create styled components like the beginning page
		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86")).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#006666")).
			Width(m.width - 6).
			Align(lipgloss.Center).
			Padding(1).
			MarginBottom(1)

		messageStyle := lipgloss.NewStyle().
			Background(lipgloss.Color("234")).
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#006666")).
			Padding(1).
			Width(m.width - 6).
			MarginBottom(1)

		controlsStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Background(lipgloss.Color("236")).
			Padding(0, 1)

		// Get ASCII art and create thank you message
		thankYouMessage := fmt.Sprintf("%s\n\n%s", getAsciiArt(), status)

		return fmt.Sprintf(
			"\n%s\n%s\n%s",
			headerStyle.Render(thankYouMessage),
			messageStyle.Render("Your Next.js project has been created with shadcn/ui components!"),
			controlsStyle.Render("Press any key to exit"),
		)
	}

	return ""
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
package main

import (
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/progress"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Set all progress bars to use most of the terminal width with some padding
		progressWidth := msg.Width - 6
		if progressWidth < 20 {
			progressWidth = 20 // Minimum width
		}
		m.progress.Width = progressWidth
		m.progress2.Width = progressWidth
		m.progress3.Width = progressWidth
		// Resize text inputs to be responsive
		inputWidth := msg.Width - 10
		if inputWidth < 20 {
			inputWidth = 20 // Minimum width
		}
		m.appName.Width = inputWidth
		m.newDirInput.Width = inputWidth
		m.searchInput.Width = inputWidth
		// Resize output viewport
		m.outputViewport.Width = msg.Width - 4
		m.outputViewport.Height = msg.Height - 12
		// Calculate list height accounting for bordered titles and controls
		listHeight := msg.Height - 12 // Account for bordered title + margins + controls
		if listHeight < 5 {
			listHeight = 5 // Minimum usable height
		}
		m.theme.SetSize(msg.Width-4, listHeight)
		m.authChoice.SetSize(msg.Width-4, listHeight)
		// Update viewport for file browser
		m.updateViewport()

	case tea.KeyMsg:
		switch m.step {
		case stepAppName:
			switch msg.String() {
			case "enter":
				if m.appName.Value() != "" {
					m.step = stepDirectory
					return m, nil
				}
			case "ctrl+c", "esc":
				return m, tea.Quit
			}
			var cmd tea.Cmd
			m.appName, cmd = m.appName.Update(msg)
			return m, cmd

		case stepDirectory:
			// Handle search mode
			if m.searching {
				switch msg.String() {
				case "enter":
					// Exit search mode and select current directory for theme selection
					m.searching = false
					m.searchInput.Blur()
					m.step = stepTheme
					return m, nil
				case "esc":
					// Exit search mode
					m.searching = false
					m.searchInput.Blur()
					m.searchInput.SetValue("")
					m.filterFiles()
				case "up", "ctrl+k":
					if m.cursor > 0 {
						m.cursor--
						m.updateViewport()
					}
				case "down", "ctrl+j":
					if m.cursor < len(m.filteredFiles)-1 {
						m.cursor++
						m.updateViewport()
					}
				case " ", "tab":
					// Space is reserved for directory selection - ignore in search mode
					return m, nil
				default:
					// Handle search input
					var cmd tea.Cmd
					m.searchInput, cmd = m.searchInput.Update(msg)
					m.filterFiles()
					return m, cmd
				}
			} else if m.creatingNewDir {
				switch msg.String() {
				case "enter":
					dirName := strings.TrimSpace(m.newDirInput.Value())
					if dirName != "" {
						fullPath := filepath.Join(m.directory, dirName)
						err := os.MkdirAll(fullPath, 0755)
						if err == nil {
							// Navigate INTO the newly created directory and go to theme selection
							m.loadDirectory(fullPath)
							m.creatingNewDir = false
							m.newDirInput.Blur()
							m.newDirInput.SetValue("")
							m.step = stepTheme
							return m, nil
						}
					}
				case "esc":
					m.creatingNewDir = false
					m.newDirInput.Blur()
					m.newDirInput.SetValue("")
				default:
					var cmd tea.Cmd
					m.newDirInput, cmd = m.newDirInput.Update(msg)
					return m, cmd
				}
			} else {
				// Normal directory navigation
				switch msg.String() {
				case "enter":
					// Check if we have a valid selection and it's a directory
					if m.cursor < len(m.filteredFiles) {
						entry := m.filteredFiles[m.cursor]
						if entry.IsDir {
							// Navigate into the selected directory AND go to theme selection (like new dir creation)
							m.loadDirectory(entry.Path)
							m.step = stepTheme
							return m, nil
						}
					}
					// If no directory selected or selection is not a directory, proceed to theme selection
					m.step = stepTheme
					return m, nil
				case "right":
					if m.cursor < len(m.filteredFiles) {
						entry := m.filteredFiles[m.cursor]
						if entry.IsDir {
							m.loadDirectory(entry.Path) // Navigate into directory
						}
					}
				case "left":
					// Go up one directory (parent directory)
					parent := filepath.Dir(m.directory)
					if parent != m.directory { // Avoid infinite loop at root
						m.loadDirectory(parent)
					}
				case "n":
					// Create new directory
					m.creatingNewDir = true
					m.newDirInput.Focus()
					return m, nil
				case "s":
					// Enter search mode
					m.searching = true
					m.searchInput.Focus()
					return m, nil
				case "up", "k":
					if m.cursor > 0 {
						m.cursor--
						m.updateViewport()
					}
				case "down", "j":
					if m.cursor < len(m.filteredFiles)-1 {
						m.cursor++
						m.updateViewport()
					}
				case "ctrl+c":
					return m, tea.Quit
				case "esc":
					// Go back to app name step
					m.step = stepAppName
					m.appName.Focus()
					return m, nil
				}
			}

		case stepTheme:
			switch msg.String() {
			case "enter":
				if _, ok := m.theme.SelectedItem().(themeItem); ok {
					m.step = stepAuthChoice
					return m, nil
				}
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				// Go back to directory step
				m.step = stepDirectory
				return m, nil
			}
			var cmd tea.Cmd
			m.theme, cmd = m.theme.Update(msg)
			return m, cmd

		case stepAuthChoice:
			switch msg.String() {
			case "enter":
				if selected, ok := m.authChoice.SelectedItem().(authItem); ok {
					switch selected.id {
					case "clerk":
						m.useClerk = true
						m.useBetterAuth = false
						m.step = stepProgress
						m.isRunning = true
						if themeSelected, ok := m.theme.SelectedItem().(themeItem); ok {
							return m, tea.Batch(
								runScript(m.appName.Value(), m.directory, themeSelected.title, m.useClerk, m.useBetterAuth),
								tickProgress(),
								tickOutputUpdate(),
							)
						}
					case "better-auth":
						m.useClerk = false
						m.useBetterAuth = true
						m.step = stepProgress
						m.isRunning = true
						if themeSelected, ok := m.theme.SelectedItem().(themeItem); ok {
							return m, tea.Batch(
								runScript(m.appName.Value(), m.directory, themeSelected.title, m.useClerk, m.useBetterAuth),
								tickProgress(),
								tickOutputUpdate(),
							)
						}
					case "none":
						m.useClerk = false
						m.useBetterAuth = false
						m.step = stepProgress
						m.isRunning = true
						if themeSelected, ok := m.theme.SelectedItem().(themeItem); ok {
							return m, tea.Batch(
								runScript(m.appName.Value(), m.directory, themeSelected.title, m.useClerk, m.useBetterAuth),
								tickProgress(),
								tickOutputUpdate(),
							)
						}
					}
					return m, nil
				}
			case "ctrl+c":
				return m, tea.Quit
			case "esc":
				// Go back to theme step
				m.step = stepTheme
				return m, nil
			}
			var cmd tea.Cmd
			m.authChoice, cmd = m.authChoice.Update(msg)
			return m, cmd


		case stepProgress:
			if msg.String() == "ctrl+c" {
				return m, tea.Quit
			}

		case stepComplete:
			return m, tea.Quit
		}

	case progressMsg:
		if m.isRunning {
			// Add delay before progress bars start animating (wait for process to actually start)
			if len(liveOutputBuf.String()) > 100 { // Only start animating after some output
				// Update all progress bars with staggered speeds
				cmd1 := m.progress.IncrPercent(progressSpeed1)
				cmd2 := m.progress2.IncrPercent(progressSpeed2)
				cmd3 := m.progress3.IncrPercent(progressSpeed3)
				return m, tea.Batch(cmd1, cmd2, cmd3, tickProgress())
			} else {
				// Just tick without incrementing until we have output
				return m, tickProgress()
			}
		}

	case outputUpdateMsg:
		if m.isRunning {
			// Update viewport with current buffer content, stripping ANSI codes
			cleanOutput := stripAnsiCodes(liveOutputBuf.String())
			m.outputViewport.SetContent(cleanOutput)
			m.outputViewport.GotoBottom() // Auto-scroll to bottom!

			return m, tickOutputUpdate()
		}

	case completeMsg:
		m.isRunning = false
		m.step = stepComplete
		m.output = msg.output
		m.err = msg.err
		m.progress.SetPercent(1.0)
		m.progress2.SetPercent(1.0)
		m.progress3.SetPercent(1.0)
		return m, nil

	case progress.FrameMsg:
		if m.step == stepProgress {
			// Update all three progress bars
			progressModel, cmd1 := m.progress.Update(msg)
			m.progress = progressModel.(progress.Model)

			progressModel2, cmd2 := m.progress2.Update(msg)
			m.progress2 = progressModel2.(progress.Model)

			progressModel3, cmd3 := m.progress3.Update(msg)
			m.progress3 = progressModel3.(progress.Model)

			return m, tea.Batch(cmd1, cmd2, cmd3)
		}
	}

	return m, nil
}
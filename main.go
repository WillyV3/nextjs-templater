package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"nextjs-templater/internal/templates"
)

type step int

const (
	stepAppName step = iota
	stepDirectory
	stepTheme
	stepAuthChoice
	stepProgress
	stepComplete
)

type fileEntry struct {
	Name  string
	Path  string
	IsDir bool
}

type model struct {
	step           step
	appName        textinput.Model
	directory      string
	files          []fileEntry
	filteredFiles  []fileEntry
	cursor         int
	viewportStart  int
	viewportEnd    int
	theme          list.Model
	authChoice     list.Model
	progress       progress.Model
	progress2      progress.Model
	progress3      progress.Model
	outputViewport viewport.Model
	newDirInput    textinput.Model
	searchInput    textinput.Model
	creatingNewDir bool
	searching      bool
	err            error
	output         string
	useClerk       bool
	useBetterAuth  bool
	isRunning      bool

	// Window size
	width  int
	height int
}

var (
	// Shared buffers for real-time output (following Gum's pattern)
	liveOutputBuf bytes.Buffer
	executing     *exec.Cmd

	// Progress bar timing controls
	progressTickInterval = time.Millisecond * 200 // Slower tick rate
	progressSpeed1       = 0.007                  // Faster speed
	progressSpeed2       = 0.004                  // Slightly slower
	progressSpeed3       = 0.003                  // Slowest

	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("86")).
		MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("237"))

	folderStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("33"))

	fileStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#008080"))
)

func initialModel() model {
	// App name input
	ti := textinput.New()
	ti.Placeholder = "Enter your app name (e.g., my-nextjs-app)"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 50

	// New directory input
	newDirInput := textinput.New()
	newDirInput.Placeholder = "Enter new directory name..."
	newDirInput.CharLimit = 100
	newDirInput.Width = 50

	// Search input
	searchInput := textinput.New()
	searchInput.Placeholder = "Search files and directories..."
	searchInput.CharLimit = 100
	searchInput.Width = 50

	// Theme list
	var items []list.Item
	for _, t := range template.NEXTJS_SHADCN_TEMPLATES {
		items = append(items, themeItem{
			id:    t.Id,
			title: t.Title,
			desc:  t.Desc,
		})
	}

	// Create custom delegate with teal highlighting
	themeDelegate := list.NewDefaultDelegate()
	themeDelegate.Styles.SelectedTitle = themeDelegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("86")). // teal
		Bold(true)
	themeDelegate.Styles.SelectedDesc = themeDelegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("86")) // teal

	themeList := list.New(items, themeDelegate, 60, 20)
	themeList.Title = "Choose a theme"
	themeList.SetShowHelp(false)

	// Auth choice list
	var authItems []list.Item
	authItems = append(authItems, authItem{
		id:    "clerk",
		title: "Clerk",
		desc:  "Complete authentication platform with social logins, MFA, and user management",
	})
	authItems = append(authItems, authItem{
		id:    "better-auth",
		title: "Better Auth",
		desc:  "Lightweight auth library with Kysely + SQLite integration",
	})
	authItems = append(authItems, authItem{
		id:    "none",
		title: "No Authentication",
		desc:  "Skip authentication setup",
	})

	// Create custom delegate with teal highlighting
	authDelegate := list.NewDefaultDelegate()
	authDelegate.Styles.SelectedTitle = authDelegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("86")). // teal
		Bold(true)
	authDelegate.Styles.SelectedDesc = authDelegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("86")) // teal

	authList := list.New(authItems, authDelegate, 60, 20)
	authList.Title = "Choose authentication"
	authList.SetShowHelp(false)

	// Three stacked progress bars - only the top one shows percentage
	prog := progress.New(
		progress.WithScaledGradient("#FF6B6B", "#4ECDC4"),
		progress.WithSpringOptions(1.0, 1.0),
	)
	prog2 := progress.New(
		progress.WithScaledGradient("#A8E6CF", "#FFD93D"),
		progress.WithSpringOptions(1.0, 1.0),
		progress.WithoutPercentage(),
	)
	prog3 := progress.New(
		progress.WithScaledGradient("#FF8A80", "#B39DDB"),
		progress.WithSpringOptions(1.0, 1.0),
		progress.WithoutPercentage(),
	)
	// Initial width (will be updated on window size changes)
	prog.Width = 60
	prog2.Width = 60
	prog3.Width = 60

	// Output viewport
	vp := viewport.New(50, 10)
	vp.Style = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#006666")).
		Padding(0, 1)

	// Start in home directory
	homeDir, _ := os.UserHomeDir()

	m := model{
		step:           stepAppName,
		appName:        ti,
		directory:      homeDir,
		theme:          themeList,
		authChoice:     authList,
		outputViewport: vp,
		newDirInput:    newDirInput,
		searchInput:    searchInput,
		progress:       prog,
		progress2:      prog2,
		progress3:      prog3,
		width:          80,
		height:         24,
	}

	m.loadDirectory(homeDir)
	return m
}

func (m *model) loadDirectory(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	m.files = []fileEntry{}
	m.directory = path

	// Add parent directory option if not root
	if path != "/" && path != filepath.Dir(path) {
		parent := filepath.Dir(path)
		m.files = append(m.files, fileEntry{
			Name:  "..",
			Path:  parent,
			IsDir: true,
		})
	}

	// Sort directories first, then files
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].IsDir() != entries[j].IsDir() {
			return entries[i].IsDir()
		}
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		// Skip hidden files
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		m.files = append(m.files, fileEntry{
			Name:  entry.Name(),
			Path:  filepath.Join(path, entry.Name()),
			IsDir: entry.IsDir(),
		})
	}

	m.filteredFiles = m.files
	m.cursor = 0
	m.searching = false // Disable search when entering new folder
	m.searchInput.Blur()
	m.searchInput.SetValue("")
	m.updateViewport()
}

// Fuzzy search function
func fuzzyMatch(pattern, text string) bool {
	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)

	if pattern == "" {
		return true
	}

	patternIdx := 0
	for _, char := range text {
		if patternIdx < len(pattern) && rune(pattern[patternIdx]) == char {
			patternIdx++
		}
	}

	return patternIdx == len(pattern)
}

func (m *model) filterFiles() {
	query := m.searchInput.Value()
	if query == "" {
		m.filteredFiles = m.files
	} else {
		m.filteredFiles = []fileEntry{}
		for _, file := range m.files {
			if fuzzyMatch(query, file.Name) {
				m.filteredFiles = append(m.filteredFiles, file)
			}
		}
	}

	// Reset cursor and update viewport
	m.cursor = 0
	m.updateViewport()
}

func (m *model) updateViewport() {
	if len(m.filteredFiles) == 0 {
		return
	}

	// Calculate available height accounting for bordered title and other elements
	// Bordered title takes ~4-5 lines, plus margins and controls
	availableHeight := m.height - 12 // More conservative for bordered title
	if m.creatingNewDir || m.searching {
		availableHeight -= 2 // Extra space for input
	}
	// Ensure minimum usable height
	if availableHeight < 3 {
		availableHeight = 3
	}

	// Ensure cursor is within bounds
	if m.cursor >= len(m.filteredFiles) {
		m.cursor = len(m.filteredFiles) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}

	// Update viewport to show cursor
	if m.cursor < m.viewportStart {
		m.viewportStart = m.cursor
	}
	if m.cursor >= m.viewportStart+availableHeight {
		m.viewportStart = m.cursor - availableHeight + 1
	}

	if m.viewportStart < 0 {
		m.viewportStart = 0
	}

	m.viewportEnd = m.viewportStart + availableHeight
	if m.viewportEnd > len(m.filteredFiles) {
		m.viewportEnd = len(m.filteredFiles)
	}
}

type themeItem struct {
	id    int
	title string
	desc  string
}

func (t themeItem) Title() string       { return t.title }
func (t themeItem) Description() string { return t.desc }
func (t themeItem) FilterValue() string { return t.title }

type authItem struct {
	id   string
	title string
	desc  string
}

func (a authItem) Title() string       { return a.title }
func (a authItem) Description() string { return a.desc }
func (a authItem) FilterValue() string { return a.title }

// stripAnsiCodes removes ANSI escape sequences from text
func stripAnsiCodes(input string) string {
	// Regular expression to match ANSI escape sequences
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return ansiRegex.ReplaceAllString(input, "")
}

// getAsciiArt reads and returns the ASCII art from file with description
func getAsciiArt() string {
	content, err := os.ReadFile("asciiArt.txt")
	if err != nil {
		// Fallback to text if file not found
		return "Next.js + Shadcn/ui Project Generator"
	}

	// Add description below the ASCII art
	developerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("23")). // dark teal
		Align(lipgloss.Center)
	developerText := developerStyle.Render("Developer: WillyV3\nMore from Willy: www.Willyv3.com")

	fullContent := string(content) + "\n" +
		"Scaffold your NextJS App with Complete Shadcn\n" +
		"Component Kit and Choice of Auth\n\n" +
		developerText

	return fullContent
}

type progressMsg float64

type outputUpdateMsg struct{}

type completeMsg struct {
	output string
	err    error
}

func runScript(appName, directory, theme string, useClerk, useBetterAuth bool) tea.Cmd {
	return func() tea.Msg {
		// Find the theme template
		var selectedTemplate template.Item
		for _, t := range template.NEXTJS_SHADCN_TEMPLATES {
			if t.Title == theme {
				selectedTemplate = t
				break
			}
		}

		// Get the absolute path to the script
		execPath, _ := os.Executable()
		scriptDir := filepath.Dir(execPath)
		scriptPath := filepath.Join(scriptDir, "create-nextjs-shadcn.sh")

		// If running with go run, the script is in the current working directory
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			wd, _ := os.Getwd()
			scriptPath = filepath.Join(wd, "create-nextjs-shadcn.sh")
		}

		// Replace placeholders in command args and add Clerk parameter
		args := strings.Replace(selectedTemplate.CommandArgs, "create-nextjs-shadcn.sh", scriptPath, 1)
		args = strings.Replace(args, "$(pwd)", directory, -1)
		args = strings.Replace(args, selectedTemplate.Title, appName, -1)

		// Add auth parameters
		if useClerk {
			args += " true false"
		} else if useBetterAuth {
			args += " false true"
		} else {
			args += " false false"
		}

		// Clear the live output buffer
		liveOutputBuf.Reset()

		var outputBuffer strings.Builder

		// Log the command being executed
		initialMsg := fmt.Sprintf("Executing: %s %s\nWorking directory: %s\n\n", selectedTemplate.Command, args, directory)
		outputBuffer.WriteString(initialMsg)
		liveOutputBuf.WriteString(initialMsg)

		executing = exec.Command(selectedTemplate.Command, strings.Fields(args)...)
		executing.Dir = directory

		// Use MultiWriter to write to both the final output buffer and the live buffer
		executing.Stdout = io.MultiWriter(&outputBuffer, &liveOutputBuf)
		executing.Stderr = io.MultiWriter(&outputBuffer, &liveOutputBuf)

		// Run the command
		err := executing.Run()

		return completeMsg{
			output: outputBuffer.String(),
			err:    err,
		}
	}
}

func tickProgress() tea.Cmd {
	return tea.Tick(progressTickInterval, func(t time.Time) tea.Msg {
		return progressMsg(progressSpeed1)
	})
}

func tickOutputUpdate() tea.Cmd {
	return tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
		return outputUpdateMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

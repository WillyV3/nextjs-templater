<h1 align="center">
  <br/>
  Next.js Templater
</h1>
<p align="center">A beautiful TUI to scaffold Next.js projects with shadcn/ui components and authentication.</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    &nbsp;
    <img src="https://img.shields.io/badge/Bubble_Tea-TUI-success?style=for-the-badge&logo=none" alt="bubble tea" />
    &nbsp;
    <img src="https://img.shields.io/badge/license-mit-green?style=for-the-badge&logo=none" alt="license" />
</p>

## Features

- **Beautiful TUI** - Built with Bubble Tea for an elegant terminal experience
- **Responsive Design** - Adapts to your terminal size automatically
- **Interactive File Browser** - Navigate directories with fuzzy search
- **Multiple Themes** - Choose from various shadcn/ui template configurations
- **Authentication Options** - Clerk, Better Auth, or no authentication
- **Real-time Progress** - Live installation output with animated progress bars
- **ASCII Art** - Beautiful branding throughout the interface

## Prerequisites

- Go 1.21+
- Node.js and npm (for the generated projects)

## Quick Start

```bash
# Clone and build
git clone https://github.com/WillyV3/nextjs-templater
cd nextjs-templater
go build -o nextjs-templater

# Run the TUI
./nextjs-templater
```

## How It Works

1. **Enter App Name** - Choose your Next.js app name
2. **Select Directory** - Browse and select where to create your project
3. **Choose Theme** - Pick from available shadcn/ui templates
4. **Select Auth** - Choose Clerk, Better Auth, or skip authentication
5. **Watch Progress** - Real-time installation with live output

## Available Templates

- **Default** - Basic Next.js with shadcn/ui
- **Dashboard** - Admin dashboard template
- **Landing Page** - Marketing site template
- **E-commerce** - Online store template
- **Blog** - Content management template

## Authentication Options

- **Clerk** - Complete authentication platform with social logins and MFA
- **Better Auth** - Lightweight auth library with Kysely + SQLite integration
- **None** - Skip authentication setup

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Developer

**WillyV3**
More from Willy: [www.Willyv3.com](https://www.willyv3.com)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
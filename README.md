<h1 align="center">
  <br/>
  Next.js Templater
</h1>
<p align="center">A beautiful TUI to scaffold Next.js projects with shadcn/ui components and authentication.</p>

<p align="center">
<img width="580" height="440" alt="Screenshot 2025-09-25 at 11 43 37 PM" src="https://github.com/user-attachments/assets/03a3d73f-624d-4b68-9866-ca38afd17acc" />
</p>

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

<p align="center">
<img width="596" height="406" alt="Screenshot 2025-09-25 at 11 44 28 PM" src="https://github.com/user-attachments/assets/d4fcaf80-7ec6-4c57-a8d8-21040a6d344e" />
</p>

- **Multiple Themes** - Choose from various shadcn/ui template configurations

<p align="center">
<img width="586" height="530" alt="Screenshot 2025-09-25 at 11 45 04 PM" src="https://github.com/user-attachments/assets/855d9852-2c21-43ab-9c13-7588453c6bb4" />
</p>

- **Authentication Options** - Clerk, Better Auth, or no authentication

<p align="center">
<img width="582" height="341" alt="Screenshot 2025-09-25 at 11 45 42 PM" src="https://github.com/user-attachments/assets/005f7145-2754-4b38-b1bc-f43ebce5efb1" />
</p>

**Installing...**

<p align="center">
<img width="599" height="533" alt="Screenshot 2025-09-25 at 11 46 28 PM" src="https://github.com/user-attachments/assets/31f688af-00db-4341-a9d3-8bceff24bb69" />
</p>

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
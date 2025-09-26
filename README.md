<h1 align="center">
  <br/>
  Next.js Templater
</h1>
<p align="center">TUI to scaffold Next.js projects with shadcn/ui components and authentication.</p>

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

- **TUI** - Built with Bubble Tea framework
- **Responsive Design** - Adapts to terminal size
- **File Browser** - Navigate directories with search

<p align="center">
  <img width="736" height="652" alt="Screenshot 2025-09-26 at 7 26 05 PM" src="https://github.com/user-attachments/assets/7af5ff2d-835c-4773-8e29-88d8752140ec" />
</p>

- **Themes** - Choose from shadcn/ui template configurations

<p align="center">
  <img width="733" height="664" alt="Screenshot 2025-09-26 at 7 26 49 PM" src="https://github.com/user-attachments/assets/2cf1e3a4-41ba-4bac-b3cd-4408b876f69d" />
</p>

- **Authentication** - Clerk, Better Auth, or no authentication (Maybe will add more later, any good ideas?)

<p align="center">
  <img width="713" height="452" alt="Screenshot 2025-09-26 at 7 27 39 PM" src="https://github.com/user-attachments/assets/941f3e9a-1a58-4a39-b830-2de89371dbb7" />
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

1. **Enter App Name** - Enter Next.js app name
2. **Select Directory** - Select directory to create project
3. **Choose Theme** - Select from shadcn/ui templates
4. **Select Auth** - Choose Clerk, Better Auth, or skip authentication
5. **Monitor Progress** - Installation with output

## Templates

- **Default** - Next.js with shadcn/ui
- **Dashboard** - Dashboard template
- **Landing Page** - Site template
- **E-commerce** - Store template
- **Blog** - Blog template

## Authentication

- **Clerk** - Authentication platform with social logins and MFA
- **Better Auth** - Auth library with Kysely + SQLite integration
- **None** - Skip authentication setup

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Developer

**WillyV3**
More from Willy: [www.Willyv3.com](https://www.willyv3.com)

## Contributing

Submit Pull Requests for contributions.

## License

This project is licensed under the MIT License.

#!/bin/bash

# NextJS + Shadcn Setup with Tweakcn Themes
# Using Next.js 15, Tailwind v4, and npm

set -e

# Check for required dependencies
check_dependencies() {
    local missing_deps=()

    # Check for node
    if ! command -v node >/dev/null 2>&1; then
        missing_deps+=("node")
    fi

    # Check for npm
    if ! command -v npm >/dev/null 2>&1; then
        missing_deps+=("npm")
    fi

    if [ ${#missing_deps[@]} -gt 0 ]; then
        echo "❌ Missing required dependencies: ${missing_deps[*]}"
        echo "Please install Node.js and npm before running this script"
        echo ""
        echo "Installation options:"
        echo "  - Ubuntu/Debian: sudo apt install nodejs npm"
        echo "  - CentOS/RHEL: sudo yum install nodejs npm"
        echo "  - macOS: brew install node"
        echo "  - Or download from: https://nodejs.org/"
        exit 127
    fi
}

# Try to use Node.js 20 if NVM is available, otherwise use system Node.js
setup_node() {
    echo "🔍 Checking Node.js setup..."

    # Try NVM first (common on macOS/Linux development environments)
    if [ -s "$HOME/.nvm/nvm.sh" ]; then
        echo "📦 Found NVM, attempting to use Node.js 20..."
        export NVM_DIR="$HOME/.nvm"
        \. "$NVM_DIR/nvm.sh"

        # Try to use Node 20, but don't fail if it's not available
        if nvm use 20 2>/dev/null; then
            echo "✅ Using Node.js 20 via NVM"
        else
            echo "⚠️  Node.js 20 not available via NVM, using system Node.js"
        fi
    else
        echo "📦 NVM not found, using system Node.js"
    fi

    # Show current Node.js version
    echo "📍 Node.js version: $(node --version)"
    echo "📍 npm version: $(npm --version)"
    echo ""
}

# Run dependency checks
check_dependencies
setup_node

PROJECT_NAME="${1:-}"
PROJECT_PATH="${2:-$(pwd)}"
THEME="${3:-}"
USE_CLERK="${4:-false}"
USE_BETTER_AUTH="${5:-false}"

if [ -z "$PROJECT_NAME" ]; then
    echo "Usage: $0 <project-name> [path] [theme]"
    echo ""
    echo "Available themes from tweakcn.com:"
    echo "  modern-minimal, violet-bloom, t3-chat, mocha-mousse, amethyst-haze,"
    echo "  doom-64, kodama-grove, cosmic-night, quantum-rose, bold-tech,"
    echo "  elegant-luxury, amber-minimal, neo-brutalism, solar-dusk, pastel-dreams,"
    echo "  clean-slate, ocean-breeze, retro-arcade, midnight-bloom, northern-lights,"
    echo "  vintage-paper, sunset-horizon, starry-night, soft-pop"
    exit 1
fi

PROJECT_NAME=$(echo "$PROJECT_NAME" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')
FULL_PATH="$PROJECT_PATH/$PROJECT_NAME"

echo "Creating: $PROJECT_NAME at $FULL_PATH"

# Create Next.js 15 app with Tailwind v4
echo "🚀 Creating Next.js app..."
cd "$PROJECT_PATH"

if ! echo "n" | npx create-next-app@latest "$PROJECT_NAME" \
    --typescript \
    --tailwind \
    --eslint \
    --app \
    --src-dir \
    --turbopack; then
    echo "❌ Failed to create Next.js app. Check your internet connection and try again."
    exit 1
fi

if [ ! -d "$FULL_PATH" ]; then
    echo "❌ Project directory was not created: $FULL_PATH"
    exit 1
fi

cd "$FULL_PATH"
echo "✅ Next.js app created successfully"

# Init shadcn and apply theme if specified
if [ ! -z "$THEME" ]; then
    echo "🎨 Initializing shadcn with $THEME theme..."
    # Run the theme command twice - first time inits shadcn, second applies theme
    if ! yes | npx shadcn@latest add "https://tweakcn.com/r/themes/${THEME}.json"; then
        echo "❌ Failed to initialize shadcn with theme. Falling back to default..."
        printf "1\n1\n" | npx shadcn@latest init || { echo "❌ Failed to initialize shadcn"; exit 1; }
    else
        echo "🎨 Applying theme configuration..."
        yes | npx shadcn@latest add "https://tweakcn.com/r/themes/${THEME}.json" || echo "⚠️  Theme reapplication failed, but continuing..."
    fi
else
    echo "🎨 Initializing shadcn with default theme..."
    if ! printf "1\n1\n" | npx shadcn@latest init; then
        echo "❌ Failed to initialize shadcn"
        exit 1
    fi
fi

# Add all components with auto-yes
echo "📦 Installing all shadcn components..."
if ! yes | npx shadcn@latest add --all; then
    echo "⚠️  Some components may have failed to install, but continuing..."
fi

# Add authentication if requested
if [ "$USE_CLERK" = "true" ]; then
    echo "Installing Clerk authentication quickstart..."
    yes | npx shadcn@latest add @clerk/nextjs-quickstart
elif [ "$USE_BETTER_AUTH" = "true" ]; then
    echo "Installing Better Auth with SQLite..."

    # Install Better Auth dependencies
    npm install better-auth better-sqlite3
    npm install --save-dev @types/better-sqlite3

    # Generate secret
    echo "Generating Better Auth secret..."
    AUTH_SECRET=$(npx @better-auth/cli@latest secret)

    # Create .env.local file
    echo "Creating .env.local file..."
    cat > .env.local << EOF
# Better Auth Configuration
BETTER_AUTH_SECRET=$AUTH_SECRET
BETTER_AUTH_URL=http://localhost:3000

# Database
DATABASE_URL=sqlite:./auth.db

# GitHub OAuth (optional)
# GITHUB_CLIENT_ID=your_github_client_id
# GITHUB_CLIENT_SECRET=your_github_client_secret
EOF

    # Create Better Auth config files manually (idiomatic approach)
    echo "Creating Better Auth configuration..."
    mkdir -p lib

    cat > lib/auth.ts << 'EOF'
import { betterAuth } from "better-auth"

export const auth = betterAuth({
  database: {
    provider: "sqlite",
    url: "./auth.db"
  },
  emailAndPassword: {
    enabled: true,
  },
})
EOF

    cat > lib/auth-client.ts << 'EOF'
import { createAuthClient } from "better-auth/react" // make sure to import from better-auth/react
export const authClient = createAuthClient({
    //you can pass client configuration here
})
EOF

    # Create auth API route (after init, using auth.handler)
    echo "Creating auth API route..."
    mkdir -p src/app/api/auth/[...all]
    cat > 'src/app/api/auth/[...all]/route.ts' << 'EOF'
import { auth } from "@/lib/auth";
import { toNextJsHandler } from "better-auth/next-js";
export const { GET, POST } = toNextJsHandler(auth.handler);
EOF

    echo "Better Auth setup complete (database will be created on first run)"
fi

# Add useful packages
echo "Adding additional packages..."
npm install lucide-react next-themes

# Setup Claude (if available)
if command -v claudenew &> /dev/null; then
    echo "Setting up Claude directory..."
    claudenew
fi

echo "✅ Done. Project at: $FULL_PATH"
if [ ! -z "$THEME" ]; then
    echo "   Theme applied: $THEME"
fi
if [ "$USE_CLERK" = "true" ]; then
    echo "   Clerk authentication: Installed"
    echo "   Don't forget to set your CLERK_SECRET_KEY and NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY"
elif [ "$USE_BETTER_AUTH" = "true" ]; then
    echo "   Better Auth: Installed with Kysely + SQLite"
    echo "   Database: SQLite (./auth.db created and migrated)"
    echo "   Config: lib/auth.ts and lib/auth-client.ts created"
    echo "   Environment: .env.local created with secrets"
    echo "   Add your GitHub OAuth credentials to .env.local for social login"
fi
echo "Run: cd $FULL_PATH && npm run dev"
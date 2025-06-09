# EnvGuard 🛡️

[![npm version](https://badge.fury.io/js/envguard.svg)](https://www.npmjs.com/package/envguard)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A powerful CLI tool that validates `.env` files against `.env.example` files and manages multiple environment configurations with ease.

## ✨ Features

✅ **Environment Validation**: Compare `.env` files against `.env.example` templates  
🌍 **Multi-Environment Management**: Manage multiple environments using `.envguard/` directory  
🔄 **Auto-Sync**: Automatically saves `.env` changes to the active environment  
🎯 **Easy Switching**: Quickly switch between development, staging, production, etc.  
🚀 **Fast & Reliable**: Built with Go for speed and reliability  
📊 **Detailed Reports**: Clear summaries showing missing, extra, and valid variables  

## 📦 Installation

### Global Installation (Recommended)

```bash
npm install -g envguard
```

### Local Installation

```bash
npm install envguard
npx envguard --help
```

## 🚀 Quick Start

```bash
# 1. Create your first environment
envguard create -e development --from-current

# 2. Create additional environments
envguard create -e staging
envguard create -e production --from-current

# 3. List all environments
envguard list

# 4. Use environments - edit .env normally!
envguard use production
# Edit .env in your editor...
envguard use development  # Previous changes auto-saved!

# 5. Check current environment status
envguard status

# 6. Validate current environment
envguard

# 7. Clean up unused environments
envguard delete -e old-staging
```

## 📋 Available Commands

| Command | Description | Auto-Sync | Example |
|---------|-------------|-----------|---------|
| `envguard` | Validate .env against .env.example | ✅ | `envguard` |
| `envguard use <env>` | Use environment + track | ✅ | `envguard use production` |
| `envguard status` | Show active environment | ✅ | `envguard status` |
| `envguard create -e <env>` | Create new environment | ✅ | `envguard create -e staging` |
| `envguard list` | List all environments | ✅ | `envguard list` |
| `envguard delete -e <env>` | Delete environment | ❌ | `envguard delete -e old-env` |

## 🎯 Example Output

```
🔍 EnvGuard - Environment File Validator

📊 Validation Results:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📁 Comparing: .env ↔ .env.example

✅ Variables found in both files (8):
   ✓ DATABASE_URL
   ✓ API_KEY
   ✓ PORT
   ✓ JWT_SECRET

⚠️  Missing variables in .env (2):
   • EMAIL_SERVICE_KEY
   • WEBHOOK_SECRET

📈 Summary:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📊 ✅ 8 variables OK • ⚠️ 2 missing
```

## 🌍 Platform Support

EnvGuard supports the following platforms:

- **macOS** (Intel & Apple Silicon)
- **Linux** (x64)
- **Windows** (x64)

## 📁 How It Works

- **`.envguard/` Directory**: All environment files are stored in this hidden directory
- **Active Environment**: The root `.env` file is always your active environment
- **Auto-Sync**: Any changes to `.env` are automatically saved to the active environment
- **Environment Usage**: `envguard use` copies the selected environment to `.env`
- **Active Tracking**: `envguard use` tracks the current environment in `.envguard/.active`
- **Status Checking**: `envguard status` shows which environment is currently active
- **Validation**: Always validates the active `.env` against `.env.example`
- **Isolation**: Each environment is completely isolated and independent

## 🛠️ Development Workflow

```bash
# Set up your project environments
envguard create -e development --from-current
envguard create -e staging 
envguard create -e production --from-current

# Work with different environments - edit .env normally!
envguard use development     # Switch to development
# Edit .env in your favorite editor...
envguard status             # Auto-saves changes, shows current environment
envguard                    # Auto-saves changes, validates environment

envguard use production     # Auto-saves dev changes, switches to production
envguard status             # Confirm you're in production
envguard                    # Validate before deployment
```

## 🐛 Troubleshooting

### Permission Issues (macOS/Linux)
```bash
# If you get permission errors, run:
chmod +x $(which envguard)
```

### Reinstallation
```bash
# If something goes wrong, try:
npm uninstall -g envguard
npm install -g envguard
```

### Platform Not Supported
If you get an "Unsupported platform" error, EnvGuard currently supports:
- macOS (darwin)
- Linux 
- Windows (win32)

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 🔗 Links

- [Source Code](https://github.com/crabest/envguard)
- [Issues](https://github.com/crabest/envguard/issues)
- [npm Package](https://www.npmjs.com/package/envguard) 
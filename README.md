# EnvGuard 🛡️

A powerful CLI tool that validates `.env` files against `.env.example` files, ensuring your environment configuration is complete and properly maintained.

## Features

✅ **Comprehensive Validation**: Compare `.env` files against `.env.example` templates  
🌍 **Multi-Environment Management**: Manage multiple environments using `.envguard/` directory  
🔄 **Auto-Sync**: Automatically saves `.env` changes to the active environment  
🎯 **Environment Switching**: Quickly switch between development, staging, production, etc.  
🚀 **Fast & Reliable**: Built with Go for speed and reliability  
📊 **Detailed Reports**: Clear summaries showing missing, extra, and valid variables  
🔧 **Flexible Configuration**: Custom file paths via command-line flags  

## Installation

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

### Build from Source

```bash
git clone https://github.com/crabest/envguard
cd envguard
go mod tidy
go build -o envguard
```

### Install Globally

```bash
go install
```

## Usage

### Basic Usage

```bash
# Validate default files (.env against .env.example)
./envguard

# Or if installed globally
envguard
```

### Environment Management

```bash
# Create a new environment
envguard create -e development
envguard create -e production --from-current

# List all environments
envguard list

# Use an environment (switch + tracking)
envguard use production
envguard use development

# Check current environment status
envguard status

# Delete an environment
envguard delete -e old-config
envguard delete -e test --no-confirm
```

### Custom File Paths

```bash
# Use custom file paths for validation
envguard --env .env.production --example .env.example

# Short flags
envguard -e .env.staging -x .env.template
```

### Help

```bash
envguard --help
```

## Example Output

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
   ✓ REDIS_URL
   ✓ LOG_LEVEL
   ✓ APP_NAME
   ✓ DEBUG

⚠️  Missing variables in .env (2):
   • EMAIL_SERVICE_KEY
   • WEBHOOK_SECRET

❌ Extra variables in .env not found in .env.example (1):
   • DEPRECATED_CONFIG

📈 Summary:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
❌ Your .env file has missing and extra variables.

📊 ✅ 8 variables OK • ⚠️ 2 missing • ❌ 1 unused
```

## Environment Management Workflow

EnvGuard supports managing multiple environment configurations using a hidden `.envguard/` directory with **automatic synchronization**:

### Quick Start with Environments

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

### How It Works

- **`.envguard/` Directory**: All environment files are stored in this hidden directory
- **Active Environment**: The root `.env` file is always your active environment  
- **Auto-Sync**: Any changes to `.env` are automatically saved to the active environment
- **Environment Usage**: `envguard use` copies the selected environment to `.env`
- **Active Tracking**: `envguard use` tracks the current environment in `.envguard/.active`
- **Status Checking**: `envguard status` shows which environment is currently active
- **Validation**: Always validates the active `.env` against `.env.example`
- **Isolation**: Each environment is completely isolated and independent

### Command Comparison

| Command | Description | Auto-Sync | Use Case |
|---------|-------------|-----------|----------|
| `envguard use <env>` | Use environment + track | ✅ Before switch | Normal workflow |
| `envguard status` | Show active environment | ✅ | Check current state |
| `envguard list` | List all environments | ✅ | See available options |
| `envguard` | Validate .env | ✅ Before validation | Check environment |


## Development

### Run Tests

```bash
go test ./...
```

### Build

```bash
go build -o envguard
```

### Format Code

```bash
go fmt ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `go fmt ./...` and `go vet ./...`
6. Submit a pull request

## License

MIT License - see LICENSE file for details 
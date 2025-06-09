# EnvGuard npm Packaging Guide 📦

This guide explains how to package and distribute EnvGuard via npm for JavaScript/Node.js developers.

## 📁 Package Structure

```
envguard-npm/
├── bin/                           # Cross-platform binaries
│   ├── envguard-darwin           # macOS binary (Intel)
│   ├── envguard-linux            # Linux binary (x64)
│   └── envguard-win.exe          # Windows binary (x64)
├── index.js                      # Main entry script
├── package.json                  # npm package configuration
├── postinstall.js               # Post-installation script
├── README.md                     # npm package documentation
├── test-install.sh              # Test script for package
└── envguard-1.0.0.tgz           # Generated package tarball
```

## 🔧 Build Process

### 1. Build Cross-Platform Binaries

```bash
# Build all binaries for npm distribution
make build-npm
```

This creates:
- `envguard-darwin` (macOS x64)
- `envguard-linux` (Linux x64) 
- `envguard-win.exe` (Windows x64)

All binaries are automatically given executable permissions on Unix systems.

### 2. Package Creation

```bash
# Create the npm package
make package-npm
```

This will:
1. Build the binaries
2. Create the npm package tarball
3. Show publishing instructions

## 📋 Key Components

### `index.js` - Main Entry Point

- **Platform Detection**: Automatically detects OS (darwin/linux/win32)
- **Binary Selection**: Chooses correct binary for current platform
- **Error Handling**: Provides helpful error messages for unsupported platforms
- **Process Management**: Handles signals and exit codes properly
- **Argument Passing**: Forwards all CLI arguments to the native binary

### `package.json` - npm Configuration

```json
{
  "name": "envguard",
  "bin": {
    "envguard": "./index.js"
  },
  "os": ["darwin", "linux", "win32"],
  "cpu": ["x64"],
  "preferGlobal": true
}
```

Key features:
- **Binary Definition**: Maps `envguard` command to `index.js`
- **Platform Restrictions**: Only allows supported platforms
- **Global Installation**: Optimized for global installation

### `postinstall.js` - Installation Script

- **Permission Setting**: Automatically sets execute permissions on Unix
- **Welcome Message**: Provides getting started instructions
- **Error Handling**: Graceful handling of permission issues

## 🚀 Usage After Installation

Once published, users can install and use EnvGuard like this:

```bash
# Global installation
npm install -g envguard

# Use the CLI
envguard --help
envguard create -e development
envguard use production
envguard status
```

## 🧪 Testing

### Local Testing

```bash
# Test the wrapper directly
cd envguard-npm
node index.js --help

# Test the postinstall script
node postinstall.js

# Test package creation
npm pack
```

### Installation Testing

```bash
# Test local installation
cd envguard-npm
npm install -g .

# Test the installed command
envguard --help

# Uninstall when done testing
npm uninstall -g envguard
```

### Automated Testing

```bash
# Run the test script
cd envguard-npm
./test-install.sh
```

## 📤 Publishing

### Prerequisites

1. **npm Account**: Create account at https://npmjs.com
2. **Login**: `npm login`
3. **Verify**: `npm whoami`

### Publishing Steps

```bash
cd envguard-npm

# 1. Final check
npm pack
tar -tzf envguard-1.0.0.tgz

# 2. Publish
npm publish

# 3. Verify
npm info envguard
```

### Version Management

```bash
# Update version
npm version patch   # 1.0.0 -> 1.0.1
npm version minor   # 1.0.0 -> 1.1.0
npm version major   # 1.0.0 -> 2.0.0

# Publish new version
npm publish
```

## 🔍 Platform Support

| Platform | Architecture | Binary Name | Status |
|----------|-------------|-------------|--------|
| macOS | x64 | `envguard-darwin` | ✅ Supported |
| Linux | x64 | `envguard-linux` | ✅ Supported |
| Windows | x64 | `envguard-win.exe` | ✅ Supported |

### Unsupported Platforms

When users try to install on unsupported platforms, they receive:

```
❌ Error: Unsupported platform 'freebsd'
EnvGuard supports: macOS (darwin), Linux, and Windows
Your platform: freebsd (x64)
```

## 🐛 Troubleshooting

### Permission Issues

**Problem**: `EACCES` error on Unix systems
**Solution**: 
```bash
chmod +x $(which envguard)
```

### Binary Not Found

**Problem**: Binary missing from installation
**Solution**:
```bash
npm uninstall -g envguard
npm install -g envguard
```

### Platform Detection Issues

**Problem**: Wrong binary selected
**Solution**: Check `os.platform()` and `os.arch()` values

## 📊 Package Size

- **Packed Size**: ~4.4 MB
- **Unpacked Size**: ~10.9 MB
- **Files Included**: 7 files (binaries + scripts)

The size is primarily due to the Go binaries, but this ensures:
- **No Dependencies**: No need for Go to be installed
- **Fast Execution**: Native performance
- **Offline Usage**: Works without internet connection

## 🔗 Resources

- **npm Documentation**: https://docs.npmjs.com/
- **Package Best Practices**: https://docs.npmjs.com/misc/developers
- **Semantic Versioning**: https://semver.org/
- **Cross-Platform Packaging**: https://nodejs.org/api/os.html 
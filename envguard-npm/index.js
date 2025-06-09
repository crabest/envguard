#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');
const os = require('os');

function getBinaryPath() {
  const platform = os.platform();
  const arch = os.arch();
  
  let binaryName;
  
  switch (platform) {
    case 'darwin':
      binaryName = 'envguard-darwin';
      break;
    case 'linux':
      binaryName = 'envguard-linux';
      break;
    case 'win32':
      binaryName = 'envguard-win.exe';
      break;
    default:
      console.error(`❌ Error: Unsupported platform '${platform}'`);
      console.error(`EnvGuard supports: macOS (darwin), Linux, and Windows`);
      console.error(`Your platform: ${platform} (${arch})`);
      process.exit(1);
  }
  
  const binaryPath = path.join(__dirname, 'bin', binaryName);
  
  // Check if binary exists
  if (!fs.existsSync(binaryPath)) {
    console.error(`❌ Error: Binary not found at ${binaryPath}`);
    console.error(`This may be a corrupted installation. Try reinstalling with:`);
    console.error(`npm uninstall -g envguard && npm install -g envguard`);
    process.exit(1);
  }
  
  return binaryPath;
}

function runEnvGuard() {
  const binaryPath = getBinaryPath();
  const args = process.argv.slice(2);
  
  // Spawn the binary with all arguments
  const child = spawn(binaryPath, args, {
    stdio: 'inherit',
    windowsHide: false
  });
  
  // Handle process exit
  child.on('close', (code) => {
    process.exit(code);
  });
  
  child.on('error', (err) => {
    if (err.code === 'ENOENT') {
      console.error(`❌ Error: Could not execute ${binaryPath}`);
      console.error(`Make sure the binary has execute permissions.`);
    } else if (err.code === 'EACCES') {
      console.error(`❌ Error: Permission denied executing ${binaryPath}`);
      console.error(`Try running: chmod +x ${binaryPath}`);
    } else {
      console.error(`❌ Error executing EnvGuard: ${err.message}`);
    }
    process.exit(1);
  });
  
  // Handle process signals
  process.on('SIGINT', () => {
    child.kill('SIGINT');
  });
  
  process.on('SIGTERM', () => {
    child.kill('SIGTERM');
  });
}

// Run the CLI
runEnvGuard(); 
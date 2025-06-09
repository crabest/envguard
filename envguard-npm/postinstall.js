#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const os = require('os');

function setExecutablePermissions() {
  const platform = os.platform();
  
  // Only set permissions on Unix-like systems
  if (platform === 'win32') {
    console.log('✅ EnvGuard installed successfully on Windows');
    return;
  }
  
  const binDir = path.join(__dirname, 'bin');
  const binaries = ['envguard-linux', 'envguard-darwin'];
  
  try {
    binaries.forEach(binary => {
      const binaryPath = path.join(binDir, binary);
      if (fs.existsSync(binaryPath)) {
        try {
          fs.chmodSync(binaryPath, 0o755);
          console.log(`✅ Set execute permissions for ${binary}`);
        } catch (err) {
          console.warn(`⚠️  Warning: Could not set permissions for ${binary}: ${err.message}`);
        }
      }
    });
    
    console.log('🎉 EnvGuard installed successfully!');
    console.log('');
    console.log('Get started:');
    console.log('  envguard --help                    # Show all commands');
    console.log('  envguard create -e development     # Create first environment');
    console.log('  envguard use development           # Switch to environment');
    console.log('  envguard status                    # Check active environment');
    console.log('');
    console.log('📖 Documentation: https://github.com/crabest/envguard');
    
  } catch (err) {
    console.error(`❌ Error during postinstall: ${err.message}`);
    console.error('EnvGuard may still work, but you might need to set permissions manually.');
    console.error(`Try: chmod +x ${path.join(binDir, '*')}`);
  }
}

// Only run if this script is being executed directly (not required)
if (require.main === module) {
  setExecutablePermissions();
}

module.exports = { setExecutablePermissions }; 
# 🌐 GoLangVPNWindow

> 🛡️ Simple L2TP/IPsec VPN Connector using Go (for Windows)

![Go](https://img.shields.io/badge/Made%20with-Go-blue?style=flat-square)
![Platform](https://img.shields.io/badge/Platform-Windows-0078D6?style=flat-square)
![License](https://img.shields.io/github/license/anh2ten/GoLangVPNWindow?style=flat-square)

---

## ✨ Overview

**GoLangVPNWindow** is a lightweight tool written in Go that allows you to connect to **L2TP/IPsec VPNs** on **Windows** easily — without going through complex UI interactions.

This is ideal for:
- Automated VPN connection setups.
- Headless environments or scripting.
- DevOps or remote access tasks.

---

## ⚙️ Features

- ✅ Connect to VPN using L2TP/IPsec protocol.
- ✅ Set pre-shared key (PSK) for IPsec.
- ✅ Auto-create VPN entry if it doesn't exist.
- ✅ Windows-native implementation via PowerShell.
- ✅ Simple and minimal Go code.

---

## 🧰 Requirements

- Windows 10/11
- Go 1.18+
- Admin privileges (required to create VPN connections)

---
 
## 🚀 Getting Started

### 1. Clone this repository

```bash
git clone https://github.com/anh2ten/GoLangVPNWindow.git
cd GoLangVPNWindow
```

### 2. Build the executable

```bash
go build -o vpnconnect.exe
```

### 3. Run the program

```bash
vpnconnect.exe
```
   

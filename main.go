package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const (
	vpnName  = "MyVPN"
	serverIP = "192.168.1.1"  
	username = "vpnuser"
	password = "vpnpass"
	psk      = "vpnsecret"
)

func run(cmd string, args ...string) {
	exec.Command(cmd, args...).Run()
}

func connectVPN() {
	fmt.Println("Connecting to VPN...")

	run("powershell", "-Command",
		"Add-VpnConnection -Name '"+vpnName+"' -ServerAddress '"+serverIP+"' -TunnelType L2tp -L2tpPsk '"+psk+"' -AuthenticationMethod Pap -Force")

	run("powershell", "-Command",
		"Set-VpnConnectionUsernamePassword -Name '"+vpnName+"' -Username '"+username+"' -Password '"+password+"'")

	run("rasdial", vpnName, username, password)
}

func disconnectVPN() {
	fmt.Println("Disconnecting VPN...")
	run("rasdial", vpnName, "/disconnect")
	run("powershell", "-Command", "Remove-VpnConnection -Name '"+vpnName+"' -Force")
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		disconnectVPN()
		os.Exit(0)
	}()

	connectVPN()
	fmt.Println("VPN connected. Press Ctrl+C to disconnect.")
	select {}
}

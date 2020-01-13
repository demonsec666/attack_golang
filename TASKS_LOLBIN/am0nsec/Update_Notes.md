### Using Hard Links to point back to attacker controlled location.

```
mklink /h C:\Windows\System32\Tasks\tasks.dll C:\Tools\Tasks.dll
Hardlink created for C:\Windows\System32\Tasks\tasks.dll <<===>> C:\Tools\Tasks.dll
```

This can redirect the search to an arbitrary location and evade tools that are looking for filemods in a particular location.

xref: https://googleprojectzero.blogspot.com/2015/12/between-rock-and-hard-link.html

In addition, you can modify the following exe's to load the CLR under the influnce of the malicious DLL.
FileHistory.exe /?

Microsoft.Uev.SyncController.exe

PresentationHost.exe

stordiag.exe

TsWpfWrp.exe

UevAgentPolicyGenerator.exe

UevAppMonitor.exe

UevTemplateBaselineGenerator.exe

UevTemplateConfigItemGenerator.exe

mmc.exe (example launch eventvwr)

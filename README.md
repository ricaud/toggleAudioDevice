# toggleAudoDevice
This small utility will toggle between two different audio devices. Made to be trigger by a single key on my keyboard (Windows only)

change the "device1" and "device2" variables to the names of the devices you want to toggle between. Build toggleDevice.go, then run toggleDevice.exe to toggle between your two defined audio devices.

This works using the nircmd.exe setdefaultaudiodevice utility. More on this utility can be found here https://nircmd.nirsoft.net/setdefaultsounddevice.html
the nircmd.exe executable is bundled with in the built Go executable and is expanded if needed. This utility uses a blank "toggle.file" in order to tell which device to toggle to between executions.

This utility requires at least go1.16 (due to "embed" package) and needs write permissions in whatever directory is resides in (for expanding nircmd.exe and creating/deleting toggle.file)

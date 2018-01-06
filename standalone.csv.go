package tapealert

// variable data generated from export of google sheets data
var sacsv = `No.,Flag,Type,Required Host Message,Cause
1,"Library
Hardware A",C,"The library  mechanism is having difficulty communicating with
the drive:
1. Turn the library off then on.
2. Restart the operation.
3. If problem persists, call the library supplier helpline.","Changer mechanism is
having trouble
communicating with the
internal drive"
2,"Library
Hardware B",W,"There is a problem with the library mechanism.
If problem persists, call the library supplier helpline.","Changer mechanism
has a hardware fault"
3,"Library
Hardware C",C,"The library has a hardware fault:
1. Reset the library
2. Restart the operation.
Check the library users manual for device specific instructions
on resetting the device.","The changer
mechanism has a
hardware fault that
requires reset to
recover."
4,"Library
Hardware D",C,"The library has a hardware fault:
1. Turn the library off and then on again.
2. Restart the operation.
3. If the problem persists, call the library supplier helpline.
Check the library users manual for device specific instructions
on turning the device power on and off.","The changer
mechanism has a
hardware fault that is
not mechanically
related, or requires a
power cycle to recover"
5,"Library
Diagnostics
Required",W,"The library mechanism may have a hardware fault.
Run extended diagnostics to verify and diagnose the problem.
Check the library users manual for device specific instructions
on running extended diagnostic tests.","The changer
mechanism may a
hardware fault which
would be identified by
extended diagnostics
(eg SCSI Send
Diagnostic)."
6,"Library
Interface",C,"The library has a problem with the host interface:
1. Check the cables and cable connections.
2. Restart the operation.","The library has
identified an interfacing
fault"
7,"Predictive
Failure",W,"A hardware failure of the library is predicted.
Call the library supplier helpline.","Predictive failure of
library hardware"
8,"Library
Maintenance",W,"Preventative maintenance of the library is required.
Check the library users manual for device specific preventative
maintenance tasks, or call your library supplier helpline.","Library preventative
maintenance required"
9,"Library
Humidity
Limits",C,"General environmental conditions inside the library are outside
the specified humidity range","Library humidity limits
exceeded"
10,"Library
Temperature
Limits",C,"General environmental conditions inside the library are outside
exceeded the specified temperature range","Library temperature
limits exceeded"
11,"Library
Voltage Limits",C,"The voltage supply to the library is outside the specified range.
There is a potential problem with the power supply or failure of
a redundant power supply","Library voltage limits
exceeded"
12,"Library Stray
Tape",C,"A cartridge has been left in a drive inside the library by a
previous hardware fault:
1. Insert an empty magazine to clear the fault.
2. If the fault does not clear, turn the library off and then on
again.
3. If the problem persists, call the library supplier helpline.","Stray cartridge left in
library drive after
previous error recovery"
13,"Library Pick
Retry",W,"There is a potential problem with a drive ejecting cartridges
short or with the library mechanism picking a cartridge from a
slot.
1. No action needs to be taken at this time.
2. If the problem persists, call the library supplier helpline.","Operation to pick a
cartridge from a slot
had to perform an
excessive number of
retries before
succeeding"
14,"Library Place
Retry",W,"There is a potential problem with the library mechanism placing
a cartridge into a slot
1. No action needs to be taken at this time.
2. If the problem persists, call the library supplier helpline.","Operation to place a
cartridge into a slot had
to perform an excessive
number of retries
before succeeding"
15,"Library Load
Retry",W,"There is a potential problem with a drive or the library
mechanism loading cartridges, or an incompatible cartridge.","Operation to load a
cartridge into a drive
had to perform an
excessive number of
retries before
succeeding"
16,Library Door,C,"The operation has failed because the library door is open:
1. Clear any obstructions from the library door.
2. Close the library door.
3. If the problem persists, call the library supplier helpline.","Changer door open
prevents library
functioning"
17,"Library
Mailslot",C,"There is a mechanical problem with the library media
import/export mailslot.","Mechanical problem
with import/export
mailslot"
18,"Library
Magazine",C,"The library cannot operate without the magazine.
1. Insert the magazine into the library
2. Restart the operation.","Library magazine not
present"
19,"Library
Security",W,Library security has been compromised,"Library door opened
then closed during
operation"
20,"Library
Security Mode",I,"The security mode of the library has been changed.
The library has either been put into secure mode, or the library
has exited the secure mode.
This is for information purposes only. No action is required.","Library security mode
changed"
21,Library Offline,I,"The library has been manually turned offline and is unavailable
for use.","Library manually
turned offline"
22,"Library Drive
Offline",I,"A drive inside the library has been taken offline.
This is for information purposes only. No action is required.","Library turned internal
drive offline"
23,"Library Scan
Retry",W,"There is a potential problem with the barcode label or the
scanner hardware in the library mechanism.
1. No action needs to be taken at this time.
2. If the problem persists, call the library supplier helpline.","Operation to scan the
barcode on a cartridge
had to perform an
excessive number of
retries before
succeeding"
24,"Library
Inventory",C,"The library has detected a inconsistency in its inventory.
1. Redo the library inventory to correct inconsistency.
2. Restart the operation
Check the applications users manual or the hardware users
manual for specific instructions on redoing the library inventory.","Inconsistent media
inventory"
25,"Library Illegal
Operation",W,"A library operation has been attempted that is invalid at this
time.","Illegal operation
detected"
26,"Dual-Port
Interface Error",W,A redundant interface port on the library has failed,"Failure of one interface
port in a dual-port
configuration, e.g.
Fibrechannel"
27,"Cooling Fan
Failure",W,A library cooling fan has failed.,"One or more fans inside
the library have failed.
Internal flag state only
cleared when all fans
are working again."
28,Power Supply,W,"A redundant power supply has failed inside the library. Check
the library users manual for instructions on replacing the failed
power supply.","Redundant PSU failure
inside the library
subsystem"
29,"Power
Consumption",W,The libary power consumption is outside the specified range,"Power consumption of
one or more devices
inside the library is
outside specified range"
30,"Pass-through
mechansim
failure",C,"A failure has occurred in the cartridge pass-through mechanism
between two library modules.","Error occurred in pass-
through mechanism
during self test or while
attempting to transfer a
cartridge between
library modules"
31,"Cartridge in
pass-through
mechanism",C,"A cartridge has been left in the library pass-through mechanism
from a previous hardware fault. Check the library users guide
for instructions on clearing this fault.","Cartridge left in the
pass-through
mechanism between
two library modules"
32,"Unreadable bar
code labels",I,The library was unable to read the bar code on a cartridge,"Unable to read a bar-
code label on a
cartridge during library
inventory/scan"`

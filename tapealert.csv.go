package tapealert

// variable data generated from export of google sheets data
var tpcsv = `No.,Flag,Type,Required Host Message,Cause
1,Read Warning,W,"The tape drive is having problems reading data. No data has
been lost, but there has been a reduction in the performance of
the tape.","The drive is having
severe trouble reading"
2,"Write
Warning",W,"The tape drive is having problems writing data. No data has
been lost, but there has been a reduction in the capacity of the
tape.","The drive is having
severe trouble writing"
3,Hard Error,W,"The operation has stopped because an error has occurred while
reading or writing data which the drive cannot correct.","The drive had a hard
read or write error"
4,Media,C,"Your data is at risk:
1. Copy any data you require from this tape.
2. Do not use this tape again.
3. Restart the operation with a different tape.","Media can no longer be
written/read, or
performance is severely
degraded"
5,Read Failure,C,"The tape is damaged or the drive is faulty. Call the tape drive
supplier helpline.","The drive can no longer
read data from the tape"
6,Write Failure,C,"The tape is from a faulty batch or the tape drive is faulty:
1. Use a good tape to test the drive.
2. If the problem persists, call the tape drive supplier helpline.","The drive can no longer
write data to the tape"
7,Media Life,W,"The tape cartridge has reached the end of its calculated useful
life:
1. Copy any data you need to another tape
2. Discard the old tape.","The media has
exceeded its specified
life"
8,"Not Data
Grade",W,"The tape cartridge is not data-grade. Any data you back up to
the tape is at risk.
Replace the cartridge with a data-grade tape.","The drive has not been
able to read the MRS
stripes"
9,Write Protect,C,"You are trying to write to a write-protected cartridge.
Remove the write-protection or use another tape.","Write command is
attempted to a write
protected tape"
10,No Removal,I,"You cannot eject the cartridge because the tape drive is in use.
Wait until the operation is complete before ejecting the
cartridge.","Manual or s/w unload
attempted when prevent
media removal on"
11,"Cleaning
Media",I,The tape in the drive is a cleaning cartridge.,"Cleaning tape loaded
into drive"
12,"Unsupported
Format",I,"You have tried to load a cartridge of a type which is not
supported by this drive.","Attempted loaded of
unsupported tape
format, e.g. DDS2 in
DDS1 drive"
13,"Recoverable
Snapped Tape",C,"The operation has failed because the tape in the drive has
snapped:
1. Discard the old tape.
2. Restart the operation with a different tape.","Tape snapped/cut in the
drive where media can
be ejected"
14,"Unrecoverable
Snapped Tape",C,"The operation has failed because the tape in the drive has
snapped:
1. Do not attempt to extract the tape cartridge.
2. Call the tape drive supplier helpline.","Tape snapped/cut in the
drive where media
cannot be ejected"
15,"Memory Chip
in Cartridge
Failure",W,"The memory in the tape cartridge has failed, which reduces
performance.
Do not use the cartridge for further backup operations.","Memory chip failed in
cartridge"
16,Forced Eject,C,"The operation has failed because the tape cartridge was
manually ejected while the tape drive was actively writing or
reading.","Manual or forced eject
while drive actively
writing or reading"
17,"Read Only
Format",W,"You have loaded a cartridge of a type that is read-only in this
drive.
The cartridge will appear as write-protected","Media loaded that is
read-only format"
18,"Tape Directory
Corrupted on
Load",W,"The directory on the tape cartridge has been corrupted.
File search performance will be degraded. The tape directory
can be rebuilt by reading all the data on the cartridge","Tape drive powered
down with tape loaded,
or permanent error
prevented the tape
directory being updated"
19,"Nearing Media
Life",I,"The tape cartridge is nearing the end of its calculated life. It is
recommended that you:
1. Use another tape cartridge for your next backup.
2. Store this tape cartridge in a safe place in case you need to
restore data from it.","Media may have
exceeded its specified
number of passes"
20,Clean Now,C,"The tape drive needs cleaning:
1. If the operation has stopped, eject the tape and clean the drive
2. If the operation has not stopped, wait for it to finish and then
clean the drive.
Check the tape drive users manual for device specific cleaning
instructions.","The drive thinks it has
a head clog, or needs
cleaning"
21,Clean Periodic,W,"The tape drive is due for routine cleaning:
1. Wait for the current operation to finish.
2. Then use a cleaning cartridge.
Check the tape drive users manual for device specific cleaning
instructions.","The drive is ready for a
periodic clean"
22,"Expired
Cleaning
Media",C,"The last cleaning cartridge used in the tape drive has worn out:
1. Discard the worn out cleaning cartridge.
2. Wait for the current operation to finish.
3. Then use a new cleaning cartridge.","The cleaning tape has
expired"
23,"Invalid
Cleaning Tape",C,"The last cleaning cartridge used in the tape drive was an invalid
type:
1. Do not use this cleaning cartridge in this drive.
2. Wait for the current operation to finish.
3. Then use a valid cleaning cartridge.","Invalid cleaning tape
type used"
24,"Retention
Requested",W,The tape drive has requested a retention operation,"The drive is having
severe trouble reading
or writing, which will
be resolved by a
retention cycle"
25,"Dual-Port
Interface Error",W,A redundant interface port on the tape drive has failed,"Failure of one interface
port in a dual-port
configuration, e.g.
Fibrechannel"
26,"Cooling Fan
Failure",W,A tape drive cooling fan has failed.,"Fan failure inside tape
drive mechanism or
tape drive enclosure"
27,Power Supply,W,"A redundant power supply has failed inside the tape drive
enclosure. Check the enclosure users manual for instructions on
replacing the failed power supply.","Redundant PSU failure
inside the tape drive
enclosure or rack
subsystem"
28,"Power
Consumption",W,The tape drive power consumption is outside the specified range,"Power consumption of
the tape drive is outside
specified range"
29,"Drive
Maintenance",W,"The drive requires
preventative
maintenance (not
cleaning).",
30,Hardware A,C,"The tape drive has a hardware fault:
1. Eject the tape or magazine.
2. Reset the drive.
3. Restart the operation.","The drive has a
hardware fault that
requires reset to
recover."
31,Hardware B,C,"The tape drive has a hardware fault:
1. Turn the tape drive off and then on again.
2. Restart the operation.
3. If the problem persists, call the tape drive supplier helpline.
Check the tape drive users manual for device specific
instructions on turning the device power on and off.","The drive has a
hardware fault which is
not read/write related
or requires a power
cycle to recover."
32,Interface,W,"The tape drive has a problem with the host interface:
1. Check the cables and cable connections.
2. Restart the operation.","The drive has identified
an interfacing fault"
33,Eject Media,C,"The operation has failed:
1. Eject the tape or magazine.
2. Insert the tape or magazine again.
3. Restart the operation.",Error recovery action
34,Download Fail,W,"The firmware download has failed because you have tried to use
the incorrect firmware for this tape drive.
Obtain the correct firmware and try again.","Firmware download
failed"
35,"Drive
Humidity",W,"Environmental conditions inside the tape drive are outside the
specified humidity range","Drive humidity limits
exceeded"
36,"Drive
Temperature",W,"Environmental conditions inside the tape drive are outside the
specified temperature range","Drive temperature
limits exceeded"
37,Drive Voltage,W,"The voltage supply to the tape drive is outside the specified
range","Drive voltage limits
exceeded"
38,"Predictive
Failure",C,"A hardware failure of the tape drive is predicted.
Call the tape drive supplier helpline.","Predictive failure of
drive hardware"
39,"Diagnostics
Required",W,"The tape drive may have a fault. Check for availability of
diagnostic information and run extended diagnostics if
applicable.
Check the tape drive users manual for instructions on running
extended diagnostic tests and retrieving diagnostic data","The drive may have had
a failure which may be
identified by stored
diagnostic information
or by running extended
diagnostics (eg Send
Diagnostic)"
40,"Loader
Hardware A",C,"The changer mechanism is having difficulty communicating
with the tape drive:
1. Turn the autoloader off then on.
2. Restart the operation.
3. If problem persists, call the tape drive supplier helpline.","Loader mechanism is
having trouble
communicating with the
tape drive"
41,"Loader Stray
Tape",C,"A tape has been left in the autoloader by a previous hardware
fault:
1. Insert an empty magazine to clear the fault.
2. If the fault does not clear, turn the autoloader off and then on
again.
3. If the problem persists, call the tape drive supplier helpline.","Stray tape left in loader
after previous error
recovery"
42,"Loader
Hardware B",W,There is a problem with the autoloader mechanism.,"Loader mechanism has
a hardware fault"
43,Loader Door,C,"The operation has failed because the autoloader door is open:
1. Clear any obstructions from the autoloader door.
2. Eject the magazine and then insert it again.
3. If the fault does not clear, turn the autoloader off and then on
again
4. If the problem persists, call the tape drive supplier helpline.",Tape changer door open
44,"Loader
Hardware C",C,"The autoloader has a hardware fault:
1. Turn the autoloader off and then on again.
2. Restart the operation.
3. If the problem persists, call the tape drive supplier helpline.
Check the autoloader users manual for device specific
instructions on turning the device power on and off.","The loader mechanism
has a hardware fault that
is not mechanically
related."
45,"Loader
Magazine",C,"The autoloader cannot operate without the magazine.
1. Insert the magazine into the autoloader
2. Restart the operation.","Loader magazine not
present"
46,"Loader
Predictive
Failure",W,"A hardware failure of the changer mechanism is predicted.
Call the tape drive supplier helpline.","Predictive failure of
loader mechanism
hardware"
50,Lost Statistics,W,Media statistics have been lost at some time in the past,"Drive or library powered
down with tape loaded"
51,"Tape directory
invalid at
unload",W,"The tape directory on the tape cartridge just unloaded has been
corrupted.
File search performance will be degraded
The tape directory can be rebuilt by reading all the data.","Error prevented the tape
directory being updated
on unload."
52,"Tape system
area write
failure",C,"The tape just unloaded could not write its system area
successfully:
1. Copy data to another tape cartridge
2. Discard the old cartridge","Write errors while
writing the system log on
unload"
53,"Tape system
area read
failure",C,"The tape system area could not be read successfully at load
time:
1. Copy data to another tape cartridge
2. Discard the old cartridge","Read errors while
reading the system area
on load"
54,No start of data,C,"The start of data could not be found on the tape:
1. Check you are using the correct format tape
2. Discard the tape or return the tape to your supplier","Tape damaged, bulk
erased, or incorrect
format"`

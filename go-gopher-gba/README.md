## TinyGo - A Go Compiler for small places
TinyGo brings the Go programming language to embedded systems and to the 
modern web by creating a new compiler based on LLVM.

With TinyGo, you can compile and run programs on over 94 different microcontroller
boards and can also produce WebAssembly code. This allows to compile programs
for web browsers, as well as for server and edge computing environments that
support WebAssembly System Interface family of interfaces.


## What does this game do?
- Display a screen with "Gopher" text and "Press START button"
- Display two gophers
- When START button is pressed: you Gopher player just appear
- With multidirectional buttons the Gopher moves left, right, top, bottom
- When A button is pressed, the Gopher jumps
- When SELECT button is pressed, it directs the player to the "Start" screen


## Volatile Memory Location (e.x. RAM, Cache)
A volatile memory location is a location in memory that can be read and written
by multiple threads simultaneously without the need for sychronization. This is
because the value of a volatile memory location is not cached by the processor
and is always read directly from memory. This makes volatile memory locations
ideal for storing data that is frequently accessed and updated by multiple threads.


## Setup and access a display-related register
> 	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

1. Memory-mapped register:
  - `0x4000004` is a hexidecimal value representing a memory-mapped register address.
  - Memory-mapped registers are special addresses in an embedded system that control
    various hardware peripherals, such as *display controllers*.

2. `uintptr()` is a conversion function that converts the memory address to an unsugned
    integer of type `uintptr`.
3. `unsafe.Pointer(...)`
    Converts the unsigned interger memory address to an unsafe pointer. This is done to
    enable low-level pointer operations, as memory-mapped registers often involve direct
    interaction with hardware.
4. `(*volatile.Register16)(...)`
    Is a type assestion that converts the unsafe pointer to a specific type, in this case,
    a `volatile.Register16` type. The use of `volatile` indicates that this type is likely 
    of a package or library dealing with memory-mapped I/O, where the data at a memory address
    may change at any time.
5. `regDISPSTAT` is a variable, possibly global, that is being assigned the memory address of
    the display status register.

In an embedded system with a display controller, `DISPSTAT` might be a register that holds status
information about the display, such as whether it is currently active, the display mode, or other
relevant information. By mapping this register to a variable, the code can read or write to
this register to control aspects of the display.


## Install dependencies
TinyDraw is a useful tool to draw gemoetric figures.
TinyFont is a font/text package for TinyGo displays.

```bash
go get tinygo.org/x/tinydraw 
go get tinygo.org/x/tinyfont
```

## Test!
To test the application, TinyGo recomments mGBA, a Game Boy Advance software emulator.

Download mGBA, untat it and coppy it in the PATH:
```bash
cp ./bin/mgba /usr/local/bin/mgba
```
Check the executable is working:
```bash
mgba --version
```

Run mGBA with the `mgba` command on the terminal.

```bash
tinygo run -target=gameboy-advance gopher.go
`

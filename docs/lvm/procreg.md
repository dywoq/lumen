# Processor's registers

Processor's registers are divided into the following categories: **general-purpose** and **control** registers.

## Shared concepts

- Registers, including both categories, can contain a 64-bit immediate values or memory addresses.

## General-purpose registers

General-purpose registers are used by a guest program. It is permitted to use them in the user
and kernel modes, therefore, the LVM's processor is forbidden from taking away this permission.

General-purpose registers are divided into **usage groups** to separate concerns and responsibilities:

- **M (Memory) Group**: These registers are used for containing memory addresses.
- **R (Regular) Group**: These registers are used for containing regular immediate values.

Following general-purpose registers exist:

- `M1`, `M2`, `M3`, `M4`, `M5`, `M6`, `M7`, `M8`
- `R1`, `R2`, `R3`, `R4`, `R5`, `R6`, `R7`, `R8`

## Control registers

Control registers are used by a guest program in the kernel. They are used to set interrupts, memory protection
mechanisms and system call management up. Unlike general-purpose registers, a guest program running in user mode
can be forbidden from using the control units if it is set explicitly. They have a `C` prefix.

Following control registers exist:

- `CMM` (Memory Management), `CSC` (System Call), `CIM` (Interrupt Management)

## Registers' identifier numbers

As stated in the [Processor's instruction set](./procis.md) document, an instruction includes a
destination register number, which is given a 8-bit region. All general-purpose and control units
are provided with a unique identifier number. These identifiers are shown below.

| **Register** | **Its identifier number (hexadecimal)** |
| ------------ | --------------------------------------- |
| `M1`         | `0x00`                                  |
| `M2`         | `0x01`                                  |
| `M3`         | `0x02`                                  |
| `M4`         | `0x03`                                  |
| `M5`         | `0x04`                                  |
| `M6`         | `0x05`                                  |
| `M7`         | `0x06`                                  |
| `M8`         | `0x07`                                  |
| `R1`         | `0x08`                                  |
| `R2`         | `0x09`                                  |
| `R3`         | `0x0A`                                  |
| `R4`         | `0x0B`                                  |
| `R5`         | `0x0C`                                  |
| `R6`         | `0x0D`                                  |
| `R7`         | `0x0E`                                  |
| `R8`         | `0x0F`                                  |
| `CMM`        | `0x10`                                  |
| `CSC`        | `0x11`                                  |
| `CIM`        | `0x12`                                  |

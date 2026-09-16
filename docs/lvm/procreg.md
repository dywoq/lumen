# Processor's registers

Processor's registers are divided into the following categories: **general-purpose** and **control** registers.

## Shared concepts

- Registers, including both categories, can contain a 64-bit immediate values or memory addresses.

## General-purpose Registers

General-purpose registers are used by a guest program. It is permitted to use them in the user
and kernel modes, therefore, the LVM's processor is forbidden from taking away this permission.

General-purpose registers are divided into **usage groups** to separate concerns and responsibilities:

- **M (Memory) Group**: Used for containing memory addresses.
- **R (Regular) Group**: Used for containing regular immediate values.

Following general-purpose registers exist:

- `M1`, `M2`, `M3`, `M4`, `M5`, `M6`, `M7`, `M8`
- `R1`, `R2`, `R3`, `R4`, `R5`, `R6`, `R7`, `R8`

## Control Registers

Control registers are used by a guest program in the kernel. They are used to set interrupts, memory protection
mechanisms and system call management up. Unlike general-purpose registers, a guest program running in user mode
can be forbidden from using the control units if it is set explicitly. They have a `C` prefix.

Following control registers exist:

- `CMM` (Control Memory Management), `CSC` (Control System Call), `CIM` (Control Interrupt Management)

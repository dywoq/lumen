# Processor's instruction set

## Instruction

A processor instruction is a command to request a specific operation from the LVM.
Its fixed size is 16 bytes. It consists of the instruction's metadata, destination register and
64-bit immediate value/source register.

## Bit form

- `0...15` - The instruction's operation code.
- `16...23` - The destination register's identifier number.
- `24...31` - The source register's identifier number.
- `32...63` - Reserved for future additions and changes.
- `64...127` - The 64-bit immediate value.

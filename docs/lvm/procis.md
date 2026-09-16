# Processor's Instruction Set

## Instruction

A processor instruction is a command to request a specific operation from the LVM.
Its fixed size is 16 bytes. It consists of the instruction metadata and an 64-bit immediate
value. Below, you can see its bit form:

- `0...15` - This region is reserved for the instruction's identifier number.
- `16...63` - This region is reserved for future additions.
- `64...127` - This region is reserved for the 64-bit immediate value.

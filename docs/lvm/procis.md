# Processor Instruction Set

## Instruction

A processor instruction is a command to request a specific operation from the LVM.
Its fixed size is 16 bytes. It consists of the instruction metadata and an 64-bit immediate
value. Below, you can see its bit form:

- `0...15` - The instruction operand.
- `64...127` - The 64-bit immediate value. This region is full of zeros when an operation
  does not require its presence.

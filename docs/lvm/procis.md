# Processor's instruction set

## Instruction

A processor instruction is a command to request a specific operation from the LVM.
Its fixed size is 16 bytes. It consists of the instruction metadata and an 64-bit immediate
value. Below, you can see its bit form:

- `0...15` - Reserved for the instruction's identifier number.
- `16...23` - Reserved to specify an instruction destination register.
- `24...63` - Reserved for future additions.
- `64...127` - Reserved for the 64-bit immediate value.

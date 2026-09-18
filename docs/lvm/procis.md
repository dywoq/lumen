# Processor's instruction set

## Instruction

A processor instruction is a command to request a specific operation from the LVM.
Its fixed size is 16 bytes. It consists of the instruction's metadata, destination register and
64-bit immediate value/source register.

### Bit form

- `0...15` - The instruction's operation code.
- `16...23` - The destination register's identifier number.
- `24...31` - The source register's identifier number.
- `32...63` - Reserved for future additions and changes.
- `64...127` - The 64-bit immediate value.

### Table

| **Name** | **Short description**                                                  |
| -------- | ---------------------------------------------------------------------- |
| `storem` | Stores a memory address into a general-purpose register                |
| `storer` | Stores an immediate value into a general-purpose register              |
| `loadm`  | Copies a memory address from source register to destination register.  |
| `loadr`  | Copies an immediate value from source register to destination register |

## Instructions

### `storem`

**Pseudo-form**:

```
storem Register, [MemoryAddress]
```

**Description**: This instruction stores the provided memory address into the destination register (`Register`).

**Invariants**:

- The provided destination register must be general-purpose and from the usage group M.

### `storer`

**Pseudo-form**:

```
storer Register, ImmediateValue
```

**Description**: This instruction stores the provided immediate value into the destination register
(`Register`).

**Invariants**:

- The provided destination register must be general-purpose and from the usage group R.

### `loadm`

**Pseudo-form**:

```
loadm DestinationRegister, SourceRegister
```

**Description**: This instruction copies a memory address of the source register (`SourceRegister`)
into the destination register (`DestinationRegister`).

**Invariants**:

- The specified registers must be general-purpose and from the usage group M.

### `loadr`

**Pseudo-form**:

```
loadr DestinationRegister, SourceRegister
```

**Description**: This instruction copies a memory address of the source register (`SourceRegister`)
into the destination register (`DestinationRegister`).

**Invariants**:

- The specified registers must be general-purpose and from the usage group R.

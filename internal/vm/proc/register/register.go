// Copyright 2926 dywoq - Apache License 2.0
// https://github.com/dywoq/lumen

package register

// Kind identifies the type of a register.
type Kind int

// GPUsageGroup identifies the usage group of a general-purpose register.
type GPUsageGroup int

// GPId specifies a unique identifier number of a general-purpose register.
// For example, if it equals to 5, and the usage group is M,
// this general-purpose register is M5.
//
// According to the LVM documentation, the identifier number must stay within
// the following range: 1 ... 8.
type GPId int

// GPInfo consists of the information, related to general-purpose registers.
type GPInfo struct {
	Id         GPId
	UsageGroup GPUsageGroup
}

// ControlType identifies the type of a control register.
type ControlType int

// ControlInfo consists of the information, related to control registers.
type ControlInfo struct {
	Type ControlType
}

// R contains the information of a register.
type R struct {
	Kind  Kind
	Value uint64
	// GPInfo must stay null if Kind is not [KindGeneralPurpose].
	GPInfo *GPInfo
	// ControlInfo must stay null If Kind is not [KindControl].
	ControlInfo *ControlInfo
}

const (
	KindGeneralPurpose Kind = iota
	KindControl
)

const (
	GPUsageGroupM GPUsageGroup = iota
	GPUsageGroupR
)

const (
	ControlTypeMm ControlType = iota
	ControlTypeSc
	ControlTypeIm
)

func (k Kind) IsValid() bool {
	switch k {
	case KindGeneralPurpose, KindControl:
		return true
	default:
		return false
	}
}

func (id GPId) IsValid() bool {
	return id >= 1 && id <= 8
}

func (ct ControlType) IsValid() bool {
	switch ct {
	case ControlTypeMm, ControlTypeSc, ControlTypeIm:
		return true
	default:
		return false
	}
}

// NewGP constructs a [R] instance, where it includes general-purpose register information,
// and the [KindGeneralPurpose] kind.
func NewGP(usageGroup GPUsageGroup) *R {
	return &R{
		Kind: KindGeneralPurpose,
		GPInfo: &GPInfo{
			UsageGroup: usageGroup,
		},
	}
}

// NewGP constructs a [R] instance, where it includes control register information,
// and the [KindGeneralPurpose] kind.
func NewControl(controlType ControlType) *R {
	return &R{
		Kind: KindControl,
		ControlInfo: &ControlInfo{
			Type: controlType,
		},
	}
}

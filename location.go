package zog

import "fmt"

type Src8 interface {
	Read8(z *Zog) (byte, error)
	String() string
}
type Dst8 interface {
	Write8(z *Zog, n byte) error
	String() string
}

type Src16 interface {
	Read16(z *Zog) (uint16, error)
	String() string
}
type Dst16 interface {
	Write16(z *Zog, nn uint16) error
	String() string
}

type Loc8 interface {
	Read8(z *Zog) (byte, error)
	Write8(z *Zog, n byte) error
	String() string
}

type Loc16 interface {
	Read16(z *Zog) (uint16, error)
	Write16(z *Zog, nn uint16) error
	String() string
}

type R8 int

const (
	A R8 = iota
	F
	B
	C
	D
	E
	H
	L
)

func (r R8) String() string {
	switch r {

	case A:
		return "A"
	case F:
		return "F"

	case B:
		return "B"
	case C:
		return "C"

	case D:
		return "D"
	case E:
		return "E"

	case H:
		return "H"
	case L:
		return "L"

	default:
		panic(fmt.Errorf("Unrecognised R8 : %d", int(r)))
	}
}

func (r R8) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, r)
	return n, nil
}
func (r R8) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", r, n)
	return nil
}

type R16 int

const (
	AF R16 = iota
	BC
	DE
	HL
	IX
	IY
	SP
	AF_PRIME
)

func (r R16) String() string {
	switch r {

	case AF:
		return "AF"
	case AF_PRIME:
		return "AF'"

	case BC:
		return "BC"

	case DE:
		return "DE"

	case HL:
		return "HL"

	case IX:
		return "IX"

	case IY:
		return "IY"

	case SP:
		return "SP"

	default:
		panic(fmt.Errorf("Unrecognised R16 : %d", int(r)))
	}
}

func (r R16) Read16(z *Zog) (uint16, error) {
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %02X <- %s\n", nn, r)
	return nn, nil
}
func (r R16) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", r, nn)
	return nil
}

type Contents struct {
	addr Src16
}

func (c Contents) String() string {
	return fmt.Sprintf("(%s)", c.addr)
}
func (c Contents) Read8(z *Zog) (byte, error) {
	// TODO: debug
	var n byte
	fmt.Printf("Z: %02X <- %s\n", n, c)
	return n, nil
}
func (c Contents) Write8(z *Zog, n byte) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %02X\n", c, n)
	return nil
}

func (c Contents) Read16(z *Zog) (uint16, error) {
	// TODO: debug
	var nn uint16
	fmt.Printf("Z: %04X <- %s\n", nn, c)
	return nn, nil
}
func (c Contents) Write16(z *Zog, nn uint16) error {
	// TODO: debug
	fmt.Printf("Z: %s <- %04X\n", c, nn)
	return nil
}

type Imm16 uint16

func (nn Imm16) String() string {
	return fmt.Sprintf("0x%04X", uint16(nn))
}
func (nn Imm16) Read16(z *Zog) (uint16, error) {
	return uint16(nn), nil
}

type Imm8 byte

func (n Imm8) String() string {
	return fmt.Sprintf("0x%02X", byte(n))
}
func (n Imm8) Read8(z *Zog) (byte, error) {
	return byte(n), nil
}

type Conditional interface {
	String() string
}

type FlagTest int

const (
	FT_Z FlagTest = iota
	FT_C
	FT_PO
	FT_PE
	FT_P
	FT_M
)

func (ft FlagTest) String() string {
	switch ft {
	case FT_Z:
		return "Z"
	case FT_C:
		return "C"
	case FT_PO:
		return "PO"
	case FT_PE:
		return "PE"
	case FT_P:
		return "P"
	case FT_M:
		return "M"
	default:
		panic(fmt.Sprintf("Unknown flag test [%d]", int(ft)))
	}
}

type LogicConstant struct{}

var True LogicConstant

func (l LogicConstant) String() string {
	panic("Attempt to render 'true' as string")
}

type Not struct {
	ft FlagTest
}

func (n Not) String() string {
	return fmt.Sprintf("N%s", n.ft)
}

type Disp int8

func (d Disp) String() string {
	return fmt.Sprintf("%d", d)
}

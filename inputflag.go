package termios

//go:generate stringer -type=InputFlag

// InputFlag describes the different input modes
type InputFlag uint32

const (
	// BRKINT  If IGNBRK is set, a BREAK is ignored.  If it is not set but
	// BRKINT is set, then a BREAK causes the input and output queues to be
	// flushed, and  if  the  terminal  is  the  controlling terminal of a
	// foreground process group, it will cause a SIGINT to be sent to this
	// foreground process group.  When neither IGNBRK nor BRKINT are set, a
	// BREAK reads as a null byte ('\0'), except when PARMRK is set, in
	// which case  it reads as the sequence \377 \0 \0.
	BRKINT InputFlag = 0x2
	// ICRNL Translate carriage return to newline on input (unless IGNCR is set).
	ICRNL InputFlag = 0x100
	// IGNBRK  Ignore BREAK condition on input.
	IGNBRK InputFlag = 0x1
	// IGNCR   Ignore carriage return on input.
	IGNCR InputFlag = 0x80
	// IGNPAR  Ignore framing errors and parity errors.
	IGNPAR InputFlag = 0x4
	// IMAXBEL (not in POSIX) Ring bell when input queue is full. Linux does not
	// implement this bit, and acts as if it is always set.
	IMAXBEL InputFlag = 0x2000
	// INLCR Translate NL to CR on input.
	INLCR InputFlag = 0x40
	// INPCK Enable input parity checking.
	INPCK InputFlag = 0x10
	// ISTRIP Strip off eighth bit.
	ISTRIP InputFlag = 0x20
	// IUCLC (not in POSIX) Map uppercase characters to lowercase on input.
	IUCLC InputFlag = 0x200
	// IUTF8 (since Linux 2.6.4) (not in POSIX) Input is UTF8; this allows
	// character-erase to be correctly performed in cooked mode.
	IUTF8 InputFlag = 0x4000
	// IXANY (XSI) Typing any character will restart stopped output. (The
	// default is to allow just the START character to restart output.)
	IXANY InputFlag = 0x800
	//IXOFF Enable XON/XOFF flow control on input.
	IXOFF InputFlag = 0x1000
	// IXON Enable XON/XOFF flow control on output.
	IXON InputFlag = 0x400
	// PARMRK If this bit is set, input bytes with parity or framing errors are
	// marked when passed to the program. This bit is meaningful only when INPCK
	// is set and IGNPAR is not set. The way erroneous bytes are marked is with
	// two preceding bytes, \377 and \0. Thus, the program actually reads three
	// bytes for one erroneous byte received from the terminal. If a valid byte
	// has the value \377, and ISTRIP (see below) is not set, the program might
	// confuse it with the prefix that marks a parity error. Therefore, a valid
	// byte \377 is passed to the program as two bytes, \377 \377, in this case.
	// If neither IGNPAR nor PARMRK is set, read a character with a parity error
	// or framing error as \0.
	PARMRK InputFlag = 0x8
)

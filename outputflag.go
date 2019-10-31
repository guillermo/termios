package termios

//go:generate stringer -type=OutputFlag

// OutputFlag describes the different output modes
type OutputFlag uint32

const (

	// OPOST Enable implementation-defined output processing.
	OPOST OutputFlag = 0000001
	// OLCUC (not in POSIX) Map lowercase characters to uppercase on
	// output.
	OLCUC OutputFlag = 0000002
	// ONLCR (XSI) Map NL to CR-NL on output.
	ONLCR OutputFlag = 0000004
	// OCRNL Map CR to NL on output.
	OCRNL OutputFlag = 0000010
	// ONOCR Don't output CR at column 0.
	ONOCR OutputFlag = 0000020
	// ONLRET Don't output CR.
	ONLRET OutputFlag = 0000040
	// OFILL Send fill characters for a delay, rather than using a timed
	// delay.
	OFILL OutputFlag = 0000100
	// OFDEL Fill character is ASCII DEL (0177).  If unset, fill character
	// is ASCII NUL ('\0').  (Not implemented on Linux.)
	OFDEL OutputFlag = 0000200
	// NLDLY Newline delay mask.  Values are NL0 and NL1.  [requires
	// _BSD_SOURCE or _SVID_SOURCE or _XOPEN_SOURCE]
	NLDLY OutputFlag = 0000400
	// CRDLY Carriage return delay mask.  Values are CR0, CR1, CR2, or CR3.
	// [requires _BSD_SOURCE or _SVID_SOURCE or _XOPEN_SOURCE]
	CRDLY OutputFlag = 0003000
	// TABDLY Horizontal  tab delay mask.  Values are TAB0, TAB1, TAB2, TAB3
	// (or XTABS).  A value of TAB3, that is, XTABS, expands tabs to spaces
	// (with tab stops every eight columns).  [requires _BSD_SOURCE or
	// _SVID_SOURCE or _XOPEN_SOURCE]
	TABDLY OutputFlag = 0014000
	// BSDLY Backspace delay mask.  Values are BS0 or BS1.  (Has never been
	// implemented.)  [requires _BSD_SOURCE or _SVID_SOURCE or
	// _XOPEN_SOURCE]
	BSDLY OutputFlag = 0020000
	// VTDLY Vertical tab delay mask.  Values are VT0 or VT1.
	VTDLY OutputFlag = 0040000
	// FFDLY Form feed delay mask.  Values are FF0 or FF1.  [requires
	// _BSD_SOURCE or _SVID_SOURCE or _XOPEN_SOURCE]
	FFDLY OutputFlag = 0100000
)

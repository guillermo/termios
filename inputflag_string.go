// Code generated by "stringer -type=InputFlag"; DO NOT EDIT.

package termios

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BRKINT-2]
	_ = x[ICRNL-256]
	_ = x[IGNBRK-1]
	_ = x[IGNCR-128]
	_ = x[IGNPAR-4]
	_ = x[IMAXBEL-8192]
	_ = x[INLCR-64]
	_ = x[INPCK-16]
	_ = x[ISTRIP-32]
	_ = x[IUCLC-512]
	_ = x[IUTF8-16384]
	_ = x[IXANY-2048]
	_ = x[IXOFF-4096]
	_ = x[IXON-1024]
	_ = x[PARMRK-8]
}

const _InputFlag_name = "IGNBRKBRKINTIGNPARPARMRKINPCKISTRIPINLCRIGNCRICRNLIUCLCIXONIXANYIXOFFIMAXBELIUTF8"

var _InputFlag_map = map[InputFlag]string{
	1:     _InputFlag_name[0:6],
	2:     _InputFlag_name[6:12],
	4:     _InputFlag_name[12:18],
	8:     _InputFlag_name[18:24],
	16:    _InputFlag_name[24:29],
	32:    _InputFlag_name[29:35],
	64:    _InputFlag_name[35:40],
	128:   _InputFlag_name[40:45],
	256:   _InputFlag_name[45:50],
	512:   _InputFlag_name[50:55],
	1024:  _InputFlag_name[55:59],
	2048:  _InputFlag_name[59:64],
	4096:  _InputFlag_name[64:69],
	8192:  _InputFlag_name[69:76],
	16384: _InputFlag_name[76:81],
}

func (i InputFlag) String() string {
	if str, ok := _InputFlag_map[i]; ok {
		return str
	}
	return "InputFlag(" + strconv.FormatInt(int64(i), 10) + ")"
}

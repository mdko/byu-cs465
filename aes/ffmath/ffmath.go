package ffmath

func xtime(in byte) (out byte) {
			overflow := (in & 0x80) == 0x80	// The msb is set, so multiplying by two will overflow into a 2nd byte
			out = in << 1
			if bool(overflow) {
					out = out ^ 0x1b
			}
			return out
}

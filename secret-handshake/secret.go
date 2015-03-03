package secret

import "encoding/binary"

const (
	_ byte = 1 << iota
	wink
	doubleBlink
	closeEyes
	jump
	reverse
)

func Handshake(code int) []string {

	var h []string
	if code > 31 || code < 0 {
		return h
	}

	b := make([]byte, 1)
	binary.PutVarint(b, int64(code))

	mask := b[0]

	if wink&mask != 0 {
		h = append(h, "wink")
	}
	if doubleBlink&mask != 0 {
		h = append(h, "double blink")
	}
	if closeEyes&mask != 0 {
		h = append(h, "close your eyes")
	}
	if jump&mask != 0 {
		h = append(h, "jump")
	}
	if reverse&mask != 0 {
		for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
			h[i], h[j] = h[j], h[i]
		}
	}
	return h

}

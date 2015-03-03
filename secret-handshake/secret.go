package secret

func Handshake(code int) []string {

	var h []string
	if code > 31 || code < 0 {
		return h
	}

	if 1&code != 0 {
		h = append(h, "wink")
	}
	if 2&code != 0 {
		h = append(h, "double blink")
	}
	if 4&code != 0 {
		h = append(h, "close your eyes")
	}
	if 8&code != 0 {
		h = append(h, "jump")
	}
	if 16&code != 0 {
		for i, j := 0, len(h)-1; i < j; i, j = i+1, j-1 {
			h[i], h[j] = h[j], h[i]
		}
	}
	return h

}

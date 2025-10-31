package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Store temporary values to avoid overwriting before using them
	tmpA := ***a
	tmpB := *b
	tmpC := *******c
	tmpD := ****d

	*******c = tmpA
	****d = tmpC
	*b = tmpD
	***a = tmpB
}

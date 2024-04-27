package iteration

import "bytes"

func Repeat(s string) string {
	var b bytes.Buffer

	for i := 0; i < 5; i++ {
		b.WriteString(s)
	}
	return b.String()
}

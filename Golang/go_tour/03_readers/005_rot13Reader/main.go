// در روت13 که روش رمزگذاری روم باستان است
// هر حرف، با 13مین حرف بعد از خود در حروف الفبا جابجا می‌شود
// چون جای حروف آ - اِم با از اِن - زِد عوض میشود
// و حروف لاتین 26 حرف است و 13مین حرف اِم است
// پس اگر حرفی قبل از اِم باشد 13+ و اگر بعدش باشد 13- میشود
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i := range b {
		//اگر حرفی قبل از اِن باشد +13 و اگر بعدش باشد -13 میشود
		if (b[i] >= 'a' && b[i] <= 'm') || (b[i] >= 'A' && b[i] <= 'M') {
			b[i] += 13
		} else if (b[i] >= 'n' && b[i] <= 'z') || (b[i] >= 'N' && b[i] <= 'Z') {
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) // os.Stdout == fmt.Print
}

//output: You cracked the code!       = Lbh penpxrq gur pbqr!

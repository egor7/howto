package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	A := 1.0
	B := 1.0
	for _ = range time.Tick(time.Second / 20) {
		A += .07
		B += .03
		b := make([]uint8, 1760)
		z := make([]float64, len(b))
		sA, cA := math.Sincos(A)
		sB, cB := math.Sincos(B)
		for k := 0; k < len(b); k++ {
			if k%80 == 79 {
				b[k] = '\n'
			} else {
				b[k] = ' '
			}
		}
		for j := 0.; j < 6.28; j += .07 {
			st, ct := math.Sincos(j)
			for i := 0.; i < 6.28; i += .02 {
				sp, cp := math.Sincos(i)
				h := ct + 2
				D := 1 / (sp*h*sA + st*cA + 5)
				t := sp*h*cA - st*
					sA
				x := int(40 + 30*D*(cp*h*cB-t*sB))
				y := int(12 + 15*
					D*(cp*h*sB+t*cB))
				o := x + 80*y
				N := int(8 * ((st*sA-sp*
					ct*cA)*cB - sp*ct*sA - st*cA - cp*ct*sB))
				if y < 22 &&
					y >= 0 && x >= 0 && x < 79 && D > z[o] {
					z[o] = D
					n := 0
					if N > 0 {
						n = N
					}
					b[o] = ".,-~:;=!*#$@"[n]
				}
			}
		}
		fmt.Print(
			"\x0c", string(b))
	}
} /****###****!
  .,~~;;;===..,--------,*/


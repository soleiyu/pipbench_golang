package pictFunc

import "math"

func Hdr(inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	p1 := Shine(inp)
	p2 := Shine(p1)
	p3 := Shine(p2)
	m1 := Dark(inp)
	m2 := Dark(m1)
	m3 := Dark(m2)

	for y := 0; y < inp.Height; y++ {
		for x := 0; x < inp.Width; x++ {

			if total(inp.Px[x][y]) < 255 {
				if total(p1.Px[x][y]) < 255 {
					res.Px[x][y] = p1.Px[x][y]
				}
				if total(p2.Px[x][y]) < 255 {
					res.Px[x][y] = p2.Px[x][y]
				}
				if total(p3.Px[x][y]) < 255 {
					res.Px[x][y] = p3.Px[x][y]
				}
			} else {
				if 255 < total(m1.Px[x][y]) {
					res.Px[x][y] = m1.Px[x][y] 
				}
				if 255 < total(m2.Px[x][y]) {
					res.Px[x][y] = m2.Px[x][y] 
				}
				if 255 < total(m3.Px[x][y]) {
					res.Px[x][y] = m3.Px[x][y] 
				}
			}
		}
	}
	
	return res
}

func Shine (inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for y := 0; y < inp.Height; y++ {
		for x := 0; x < inp.Width; x++ {
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = (uint8)(255 * math.Sin(0.5 * 3.1415926 * (float64)(inp.Px[x][y][1]) / 255.0))
			res.Px[x][y][2] = (uint8)(255 * math.Sin(0.5 * 3.1415926 * (float64)(inp.Px[x][y][2]) / 255.0))
			res.Px[x][y][3] = (uint8)(255 * math.Sin(0.5 * 3.1415926 * (float64)(inp.Px[x][y][3]) / 255.0))
		}
	}

	return res
}

func Dark (inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for y := 0; y < inp.Height; y++ {
		for x := 0; x < inp.Width; x++ {
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = 
				(uint8)(255 - 255 * math.Cos(0.5 * 3.1415926 * (float64)(inp.Px[x][y][1]) / 255.0))
			res.Px[x][y][2] = 
				(uint8)(255 - 255 * math.Cos(0.5 * 3.1415926 * (float64)(inp.Px[x][y][2]) / 255.0))
			res.Px[x][y][3] = 
				(uint8)(255 - 255 * math.Cos(0.5 * 3.1415926 * (float64)(inp.Px[x][y][3]) / 255.0))
		}
	}

	return res
}

func total(vals []uint8) int {
	t := 0
	t += (int)(vals[1])
	t += (int)(vals[2])
	t += (int)(vals[3])
	return t
}


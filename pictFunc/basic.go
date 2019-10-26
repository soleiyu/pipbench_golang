package pictFunc

func NGPS(a Pict) Pict {
	res := MkPict(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = 255 - a.Px[x][y][i]
				}
			}
		}

	return res
}

func AMul(a, b Pict) Pict {
	res := MkPict(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = (uint8)(255.0 * (1.0 - ((float64)(255 - a.Px[x][y][i]) / 255.0) * ((float64)(255 - b.Px[x][y][i]) / 255.0)))
				}
			}
		}

	return res
}

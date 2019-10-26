package pictFunc

func Sub(a, b Pict) Pict {
	res := MkPict(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
				if a.Px[x][y][i] - b.Px[x][y][i] < 0 {
					res.Px[x][y][i] = 0
				} else {
					res.Px[x][y][i] = a.Px[x][y][i] - b.Px[x][y][i]
				}
			}
		}
	}

	return res
}

func Mul(a, b Pict) Pict {
	res := MkPict(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = (uint8)(
						255.0 * 
						(float32)(a.Px[x][y][i]) * 
						(float32)(b.Px[x][y][i]) / 
						(255.0 * 255.0))
			}
		}
	}

	return res
}

package pictFunc

func FastSobel(inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		res.Px[x][0][0] = 255
		res.Px[x][0][1] = 0
		res.Px[x][0][2] = 0
		res.Px[x][0][3] = 0
		res.Px[x][inp.Height - 1][0] = 255
		res.Px[x][inp.Height - 1][1] = 0
		res.Px[x][inp.Height - 1][2] = 0
		res.Px[x][inp.Height - 1][3] = 0
	}
	for y := 1; y < inp.Height - 1; y++ {
		res.Px[0][y][0] = 255
		res.Px[0][y][1] = 0
		res.Px[0][y][2] = 0
		res.Px[0][y][3] = 0
		res.Px[inp.Width - 1][y][0] = 255
		res.Px[inp.Width - 1][y][1] = 0
		res.Px[inp.Width - 1][y][2] = 0
		res.Px[inp.Width - 1][y][3] = 0
	}

	return res
}


package pictFunc

func Wele2 (inp Pict) Pict {
	res:= MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for c := 0; c < 4; c++ {
				res.Px[x][y][c] = inp.Px[x][y][c]
			}
		}
	}

	for x := 1; x < inp.Width - 1; x++ {
		for y := 1; y < inp.Height - 1; y++ {
			ct := 0
			if(inp.Px[x - 1][y - 1][1] == 255){
				ct ++
			}
			if(inp.Px[x - 1][y][1] == 255){
				ct ++
			}
			if(inp.Px[x - 1][y + 1][1] == 255){
				ct ++
			}
			if(inp.Px[x][y - 1][1] == 255){
				ct ++
			}
			if(inp.Px[x][y + 1][1] == 255){
				ct ++
			}
			if(inp.Px[x + 1][y - 1][1] == 255){
				ct ++
			}
			if(inp.Px[x + 1][y][1] == 255){
				ct ++
			}
			if(inp.Px[x + 1][y + 1][1] == 255){
				ct ++
			}

			if(ct < 2){
				res.Px[x][y][1] = 0
				res.Px[x][y][2] = 0
				res.Px[x][y][3] = 0
			}
		}
	}

	return res
}

func Wfat (inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			res.Px[x][y][0] = 255
			res.Px[x][y][1] = 0
			res.Px[x][y][2] = 0
			res.Px[x][y][3] = 0
		}
	}

	for y := 0; y < inp.Height; y++ {
		for c := 1; c < 4; c++ {
			if(inp.Px[0][y][c] != 0){
				res.Px[0][y][c] = 255
				res.Px[1][y][c] = 255
			}
			if(inp.Px[inp.Width - 1][y][c] != 0){
				res.Px[inp.Width - 1][y][c] = 255
				res.Px[inp.Width - 2][y][c] = 255
			}
			for x := 1; x < inp.Width - 1; x++ {
				if(inp.Px[x][y][c] != 0){
					res.Px[x][y][c] = 255
					res.Px[x + 1][y][c] = 255
					res.Px[x - 1][y][c] = 255
				}
			}
		}
	}

	for x := 0; x < inp.Width; x++ {
		for c := 1; c < 4; c++ {
			if(inp.Px[x][0][c] != 0){
				res.Px[x][0][c] = 255
				res.Px[x][1][c] = 255
			}
			if(inp.Px[x][inp.Height - 1][c] != 0){
				res.Px[x][inp.Height - 1][c] = 255
				res.Px[x][inp.Height - 2][c] = 255
			}
			for y := 1; y < inp.Height - 1; y++ {
				if(inp.Px[x][y][c] != 0){
					res.Px[x][y][c] = 255
					res.Px[x][y + 1][c] = 255
					res.Px[x][y - 1][c] = 255
				}
			}
		}
	}

	return res
}


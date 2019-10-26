package pictFunc

func BinColor(inp Pict, hol int) Pict {
	res := MkPict(inp.Width, inp.Height)
	
	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			res.Px[x][y][0] = 255

			rv := (int)(inp.Px[x][y][1]) + (int)(inp.Px[x][y][2]) + (int)(inp.Px[x][y][3]) 
			
			if(rv < hol) {
				res.Px[x][y][1] = 0
				res.Px[x][y][2] = 0
				res.Px[x][y][3] = 0
			}	else {
				res.Px[x][y][1] = 255
				res.Px[x][y][2] = 255
				res.Px[x][y][3] = 255
			}
		}
	}

	return res
}

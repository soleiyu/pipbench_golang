package pictFunc

func Xhsample(a Pict) Pict {
	res := MkPict(a.Width / 2, a.Height / 2)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = 
						a.Px[2 * x][2 * y][i] / 2 + a.Px[2 * x + 1][2 * y + 1][i] / 2
				}
			}
		}

	return res
}

func X2sample(a Pict) Pict {
	res := MkPict(a.Width * 2, a.Height * 2)

	hc := make([][][]uint, res.Width)
	for x := 0; x < res.Width; x++ {
		hc[x] = make([][]uint, res.Height)
		for y := 0; y < res.Height; y++ {
			hc[x][y] = make([]uint, 4)
		}
	}

	for y := 0; y < res.Height; y++ {
		for x := 0; x < res.Width; x++ {
			res.Px[x][y][0] = 255

			if y == 0 {
				if x == 0 {
					hc[x][y][1] = uint(a.Px[x][y][1]) * 9
					hc[x][y][2] = uint(a.Px[x][y][2]) * 9
					hc[x][y][3] = uint(a.Px[x][y][3]) * 9	
				} else if x == res.Width - 1 {
					hc[x][y][1] = uint(a.Px[a.Width - 1][y][1]) * 9
					hc[x][y][2] = uint(a.Px[a.Width - 1][y][2]) * 9
					hc[x][y][3] = uint(a.Px[a.Width - 1][y][3]) * 9	
				} else if x % 2 == 1 {
					hc[x][y][1] = uint(a.Px[x / 2][y][1]) * 6 + uint(a.Px[x / 2 + 1][y][1]) * 3
					hc[x][y][2] = uint(a.Px[x / 2][y][2]) * 6 + uint(a.Px[x / 2 + 1][y][2]) * 3
					hc[x][y][3] = uint(a.Px[x / 2][y][3]) * 6 + uint(a.Px[x / 2 + 1][y][3]) * 3
				} else{
					hc[x][y][1] = uint(a.Px[x / 2 - 1][y][1]) * 3 + uint(a.Px[x / 2][y][1]) * 6
					hc[x][y][2] = uint(a.Px[x / 2 - 1][y][2]) * 3 + uint(a.Px[x / 2][y][2]) * 6
					hc[x][y][3] = uint(a.Px[x / 2 - 1][y][3]) * 3 + uint(a.Px[x / 2][y][3]) * 6
				}
			} else if y == res.Height - 1 {
				if x == 0 {
					hc[x][y][1] = uint(a.Px[x][a.Height - 1][1]) * 9
					hc[x][y][2] = uint(a.Px[x][a.Height - 1][2]) * 9
					hc[x][y][3] = uint(a.Px[x][a.Height - 1][3]) * 9
				} else if x == res.Width - 1 {
					hc[x][y][1] = uint(a.Px[a.Width - 1][a.Height - 1][1]) * 9
					hc[x][y][2] = uint(a.Px[a.Width - 1][a.Height - 1][2]) * 9
					hc[x][y][3] = uint(a.Px[a.Width - 1][a.Height - 1][3]) * 9	
				} else if x % 2 == 1 {
					hc[x][y][1] = uint(a.Px[x / 2][a.Height - 1][1]) * 6 + uint(a.Px[x / 2 + 1][a.Height - 1][1]) * 3
					hc[x][y][2] = uint(a.Px[x / 2][a.Height - 1][2]) * 6 + uint(a.Px[x / 2 + 1][a.Height - 1][2]) * 3
					hc[x][y][3] = uint(a.Px[x / 2][a.Height - 1][3]) * 6 + uint(a.Px[x / 2 + 1][a.Height - 1][3]) * 3
				} else{
					hc[x][y][1] = uint(a.Px[x / 2 - 1][a.Height - 1][1]) * 3 + uint(a.Px[x / 2][a.Height - 1][1]) * 6
					hc[x][y][2] = uint(a.Px[x / 2 - 1][a.Height - 1][2]) * 3 + uint(a.Px[x / 2][a.Height - 1][2]) * 6
					hc[x][y][3] = uint(a.Px[x / 2 - 1][a.Height - 1][3]) * 3 + uint(a.Px[x / 2][a.Height - 1][3]) * 6
				}
			} else if y % 2 == 1 {
				if x == 0 {
					hc[x][y][1] = uint(a.Px[x][y / 2][1]) * 6 + uint(a.Px[x][y / 2 + 1][1]) * 3
					hc[x][y][2] = uint(a.Px[x][y / 2][2]) * 6 + uint(a.Px[x][y / 2 + 1][2]) * 3
					hc[x][y][3] = uint(a.Px[x][y / 2][3]) * 6 + uint(a.Px[x][y / 2 + 1][3]) * 3
				} else if x == res.Width - 1 {
					hc[x][y][1] = uint(a.Px[a.Width - 1][y / 2][1]) * 6 + uint(a.Px[a.Width - 1][y / 2 + 1][1]) * 3
					hc[x][y][2] = uint(a.Px[a.Width - 1][y / 2][2]) * 6 + uint(a.Px[a.Width - 1][y / 2 + 1][2]) * 3
					hc[x][y][3] = uint(a.Px[a.Width - 1][y / 2][3]) * 6 + uint(a.Px[a.Width - 1][y / 2 + 1][3]) * 3
				} else if x % 2 == 1 {
					for i := 1; i < 4; i++ {
						hc[x][y][i] = 
							uint(a.Px[x / 2][y / 2][i]) * 4 + 
							uint(a.Px[x / 2 + 1][y / 2][i]) * 2 + 
							uint(a.Px[x / 2][y / 2 + 1][i]) * 2 + 
							uint(a.Px[x / 2 + 1][y / 2 + 1][i]) * 1
					}
				} else{
					for i := 1; i < 4; i++ {
						hc[x][y][i] = 
							uint(a.Px[x / 2 - 1][y / 2][i]) * 2 + 
							uint(a.Px[x / 2][y / 2][i]) * 4 + 
							uint(a.Px[x / 2 - 1][y / 2 + 1][i]) * 1 + 
							uint(a.Px[x / 2][y / 2 + 1][i]) * 2
					}
				}
			} else {
				if x == 0 {
					hc[x][y][1] = uint(a.Px[x][y / 2 - 1][1]) * 3 + uint(a.Px[x][y / 2][1]) * 6
					hc[x][y][2] = uint(a.Px[x][y / 2 - 1][2]) * 3 + uint(a.Px[x][y / 2][2]) * 6
					hc[x][y][3] = uint(a.Px[x][y / 2 - 1][3]) * 3 + uint(a.Px[x][y / 2][3]) * 6
				} else if x == res.Width - 1 {
					hc[x][y][1] = uint(a.Px[a.Width - 1][y / 2 - 1][1]) * 3 + uint(a.Px[a.Width - 1][y / 2][1]) * 6
					hc[x][y][2] = uint(a.Px[a.Width - 1][y / 2 - 1][2]) * 3 + uint(a.Px[a.Width - 1][y / 2][2]) * 6
					hc[x][y][3] = uint(a.Px[a.Width - 1][y / 2 - 1][3]) * 3 + uint(a.Px[a.Width - 1][y / 2][3]) * 6
				} else if x % 2 == 1 {
					for i := 1; i < 4; i++ {
						hc[x][y][i] = 
							uint(a.Px[x / 2][y / 2 - 1][i]) * 2 + 
							uint(a.Px[x / 2 + 1][y / 2 - 1][i]) * 1 + 
							uint(a.Px[x / 2][y / 2][i]) * 4 + 
							uint(a.Px[x / 2 + 1][y / 2][i]) * 2
					}
				} else{
					for i := 1; i < 4; i++ {
						hc[x][y][i] = 
							uint(a.Px[x / 2 - 1][y / 2 - 1][i]) * 1 + 
							uint(a.Px[x / 2][y / 2 - 1][i]) * 2 + 
							uint(a.Px[x / 2 - 1][y / 2][i]) * 2 + 
							uint(a.Px[x / 2][y / 2][i]) * 4
					}
				}
			}	
		}
	}

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			for i := 1; i < 4; i++ {
				res.Px[x][y][i] = uint8(hc[x][y][i] / 9)
			}
		}
	}

	return res
}


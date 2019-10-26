package pictFunc

import ( "fmt" )

func Exposure(inp []Pict) Pict {
	res := MkPict(inp[0].Width, inp[0].Height)

	fmt.Println("num, ", len(inp));

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			for i := 0; i < len(inp); i++ {
				for c := 0; c < 4; c++ {
					if(i == 0) {
						res.Px[x][y][c] = inp[i].Px[x][y][c]
					}
					if (res.Px[x][y][c] < inp[i].Px[x][y][c]) {
						res.Px[x][y][c] = inp[i].Px[x][y][c]
					}

				}
			}
		}
	}

	return res
}

func ExposureAve(inp []Pict) Pict {
	res := MkPict(inp[0].Width, inp[0].Height)

	fmt.Println("num, ", len(inp));

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			for c := 0; c < 4; c++ {
				rv := (int)(inp[0].Px[x][y][c])
				for i := 1; i < len(inp); i++ {
					rv += (int)(inp[i].Px[x][y][c])	
				}

				rv /= len(inp)
				res.Px[x][y][c] = uint8(rv)
			}
		}
	}

	return res
}

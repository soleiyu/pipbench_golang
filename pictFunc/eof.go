package pictFunc

//EseOpticalFlow

func EOF(bef, aft Pict) Pict {
	blpls := BinColor(Mylpls(bef), 150)
	alpls := BinColor(Mylpls(aft), 150)


	elealpls := Wele2(alpls)
	dif := difaft(blpls, Wfat(elealpls))

	rcache := div(blpls, dif, 10)
	rcache = Wfat(rcache)

	res := overWrite(aft, rcache)	

	return res
}

func div (blpls, alpls Pict, rv int) Pict {
	res := MkPict(blpls.Width, blpls.Height)

	for x := 0; x < blpls.Width; x++ {
		for y := 0; y < blpls.Height; y ++ {
			res.Px[x][y][0] = 255
			for c := 1; c < 4; c ++ {
				res.Px[x][y][c] = 0
			}
		}
	}

	//X SEARCH
	for y := 0; y < blpls.Height; y++	{
		for x := rv; x < blpls.Width - rv; x++ {
			
			fc := 0
			bc := 0

			for ix := 1; ix < rv; ix++ {
				if(alpls.Px[x][y][1] != 0) {
					if(blpls.Px[ix + x][y][1] != 0) {
						fc ++
					}
					if(blpls.Px[x - ix][y][1] != 0) {
						bc ++
					}
				}

				if(0 < fc) {	
					res.Px[x][y][1] = 255
				}
				if(0 < bc) {
					res.Px[x][y][2] = 255
				}
			}
		}
	}

	return res
}

func overWrite(inp, over Pict) Pict {
	res := MkPict(inp.Width, inp.Height)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			res.Px[x][y][0] = inp.Px[x][y][0]
			
			if (inp.Px[x][y][1] < over.Px[x][y][1]) {
				res.Px[x][y][1] = over.Px[x][y][1]
			} else {
				res.Px[x][y][1] = inp.Px[x][y][1]
			}
			
			if (inp.Px[x][y][2] < over.Px[x][y][2]) {
				res.Px[x][y][2] = over.Px[x][y][2]
			} else {
				res.Px[x][y][2] = inp.Px[x][y][2]
			}
			
			if (inp.Px[x][y][3] < over.Px[x][y][3]) {
				res.Px[x][y][3] = over.Px[x][y][3]
			} else {
				res.Px[x][y][3] = inp.Px[x][y][3]
			}
		}
	}

	return res
}

func difaft(bef, aft Pict) Pict {
	res := MkPict(bef.Width, bef.Height)

	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			res.Px[x][y][0] = 255

			if(bef.Px[x][y][1] != aft.Px[x][y][1]) {
				res.Px[x][y][1] = aft.Px[x][y][1]
			} else {
				res.Px[x][y][1] = 0
			}	
			if(bef.Px[x][y][2] != aft.Px[x][y][2]) {
				res.Px[x][y][2] = aft.Px[x][y][2]
			} else {
				res.Px[x][y][2] = 0
			}	
			if(bef.Px[x][y][3] != aft.Px[x][y][3]) {
				res.Px[x][y][3] = aft.Px[x][y][3]
			} else {
				res.Px[x][y][3] = 0
			}
		}
	}

	return res
}

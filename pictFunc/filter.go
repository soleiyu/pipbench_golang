package pictFunc

func Mylpls(inp Pict) Pict {
	return lpls(pad(inp, 1))
}

func pad(inp Pict, ofs int) Pict {
	res := MkPict(inp.Width + 2 * ofs, inp.Height + 2 * ofs)

	for x := 0; x < ofs; x++ {
		for y := 0; y < inp.Height; y++ {
			for c := 0; c < 4; c++ {
				res.Px[x][y + ofs][c] = inp.Px[0][y][c]
				res.Px[ofs + inp.Width + x][y + ofs][c] = inp.Px[inp.Width - 1][y][c]
			}
		}
	}
	for x := 0; x < res.Width; x++ {
		for y := 0; y < ofs; y++ {
			for c := 0; c < 4; c++ {
				res.Px[x][y][c] = res.Px[x][0][c]
				res.Px[x][ofs + inp.Height + y][c] = res.Px[x][inp.Height - 1][c]
			}
		}
	}

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for c := 0; c < 4; c++ {
				res.Px[x + ofs][y + ofs][c] = inp.Px[x][y][c]
			}
		}
	}

	return res
}

func lpls(inp Pict) Pict {
	res := MkPict(inp.Width - 2, inp.Height - 2)

	for x := 0; x < inp.Width - 2; x++ {
		for y := 0; y < inp.Height - 2; y++ {
		 	rval := int(inp.Px[x][y][1]) 
			rval += int(inp.Px[x + 1][y][1])
			rval += int(inp.Px[x + 2][y][1])
			rval += int(inp.Px[x][y + 1][1])
			rval -= 8 * int(inp.Px[x + 1][y + 1][1])
			rval += int(inp.Px[x + 2][y + 1][1])
			rval += int(inp.Px[x][y + 2][1])
			rval += int(inp.Px[x + 1][y + 2][1])
			rval += int(inp.Px[x + 2][y + 2][1])
		
			gval := int(inp.Px[x][y][2]) 
			gval += int(inp.Px[x + 1][y][2])
			gval += int(inp.Px[x + 2][y][2])
			gval += int(inp.Px[x][y + 1][2])
			gval -= 8 * int(inp.Px[x + 1][y + 1][2])
			gval += int(inp.Px[x + 2][y + 1][2])
			gval += int(inp.Px[x][y + 2][2])
			gval += int(inp.Px[x + 1][y + 2][2])
			gval += int(inp.Px[x + 2][y + 2][2])
	 	
			bval := int(inp.Px[x][y][3]) 
			bval += int(inp.Px[x + 1][y][3])
			bval += int(inp.Px[x + 2][y][3])
			bval += int(inp.Px[x][y + 1][3])
			bval -= 8 * int(inp.Px[x + 1][y + 1][3])
			bval += int(inp.Px[x + 2][y + 1][3])
			bval += int(inp.Px[x][y + 2][3])
			bval += int(inp.Px[x + 1][y + 2][3])
			bval += int(inp.Px[x + 2][y + 2][3])

			res.Px[x][y][0] = 255

			if (rval < 0){
				res.Px[x][y][1] = uint8(0)
			} else if (255 < rval) {
				res.Px[x][y][1] = uint8(255)
			} else {
				res.Px[x][y][1] = uint8(rval)
			}

			if (gval < 0){
				res.Px[x][y][2] = uint8(0)
			} else if (255 < gval) {
				res.Px[x][y][2] = uint8(255)
			} else {
				res.Px[x][y][2] = uint8(gval)
			}

			if (bval < 0){
				res.Px[x][y][3] = uint8(0)
			} else if (255 < bval) {
				res.Px[x][y][3] = uint8(255)
			} else {
				res.Px[x][y][3] = uint8(bval)
			}
		}
	}
		return res;
}

func anti(inp Pict) Pict {
	res := MkPict(inp.Width, inp.Height)
	
	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			res.Px[x][y][0] = 255
			res.Px[x][y][1] = 255 - inp.Px[x][y][1]
			res.Px[x][y][2] = 255 - inp.Px[x][y][2]
			res.Px[x][y][3] = 255 - inp.Px[x][y][3]
		}
	}

	return res;
}

func Average(a, b Pict) Pict {
	res := MkPict(a.Width, a.Height)
	
	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255
			res.Px[x][y][1] = (uint8)((a.Px[x][y][1] + b.Px[x][y][1]) / 2)
			res.Px[x][y][2] = (uint8)((a.Px[x][y][2] + b.Px[x][y][2]) / 2)
			res.Px[x][y][3] = (uint8)((a.Px[x][y][3] + b.Px[x][y][3]) / 2)
		}
	}

	return res;
}
func MyGaus(inp Pict, ks int) Pict {
	return MyGausX(MyGausY(inp, ks), ks)
}

func MyGausX(inp Pict, ks int) Pict {
	res := MkPict(inp.Width , inp.Height)
	st := int(ks / 2)
	cache := PushX(inp, st)
	sum := uint32(0)
	ker := make([]uint32, ks)

	for i := 0; i < st + 1; i++ {
		ker[i] = uint32(i + 1)
		ker[ks - 1 - i] = uint32(i + 1)
	}

	for i := 0; i < ks; i++ {
		sum += ker[i]
	}

  for x := 0; x < inp.Width; x++ {
  	for y := 0; y < inp.Height; y++ {
			pa := uint32(0)
			pr := uint32(0)
			pg := uint32(0)
			pb := uint32(0)

			for w := 0; w < ks; w++ {
 				pa += uint32(ker[w] * uint32(cache.Px[x + w][y][0]))
				pr += uint32(ker[w] * uint32(cache.Px[x + w][y][1]))
				pg += uint32(ker[w] * uint32(cache.Px[x + w][y][2]))
				pb += uint32(ker[w] * uint32(cache.Px[x + w][y][3]))
			}

			res.Px[x][y][0] = uint8(pa / sum)
			res.Px[x][y][1] = uint8(pr / sum)
			res.Px[x][y][2] = uint8(pg / sum)
			res.Px[x][y][3] = uint8(pb / sum)
			
		}
	}

	return res 
}

func MyGausY(inp Pict, ks int) Pict {
	res := MkPict(inp.Width , inp.Height)
	st := int(ks / 2)
	cache := PushY(inp, st)
	sum := uint32(0)
	ker := make([]uint32, ks)

	for i := 0; i < st + 1; i++ {
		ker[i] = uint32(i + 1)
		ker[ks - 1 - i] = uint32(i + 1)
	}

	for i := 0; i < ks; i++ {
		sum += ker[i]
	}

  for x := 0; x < inp.Width; x++ {
  	for y := 0; y < inp.Height; y++ {
			var pa uint32 = uint32(0)
			var pr uint32 = uint32(0)
			var pg uint32 = uint32(0)
			var pb uint32 = uint32(0)

			for w := 0; w < ks; w++ {
 				pa = pa + uint32(ker[w] * uint32(cache.Px[x][y + w][0]))
				pr = pr + uint32(ker[w] * uint32(cache.Px[x][y + w][1]))
				pg = pg + uint32(ker[w] * uint32(cache.Px[x][y + w][2]))
				pb = pb + uint32(ker[w] * uint32(cache.Px[x][y + w][3]))
			}

			res.Px[x][y][0] = uint8(pa / sum)
			res.Px[x][y][1] = uint8(pr / sum)
			res.Px[x][y][2] = uint8(pg / sum)
			res.Px[x][y][3] = uint8(pb / sum)
			
		}
	}
	
	return res 
}

func PushX(inp Pict, ofs int) Pict {
	res := MkPict(inp.Width + 2 * ofs, inp.Height)

	for x := 0; x < ofs; x++ {
		for y := 0; y < inp.Height; y++ {
			for i := 0; i < 4; i++ {
				res.Px[x][y][i] = inp.Px[0][y][i]
				res.Px[x + ofs + inp.Width][y][i] = inp.Px[inp.Width - 1][y][i]
			}
		}
	}

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for i := 0; i < 4; i++ {
				res.Px[ofs + x][y][i] = inp.Px[x][y][i]
			}
		}
	}

	return res
}

func PushY(inp Pict, ofs int) Pict {
	res := MkPict(inp.Width, inp.Height + 2 * ofs)

	for y := 0; y < ofs; y++ {
		for x := 0; x < inp.Width; x++ {
			for i := 0; i < 4; i++ {
				res.Px[x][y][i] = inp.Px[x][0][i]
				res.Px[x][y + ofs + inp.Height][i] = inp.Px[x][inp.Height - 1][i]
			}
		}
	}

	for y := 0; y < inp.Height; y++ {
		for x := 0; x < inp.Width; x++ {
			for i := 0; i < 4; i++ {
				res.Px[x][y + ofs][i] = inp.Px[x][y][i]
			}
		}
	}

	return res
}

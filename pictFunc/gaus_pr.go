package pictFunc

import(
	"sync"
	"runtime"
	"fmt"
)

func MyGaus_pr(inp Pict, ks int) Pict {
	cache := MyGausX_pr(MyGausY_pr(inp, ks), ks)

	return cache
}

func gKernelX(inp Pict, x, y, ks int, sum uint32, ker []uint32) (uint8, uint8, uint8) {
	pr := uint32(0)
	pg := uint32(0)
	pb := uint32(0)

	for w := 0; w < ks; w++ {
		pr += uint32(ker[w] * uint32(inp.Px[x + w][y][1]))
		pg += uint32(ker[w] * uint32(inp.Px[x + w][y][2]))
		pb += uint32(ker[w] * uint32(inp.Px[x + w][y][3]))
	}

	return uint8(pr / sum), uint8(pg / sum), uint8(pb / sum)	
}

func MyGausX_pr(inp Pict, ks int) Pict {
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
	
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	
	var wg sync.WaitGroup
	wg.Add(cpus)

	endpoint := make([]int, cpus + 1)
	for i := 0; i < cpus + 1; i++ {
		endpoint[i] = inp.Height * i / cpus
	}

	for prn := 0; prn < cpus; prn ++ {
		myn := prn
		go func() {
			fmt.Println(myn)
			for y := endpoint[myn]; y < endpoint[myn + 1]; y++ {
  			for x := 0; x < inp.Width; x++ {
					rr, rg, rb := gKernelX(cache, x, y, ks, sum, ker)
					res.Px[x][y][0] = cache.Px[x][y][0]
					res.Px[x][y][1] = uint8(rr)
					res.Px[x][y][2] = uint8(rg)
					res.Px[x][y][3] = uint8(rb)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return res 
}

func gKernelY(inp Pict, x, y, ks int, sum uint32, ker []uint32) (uint8, uint8, uint8) {
	pr := uint32(0)
	pg := uint32(0)
	pb := uint32(0)

	for w := 0; w < ks; w++ {
		pr += uint32(ker[w] * uint32(inp.Px[x][y + w][1]))
		pg += uint32(ker[w] * uint32(inp.Px[x][y + w][2]))
		pb += uint32(ker[w] * uint32(inp.Px[x][y + w][3]))
	}

	return uint8(pr / sum), uint8(pg / sum), uint8(pb / sum)	
}


func MyGausY_pr(inp Pict, ks int) Pict {
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
	
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	
	var wg sync.WaitGroup
	wg.Add(cpus)

	endpoint := make([]int, cpus + 1)
	for i := 0; i < cpus + 1; i++ {
		endpoint[i] = inp.Height * i / cpus
	}

	for prn := 0; prn < cpus; prn ++ {
		myn := prn
		go func() {
			fmt.Println(myn)
			for y := endpoint[myn]; y < endpoint[myn + 1]; y++ {
  			for x := 0; x < inp.Width; x++ {
					rr, rg, rb := gKernelY(cache, x, y, ks, sum, ker)
					res.Px[x][y][0] = cache.Px[x][y][0]
					res.Px[x][y][1] = uint8(rr)
					res.Px[x][y][2] = uint8(rg)
					res.Px[x][y][3] = uint8(rb)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	
	return res 
}

func PushX_pr(inp Pict, ofs int) Pict {
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

func PushY_pr(inp Pict, ofs int) Pict {
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

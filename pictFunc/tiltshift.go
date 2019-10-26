package pictFunc

import(
	"sync"
	"runtime"
	"fmt"
	"math"
)

func TiltShiftM(inp Pict) Pict {
	mbdst := getDist(0, 0, inp.Width, inp.Height)
	cache := PushX(PushY(inp, (int)(mbdst)), (int)(mbdst))
	res := MkPict(inp.Width , inp.Height)

	for y := 0; y < inp.Height; y++ {
		fmt.Printf("%d / %d \n", y + 1, inp.Height)
		for x := 0; x < inp.Width; x++ {
			ks := getDistKS(x, y, inp.Width, inp.Height, (int)(mbdst), mbdst)
			rr, rg, rb := csDot(cache, x + ks / 2, y + ks / 2, ks)
			res.Px[x][y][0] = inp.Px[x][y][0]
			res.Px[x][y][1] = rr
			res.Px[x][y][2] = rg
			res.Px[x][y][3] = rb
		}
	}
	return res
}

func TiltShiftP(inp Pict, ksm, xp, yp int) Pict {
	cache := PushX(PushY(inp, ksm), ksm)
	res := MkPict(inp.Width , inp.Height)
	mbdst := getDistP(0, 0, xp, yp)
	
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
					ks := getDistKS(x, y, xp, yp, ksm, mbdst)
					rr, rg, rb := csDot(cache, x + ksm, y + ksm, ks)
					res.Px[x][y][0] = inp.Px[x][y][0]
					res.Px[x][y][1] = rr
					res.Px[x][y][2] = rg
					res.Px[x][y][3] = rb
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return res
}

func TiltShift(inp Pict, ksm int) Pict {
	cache := PushX(PushY(inp, ksm), ksm)
	res := MkPict(inp.Width , inp.Height)
	mbdst := getDist(0, 0, inp.Width, inp.Height)
	
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
					ks := getDistKS(x, y, inp.Width, inp.Height, ksm, mbdst)
					rr, rg, rb := csDot(cache, x + ksm, y + ksm, ks)
					res.Px[x][y][0] = inp.Px[x][y][0]
					res.Px[x][y][1] = rr
					res.Px[x][y][2] = rg
					res.Px[x][y][3] = rb
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	return res
}

func getDistP(x, y, xp, yp int) float64 {
	return math.Sqrt(math.Pow((float64)(xp - x), 2) + 
		math.Pow((float64)(yp - y), 2))
}

func getDistKSP(x, y, xp, yp, ksm int, mbdst float64) int {
	dst := getDistP(x, y, xp, yp)
	ratio := dst / mbdst
	ks := (int)((float64)(ksm) * ratio *(ratio* 0.3 + 0.7))
	if ks % 2 == 0 {
		ks += 1
	}

	return ks
}


func getDist(x, y, w, h int) float64 {
	cx := w / 2
	cy := h / 2

	return math.Sqrt(math.Pow((float64)(cx -x), 2) + 
		math.Pow((float64)(cy - y), 2))
}

func getDistKS(x, y, w, h, ksm int, mbdst float64) int {
	dst := getDist(x, y, w, h)
	ratio := dst / mbdst
	ks := (int)((float64)(ksm) * ratio *(ratio* 0.3 + 0.7))
	if ks % 2 == 0 {
		ks += 1
	}

	return ks
}

func csDot(inp Pict, x, y, ks int) (uint8, uint8, uint8) {
	pr := uint32(0)
	pg := uint32(0)
	pb := uint32(0)
	sum := uint32(0)
	ker := make([]uint32, ks)
	cr := make([]uint32, ks)
	cg := make([]uint32, ks)
	cb := make([]uint32, ks)

	hks := ks / 2
	
	for i := 0; i < hks + 1; i++ {
		ker[i] = uint32(i + 1)
		ker[ks - 1 - i] = uint32(i + 1)
	}

	for i := 0; i < ks; i++ {
		sum += ker[i]
	}

	// X
	for ix := 0; ix < ks; ix ++ {
		for w := 0; w < ks; w++ {
			cr[ix] += uint32(ker[w] * uint32(inp.Px[x + w - hks][y + ix -hks][1]))
			cg[ix] += uint32(ker[w] * uint32(inp.Px[x + w - hks][y + ix -hks][2]))
			cb[ix] += uint32(ker[w] * uint32(inp.Px[x + w - hks][y + ix -hks][3]))
		}
		cr[ix] /= sum
		cg[ix] /= sum
		cb[ix] /= sum
	}

	//Y
	for h := 0; h < ks; h++ {
		pr += uint32(ker[h] * cr[h])
		pg += uint32(ker[h] * cg[h])
		pb += uint32(ker[h] * cb[h])
	}

	return uint8(pr / sum), uint8(pg / sum), uint8(pb / sum)	
}


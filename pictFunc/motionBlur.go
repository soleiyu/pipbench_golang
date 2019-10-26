package pictFunc

import(
	"sync"
	"runtime"
	"fmt"
)

func MotionBlur_pr(inp Pict, size, pnum int) Pict {
	res := MkPict(inp.Width , inp.Height)
	
	cpus := runtime.NumCPU()
	println("Num CPU :", cpus)

	runtime.GOMAXPROCS(cpus)
	var wg sync.WaitGroup
	wg.Add(cpus)

	pxs_total := inp.Width * inp.Height
	pxs_perc := pxs_total / cpus
	pxs_left := pxs_total % cpus

	for prn := 0; prn < cpus; prn ++ {
		myn := prn

		go func() {
			fmt.Println(myn)

			rnum := pxs_perc
			if myn < pxs_left {
				rnum ++
			}

			for i := 0; i < rnum; i++ {
				pvs := []int{0, 0, 0, 0}
				sum := 0

					x := (cpus * i + myn) % inp.Width
					y := (cpus * i + myn) / inp.Width

				for n := 0; n < size; n++ {
					cac := mbc(inp, x, y, (n + 1) * 10)
					sum += (size - n)

					for j := 0; j < 4; j++ {
						pvs[j] += (size - n) * int(cac[j])
					}
				}

				for j := 1; j < 4; j++ {
					res.Px[x][y][j] = uint8(pvs[j] / sum)
				}
				res.Px[x][y][0] = uint8(255)

			}
			wg.Done()
		}()
	}

	wg.Wait()

	return res
}

func mbc(inp Pict, x, y, s int) []uint8 {

	ratio := float32(inp.Height) / float32(inp.Height + s) 

	cx := inp.Width / 2
	cy := inp.Height / 2

	return gfp(inp,
		float32(cx) + float32(x - cx) * ratio,
		float32(cy) + float32(y - cy) * ratio)
}

func gfp(inp Pict, fx, fy float32) []uint8 {
	xi := int(fx)
	yi := int(fy)
	rx := fx - float32(xi)
	ry := fy - float32(yi)
	nx := float32(1) - rx
	ny := float32(1) - ry

	res := make([]uint8, 4)

	for i := 0; i < 4; i++ {
		res[i] = uint8(
			float32(inp.Px[xi][yi][i]) * rx * ry +
			float32(inp.Px[xi + 1][yi][i]) * nx * ry +
			float32(inp.Px[xi][yi + 1][i]) * rx * ny +
			float32(inp.Px[xi + 1][yi + 1][i]) * nx * ny)
	}

	return res
}

package pictFunc

import(
	"sync"
	"runtime"
)

type Complex struct {
  Rn float64
  In float64
}

func MkComp(rn , in float64) Complex {
  var c Complex
  c.Rn = rn
  c.In = in
  return c
}

func MandelFunc (zn, c Complex) Complex {
  return MkComp(
    zn.Rn * zn.Rn - zn.In * zn.In + c.Rn,
    2.0 * zn.In * zn.Rn + c.In)
}

func (zn Complex) DistSq () float64 {
  return zn.Rn * zn.Rn + zn.In * zn.In
}

func IsMandel(x, y ,zx, zy float64, lim int) int {
  z := MkComp(zx, zy)
  c := MkComp(x, y)

  for i := 0; i < lim; i++ {
    z = MandelFunc(z, c)
    if 4 < z.DistSq() {
      return i
    }
  }

  return lim
}

func MkMandelMapQ_pr(w, h, cnum int, rn, cpr, in, zx, zy, cpy float64) Pict {
	p := MkPict(w, h)
  qw := w * 5
  qh := h * 5
  hw := qw / 2
  hh := qh / 2

	cpus := runtime.NumCPU()

	if cnum < cpus {
		cpus = cnum
	}
//	println("Num CPU :", cpus)

	runtime.GOMAXPROCS(cpus)
	var wg sync.WaitGroup
	wg.Add(cpus)

	pxs_total := w * h
	pxs_perc := pxs_total / cpus
	pxs_left := pxs_total % cpus

	for prn := 0; prn < cpus; prn ++ {
		myn := prn

		go func() {
			rnum := pxs_perc

			if myn < pxs_left {
				rnum ++
			}

			for i := 0; i < rnum; i++ {
				x := (cpus * i + myn) % w
				y := (cpus * i + myn) / w

				cp := MkPict(5, 5)

				for xx := 0; xx < 5; xx++ {
					for yy := 0; yy < 5; yy++ {
			      pxv := IsMandel(
							rn * (float64)(5 * x + xx - hw) / (float64)(hw) +cpr,
							in * (float64)(5 * y + yy - hh) / (float64)(hh) +cpy,
							zx, zy, 1791)
						cp.Px[xx][yy] = colorMap(pxv)
					}
				}

				p.Px[x][y] = gaus5(cp)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return p

}

func gaus5(cp Pict) []uint8 {
	res := make([]uint8, 4)
	res[0] = 255

	for i := 1; i < 4; i++ {
		cache := int(cp.Px[0][0][i]) +
			4  * int(cp.Px[1][0][i]) +
			6  * int(cp.Px[2][0][i]) +
			4  * int(cp.Px[3][0][i]) +
			1  * int(cp.Px[4][0][i]) +
			4  * int(cp.Px[0][1][i]) +
			16 * int(cp.Px[1][1][i]) +
			24 * int(cp.Px[2][1][i]) +
			16 * int(cp.Px[3][1][i]) +
			4  * int(cp.Px[4][1][i]) +
			6  * int(cp.Px[0][2][i]) +
			24 * int(cp.Px[1][2][i]) +
			36 * int(cp.Px[2][2][i]) +
			24 * int(cp.Px[3][2][i]) +
			6  * int(cp.Px[4][2][i]) +
			4  * int(cp.Px[0][3][i]) +
			16 * int(cp.Px[1][3][i]) +
			24 * int(cp.Px[2][3][i]) +
			16 * int(cp.Px[3][3][i]) +
			4  * int(cp.Px[4][3][i]) +
			1  * int(cp.Px[0][4][i]) +
			4  * int(cp.Px[1][4][i]) +
			6  * int(cp.Px[2][4][i]) +
			4  * int(cp.Px[3][4][i]) +
			1  * int(cp.Px[4][4][i])
		cache /= 256
		res[i] = uint8(cache)
	}

	return res
}

func colorMap(v int) []uint8 {
	res := make([]uint8, 4)
	res[0] = 255

	if v < 256 {
		res[1] = uint8(v)
		res[2] = uint8(0)
		res[3] = uint8(0)
	} else if v < 512 {
		res[1] = uint8(255)
		res[2] = uint8(v - 255)
		res[3] = uint8(0)
	} else if v < 768 {
		res[1] = uint8(767 - v)
		res[2] = uint8(255)
		res[3] = uint8(0)
	} else if v < 1024 {
		res[1] = uint8(0)
		res[2] = uint8(255)
		res[3] = uint8(v - 767)
	} else if v < 1280 {
		res[1] = uint8(0)
		res[2] = uint8(1279 - v)
		res[3] = uint8(255)
	} else if v < 1536 {
		res[1] = uint8(v - 1279)
		res[2] = uint8(v - 1279)
		res[3] = uint8(255)
	} else {
		res[1] = uint8(255)
		res[2] = uint8(255)
		res[3] = uint8(255)
	}

	return res
}


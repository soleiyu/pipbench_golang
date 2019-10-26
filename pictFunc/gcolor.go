package pictFunc

type Gcolor struct{
	A int64
	R int64
	G int64
	B int64
}

func NewGcolor(a int64, r int64, g int64, b int64) *Gcolor{
	res := new(Gcolor)

	res.A = a
	res.R = r
	res.G = g
	res.B = b

	return res
}

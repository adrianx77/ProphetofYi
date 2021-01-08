package api

type YaoShu int

const (
	Laoyin   YaoShu = 6
	Shaoyang YaoShu = 7
	Shaoyin  YaoShu = 8
	Laoyang  YaoShu = 9
)

type Yao struct {
	Yang bool
	Bian bool
}

type ResultOfQigua struct {
	Bengua []YaoShu
}

type ZhanbuName string

const (
	Chouce ZhanbuName = "筹策"
)

type Zhanbu interface {
	CreateStage()
	Qigua()
	Dump()
	DumpGuaXiang()
}

type ZhanbuFactory interface {
	CreateZhanbu(name ZhanbuName) Zhanbu
}

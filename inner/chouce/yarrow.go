package chouce

import (
	"ProphetOfYi/inner/api"
	"ProphetOfYi/inner/util"
	"fmt"
	"strconv"
)

type Chouce string

type ChouceGroup struct {
	cc    [50]string
	count int
}
type ChouceStage struct {
	Using        ChouceGroup
	Taiji        string
	Tian         ChouceGroup
	Di           ChouceGroup
	Xiao_Wuming  ChouceGroup
	Wuming_Zhong ChouceGroup
	Zhong_Shi    ChouceGroup
	BianArray    [3]ChouceGroup
	YaoShu       [6]api.YaoShu
	Yao          [6]api.Yao
}

func (stage *ChouceStage) Dump() {
	dumCG("Using", &stage.Using)
	fmt.Println("taiji:", stage.Taiji)
	dumCG("Tian", &stage.Tian)
	dumCG("Di", &stage.Di)
	dumCG("Xiao_Wuming", &stage.Xiao_Wuming)
	dumCG("Wuming_Zhong", &stage.Wuming_Zhong)
	dumCG("Zhong_Shi", &stage.Zhong_Shi)
	dumCG("BianArray[0]", &stage.BianArray[0])
	dumCG("BianArray[1]", &stage.BianArray[1])
	dumCG("BianArray[2]", &stage.BianArray[2])

	fmt.Println("YaoShu6:", stage.YaoShu[5])
	fmt.Println("YaoShu5:", stage.YaoShu[4])
	fmt.Println("YaoShu4:", stage.YaoShu[3])
	fmt.Println("YaoShu3:", stage.YaoShu[2])
	fmt.Println("YaoShu2:", stage.YaoShu[1])
	fmt.Println("YaoShu1:", stage.YaoShu[0])
	for i := 5; i >= 0; i-- {
		dumpGuaXiang(stage.YaoShu[i])
	}
}

func dumCG(name string, cg *ChouceGroup) {
	fmt.Println(name, ".cc", cg.cc)
	fmt.Println(name, ".count", cg.count)
}

func (stage * ChouceStage) DumpGuaXiang(){
	for i := 5; i >= 0; i-- {
		dumpGuaXiang(stage.YaoShu[i])
	}
}



func dumpGuaXiang(shu api.YaoShu) {
	if shu == api.Laoyin {
		fmt.Println("▆▆  ▆▆✕")
	} else if shu == api.Shaoyang {
		fmt.Println("▆▆▆▆▆")
	} else if shu == api.Shaoyin {
		fmt.Println("▆▆  ▆▆")
	} else {
		fmt.Println("▆▆▆▆▆◯")
	}
}

func (stage *ChouceStage) CreateStage() {
	stage.createDayan()
	shuffle(&stage.Using, util.Random)
}

func initChoueGroup(cg *ChouceGroup) {
	cg.count = 0
}

func pushChouce(cg *ChouceGroup, cc string) {
	cg.cc[cg.count] = cc
	cg.count++
}

func popChouce(target *ChouceGroup, index int) string {
	if index >= target.count {
		return ""
	}
	outString := target.cc[index]
	j := index
	for ; j < target.count-1; j++ {
		target.cc[j] = target.cc[j+1]
	}
	target.cc[j] = ""
	target.count--
	return outString
}

func moveTo(from *ChouceGroup, to *ChouceGroup) {
	inputCount := from.count
	for i := 0; i < inputCount; i++ {
		pushChouce(to, popChouce(from, 0))
	}
}

func (stage *ChouceStage) createDayan() {
	for i := 0; i < 50; i++ {
		pushChouce(&stage.Using, strconv.Itoa(i+1))
	}
}

func shuffle(input *ChouceGroup, RanFunc func(int, int) int) {
	cgTemp := &ChouceGroup{}
	inputCount := input.count
	for i := 0; i < inputCount; i++ {
		pushChouce(cgTemp, popChouce(input, 0))
	}

	for i := inputCount; i > 0; i-- {
		r := RanFunc(0, i)
		pushChouce(input, popChouce(cgTemp, r))
	}
}

func (stage *ChouceStage) YaoBian() {
	//Tian
	tianCount := util.Random(4, stage.Using.count-4)
	for i := 0; i < tianCount; i++ {
		r := util.Random(0, tianCount-i)
		pushChouce(&stage.Tian, popChouce(&stage.Using, r))
	}
	//Di
	diCount := stage.Using.count
	for i := 0; i < diCount; i++ {
		r := util.Random(0, diCount-i)
		pushChouce(&stage.Di, popChouce(&stage.Using, r))
	}

	// Di
	tempFromDi := util.Random(0, diCount)
	pushChouce(&stage.Xiao_Wuming, popChouce(&stage.Di, tempFromDi))

	// Tian
	DivBy4(&stage.Tian, &stage.Using)
	moveTo(&stage.Tian, &stage.Wuming_Zhong)
	//Di
	DivBy4(&stage.Di, &stage.Using)
	moveTo(&stage.Di, &stage.Zhong_Shi)
}

func (stage *ChouceStage) Qigua() {
	taijiIndex := util.Random(0, 50)
	stage.Taiji = popChouce(&stage.Using, taijiIndex)
	for i := 0; i < 6; i++ {
		stage.YaoShu[i] = stage.DeYao()
		stage.Yao[i].Bian = stage.YaoShu[i] == api.Laoyin || stage.YaoShu[i] == api.Laoyang
		stage.Yao[i].Yang = (stage.YaoShu[i] % 2) == 1
	}
}

func (stage *ChouceStage) DeYao() api.YaoShu {
	stage.YaoBian()
	HeBian(stage)

	stage.YaoBian()
	HeBian(stage)

	stage.YaoBian()
	HeBian(stage)
	poolCount := stage.Using.count
	moveTo(&stage.BianArray[0], &stage.Using)
	moveTo(&stage.BianArray[1], &stage.Using)
	moveTo(&stage.BianArray[2], &stage.Using)
	shuffle(&stage.Using, util.Random)
	return api.YaoShu(poolCount / 4)
}

func DivBy4(cg *ChouceGroup, pool *ChouceGroup) {
	temp := ""
	for cg.count > 4 {
		for i := 0; i < 4; i++ {
			temp = popChouce(cg, 0)
			pushChouce(pool, temp)
		}
	}
}

func HeBian(stage *ChouceStage) {
	index := 0
	for stage.BianArray[index].count > 0 {
		index++
	}
	moveTo(&stage.Xiao_Wuming, &stage.BianArray[index])
	moveTo(&stage.Wuming_Zhong, &stage.BianArray[index])
	moveTo(&stage.Zhong_Shi, &stage.BianArray[index])
}

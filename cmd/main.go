package main

import (
	"ProphetOfYi/inner/api"
	"ProphetOfYi/inner/chouce"
	"ProphetOfYi/inner/util"
)

func main() {
	util.Initialize()
	chouceInstance := chouce.GetFactory().CreateZhanbu(api.Chouce)
	chouceInstance.Qigua()
	chouceInstance.DumpGuaXiang()
}

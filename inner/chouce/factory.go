package chouce

import "ProphetOfYi/inner/api"

type ZhanbuFactoryImp struct {
}

func (*ZhanbuFactoryImp) CreateZhanbu(name api.ZhanbuName) api.Zhanbu {
	if name == api.Chouce {
		stage := &ChouceStage{}
		stage.CreateStage()
		return stage
	}
	return nil
}

var factory *ZhanbuFactoryImp = nil

func GetFactory() *ZhanbuFactoryImp {
	if factory == nil {
		factory = &ZhanbuFactoryImp{}
	}
	return factory
}

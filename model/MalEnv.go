package model

type MalEnv struct {
	outer *MalEnv
	data  map[string]MalForm
}

func (env *MalEnv) Set(key string, value MalForm) {
	env.data[key] = value
}

func (env *MalEnv) Find(key string) *MalEnv {
	if _, exists := env.data[key]; exists {
		return env
	}
	if env.outer != nil {
		return env.outer.Find(key)
	}
	return nil
}

func (env *MalEnv) Get(key string) MalForm {
	envWithKey := env.Find(key)
	if envWithKey == nil {
		return MalForm{
			Err: MalErr("symbol " + key + " not found"),
		}
	}
	return envWithKey.data[key]
}

func NewMalEnv(outer *MalEnv, init map[string]MalForm) *MalEnv {
	initData := map[string]MalForm{}
	for key, value := range init {
		initData[key] = value
	}
	return &MalEnv{outer, initData}
}

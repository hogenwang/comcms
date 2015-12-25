package models

//获取系统设置
func GetConfig() *Config {
	config := &Config{Id: 1}
	has, err := x.Get(config)
	if has && err == nil {
		return config
	} else {
		return nil
	}
}

//编辑系统设置
func EditConfig(cfg *Config) error {
	if _, err := x.Id(cfg.Id).Update(cfg); err != nil {
		return err
	} else {
		return nil
	}
}

package config

type UserConfig struct {
	Port     string
	DbConfig struct {
		DBName  string
		ColName string
	}
}

var cfgs = map[string]UserConfig{
	"prod": {
		Port: ":8005",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "user",
		},
	},
	"qa": {
		Port: ":8005",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "user",
		},
	},
	"dev": {
		Port: ":8005",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "user",
		},
	},
}

func GetUserConfig(env string) *UserConfig {
	config, isExist := cfgs[env]
	if !isExist {
		panic("config does not exist")
	}
	return &config
}

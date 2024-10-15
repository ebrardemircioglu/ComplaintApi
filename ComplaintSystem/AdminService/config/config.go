package config

type AdminConfig struct {
	Port     string
	DbConfig struct {
		DBName  string
		ColName string
	}
}

var cfgs = map[string]AdminConfig{
	"prod": {
		Port: ":8006",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "admin",
		},
	},
	"qa": {
		Port: ":8006",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "admin",
		},
	},
	"dev": {
		Port: ":8006",
		DbConfig: struct {
			DBName  string
			ColName string
		}{
			DBName:  "complaintSystem",
			ColName: "admin",
		},
	},
}

func GetAdminConfig(env string) *AdminConfig {
	config, isExist := cfgs[env]
	if !isExist {
		panic("config does not exist")
	}
	return &config
}

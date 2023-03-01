package config

const ConfigApp = `{
	"env":"local",
	"port":"1909",
	"mysql":{
		"host":"127.0.0.1",
        "port":"3306",
        "user":"admin",
        "password":"admin",
        "db":"pth"
	},
	"monggodb":{
		"host":"mongodb://localhost:27017%s%s",
        "user":"",
        "password":"",
        "db":"pth"
	},
	"redis":{
        "host":"localhost",
		"port":"6379",
        "db":"0",
		"password":""
    }	
	
	
}`

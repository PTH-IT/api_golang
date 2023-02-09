package config

const ConfigApp = `{
	"env":"local",
	"port":":1909",
	"mysql":{
		"host":"127.0.0.1",
        "port":"3306",
        "user":"admin",
        "password":"admin",
        "db":"pth"
	},
	"monggodb":{
		"host":"mongodb+srv://%s:%s@cluster0.gqs073o.mongodb.net/?retryWrites=true&w=majority",
        "user":"haupham",
        "password":"haupham809",
        "db":"pth"
	},
	"redis":{
        "host":"localhost",
		"port":"6379",
        "db":0,
		"password":""
    }	
	
	
}`

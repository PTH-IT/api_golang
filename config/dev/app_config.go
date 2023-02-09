package config

const ConfigApp = `{
	"env":"dev",
	"port":":1909",
	"mysql":{
		"host":"",
        "port":"",
        "user":"",
        "password":"",
        "db":""
	},
	"monggodb":{
		"host":'mongodb+srv://%s:%s@cluster0.gqs073o.mongodb.net/?retryWrites=true&w=majority',
        "user":"haupham",
        "password":"haupham809",
        "db":"pth"
	}
}`

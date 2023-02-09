package config

const ConfigApp = `{
	"env":"prod",
	"port":":80",
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

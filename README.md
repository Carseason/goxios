### GET
    goxios.Get(goxios.Config{
        Url: "",
    }).Then(func(res goxios.Axios) {
    }).Catch(func(err error) {
    })

### POST
    goxios.Post(goxios.Config{
        Url: "",
        Headers: Headers{
			UserAgent:   "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Mobile Safari/537.36",
			Cookie:      "",
			ContentType: "application/json",
		},
        Data:map[string]interface{}{},
    }).Then(func(res goxios.Axios) {
    }).Catch(func(err error) {
    })
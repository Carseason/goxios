### GET
    goxios.Get(goxios.Config{
        Url: "",
    }).Then(func(res goxios.Goxios) {
    }).Catch(func(err error) {
    })

### POST
    goxios.Post(goxios.Config{
        Url: "",
        Headers: goxios.Headers{
			UserAgent:   "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Mobile Safari/537.36",
			Cookie:      "",
			ContentType: "application/json",
		},
        Data:nil,
    }).Then(func(res goxios.Goxios) {
    }).Catch(func(err error) {
    })
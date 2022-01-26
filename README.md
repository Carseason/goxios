### GET
    goxios.Do(RequestConfig{
		Method: goxios.GetMethod,
		Url:    "https://www.baidu.com",
	}).Then(func(responseData Response) {
	}).Catch(func(responseError error) {
	})

### POST
     goxios.Do(RequestConfig{
		Method: goxios.PostMethod,
		Url:    "https://www.baidu.com",
        Data:   nil
	}).Then(func(responseData Response) {
	}).Catch(func(responseError error) {
	})
package models

type RespResult struct {
	Url         string
	Available   bool
	ReqDuration int64
	Err         error
}

func SetResult(reqDuration int64, url string, err error) RespResult {
	if err != nil {
		return RespResult{
			Url:       url,
			Available: false,
			Err:       err,
		}
	}
	return RespResult{
		Url:         url,
		Available:   true,
		ReqDuration: reqDuration,
		Err:         err,
	}
}

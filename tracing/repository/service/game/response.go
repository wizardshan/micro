package game

type Response struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
	Data any     `json:"data"`
}

func (resp *Response) Success() bool {
	return resp.Code == 200
}

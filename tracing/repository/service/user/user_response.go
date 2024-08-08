package user

type Response struct {
	Code int     `json:"code"`
	Msg  *string `json:"msg"`
}

func (resp *Response) Success() bool {
	return resp.Code == 200
}

type Users []*User

type User struct {
	ID       int    `json:"uid"`
	Nickname string `json:"uname"`
	Phone    string `json:"phone"`
}

type UserByIDMapping map[int]*User

func (*UserByIDMapping) name() {

}

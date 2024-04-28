package response

type BaseResponse[R any] struct {
	Msg  string `json:"msg"`
	Item R      `json:"item"`
}

type OnlyIdAndMsg struct {
	Msg string `json:"msg"`
	ID  string `json:"id"`
}

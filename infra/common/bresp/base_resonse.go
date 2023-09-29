package bresp

type BaseRespone[T any] struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    T      `json:"data,omitempty"`
	Success bool   `json:"success"`
}

func IsSuccess[T any](rsp *BaseRespone[T]) bool {
	return rsp != nil && rsp.Success
}

func OfSuccess[T any](data T) *BaseRespone[T] {
	return &BaseRespone[T]{
		Success: true,
		Msg:     "SUCCESS",
		Data:    data,
	}
}

func OfErr[T any](code int, msg string) *BaseRespone[T] {
	return &BaseRespone[T]{
		Code: code,
		Msg:  msg,
	}
}

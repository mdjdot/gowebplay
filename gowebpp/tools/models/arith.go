package models

type Arith struct {
}
type Request struct {
	A, B int
}
type Response struct {
	C int
}

func (a *Arith) Sum(req *Request, resp *Response) error {
	resp.C = req.A + req.B
	return nil
}

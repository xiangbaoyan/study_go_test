package rpcdemo

import "github.com/pkg/errors"

type DemoService struct{}
type Args struct {
	A, B int
}

//直接在cmd 上 用这个TCP 链接进行输出 {"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("divisor can't be zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}

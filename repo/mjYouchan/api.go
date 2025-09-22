package mjyouchan

type MjYouchuanClient interface {
}

func NewMjYouchuanClient() MjYouchuanClient {
	return mjYouchuanImpl{}
}

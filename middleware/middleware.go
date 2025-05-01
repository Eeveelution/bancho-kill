package middleware

type Middleware interface {
	Process(data []byte)
}

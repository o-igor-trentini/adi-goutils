package uchan

type ChannelResponse[T any] struct {
	Data  T
	Error error
}

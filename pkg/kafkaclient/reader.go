package kafkaclient

import "context"

type reader struct {
}

func (r *reader) start(ctx context.Context, output chan<- *StreamingMessage) {

}

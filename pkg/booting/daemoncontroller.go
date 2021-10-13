package booting

import "context"

type daemonController struct {
	ctx     context.Context
	cancel  context.CancelFunc
	cleanup func()
}

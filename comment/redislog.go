/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 13:23:38
 */
package comment

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisLog struct {
	Id int
}

func (rl RedisLog) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return context.Background(), nil
}

func (rl RedisLog) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	fmt.Println(cmd.String())
	return nil
}

func (rl RedisLog) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return context.Background(), nil
}

func (rl RedisLog) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	return nil
}

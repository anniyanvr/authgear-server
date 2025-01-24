package webapp

import (
	"context"
	"encoding/json"

	goredis "github.com/redis/go-redis/v9"

	"github.com/authgear/authgear-server/pkg/auth/webapp"
	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	"github.com/authgear/authgear-server/pkg/util/pubsub"
)

type Publisher struct {
	AppID       config.AppID
	RedisHandle *appredis.Handle
	Publisher   *pubsub.Publisher
}

func NewPublisher(appID config.AppID, handle *appredis.Handle) *Publisher {
	p := &Publisher{
		AppID:       appID,
		RedisHandle: handle,
	}
	p.Publisher = &pubsub.Publisher{
		RedisPool: p,
	}
	return p
}

func (p *Publisher) Get() *goredis.Client {
	return p.RedisHandle.Client()
}

func (p *Publisher) Publish(ctx context.Context, s *webapp.Session, msg *WebsocketMessage) error {
	channelName := WebsocketChannelName(string(p.AppID), s.ID)

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = p.Publisher.Publish(ctx, channelName, b)
	if err != nil {
		return err
	}

	return nil
}

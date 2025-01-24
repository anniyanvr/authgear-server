package challenge

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	goredis "github.com/redis/go-redis/v9"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis"
	"github.com/authgear/authgear-server/pkg/lib/infra/redis/appredis"
	"github.com/authgear/authgear-server/pkg/util/clock"
)

type Provider struct {
	Redis *appredis.Handle
	AppID config.AppID
	Clock clock.Clock
}

func (p *Provider) Create(ctx context.Context, purpose Purpose) (*Challenge, error) {
	now := p.Clock.NowUTC()
	ttl := purpose.ValidityPeriod()
	c := &Challenge{
		Token:     GenerateChallengeToken(),
		Purpose:   purpose,
		CreatedAt: now,
		ExpireAt:  now.Add(ttl),
	}

	key := challengeKey(p.AppID, c.Token)
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	err = p.Redis.WithConnContext(ctx, func(ctx context.Context, conn redis.Redis_6_0_Cmdable) error {
		_, err = conn.SetNX(ctx, key, data, ttl).Result()
		if errors.Is(err, goredis.Nil) {
			return errors.New("fail to create new challenge")
		} else if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (p *Provider) Get(ctx context.Context, token string) (*Challenge, error) {
	key := challengeKey(p.AppID, token)

	c := &Challenge{}

	err := p.Redis.WithConnContext(ctx, func(ctx context.Context, conn redis.Redis_6_0_Cmdable) error {
		data, err := conn.Get(ctx, key).Bytes()
		if errors.Is(err, goredis.Nil) {
			return ErrInvalidChallenge
		} else if err != nil {
			return err
		}

		err = json.Unmarshal(data, c)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (p *Provider) Consume(ctx context.Context, token string) (*Purpose, error) {
	key := challengeKey(p.AppID, token)

	c := &Challenge{}

	err := p.Redis.WithConnContext(ctx, func(ctx context.Context, conn redis.Redis_6_0_Cmdable) error {
		data, err := conn.Get(ctx, key).Bytes()
		if errors.Is(err, goredis.Nil) {
			return ErrInvalidChallenge
		} else if err != nil {
			return err
		}

		err = json.Unmarshal(data, c)
		if err != nil {
			return err
		}

		_, err = conn.Del(ctx, key).Result()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &c.Purpose, nil
}

func challengeKey(appID config.AppID, token string) string {
	return fmt.Sprintf("app:%s:challenge:%s", appID, token)
}

package idpsession

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/session"
	"github.com/authgear/authgear-server/pkg/lib/session/access"

	"github.com/authgear/authgear-server/pkg/util/clock"
)

func TestProvider(t *testing.T) {
	Convey("Provider", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := NewMockStore(ctrl)
		accessEvents := NewMockAccessEventProvider(ctrl)

		clock := clock.NewMockClockAt("2020-01-01T00:00:00Z")
		initialTime := clock.Time

		accessEvent := access.Event{
			Timestamp: initialTime,
		}

		disabled := false
		provider := &Provider{
			Store:        store,
			AccessEvents: accessEvents,
			TrustProxy:   true,
			Config: &config.SessionConfig{
				IdleTimeoutEnabled: &disabled,
			},
			Clock:  clock,
			Random: rand.New(rand.NewSource(0)),
		}

		Convey("creating session", func() {
			Convey("should be successful", func() {
				store.EXPECT().Create(gomock.Any(), gomock.Any(), initialTime).Return(nil)
				accessEvents.EXPECT().InitStream(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				s, token := provider.MakeSession(&session.Attrs{
					UserID: "user-id",
				})
				ctx := context.Background()
				err := provider.Create(ctx, s)

				So(err, ShouldBeNil)
				So(token, ShouldNotBeEmpty)
				So(s, ShouldResemble, &IDPSession{
					ID: s.ID,
					Attrs: session.Attrs{
						UserID: "user-id",
					},
					AccessInfo: access.Info{
						InitialAccess: accessEvent,
						LastAccess:    accessEvent,
					},
					CreatedAt:                  initialTime,
					AuthenticatedAt:            initialTime,
					TokenHash:                  s.TokenHash,
					ExpireAtForResolvedSession: initialTime,
				})
			})
		})

		Convey("getting session", func() {
			fixtureSession := IDPSession{
				ID: "session-id",
				Attrs: session.Attrs{
					UserID: "user-id",
				},
				CreatedAt:       initialTime,
				AuthenticatedAt: initialTime,
				TokenHash:       "15be5b9c05673532b445d3295a86afd6b2615775e0233e9798cbe3c846a08d05",
			}
			store.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, id string) (*IDPSession, error) {
				if id == fixtureSession.ID {
					return &fixtureSession, nil
				}
				return nil, ErrSessionNotFound
			})

			Convey("should be successful using session token", func() {
				ctx := context.Background()
				session, err := provider.GetByToken(ctx, "session-id.token")
				So(err, ShouldBeNil)
				So(session, ShouldResemble, &fixtureSession)
			})

			Convey("should reject non-existent session", func() {
				ctx := context.Background()
				session, err := provider.GetByToken(ctx, "session-id-unknown.token")
				So(err, ShouldBeError, ErrSessionNotFound)
				So(session, ShouldBeNil)
			})

			Convey("should reject incorrect token", func() {
				ctx := context.Background()
				session, err := provider.GetByToken(ctx, "session-id.incorrect-token")
				So(err, ShouldBeError, ErrSessionNotFound)
				So(session, ShouldBeNil)

				session, err = provider.GetByToken(ctx, "invalid-token")
				So(err, ShouldBeError, ErrSessionNotFound)
				So(session, ShouldBeNil)
			})
			Convey("should reject if session is expired", func() {
				clock.AdvanceSeconds(1000000)
				ctx := context.Background()
				session, err := provider.GetByToken(ctx, "session-id.token")
				So(err, ShouldBeError, ErrSessionNotFound)
				So(session, ShouldBeNil)
			})
		})
	})

	enabled := true
	disabled := false

	Convey("checkSessionExpired", t, func() {

		session := &IDPSession{
			ID:        "session-id",
			CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			AccessInfo: access.Info{
				LastAccess: access.Event{
					Timestamp: time.Date(2020, 1, 1, 0, 0, 25, 0, time.UTC),
				},
			},
		}
		var cfg *config.SessionConfig
		check := func(cfg *config.SessionConfig, mins, secs int) bool {
			clock := clock.NewMockClockAtTime(time.Date(2020, 1, 1, 0, mins, secs, 0, time.UTC))
			provider := &Provider{
				Config: cfg,
				Clock:  clock,
			}
			expireAtForResolvedSession_before := session.ExpireAtForResolvedSession
			actual := !provider.CheckSessionExpired(session)
			expireAtForResolvedSession_after := session.ExpireAtForResolvedSession
			So(expireAtForResolvedSession_before.Equal(expireAtForResolvedSession_after), ShouldBeTrue)
			return actual
		}

		Convey("check session lifetime", func() {
			cfg = &config.SessionConfig{
				Lifetime:           120,
				IdleTimeoutEnabled: &disabled,
				IdleTimeout:        30,
			}

			So(check(cfg, 0, 0), ShouldBeTrue)
			So(check(cfg, 0, 56), ShouldBeTrue)
			So(check(cfg, 2, 0), ShouldBeTrue)
			So(check(cfg, 2, 1), ShouldBeFalse)
		})

		Convey("check idle timeout", func() {
			cfg = &config.SessionConfig{
				Lifetime:           120,
				IdleTimeoutEnabled: &enabled,
				IdleTimeout:        30,
			}

			So(check(cfg, 0, 0), ShouldBeTrue)
			So(check(cfg, 0, 55), ShouldBeTrue)
			So(check(cfg, 0, 56), ShouldBeFalse)
			So(check(cfg, 2, 1), ShouldBeFalse)
		})
	})
}

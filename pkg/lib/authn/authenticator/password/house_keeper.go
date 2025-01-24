// Copyright 2015-present Oursky Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package password

import (
	"context"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/util/log"
)

type HousekeeperLogger struct {
	*log.Logger
}

func NewHousekeeperLogger(lf *log.Factory) HousekeeperLogger {
	return HousekeeperLogger{lf.New("password-housekeeper")}
}

type Housekeeper struct {
	Store  *HistoryStore
	Logger HousekeeperLogger
	Config *config.AuthenticatorPasswordConfig
}

func (p *Housekeeper) Housekeep(ctx context.Context, authID string) (err error) {
	if !p.Config.Policy.IsEnabled() {
		return
	}

	p.Logger.Debug("remove password history")
	err = p.Store.RemovePasswordHistory(ctx, authID, p.Config.Policy.HistorySize, p.Config.Policy.HistoryDays)
	if err != nil {
		p.Logger.WithError(err).Error("unable to housekeep password history")
	}

	return
}

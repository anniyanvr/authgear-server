package web

import (
	"github.com/authgear/authgear-server/pkg/lib/config"
)

type UIImplementationService struct {
	UIConfig                       *config.UIConfig
	GlobalUIImplementation         config.GlobalUIImplementation
	GlobalUISettingsImplementation config.GlobalUISettingsImplementation
}

func (s *UIImplementationService) GetUIImplementation() config.UIImplementation {
	switch s.UIConfig.Implementation {
	case config.UIImplementationAuthflowV2:
		// authflowv2 is authflowv2
		return config.UIImplementationAuthflowV2
	case config.Deprecated_UIImplementationAuthflow:
		// Treat authflow as authflowv2
		return config.UIImplementationAuthflowV2
	case config.UIImplementationInteraction:
		// interaction is interaction
		// In case a project wants to use the legacy implementation.
		return config.UIImplementationInteraction
	default:
		// When it is unspecified in the config,
		// we use the env var to determine.
		switch s.GlobalUIImplementation {
		case config.GlobalUIImplementation(config.UIImplementationAuthflowV2):
			// authflowv2 is authflowv2
			return config.UIImplementationAuthflowV2
		case config.GlobalUIImplementation(config.Deprecated_UIImplementationAuthflow):
			// Treat authflow as authflowv2
			return config.UIImplementationAuthflowV2
		case config.GlobalUIImplementation(config.UIImplementationInteraction):
			// interaction is interaction
			// In case a project wants to use the legacy implementation.
			return config.UIImplementationInteraction
		default:
			// The ultimate default is now authflowv2
			// When you change this, you also need to change portal/src/system-config.ts
			return config.UIImplementationAuthflowV2
		}
	}
}

func (s *UIImplementationService) GetSettingsUIImplementation() config.SettingsUIImplementation {
	switch s.UIConfig.SettingsImplementation {
	case config.SettingsUIImplementationV1:
		return config.SettingsUIImplementationV1
	case config.SettingsUIImplementationV2:
		return config.SettingsUIImplementationV2
	default:
		// When it is unspecified in the config,
		// we use the env var to determine.
		switch s.GlobalUISettingsImplementation {
		case config.GlobalUISettingsImplementation(config.SettingsUIImplementationV1):
			return config.SettingsUIImplementationV1
		case config.GlobalUISettingsImplementation(config.SettingsUIImplementationV2):
			return config.SettingsUIImplementationV2
		default:
			// The ultimate default is now v2.
			// When you change this, you also need to change portal/src/system-config.ts
			return config.SettingsUIImplementationV2
		}
	}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package smtp

import (
	"github.com/authgear/authgear-server/pkg/lib/translation"
	"github.com/authgear/authgear-server/pkg/portal/model"
	"github.com/authgear/authgear-server/pkg/util/template"
)

// Injectors from wire.go:

func NewTranslationService(app *model.App) *translation.Service {
	manager := ProvideResourceManager(app)
	defaultLanguageTag := ProvideDefaultLanguageTag(app)
	supportedLanguageTags := ProvideSupportedLanguageTags(app)
	resolver := &template.Resolver{
		Resources:             manager,
		DefaultLanguageTag:    defaultLanguageTag,
		SupportedLanguageTags: supportedLanguageTags,
	}
	engine := &template.Engine{
		Resolver: resolver,
	}
	noopStaticAssetResolver := ProvideStaticAssetResolver()
	service := &translation.Service{
		TemplateEngine: engine,
		StaticAssets:   noopStaticAssetResolver,
	}
	return service
}

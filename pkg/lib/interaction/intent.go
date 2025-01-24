package interaction

import (
	"context"
	"reflect"
)

type Intent interface {
	InstantiateRootNode(goCtx context.Context, ctx *Context, graph *Graph) (Node, error)
	DeriveEdgesForNode(goCtx context.Context, graph *Graph, node Node) ([]Edge, error)
}

type IntentFactory func() Intent

var intentRegistry = map[string]IntentFactory{}

func RegisterIntent(intent Intent) {
	intentType := reflect.TypeOf(intent).Elem()

	intentKind := intentType.Name()
	factory := IntentFactory(func() Intent {
		return reflect.New(intentType).Interface().(Intent)
	})

	if _, hasKind := intentRegistry[intentKind]; hasKind {
		panic("interaction: duplicated intent kind: " + intentKind)
	}
	intentRegistry[intentKind] = factory
}

func IntentKind(intent Intent) string {
	intentType := reflect.TypeOf(intent).Elem()
	return intentType.Name()
}

func InstantiateIntent(kind string) Intent {
	factory, ok := intentRegistry[kind]
	if !ok {
		panic("interaction: unknown intent kind: " + kind)
	}
	return factory()
}

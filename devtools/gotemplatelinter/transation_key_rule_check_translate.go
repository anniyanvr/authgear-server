package main

import (
	"text/template/parse"
)

// validate `translate` command
//
// e.g. (translate "app.name" nil)
func CheckCommandTranslate(translateNode *parse.CommandNode) (err error) {
	// 2nd arg should be translation key
	for idx, arg := range translateNode.Args {
		if idx == 1 {
			err = CheckTranslationKeyNode(arg)
			if err != nil {
				return err
			}
		}

	}
	return
}

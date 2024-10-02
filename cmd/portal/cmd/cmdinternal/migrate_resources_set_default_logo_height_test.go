package cmdinternal

import (
	"regexp"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMigrateSetDefaultLogoHeight(t *testing.T) {

	Convey("getMatchingConfigSourcePaths", t, func() {
		test := func(regex *regexp.Regexp, configSourceData map[string]string, expectedMatched []string, expectedOK bool) {
			matched, ok := getMatchingConfigSourcePaths(regex, configSourceData)

			sortedMatched := sort.StringSlice(matched)
			sortedMatched.Sort()

			sortedExpectedMatched := sort.StringSlice(expectedMatched)
			sortedExpectedMatched.Sort()
			So(sortedMatched, ShouldResemble, sortedExpectedMatched)
			So(ok, ShouldEqual, expectedOK)
		}

		configSourceFixture := map[string]string{
			// light logos
			"static_2f_zh-HK_2f_app_5f_logo.png": "base64-encoded",
			"static_2f_en_2f_app_5f_logo.png":    "base64-encoded",
			"static_2f_es-ES_2f_app_5f_logo.png": "base64-encoded",

			// dark logos
			"static_2f_zh-HK_2f_app_5f_logo_5f_dark.png": "base64-encoded",
			"static_2f_en_2f_app_5f_logo_5f_dark.png":    "base64-encoded",
			"static_2f_es-ES_2f_app_5f_logo_5f_dark.png": "base64-encoded",

			// light theme css
			"static_2f_authgear-authflowv_32_-light-theme.css": "base64-encoded",

			// dark theme css
			"static_2f_authgear-authflowv_32_-dark-theme.css": "base64-encoded",

			// true-negatives
			"authgear.yaml":                          "base64-encoded",
			"authgear.secrets.yaml":                  "base64-encoded",
			"templates_2f_zh-HK_2f_translation.json": "base64-encoded",
		}

		Convey("nothing matched", func() {
			var expectedMatched []string
			expectedOK := false

			test(regexp.MustCompile(`^wrong regex$`), configSourceFixture, expectedMatched, expectedOK)
		})

		Convey("matched 3 light logos", func() {
			expectedMatched := []string{
				"static_2f_en_2f_app_5f_logo.png",
				"static_2f_es-ES_2f_app_5f_logo.png",
				"static_2f_zh-HK_2f_app_5f_logo.png",
			}
			expectedOK := true
			test(regexp.MustCompile(`^static/([a-zA-Z-]+)/app_logo\.(png|jpe|jpeg|jpg|gif)$`), configSourceFixture, expectedMatched, expectedOK)
		})

		Convey("matched 3 dark logos", func() {
			expectedMatched := []string{
				"static_2f_zh-HK_2f_app_5f_logo_5f_dark.png",
				"static_2f_en_2f_app_5f_logo_5f_dark.png",
				"static_2f_es-ES_2f_app_5f_logo_5f_dark.png",
			}
			expectedOK := true
			test(regexp.MustCompile(`^static/([a-zA-Z-]+)/app_logo_dark\.(png|jpe|jpeg|jpg|gif)$`), configSourceFixture, expectedMatched, expectedOK)
		})

		Convey("matched 1 light theme css", func() {
			expectedMatched := []string{
				"static_2f_authgear-authflowv_32_-light-theme.css",
			}
			expectedOK := true
			test(regexp.MustCompile(`^static/authgear-authflowv2-light-theme.css$`), configSourceFixture, expectedMatched, expectedOK)
		})

		Convey("matched 1 dark theme css", func() {
			expectedMatched := []string{
				"static_2f_authgear-authflowv_32_-dark-theme.css",
			}
			expectedOK := true
			test(regexp.MustCompile(`^static/authgear-authflowv2-dark-theme.css$`), configSourceFixture, expectedMatched, expectedOK)
		})

	})
}

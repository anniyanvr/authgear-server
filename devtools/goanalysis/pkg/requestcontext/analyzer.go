package requestcontext

import (
	"go/ast"
	"go/types"
	"slices"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var vettedPositions = []string{
	"/pkg/util/httproute/httproute.go:109:38",
	"/pkg/util/httputil/csp.go:158:42",
	"/pkg/util/httputil/csp.go:164:42",
	"/pkg/util/httputil/file_server.go:147:11",
	"/pkg/util/graphqlutil/graphiql.go:84:32",
	"/pkg/util/template/engine.go:307:57",
	"/pkg/lib/dpop/middleware.go:45:35",
	"/pkg/lib/session/middleware.go:48:23",
	"/pkg/lib/session/middleware.go:57:41",
	"/pkg/lib/session/middleware.go:61:34",
	"/pkg/lib/oauth/response_mode.go:105:40",
	"/pkg/lib/oauth/response_mode.go:117:40",
	"/pkg/lib/oauth/scope.go:44:34",
	"/pkg/admin/transport/handler_graphql.go:47:9",
	"/pkg/admin/transport/handler_user_export_create.go:55:9",
	"/pkg/admin/transport/handler_user_export_get.go:37:9",
	"/pkg/admin/transport/handler_user_import_create.go:28:9",
	"/pkg/admin/transport/handler_user_import_get.go:29:9",
	"/pkg/admin/transport/ui_param.go:35:10",
	"/pkg/lib/admin/authz/middleware.go:85:18",
	"/pkg/auth/handler/saml/login.go:255:5",
	"/pkg/auth/handler/saml/login.go:295:3",
	"/auth/handler/saml/login_finish.go:41:9",
	"/pkg/auth/handler/saml/logout.go:86:9",
	"/pkg/util/pubsub/http_handler.go:51:40",
	"/pkg/lib/authenticationflow/intl_middleware.go:16:41",
	"/pkg/lib/authenticationflow/rate_limit_middleware.go:31:49",
	"/pkg/auth/webapp/context_holder_middleware.go:26:32",
	"/pkg/auth/webapp/intl_middleware.go:12:41",
	"/pkg/auth/webapp/require_authenticated_middleware.go:14:10",
	"/pkg/auth/webapp/session_middleware.go:59:55",
	"/pkg/auth/webapp/session_middleware.go:68:34",
	"/pkg/auth/webapp/session_middleware.go:70:54",
	"/pkg/auth/webapp/session_middleware.go:79:34",
	"/pkg/auth/webapp/session_middleware.go:82:34",
	"/pkg/auth/webapp/session_middleware.go:95:35",
	"/pkg/auth/webapp/settings_sub_routes_middleware.go:26:10",
	"/pkg/auth/webapp/success_page_middleware.go:41:10",
	"/pkg/auth/webapp/ui_param.go:55:28",
	"/pkg/auth/webapp/ui_param.go:58:43",
	"/pkg/auth/webapp/ui_param.go:65:42",
	"/pkg/auth/webapp/ui_param.go:86:10",
	"/pkg/auth/webapp/visitor_id_middleware.go:26:24",
	"/pkg/auth/webapp/wechat_redirect_uri_middleware.go:59:31",
	"/pkg/auth/webapp/wechat_redirect_uri_middleware.go:115:39",
	"/pkg/auth/webapp/x_color_scheme.go:74:27",
	"/pkg/auth/handler/webapp/viewmodels/base.go:213:32",
	"/pkg/auth/handler/webapp/viewmodels/base.go:228:35",
	"/pkg/auth/handler/webapp/viewmodels/base.go:230:57",
	"/pkg/auth/handler/webapp/viewmodels/base.go:282:38",
	"/pkg/auth/handler/webapp/viewmodels/base.go:288:24",
	"/pkg/auth/handler/webapp/viewmodels/base.go:300:43",
	"/pkg/auth/handler/webapp/viewmodels/base.go:376:47",
	"/pkg/auth/handler/webapp/viewmodels/base.go:381:28",
	"/pkg/auth/handler/webapp/viewmodels/base.go:394:33",
	"/pkg/lib/accountmanagement/rate_limit_middleware.go:41:10",
	"/pkg/lib/workflow/intl_middleware.go:16:41",
	"/pkg/auth/handler/webapp/auth_entry_point_middleware.go:31:31",
	"/pkg/auth/handler/webapp/auth_entry_point_middleware.go:32:35",
	"/pkg/auth/handler/webapp/authflow_change_password.go:96:26",
	"/pkg/auth/handler/webapp/authflow_controller.go:978:30",
	"/pkg/auth/handler/webapp/authflow_controller.go:983:24",
	"/pkg/auth/handler/webapp/authflow_controller.go:991:19",
	"/pkg/auth/handler/webapp/authflow_controller.go:1000:19",
	"/pkg/auth/handler/webapp/authflow_create_password.go:132:26",
	"/pkg/auth/handler/webapp/authflow_enter_oob_otp.go:156:26",
	"/pkg/auth/handler/webapp/authflow_enter_password.go:138:26",
	"/pkg/auth/handler/webapp/authflow_enter_recovery_code.go:90:26",
	"/pkg/auth/handler/webapp/authflow_enter_totp.go:88:26",
	"/pkg/auth/handler/webapp/authflow_finish_flow.go:35:33",
	"/pkg/auth/handler/webapp/authflow_forgot_password.go:227:33",
	"/pkg/auth/handler/webapp/authflow_forgot_password_otp.go:213:26",
	"/pkg/auth/handler/webapp/authflow_forgot_password_success.go:52:26",
	"/pkg/auth/handler/webapp/authflow_login.go:197:33",
	"/pkg/auth/handler/webapp/authflow_oob_otp_link.go:136:26",
	"/pkg/auth/handler/webapp/authflow_promote.go:128:33",
	"/pkg/auth/handler/webapp/authflow_prompt_create_passkey.go:102:26",
	"/pkg/auth/handler/webapp/authflow_reauth.go:41:33",
	"/pkg/auth/handler/webapp/authflow_reset_password.go:128:35",
	"/pkg/auth/handler/webapp/authflow_reset_password.go:132:27",
	"/pkg/auth/handler/webapp/authflow_reset_password_success.go:59:33",
	"/pkg/auth/handler/webapp/authflow_setup_oob_otp.go:131:26",
	"/pkg/auth/handler/webapp/authflow_setup_totp.go:111:26",
	"/pkg/auth/handler/webapp/authflow_signup.go:142:33",
	"/pkg/auth/handler/webapp/authflow_terminate_other_sessions.go:73:26",
	"/pkg/auth/handler/webapp/authflow_use_passkey.go:108:26",
	"/pkg/auth/handler/webapp/authflow_view_recovery_code.go:85:26",
	"/pkg/auth/handler/webapp/authflow_wechat.go:168:26",
	"/pkg/auth/handler/webapp/authflow_whatsapp_otp.go:147:26",
	"/pkg/auth/handler/webapp/confirm_web3_account.go:104:28",
	"/pkg/auth/handler/webapp/controller.go:98:30",
	"/pkg/auth/handler/webapp/controller.go:130:27",
	"/pkg/auth/handler/webapp/controller.go:136:30",
	"/pkg/auth/handler/webapp/controller.go:145:28",
	"/pkg/auth/handler/webapp/controller.go:154:30",
	"/pkg/auth/handler/webapp/controller.go:209:27",
	"/pkg/auth/handler/webapp/controller.go:211:20",
	"/pkg/auth/handler/webapp/controller.go:216:18",
	"/pkg/auth/handler/webapp/controller.go:228:25",
	"/pkg/auth/handler/webapp/controller.go:247:30",
	"/pkg/auth/handler/webapp/controller.go:255:31",
	"/pkg/auth/handler/webapp/controller.go:287:25",
	"/pkg/auth/handler/webapp/controller.go:304:20",
	"/pkg/auth/handler/webapp/controller.go:312:20",
	"/pkg/auth/handler/webapp/controller.go:325:30",
	"/pkg/auth/handler/webapp/controller.go:336:30",
	"/pkg/auth/handler/webapp/controller.go:345:30",
	"/pkg/auth/handler/webapp/csrf_error_instruction.go:49:33",
	"/pkg/auth/handler/webapp/implementation_switcher.go:50:42",
	"/pkg/auth/handler/webapp/implementation_switcher.go:62:30",
	"/pkg/auth/handler/webapp/login.go:115:28",
	"/pkg/auth/handler/webapp/panic_middleware.go:68:61",
	"/pkg/auth/handler/webapp/passkey_creation_options.go:43:28",
	"/pkg/auth/handler/webapp/passkey_request_options.go:52:28",
	"/pkg/auth/handler/webapp/promote.go:110:28",
	"/pkg/auth/handler/webapp/reauth.go:31:34",
	"/pkg/auth/handler/webapp/select_account.go:110:32",
	"/pkg/auth/handler/webapp/select_account.go:111:34",
	"/pkg/auth/handler/webapp/setting_implementation_switcher.go:50:50",
	"/pkg/auth/handler/webapp/setting_implementation_switcher.go:61:38",
	"/pkg/auth/handler/webapp/settings_delete_account.go:94:39",
	"/pkg/auth/handler/webapp/settings_delete_account.go:96:34",
	"/pkg/auth/handler/webapp/settings_delete_account_success.go:69:34",
	"/pkg/auth/handler/webapp/settings_profile.go:40:31",
	"/pkg/auth/handler/webapp/settings_sessions.go:113:39",
	"/pkg/auth/handler/webapp/signup.go:101:28",
	"/pkg/auth/handler/webapp/sso_callback.go:42:53",
	"/pkg/auth/handler/webapp/sso_callback.go:48:26",
	"/pkg/auth/handler/webapp/sso_callback.go:56:58",
	"/pkg/auth/handler/webapp/sso_callback.go:61:44",
	"/pkg/auth/handler/webapp/sso_callback.go:74:44",
	"/pkg/auth/handler/webapp/websocket.go:45:25",
	"/pkg/auth/handler/webapp/websocket.go:56:25",
	"/pkg/auth/handler/webapp/websocket.go:80:30",
	"/pkg/auth/handler/webapp/authflowv2/account_linking.go:151:26",
	"/pkg/auth/handler/webapp/authflowv2/change_password.go:115:26",
	"/pkg/auth/handler/webapp/authflowv2/change_password_success.go:65:33",
	"/pkg/auth/handler/webapp/authflowv2/create_password.go:179:26",
	"/pkg/auth/handler/webapp/authflowv2/enter_oob_otp.go:222:26",
	"/pkg/auth/handler/webapp/authflowv2/enter_password.go:183:26",
	"/pkg/auth/handler/webapp/authflowv2/enter_recovery_code.go:113:26",
	"/pkg/auth/handler/webapp/authflowv2/enter_totp.go:132:26",
	"/pkg/auth/handler/webapp/authflowv2/error.go:65:33",
	"/pkg/auth/handler/webapp/authflowv2/finish_flow.go:36:33",
	"/pkg/auth/handler/webapp/authflowv2/forgot_password.go:286:33",
	"/pkg/auth/handler/webapp/authflowv2/forgot_password_link_sent.go:86:26",
	"/pkg/auth/handler/webapp/authflowv2/forgot_password_otp.go:204:26",
	"/pkg/auth/handler/webapp/authflowv2/internal_signup_login.go:234:33",
	"/pkg/auth/handler/webapp/authflowv2/ldap_login.go:109:26",
	"/pkg/auth/handler/webapp/authflowv2/login.go:247:33",
	"/pkg/auth/handler/webapp/authflowv2/oob_otp_link.go:170:26",
	"/pkg/auth/handler/webapp/authflowv2/promote.go:139:33",
	"/pkg/auth/handler/webapp/authflowv2/prompt_create_passkey.go:126:26",
	"/pkg/auth/handler/webapp/authflowv2/reauth.go:42:33",
	"/pkg/auth/handler/webapp/authflowv2/reset_password.go:227:35",
	"/pkg/auth/handler/webapp/authflowv2/reset_password.go:231:27",
	"/pkg/auth/handler/webapp/authflowv2/reset_password_success.go:60:33",
	"/pkg/auth/handler/webapp/authflowv2/select_account.go:111:32",
	"/pkg/auth/handler/webapp/authflowv2/settings_delete_account.go:80:39",
	"/pkg/auth/handler/webapp/authflowv2/settings_delete_account.go:82:34",
	"/pkg/auth/handler/webapp/authflowv2/settings_delete_account_success.go:56:34",
	"/pkg/auth/handler/webapp/authflowv2/settings_sessions.go:122:39",
	"/pkg/auth/handler/webapp/authflowv2/setup_oob_otp.go:163:26",
	"/pkg/auth/handler/webapp/authflowv2/setup_totp.go:122:26",
	"/pkg/auth/handler/webapp/authflowv2/terminate_other_sessions.go:74:26",
	"/pkg/auth/handler/webapp/authflowv2/use_passkey.go:150:26",
	"/pkg/auth/handler/webapp/authflowv2/verify_bot_protection.go:85:26",
	"/pkg/auth/handler/webapp/authflowv2/view_recovery_code.go:86:26",
	"/pkg/auth/handler/webapp/authflowv2/wechat.go:169:26",
	"/pkg/lib/healthz/healthz.go:27:23",
	"/pkg/util/sentry/request.go:25:9",
	"/pkg/lib/deps/context.go:23:10",
	"/pkg/lib/deps/middleware_request.go:38:48",
	"/pkg/lib/deps/middleware_request.go:54:39",
	"/pkg/lib/deps/middleware_request.go:55:34",
	"/pkg/lib/deps/providers.go:166:25",
	"/pkg/lib/infra/middleware/sentry.go:21:47",
	"/pkg/auth/api/require_authenticated_middleware.go:21:31",
	"/pkg/auth/handler/api/accountmanagement_v1_identification.go:69:9",
	"/pkg/auth/handler/api/accountmanagement_v1_identification_oauth.go:60:9",
	"/pkg/auth/handler/api/anonymous_user_promotion_code.go:91:9",
	"/pkg/auth/handler/api/anonymous_user_signup.go:101:9",
	"/pkg/auth/handler/api/authenticationflow_v1_create.go:90:9",
	"/pkg/auth/handler/api/authenticationflow_v1_get.go:49:9",
	"/pkg/auth/handler/api/authenticationflow_v1_input.go:77:9",
	"/pkg/auth/handler/api/presign_images_upload.go:59:9",
	"/pkg/auth/handler/api/workflow_get.go:35:9",
	"/pkg/auth/handler/api/workflow_input.go:60:9",
	"/pkg/auth/handler/api/workflow_new.go:92:9",
	"/pkg/auth/handler/api/workflow_v2.go:212:9",
	"/pkg/auth/handler/api/workflow_websocket.go:55:9",
	"/pkg/auth/handler/oauth/app_session_token.go:69:27",
	"/pkg/auth/handler/oauth/authorize.go:52:26",
	"/pkg/auth/handler/oauth/challenge.go:91:27",
	"/pkg/auth/handler/oauth/consent.go:85:27",
	"/pkg/auth/handler/oauth/consent.go:105:28",
	"/pkg/auth/handler/oauth/consent.go:114:28",
	"/pkg/auth/handler/oauth/end_session.go:48:26",
	"/pkg/auth/handler/oauth/revoke.go:47:26",
	"/pkg/auth/handler/oauth/token.go:50:26",
	"/pkg/auth/handler/oauth/userinfo.go:45:26",
	"/pkg/auth/handler/oauth/userinfo.go:48:27",
	"/pkg/auth/handler/siwe/nonce.go:46:38",
	"/pkg/images/deps/context.go:23:10",
	"/pkg/images/deps/middleware_request.go:26:48",
	"/pkg/images/deps/middleware_request.go:37:39",
	"/pkg/images/deps/middleware_request.go:38:34",
	"/pkg/images/handler/get.go:95:43",
	"/pkg/images/handler/post.go:162:83",
	"/pkg/images/handler/post.go:193:10",
	"/pkg/lib/session/test/context.go:85:35",
	"/pkg/portal/session/middleware_session_info.go:19:37",
	"/pkg/portal/session/middleware_session_required.go:12:38",
	"/pkg/portal/transport/admin_api_handler.go:54:45",
	"/pkg/portal/transport/admin_api_handler.go:61:44",
	"/pkg/portal/transport/admin_api_handler.go:81:39",
	"/pkg/portal/transport/graphql_handler.go:44:29",
	"/pkg/portal/transport/stripe_webhook_handler.go:82:47",
	"/pkg/portal/transport/stripe_webhook_handler.go:85:4",
	"/pkg/portal/transport/stripe_webhook_handler.go:91:4",
	"/pkg/portal/transport/stripe_webhook_handler.go:96:4",
	"/pkg/portal/csp_middleware.go:16:39",
	"/pkg/resolver/handler/resolve.go:61:9",
}

var Analyzer = &analysis.Analyzer{
	Name: "requestcontext",
	Doc:  "requestcontext forbids (*net/http.Request).Context, except those locations explicitly hard-coded in this analyzer.",
	Run:  run,
	// See https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/inspect
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func isATestFile(pass *analysis.Pass, n ast.Node) bool {
	position := pass.Fset.Position(n.Pos())
	// position is valid if Line > 0
	// See https://pkg.go.dev/go/token#Position
	if position.Line > 0 {
		if strings.HasSuffix(position.Filename, "_test.go") {
			return true
		}
	}
	return false
}

func IsVettedPos(pass *analysis.Pass, n ast.Node) bool {
	position := pass.Fset.Position(n.Pos())

	// position is valid if Line > 0
	// See https://pkg.go.dev/go/token#Position
	if position.Line > 0 {
		b := slices.ContainsFunc(vettedPositions, func(s string) bool {
			return strings.HasSuffix(position.String(), s)
		})
		if b {
			return true
		}
	}

	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	traverse := func(n ast.Node) {
		isTestFile := isATestFile(pass, n)

		if !isTestFile {
			if n, ok := n.(*ast.SelectorExpr); ok {
				if selObj := pass.TypesInfo.ObjectOf(n.Sel); selObj != nil {
					if f, ok := selObj.(*types.Func); ok {
						fullName := f.FullName()
						switch fullName {
						case "(*net/http.Request).Context":
							isVetted := IsVettedPos(pass, n)
							if !isVetted {
								pass.Reportf(n.Pos(), "Unvetted usage of request.Context is forbidden.")
							}
						default:
							break
						}
					}
				}
			}
		}
	}

	inspect.Preorder(nil, traverse)
	return nil, nil
}

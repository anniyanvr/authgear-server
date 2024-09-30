package authflowv2

import (
	"github.com/google/wire"
)

var DependencySet = wire.NewSet(
	wire.Struct(new(AuthflowV2Navigator), "*"),
	wire.Struct(new(InternalAuthflowV2SignupLoginHandler), "*"),
	wire.Struct(new(AuthflowV2LoginHandler), "*"),
	wire.Struct(new(AuthflowV2SignupHandler), "*"),
	wire.Struct(new(AuthflowV2ReauthHandler), "*"),
	wire.Struct(new(AuthflowV2EnterPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2EnterOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SetupOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowV2ViewRecoveryCodeHandler), "*"),
	wire.Struct(new(AuthflowV2ErrorHandler), "*"),
	wire.Struct(new(AuthflowV2NoAuthenticatorHandler), "*"),
	wire.Struct(new(AuthflowV2CreatePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2AccountStatusHandler), "*"),
	wire.Struct(new(AuthflowV2NotFoundHandler), "*"),
	wire.Struct(new(AuthflowV2SelectAccountHandler), "*"),
	wire.Struct(new(AuthflowV2VerifyBotProtectionHandler), "*"),
	wire.Struct(new(AuthflowV2EnterRecoveryCodeHandler), "*"),
	wire.Struct(new(AuthflowV2ChangePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ChangePasswordSuccessHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordLinkSentHandler), "*"),
	wire.Struct(new(AuthflowV2ForgotPasswordOTPHandler), "*"),
	wire.Struct(new(AuthflowV2ResetPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2ResetPasswordSuccessHandler), "*"),
	wire.Struct(new(AuthflowV2SetupTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2EnterTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2OOBOTPLinkHandler), "*"),
	wire.Struct(new(AuthflowV2VerifyLoginLinkOTPHandler), "*"),
	wire.Struct(new(AuthflowV2PromptCreatePasskeyHandler), "*"),
	wire.Struct(new(AuthflowV2UsePasskeyHandler), "*"),
	wire.Struct(new(AuthflowV2TerminateOtherSessionsHandler), "*"),
	wire.Struct(new(AuthflowV2PromoteHandler), "*"),
	wire.Struct(new(AuthflowV2FinishFlowHandler), "*"),
	wire.Struct(new(AuthflowV2WechatHandler), "*"),
	wire.Struct(new(AuthflowV2AccountLinkingHandler), "*"),
	wire.Struct(new(AuthflowV2LDAPLoginHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsProfileHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsProfileEditHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsChangePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsChangePasskeyHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsBiometricHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsSessionsHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsMFAHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsAdvancedSettingsHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsDeleteAccountHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsDeleteAccountSuccessHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsMFACreatePasswordHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsMFAPasswordHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsMFACreateTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsMFAEnterTOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsOOBOTPHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityAddEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityEditEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityVerifyEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityListEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityViewEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityChangePrimaryEmailHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityListPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityViewPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityAddPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityEditPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityChangePrimaryPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityVerifyPhoneHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityListUsernameHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityNewUsernameHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityViewUsernameHandler), "*"),
	wire.Struct(new(AuthflowV2SettingsIdentityEditUsernameHandler), "*"),
)

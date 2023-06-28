package otp

// Otp Type
var (
	SignIn           = `sign_in`
	SignUp           = `sign_up`
	SignInView       = `Sign In`
	SignUpView       = `Sign Up`
	OTP_MESSAGE_BASE = "%s this is your OTP code for (%s) to Kurirmoo. Please do not reveal your OTP code to other."
	FormatDatetime   = "2006-01-02T15:04:05Z"

	OtpTypeAllowed = map[string]bool{
		SignIn: true,
		SignUp: true,
	}

	OtpTypes = map[string]*string{
		SignIn: &SignIn,
		SignUp: &SignUp,
	}

	OtpTypesView = map[string]*string{
		SignIn: &SignInView,
		SignUp: &SignUpView,
	}
)

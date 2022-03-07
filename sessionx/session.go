package sessionx

import "github.com/aws/aws-sdk-go/aws/session"

// New returns a session based on sso credentials from ~/.aws/config
func New(profile string) (*session.Session, error) {
	return session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           profile,
	})
}

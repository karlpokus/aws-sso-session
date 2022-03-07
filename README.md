# aws-sso-session
Put your local aws sso credentials into a go aws sdk session. Useful for local test env.

# usage
````go
import "github.com/karlpokus/aws-sso-session/sessionx"

func main() {
  profile := "my_profile" // defined in ~/.aws/config
  sess, err := sessionx.New(profile)
}
````

# license
MIT

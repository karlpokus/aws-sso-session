# aws-sso-session
Use your local aws sso credentials with aws-go-sdk. Useful for local test env.

Since the aws-go-sdk has moved on to version 2, I'll maintain multiple versions of this lib, so make sure you import by version.

# usage

````go
// v1
import "github.com/karlpokus/aws-sso-session/sessionx"

func main() {
  profile := "my_profile" // defined in ~/.aws/config
  sess, err := sessionx.New(profile)
}
````

````go
// v2
import "github.com/karlpokus/aws-sso-session/sessionx/v2"

func main() {
  profile := "my_profile" // defined in ~/.aws/config
  region := "eu-north-1"
  cfg, err := sessionx.New(context.Background(), profile, region)
}
````

# license
MIT

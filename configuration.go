package goshikimori

type Configuration struct {
  Application, AccessToken string
}

type FastId struct {
  Id   int
  Conf Configuration
  Err  error
}

// You need to enter the application and the private key.
//
// To register the application, follow the link from [OAuth].
//
// [OAuth]: https://github.com/heycatch/goshikimori#shikimori-documentation
func Add(appname, token string) *Configuration {
  return &Configuration{Application: appname, AccessToken: token}
}

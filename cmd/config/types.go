package config

type Config struct {
	General              general
	Client               client
	Instance             instance
	SSH                  ssh
	IMAP                 imap
	InviteCode           string
	CurrentRemoteAddress string
}

type Env struct {
	Config Config
	Dist   DistData
}

var config Config

type general struct {
	Protocol string
	Domain   string
}

type client struct {
	IdleConnTimeoutMs int
	TimeoutMs         int
}

type instance struct {
	InitializationOnStartupEnabled bool
	UnitSystemD                    string
}

type ssh struct {
	Port     int
	UserName string
}

type imap struct {
	Host  string
	Port  int
	Users struct {
		Username []string
		Password []string
	}
}

type DistData struct {
	Version   string
	BuildTime string
	Commit    string
	Compiler  string
}

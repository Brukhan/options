package flags

import (
	"github.com/Brukhan/options/option"
)

type Options struct {
	CertifPath string
	KeyPath string
	RedisServer string
	RedisQueue string
}

func Register(place *Options) {
	set := option.NewOptionSet("ApnsOps")
	set.String(&place.CertifPath, "certif-path", "CertifPath", "production-push.crt", "Path to certificate pem file")
	set.String(&place.Key, "key-path", "KeyPath", "production-push.key", "Path to key pem file")
	set.String(&place.RedisServer, "redis-server", "RedisServer", "production-push.key", "IP or URL to redis server")
	set.String(&place.RedisQueue, "redis-queue", "RedisQueue", "production-push.key", "name of the list that it should pop from (using BRPOP command)")

}

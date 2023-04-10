package configs

// Redis ..
type Redis struct {
	IsCluster  bool
	Addresses  []string
	Prefix     string
	Expiration int64
	PoolSize   int
	// 以下两个在集群下无效
	Db       int
	Password string
}

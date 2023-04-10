package configs

// Prometheus ..
type Prometheus struct {
	IsEnable         bool
	RoutePath        string   // 注册路由，供prometheus拉取数据
	RequestsTotal    string   // 默认requests_total
	RequestsCostTime string   // 默认requests_cost_time
	Tags             []string // 默认prometheus
}

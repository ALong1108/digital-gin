package configs

// GrpcClient ..
type GrpcClient struct {
	//必填，必须保证唯一，且符合证书域名规则(如果使用证书)
	//如果采用服务发现，则用于服务名
	ServerName string

	//指定tag
	Tag string

	//服务发现类型，目前可选static、consul，默认是static
	ResolverType string
	//默认ResolverType+ServerName，必须保证在同个项目里所有外部服务都是唯一
	ResolverScheme string
	//服务发现的地址，如consul、etcd地址
	ResolverAddresses []string
	//负载均衡策略名称，支持round_robin、pick_first、p2c，默认是p2c
	BalancerName string

	//服务地址，如果ResolverType是static，必填
	Addresses []string

	//调用具有证书的grpc服务，必须要指定客户端证书
	CertFile string

	//是否需要Auth验证
	IsAuth bool
}

// GrpcClients grpc 客户端连接维护
type GrpcClients struct {
	Proxy GrpcClient

	// 算法
	IndustryNorm           GrpcClient
	SkillExtract           GrpcClient
	EduNorm                GrpcClient
	PositionRankNorm       GrpcClient
	DepartmentNorm         GrpcClient
	CompanyNorm            GrpcClient
	FunctionNorm           GrpcClient
	SchoolNorm             GrpcClient
	MajorNorm              GrpcClient
	FunctionPred           GrpcClient
	ManagementPre          GrpcClient
	MajorNerNorm           GrpcClient
	JdManagementPre        GrpcClient
	IndustryExtract        GrpcClient
	JdSimilar              GrpcClient
	SentimentAnalysis      GrpcClient
	LimitsExtract          GrpcClient
	PrerequiresExtract     GrpcClient
	CompanySeg             GrpcClient
	CompanyPref            GrpcClient
	ExperienceExtract      GrpcClient
	CertificateNorm        GrpcClient // 证书归一化服务
	JdCvTextSimilar        GrpcClient // JdCv相关性服务
	CompanyTypeNorm        GrpcClient // 公司类型归一化
	JieBaCutFull           GrpcClient // 结巴分词服务
	EsEtl                  GrpcClient // EsEtl 分词,向量服务
	EsRecallWords          GrpcClient // 分词召回服务
	JdSimilarRecall        GrpcClient // 职位相似召回服务
	EduLanguageExtract     GrpcClient
	FunctionExtract        GrpcClient
	IndustryProjectExtract GrpcClient
	PostRelationExtraction GrpcClient
	PostTokenWeight        GrpcClient
	CompanyDomain          GrpcClient

	// Rank
	RecommendRank                GrpcClient // 推荐服务
	RecommendRankPreTest         GrpcClient // 推荐服务旧版
	JobStability                 GrpcClient // 离职意愿度服务
	MatchScoreCalculation        GrpcClient // 职位简历匹配分计算服务
	MatchScoreCalculationPreTest GrpcClient // 职位简历匹配分计算服务
	RecommendMultipleRank        GrpcClient // 推荐rank多路召回结果服务
	RecommendCompetitionPower    GrpcClient // 推荐多路召回竞争力服务
	CvSimilarRecall              GrpcClient // 简历相似召回服务

	// 搜索推荐
	SearchRecommendation         GrpcClient // 搜索推荐服务
	SearchRecommendationCompared GrpcClient // 搜索推荐对比服务
	MatchScore                   GrpcClient // 搜索推荐对比服务
	BdDownload                   GrpcClient // bd公司成功率

	// 人才地图简历查重服务
	TalentmapDuplicate GrpcClient // 人才地图简历查重服务

	// 爬虫
	ProxyManager GrpcClient

	// 大数据
	BigdataWeb GrpcClient

	// 解析服务
	Etl GrpcClient

	// 搜索服务
	Search   GrpcClient
	SoApollo GrpcClient

	// 数据流
	Artery GrpcClient
}

// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConsumerGroup struct {
	// 消费者名称。
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否有序消费
	Order *bool `json:"order,omitempty" xml:"order,omitempty"`
	// 消费超时时长，单位为妙
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s ConsumerGroup) String() string {
	return tea.Prettify(s)
}

func (s ConsumerGroup) GoString() string {
	return s.String()
}

func (s *ConsumerGroup) SetName(v string) *ConsumerGroup {
	s.Name = &v
	return s
}

func (s *ConsumerGroup) SetOrder(v bool) *ConsumerGroup {
	s.Order = &v
	return s
}

func (s *ConsumerGroup) SetTimeout(v int32) *ConsumerGroup {
	s.Timeout = &v
	return s
}

type EncryptConf struct {
	// enable
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 加密算法，只支持default和m4。当 enable 为 true 时，此项必选。
	EncryptType *string             `json:"encrypt_type,omitempty" xml:"encrypt_type,omitempty"`
	UserCmkInfo *EncryptUserCmkConf `json:"user_cmk_info,omitempty" xml:"user_cmk_info,omitempty"`
}

func (s EncryptConf) String() string {
	return tea.Prettify(s)
}

func (s EncryptConf) GoString() string {
	return s.String()
}

func (s *EncryptConf) SetEnable(v bool) *EncryptConf {
	s.Enable = &v
	return s
}

func (s *EncryptConf) SetEncryptType(v string) *EncryptConf {
	s.EncryptType = &v
	return s
}

func (s *EncryptConf) SetUserCmkInfo(v *EncryptUserCmkConf) *EncryptConf {
	s.UserCmkInfo = v
	return s
}

type EncryptUserCmkConf struct {
	// arn
	Arn *string `json:"arn,omitempty" xml:"arn,omitempty"`
	// cmk_key_id
	CmkKeyId *string `json:"cmk_key_id,omitempty" xml:"cmk_key_id,omitempty"`
	// region_id
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
}

func (s EncryptUserCmkConf) String() string {
	return tea.Prettify(s)
}

func (s EncryptUserCmkConf) GoString() string {
	return s.String()
}

func (s *EncryptUserCmkConf) SetArn(v string) *EncryptUserCmkConf {
	s.Arn = &v
	return s
}

func (s *EncryptUserCmkConf) SetCmkKeyId(v string) *EncryptUserCmkConf {
	s.CmkKeyId = &v
	return s
}

func (s *EncryptUserCmkConf) SetRegionId(v string) *EncryptUserCmkConf {
	s.RegionId = &v
	return s
}

type LogtailConfig struct {
	// logtail 配置的名称。
	ConfigName *string `json:"configName,omitempty" xml:"configName,omitempty"`
	// 创建时间，unix 时间戳。
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// logtail 输入的详细配置。
	InputDetail map[string]interface{} `json:"inputDetail,omitempty" xml:"inputDetail,omitempty"`
	// logtail 读取日志的输入类型。
	InputType *string `json:"inputType,omitempty" xml:"inputType,omitempty"`
	// 最后一次修改时间，unix 时间戳。
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// 日志样例，可以用于自动生成正则捕获字段。
	LogSample *string `json:"logSample,omitempty" xml:"logSample,omitempty"`
	// logtail 输出的详细配置。
	OutputDetail *LogtailConfigOutputDetail `json:"outputDetail,omitempty" xml:"outputDetail,omitempty" type:"Struct"`
	// logtail 输出的目标类型。这里固定选择 LogService。
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
}

func (s LogtailConfig) String() string {
	return tea.Prettify(s)
}

func (s LogtailConfig) GoString() string {
	return s.String()
}

func (s *LogtailConfig) SetConfigName(v string) *LogtailConfig {
	s.ConfigName = &v
	return s
}

func (s *LogtailConfig) SetCreateTime(v int64) *LogtailConfig {
	s.CreateTime = &v
	return s
}

func (s *LogtailConfig) SetInputDetail(v map[string]interface{}) *LogtailConfig {
	s.InputDetail = v
	return s
}

func (s *LogtailConfig) SetInputType(v string) *LogtailConfig {
	s.InputType = &v
	return s
}

func (s *LogtailConfig) SetLastModifyTime(v int64) *LogtailConfig {
	s.LastModifyTime = &v
	return s
}

func (s *LogtailConfig) SetLogSample(v string) *LogtailConfig {
	s.LogSample = &v
	return s
}

func (s *LogtailConfig) SetOutputDetail(v *LogtailConfigOutputDetail) *LogtailConfig {
	s.OutputDetail = v
	return s
}

func (s *LogtailConfig) SetOutputType(v string) *LogtailConfig {
	s.OutputType = &v
	return s
}

type LogtailConfigOutputDetail struct {
	// 日志项目的 endpoint。
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// 输出的目标 logstore 名称。
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// 地域。
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s LogtailConfigOutputDetail) String() string {
	return tea.Prettify(s)
}

func (s LogtailConfigOutputDetail) GoString() string {
	return s.String()
}

func (s *LogtailConfigOutputDetail) SetEndpoint(v string) *LogtailConfigOutputDetail {
	s.Endpoint = &v
	return s
}

func (s *LogtailConfigOutputDetail) SetLogstoreName(v string) *LogtailConfigOutputDetail {
	s.LogstoreName = &v
	return s
}

func (s *LogtailConfigOutputDetail) SetRegion(v string) *LogtailConfigOutputDetail {
	s.Region = &v
	return s
}

type SavedSearch struct {
	// displayName
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// logstore
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// savedsearchName
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	// searchQuery
	SearchQuery *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s SavedSearch) String() string {
	return tea.Prettify(s)
}

func (s SavedSearch) GoString() string {
	return s.String()
}

func (s *SavedSearch) SetDisplayName(v string) *SavedSearch {
	s.DisplayName = &v
	return s
}

func (s *SavedSearch) SetLogstore(v string) *SavedSearch {
	s.Logstore = &v
	return s
}

func (s *SavedSearch) SetSavedsearchName(v string) *SavedSearch {
	s.SavedsearchName = &v
	return s
}

func (s *SavedSearch) SetSearchQuery(v string) *SavedSearch {
	s.SearchQuery = &v
	return s
}

func (s *SavedSearch) SetTopic(v string) *SavedSearch {
	s.Topic = &v
	return s
}

type Chart struct {
	// action
	Action map[string]interface{} `json:"action,omitempty" xml:"action,omitempty"`
	// 图表的显示配置
	Display *ChartDisplay `json:"display,omitempty" xml:"display,omitempty" type:"Struct"`
	// 查询配置
	Search *ChartSearch `json:"search,omitempty" xml:"search,omitempty" type:"Struct"`
	// 图表标题。支持大小写英文字母、数字、下划线_、连字符-，连字符与下划线不能作为名称开头与结尾，长度必须在[2,64] 之间。
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 图表的类型。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Chart) String() string {
	return tea.Prettify(s)
}

func (s Chart) GoString() string {
	return s.String()
}

func (s *Chart) SetAction(v map[string]interface{}) *Chart {
	s.Action = v
	return s
}

func (s *Chart) SetDisplay(v *ChartDisplay) *Chart {
	s.Display = v
	return s
}

func (s *Chart) SetSearch(v *ChartSearch) *Chart {
	s.Search = v
	return s
}

func (s *Chart) SetTitle(v string) *Chart {
	s.Title = &v
	return s
}

func (s *Chart) SetType(v string) *Chart {
	s.Type = &v
	return s
}

type ChartDisplay struct {
	// 高度
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 宽度
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
	// x 轴
	XAxis []*string `json:"xAxis,omitempty" xml:"xAxis,omitempty" type:"Repeated"`
	// x 坐标
	XPos *int64 `json:"xPos,omitempty" xml:"xPos,omitempty"`
	// y 轴
	YAxis []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
	// y 坐标
	YPos *int64 `json:"yPos,omitempty" xml:"yPos,omitempty"`
}

func (s ChartDisplay) String() string {
	return tea.Prettify(s)
}

func (s ChartDisplay) GoString() string {
	return s.String()
}

func (s *ChartDisplay) SetHeight(v int64) *ChartDisplay {
	s.Height = &v
	return s
}

func (s *ChartDisplay) SetWidth(v int64) *ChartDisplay {
	s.Width = &v
	return s
}

func (s *ChartDisplay) SetXAxis(v []*string) *ChartDisplay {
	s.XAxis = v
	return s
}

func (s *ChartDisplay) SetXPos(v int64) *ChartDisplay {
	s.XPos = &v
	return s
}

func (s *ChartDisplay) SetYAxis(v []*string) *ChartDisplay {
	s.YAxis = v
	return s
}

func (s *ChartDisplay) SetYPos(v int64) *ChartDisplay {
	s.YPos = &v
	return s
}

type ChartSearch struct {
	// 结束时间
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// logstore 名称
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// 查询语句
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 开始时间
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s ChartSearch) String() string {
	return tea.Prettify(s)
}

func (s ChartSearch) GoString() string {
	return s.String()
}

func (s *ChartSearch) SetEnd(v string) *ChartSearch {
	s.End = &v
	return s
}

func (s *ChartSearch) SetLogstore(v string) *ChartSearch {
	s.Logstore = &v
	return s
}

func (s *ChartSearch) SetQuery(v string) *ChartSearch {
	s.Query = &v
	return s
}

func (s *ChartSearch) SetStart(v string) *ChartSearch {
	s.Start = &v
	return s
}

func (s *ChartSearch) SetTopic(v string) *ChartSearch {
	s.Topic = &v
	return s
}

type Dashboard struct {
	// 属性值，可用于修改仪表盘的布局等属性，例如 "type": "free" 自由布局， "type":"grid" 网格布局。
	Attribute map[string]*string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 仪表盘内包含的图表。
	Charts []*Chart `json:"charts,omitempty" xml:"charts,omitempty" type:"Repeated"`
	// 仪表盘ID。同一个Project下，仪表盘ID唯一，不可重复。
	DashboardName *string `json:"dashboardName,omitempty" xml:"dashboardName,omitempty"`
	// 描述信息。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 仪表盘的展示名称。
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
}

func (s Dashboard) String() string {
	return tea.Prettify(s)
}

func (s Dashboard) GoString() string {
	return s.String()
}

func (s *Dashboard) SetAttribute(v map[string]*string) *Dashboard {
	s.Attribute = v
	return s
}

func (s *Dashboard) SetCharts(v []*Chart) *Dashboard {
	s.Charts = v
	return s
}

func (s *Dashboard) SetDashboardName(v string) *Dashboard {
	s.DashboardName = &v
	return s
}

func (s *Dashboard) SetDescription(v string) *Dashboard {
	s.Description = &v
	return s
}

func (s *Dashboard) SetDisplayName(v string) *Dashboard {
	s.DisplayName = &v
	return s
}

type EtlJob struct {
	// 是否启用
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// 任务名称
	EtlJobName *string `json:"etlJobName,omitempty" xml:"etlJobName,omitempty"`
	// 运行函数配置
	FunctionConfig *EtlJobFunctionConfig `json:"functionConfig,omitempty" xml:"functionConfig,omitempty" type:"Struct"`
	// 参数列表
	FunctionParameter map[string]interface{} `json:"functionParameter,omitempty" xml:"functionParameter,omitempty"`
	// 日志配置
	LogConfig *EtlJobLogConfig `json:"logConfig,omitempty" xml:"logConfig,omitempty" type:"Struct"`
	// 配置数据来源
	SourceConfig *EtlJobSourceConfig `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty" type:"Struct"`
	// 触发器配置
	TriggerConfig *EtlJobTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Struct"`
}

func (s EtlJob) String() string {
	return tea.Prettify(s)
}

func (s EtlJob) GoString() string {
	return s.String()
}

func (s *EtlJob) SetEnable(v bool) *EtlJob {
	s.Enable = &v
	return s
}

func (s *EtlJob) SetEtlJobName(v string) *EtlJob {
	s.EtlJobName = &v
	return s
}

func (s *EtlJob) SetFunctionConfig(v *EtlJobFunctionConfig) *EtlJob {
	s.FunctionConfig = v
	return s
}

func (s *EtlJob) SetFunctionParameter(v map[string]interface{}) *EtlJob {
	s.FunctionParameter = v
	return s
}

func (s *EtlJob) SetLogConfig(v *EtlJobLogConfig) *EtlJob {
	s.LogConfig = v
	return s
}

func (s *EtlJob) SetSourceConfig(v *EtlJobSourceConfig) *EtlJob {
	s.SourceConfig = v
	return s
}

func (s *EtlJob) SetTriggerConfig(v *EtlJobTriggerConfig) *EtlJob {
	s.TriggerConfig = v
	return s
}

type EtlJobFunctionConfig struct {
	// 账户 id
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// 函数名
	FunctionName *string `json:"functionName,omitempty" xml:"functionName,omitempty"`
	// 函数 provider，可选值为 FunctionCompute 、CloudProdLogDispatch。当值为 FunctionCompute 时，endpoint、accountid 、regionName 、serviceName 、functionName 必选。
	FunctionProvider *string `json:"functionProvider,omitempty" xml:"functionProvider,omitempty"`
	// 地域
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// 角色授权
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// 服务名
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s EtlJobFunctionConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobFunctionConfig) GoString() string {
	return s.String()
}

func (s *EtlJobFunctionConfig) SetAccountId(v string) *EtlJobFunctionConfig {
	s.AccountId = &v
	return s
}

func (s *EtlJobFunctionConfig) SetEndpoint(v string) *EtlJobFunctionConfig {
	s.Endpoint = &v
	return s
}

func (s *EtlJobFunctionConfig) SetFunctionName(v string) *EtlJobFunctionConfig {
	s.FunctionName = &v
	return s
}

func (s *EtlJobFunctionConfig) SetFunctionProvider(v string) *EtlJobFunctionConfig {
	s.FunctionProvider = &v
	return s
}

func (s *EtlJobFunctionConfig) SetRegionName(v string) *EtlJobFunctionConfig {
	s.RegionName = &v
	return s
}

func (s *EtlJobFunctionConfig) SetRoleArn(v string) *EtlJobFunctionConfig {
	s.RoleArn = &v
	return s
}

func (s *EtlJobFunctionConfig) SetServiceName(v string) *EtlJobFunctionConfig {
	s.ServiceName = &v
	return s
}

type EtlJobLogConfig struct {
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// logstore 名称
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// project 名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s EtlJobLogConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobLogConfig) GoString() string {
	return s.String()
}

func (s *EtlJobLogConfig) SetEndpoint(v string) *EtlJobLogConfig {
	s.Endpoint = &v
	return s
}

func (s *EtlJobLogConfig) SetLogstoreName(v string) *EtlJobLogConfig {
	s.LogstoreName = &v
	return s
}

func (s *EtlJobLogConfig) SetProjectName(v string) *EtlJobLogConfig {
	s.ProjectName = &v
	return s
}

type EtlJobSourceConfig struct {
	// logstore 名称
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
}

func (s EtlJobSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobSourceConfig) GoString() string {
	return s.String()
}

func (s *EtlJobSourceConfig) SetLogstoreName(v string) *EtlJobSourceConfig {
	s.LogstoreName = &v
	return s
}

type EtlJobTriggerConfig struct {
	// 最大重试次数，必须在[0,100] 之间
	MaxRetryTime *int32 `json:"maxRetryTime,omitempty" xml:"maxRetryTime,omitempty"`
	// 角色授权配置
	RoleArn *string `json:"roleArn,omitempty" xml:"roleArn,omitempty"`
	// 开始位置，可选 latest、at-unixtime， 默认 latest。
	StartingPosition *string `json:"startingPosition,omitempty" xml:"startingPosition,omitempty"`
	// 开始时间
	StartingUnixtime *int64 `json:"startingUnixtime,omitempty" xml:"startingUnixtime,omitempty"`
	// 触发间隔，单位为秒，必须在 [3,600] 之间
	TriggerInterval *int32 `json:"triggerInterval,omitempty" xml:"triggerInterval,omitempty"`
}

func (s EtlJobTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s EtlJobTriggerConfig) GoString() string {
	return s.String()
}

func (s *EtlJobTriggerConfig) SetMaxRetryTime(v int32) *EtlJobTriggerConfig {
	s.MaxRetryTime = &v
	return s
}

func (s *EtlJobTriggerConfig) SetRoleArn(v string) *EtlJobTriggerConfig {
	s.RoleArn = &v
	return s
}

func (s *EtlJobTriggerConfig) SetStartingPosition(v string) *EtlJobTriggerConfig {
	s.StartingPosition = &v
	return s
}

func (s *EtlJobTriggerConfig) SetStartingUnixtime(v int64) *EtlJobTriggerConfig {
	s.StartingUnixtime = &v
	return s
}

func (s *EtlJobTriggerConfig) SetTriggerInterval(v int32) *EtlJobTriggerConfig {
	s.TriggerInterval = &v
	return s
}

type EtlMeta struct {
	// 是否启用
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// key
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// tag
	EtlMetaTag *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	// value
	EtlMetaValue *string `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s EtlMeta) String() string {
	return tea.Prettify(s)
}

func (s EtlMeta) GoString() string {
	return s.String()
}

func (s *EtlMeta) SetEnable(v bool) *EtlMeta {
	s.Enable = &v
	return s
}

func (s *EtlMeta) SetEtlMetaKey(v string) *EtlMeta {
	s.EtlMetaKey = &v
	return s
}

func (s *EtlMeta) SetEtlMetaName(v string) *EtlMeta {
	s.EtlMetaName = &v
	return s
}

func (s *EtlMeta) SetEtlMetaTag(v string) *EtlMeta {
	s.EtlMetaTag = &v
	return s
}

func (s *EtlMeta) SetEtlMetaValue(v string) *EtlMeta {
	s.EtlMetaValue = &v
	return s
}

type ExternalStore struct {
	// 外部存储的名称。
	ExternalStoreName *string `json:"externalStoreName,omitempty" xml:"externalStoreName,omitempty"`
	// 参数
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// 类型。可选 rds-vpc 或者 oss
	StoreType *string `json:"storeType,omitempty" xml:"storeType,omitempty"`
}

func (s ExternalStore) String() string {
	return tea.Prettify(s)
}

func (s ExternalStore) GoString() string {
	return s.String()
}

func (s *ExternalStore) SetExternalStoreName(v string) *ExternalStore {
	s.ExternalStoreName = &v
	return s
}

func (s *ExternalStore) SetParameter(v map[string]interface{}) *ExternalStore {
	s.Parameter = v
	return s
}

func (s *ExternalStore) SetStoreType(v string) *ExternalStore {
	s.StoreType = &v
	return s
}

type Logging struct {
	// logging 配置项
	LoggingDetails []*LoggingLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// project 名称。
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s Logging) String() string {
	return tea.Prettify(s)
}

func (s Logging) GoString() string {
	return s.String()
}

func (s *Logging) SetLoggingDetails(v []*LoggingLoggingDetails) *Logging {
	s.LoggingDetails = v
	return s
}

func (s *Logging) SetLoggingProject(v string) *Logging {
	s.LoggingProject = &v
	return s
}

type LoggingLoggingDetails struct {
	// logstore 名称。
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// logging 类型。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s LoggingLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s LoggingLoggingDetails) GoString() string {
	return s.String()
}

func (s *LoggingLoggingDetails) SetLogstore(v string) *LoggingLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *LoggingLoggingDetails) SetType(v string) *LoggingLoggingDetails {
	s.Type = &v
	return s
}

type Logstore struct {
	// 接收日志后，自动添加客户端外网IP和日志到达时间
	AppendMeta *bool `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	// 是否开启 shard 自动分裂。当写入数据量超过已有分区（Shard）写入服务能力且持续5分钟以上时，开启自动分裂功能可自动根据数据量增加分区数量
	AutoSplit *bool `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// 创建时间。
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// WebTracking功能支持快速采集各种浏览器以及iOS/Android/APP访问信息，默认关闭
	EnableTracking *bool `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	// Encrypt configuration
	EncryptConf *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	// 必须在 (30, ttl) 之间
	HotTtl *int32 `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	// 最后修改时间。
	LastModifyTime *int32 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// logstore 的名称。
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	// 最大 shard 数量。
	MaxSplitShard *int32 `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	// shard 数量。
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	// telemetryType
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	// 数据保存的天数。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s Logstore) String() string {
	return tea.Prettify(s)
}

func (s Logstore) GoString() string {
	return s.String()
}

func (s *Logstore) SetAppendMeta(v bool) *Logstore {
	s.AppendMeta = &v
	return s
}

func (s *Logstore) SetAutoSplit(v bool) *Logstore {
	s.AutoSplit = &v
	return s
}

func (s *Logstore) SetCreateTime(v int32) *Logstore {
	s.CreateTime = &v
	return s
}

func (s *Logstore) SetEnableTracking(v bool) *Logstore {
	s.EnableTracking = &v
	return s
}

func (s *Logstore) SetEncryptConf(v *EncryptConf) *Logstore {
	s.EncryptConf = v
	return s
}

func (s *Logstore) SetHotTtl(v int32) *Logstore {
	s.HotTtl = &v
	return s
}

func (s *Logstore) SetLastModifyTime(v int32) *Logstore {
	s.LastModifyTime = &v
	return s
}

func (s *Logstore) SetLogstoreName(v string) *Logstore {
	s.LogstoreName = &v
	return s
}

func (s *Logstore) SetMaxSplitShard(v int32) *Logstore {
	s.MaxSplitShard = &v
	return s
}

func (s *Logstore) SetShardCount(v int32) *Logstore {
	s.ShardCount = &v
	return s
}

func (s *Logstore) SetTelemetryType(v string) *Logstore {
	s.TelemetryType = &v
	return s
}

func (s *Logstore) SetTtl(v int32) *Logstore {
	s.Ttl = &v
	return s
}

type Machine struct {
	// 机器 ip 地址。
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 最后一次心跳时间。Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	LastHeartbeatTime *int64 `json:"lastHeartbeatTime,omitempty" xml:"lastHeartbeatTime,omitempty"`
	// 机器的唯一标识。
	MachineUniqueid *string `json:"machine-uniqueid,omitempty" xml:"machine-uniqueid,omitempty"`
	// 机器的用户自定义标识。
	UserdefinedId *string `json:"userdefined-id,omitempty" xml:"userdefined-id,omitempty"`
}

func (s Machine) String() string {
	return tea.Prettify(s)
}

func (s Machine) GoString() string {
	return s.String()
}

func (s *Machine) SetIp(v string) *Machine {
	s.Ip = &v
	return s
}

func (s *Machine) SetLastHeartbeatTime(v int64) *Machine {
	s.LastHeartbeatTime = &v
	return s
}

func (s *Machine) SetMachineUniqueid(v string) *Machine {
	s.MachineUniqueid = &v
	return s
}

func (s *Machine) SetUserdefinedId(v string) *Machine {
	s.UserdefinedId = &v
	return s
}

type MachineGroup struct {
	// 机器组属性。
	GroupAttribute *MachineGroupGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// 机器组名称。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 机器组种类。目前固定为空字符串。
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 机器组标识种类，支持 IP 标识或者用户自定义标识，即 ip 、userdefined。
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// 机器组标识列表。
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s MachineGroup) String() string {
	return tea.Prettify(s)
}

func (s MachineGroup) GoString() string {
	return s.String()
}

func (s *MachineGroup) SetGroupAttribute(v *MachineGroupGroupAttribute) *MachineGroup {
	s.GroupAttribute = v
	return s
}

func (s *MachineGroup) SetGroupName(v string) *MachineGroup {
	s.GroupName = &v
	return s
}

func (s *MachineGroup) SetGroupType(v string) *MachineGroup {
	s.GroupType = &v
	return s
}

func (s *MachineGroup) SetMachineIdentifyType(v string) *MachineGroup {
	s.MachineIdentifyType = &v
	return s
}

func (s *MachineGroup) SetMachineList(v []*string) *MachineGroup {
	s.MachineList = v
	return s
}

type MachineGroupGroupAttribute struct {
	// 机器组所依赖的外部管理系统标识。
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// 机器组的日志主题。
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s MachineGroupGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s MachineGroupGroupAttribute) GoString() string {
	return s.String()
}

func (s *MachineGroupGroupAttribute) SetExternalName(v string) *MachineGroupGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *MachineGroupGroupAttribute) SetGroupTopic(v string) *MachineGroupGroupAttribute {
	s.GroupTopic = &v
	return s
}

type Project struct {
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 最后更新时间
	LastModifyTime *string `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// owner
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// Project名称
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 所在区域
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Project) String() string {
	return tea.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) SetCreateTime(v string) *Project {
	s.CreateTime = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetLastModifyTime(v string) *Project {
	s.LastModifyTime = &v
	return s
}

func (s *Project) SetOwner(v string) *Project {
	s.Owner = &v
	return s
}

func (s *Project) SetProjectName(v string) *Project {
	s.ProjectName = &v
	return s
}

func (s *Project) SetRegion(v string) *Project {
	s.Region = &v
	return s
}

func (s *Project) SetStatus(v string) *Project {
	s.Status = &v
	return s
}

type Shard struct {
	// Shard的创建时间。Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	CreateTime *int32 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 指定Shard范围的结束值，Shard范围中不包含该值。即 shard 包含MD5值在 [inclusiveBeginKey, exclusiveEndKey) 之间的日志。
	ExclusiveEndKey *string `json:"exclusiveEndKey,omitempty" xml:"exclusiveEndKey,omitempty"`
	// 指定Shard范围的起始值，Shard范围中包含该值。即 shard 包含MD5值在 [inclusiveBeginKey, exclusiveEndKey) 之间的日志。
	InclusiveBeginKey *string `json:"inclusiveBeginKey,omitempty" xml:"inclusiveBeginKey,omitempty"`
	// shard id
	ShardID *int32 `json:"shardID,omitempty" xml:"shardID,omitempty"`
	// shard 的读写状态，readwrite 或者 readonly。
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s Shard) String() string {
	return tea.Prettify(s)
}

func (s Shard) GoString() string {
	return s.String()
}

func (s *Shard) SetCreateTime(v int32) *Shard {
	s.CreateTime = &v
	return s
}

func (s *Shard) SetExclusiveEndKey(v string) *Shard {
	s.ExclusiveEndKey = &v
	return s
}

func (s *Shard) SetInclusiveBeginKey(v string) *Shard {
	s.InclusiveBeginKey = &v
	return s
}

func (s *Shard) SetShardID(v int32) *Shard {
	s.ShardID = &v
	return s
}

func (s *Shard) SetStatus(v string) *Shard {
	s.Status = &v
	return s
}

type ApplyConfigToMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ApplyConfigToMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyConfigToMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *ApplyConfigToMachineGroupResponse) SetHeaders(v map[string]*string) *ApplyConfigToMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *ApplyConfigToMachineGroupResponse) SetStatusCode(v int32) *ApplyConfigToMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type BatchCreateEtlMetaRequest struct {
	EtlMetaList []*BatchCreateEtlMetaRequestEtlMetaList `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
}

func (s BatchCreateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaRequest) SetEtlMetaList(v []*BatchCreateEtlMetaRequestEtlMetaList) *BatchCreateEtlMetaRequest {
	s.EtlMetaList = v
	return s
}

type BatchCreateEtlMetaRequestEtlMetaList struct {
	// 是否启用
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,255]之间。
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字。由数字、大小写字母、下划线_、连字符-组成，长度需要在[2,64]之间。
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,128]之间。
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s BatchCreateEtlMetaRequestEtlMetaList) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaRequestEtlMetaList) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEnable(v bool) *BatchCreateEtlMetaRequestEtlMetaList {
	s.Enable = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaKey(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaKey = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaName(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaName = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaTag(v string) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchCreateEtlMetaRequestEtlMetaList) SetEtlMetaValue(v map[string]interface{}) *BatchCreateEtlMetaRequestEtlMetaList {
	s.EtlMetaValue = v
	return s
}

type BatchCreateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchCreateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateEtlMetaResponse) SetHeaders(v map[string]*string) *BatchCreateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateEtlMetaResponse) SetStatusCode(v int32) *BatchCreateEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type BatchModifyEtlMetaStatusRequest struct {
	// 当 range 的值为 "list" 时有效，匹配list中的 metaKey
	EtlMetaKeyList []*string `json:"etlMetaKeyList,omitempty" xml:"etlMetaKeyList,omitempty" type:"Repeated"`
	EtlMetaName    *string   `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// 匹配的 tag，当 tag 为 "__all_etl_meta_tag_match__" 时表示全部匹配。
	EtlMetaTag *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	// 操作作用的范围，可选 all 代表匹配全部，list 按名单列表匹配 key 两种模式。
	Range *string `json:"range,omitempty" xml:"range,omitempty"`
	// 操作类型，支持启用、禁用、删除三种，即 batch_enable、batch_disable、batch_delete。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s BatchModifyEtlMetaStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyEtlMetaStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaKeyList(v []*string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaKeyList = v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaName(v string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaName = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetEtlMetaTag(v string) *BatchModifyEtlMetaStatusRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetRange(v string) *BatchModifyEtlMetaStatusRequest {
	s.Range = &v
	return s
}

func (s *BatchModifyEtlMetaStatusRequest) SetType(v string) *BatchModifyEtlMetaStatusRequest {
	s.Type = &v
	return s
}

type BatchModifyEtlMetaStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchModifyEtlMetaStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchModifyEtlMetaStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchModifyEtlMetaStatusResponse) SetHeaders(v map[string]*string) *BatchModifyEtlMetaStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchModifyEtlMetaStatusResponse) SetStatusCode(v int32) *BatchModifyEtlMetaStatusResponse {
	s.StatusCode = &v
	return s
}

type BatchUpdateEtlMetaRequest struct {
	EtlMetaList *BatchUpdateEtlMetaRequestEtlMetaList `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Struct"`
}

func (s BatchUpdateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaRequest) SetEtlMetaList(v *BatchUpdateEtlMetaRequestEtlMetaList) *BatchUpdateEtlMetaRequest {
	s.EtlMetaList = v
	return s
}

type BatchUpdateEtlMetaRequestEtlMetaList struct {
	// 是否启用。etlMetaTag、etlMetaValue、enable 至少需要存在一个。
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,255]之间。
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字。由数字、大小写字母、下划线_、连字符-组成，长度需要在[2,64]之间。
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,128]之间。
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s BatchUpdateEtlMetaRequestEtlMetaList) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaRequestEtlMetaList) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEnable(v bool) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.Enable = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaKey(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaKey = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaName(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaName = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaTag(v string) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaTag = &v
	return s
}

func (s *BatchUpdateEtlMetaRequestEtlMetaList) SetEtlMetaValue(v map[string]interface{}) *BatchUpdateEtlMetaRequestEtlMetaList {
	s.EtlMetaValue = v
	return s
}

type BatchUpdateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s BatchUpdateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateEtlMetaResponse) SetHeaders(v map[string]*string) *BatchUpdateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateEtlMetaResponse) SetStatusCode(v int32) *BatchUpdateEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type CreateConsumerGroupRequest struct {
	ConsumerGroup *string `json:"consumerGroup,omitempty" xml:"consumerGroup,omitempty"`
	Order         *bool   `json:"order,omitempty" xml:"order,omitempty"`
	Timeout       *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetConsumerGroup(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroup = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetOrder(v bool) *CreateConsumerGroupRequest {
	s.Order = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetTimeout(v int32) *CreateConsumerGroupRequest {
	s.Timeout = &v
	return s
}

type CreateConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetStatusCode(v int32) *CreateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateDomainRequest struct {
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

type CreateEtlMetaRequest struct {
	// 是否启用
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,255]之间。
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字。由数字、大小写字母、下划线_、连字符-组成，长度需要在[2,64]之间。
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,128]之间。
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s CreateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateEtlMetaRequest) SetEnable(v bool) *CreateEtlMetaRequest {
	s.Enable = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaKey(v string) *CreateEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaName(v string) *CreateEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaTag(v string) *CreateEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *CreateEtlMetaRequest) SetEtlMetaValue(v map[string]interface{}) *CreateEtlMetaRequest {
	s.EtlMetaValue = v
	return s
}

type CreateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateEtlMetaResponse) SetHeaders(v map[string]*string) *CreateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateEtlMetaResponse) SetStatusCode(v int32) *CreateEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type CreateIndexRequest struct {
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// 配置全文索引
	Line *CreateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// 开启日志聚类，开启后白名单与黑名单至多生效其中一个。
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// 日志聚类的聚类字段黑名单
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// 日志聚类的聚类字段白名单
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// 统计字段的最大长度
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// 保存时间，单位为天
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetKeys(v map[string]*KeysValue) *CreateIndexRequest {
	s.Keys = v
	return s
}

func (s *CreateIndexRequest) SetLine(v *CreateIndexRequestLine) *CreateIndexRequest {
	s.Line = v
	return s
}

func (s *CreateIndexRequest) SetLogReduce(v bool) *CreateIndexRequest {
	s.LogReduce = &v
	return s
}

func (s *CreateIndexRequest) SetLogReduceBlackList(v []*string) *CreateIndexRequest {
	s.LogReduceBlackList = v
	return s
}

func (s *CreateIndexRequest) SetLogReduceWhiteList(v []*string) *CreateIndexRequest {
	s.LogReduceWhiteList = v
	return s
}

func (s *CreateIndexRequest) SetMaxTextLen(v int32) *CreateIndexRequest {
	s.MaxTextLen = &v
	return s
}

func (s *CreateIndexRequest) SetTtl(v int32) *CreateIndexRequest {
	s.Ttl = &v
	return s
}

type CreateIndexRequestLine struct {
	// 大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// 包含中文
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// 排除的字段列表，不能与include_keys同时指定。
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// 包含的字段列表，不能与exclude_keys同时指定。
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// 分词符列表。可以设置一个分词参数，指定这个字段按照哪一种方式分词。
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s CreateIndexRequestLine) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestLine) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestLine) SetCaseSensitive(v bool) *CreateIndexRequestLine {
	s.CaseSensitive = &v
	return s
}

func (s *CreateIndexRequestLine) SetChn(v bool) *CreateIndexRequestLine {
	s.Chn = &v
	return s
}

func (s *CreateIndexRequestLine) SetExcludeKeys(v []*string) *CreateIndexRequestLine {
	s.ExcludeKeys = v
	return s
}

func (s *CreateIndexRequestLine) SetIncludeKeys(v []*string) *CreateIndexRequestLine {
	s.IncludeKeys = v
	return s
}

func (s *CreateIndexRequestLine) SetToken(v []*string) *CreateIndexRequestLine {
	s.Token = v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

type CreateLogStoreRequest struct {
	AppendMeta     *bool        `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	AutoSplit      *bool        `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	LogstoreName   *string      `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	MaxSplitShard  *int32       `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	ShardCount     *int32       `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType  *string      `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	Ttl            *int32       `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateLogStoreRequest) SetAppendMeta(v bool) *CreateLogStoreRequest {
	s.AppendMeta = &v
	return s
}

func (s *CreateLogStoreRequest) SetAutoSplit(v bool) *CreateLogStoreRequest {
	s.AutoSplit = &v
	return s
}

func (s *CreateLogStoreRequest) SetEnableTracking(v bool) *CreateLogStoreRequest {
	s.EnableTracking = &v
	return s
}

func (s *CreateLogStoreRequest) SetEncryptConf(v *EncryptConf) *CreateLogStoreRequest {
	s.EncryptConf = v
	return s
}

func (s *CreateLogStoreRequest) SetHotTtl(v int32) *CreateLogStoreRequest {
	s.HotTtl = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogstoreName(v string) *CreateLogStoreRequest {
	s.LogstoreName = &v
	return s
}

func (s *CreateLogStoreRequest) SetMaxSplitShard(v int32) *CreateLogStoreRequest {
	s.MaxSplitShard = &v
	return s
}

func (s *CreateLogStoreRequest) SetShardCount(v int32) *CreateLogStoreRequest {
	s.ShardCount = &v
	return s
}

func (s *CreateLogStoreRequest) SetTelemetryType(v string) *CreateLogStoreRequest {
	s.TelemetryType = &v
	return s
}

func (s *CreateLogStoreRequest) SetTtl(v int32) *CreateLogStoreRequest {
	s.Ttl = &v
	return s
}

type CreateLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateLogStoreResponse) SetHeaders(v map[string]*string) *CreateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateLogStoreResponse) SetStatusCode(v int32) *CreateLogStoreResponse {
	s.StatusCode = &v
	return s
}

type CreateLoggingRequest struct {
	// 服务日志配置列表。
	LoggingDetails []*CreateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// 服务日志要保存到的 project 名称。
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s CreateLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingRequest) GoString() string {
	return s.String()
}

func (s *CreateLoggingRequest) SetLoggingDetails(v []*CreateLoggingRequestLoggingDetails) *CreateLoggingRequest {
	s.LoggingDetails = v
	return s
}

func (s *CreateLoggingRequest) SetLoggingProject(v string) *CreateLoggingRequest {
	s.LoggingProject = &v
	return s
}

type CreateLoggingRequestLoggingDetails struct {
	// 该种类服务日志要保存到的 logstore 名称。
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// 服务日志的种类。可选 "consumergroup_log"、 "logtail_alarm"、"operation_log"、"logtail_profile"、"metering"、"logtail_status"、"scheduled_sql_alert"、 "etl_alert" 等。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateLoggingRequestLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingRequestLoggingDetails) GoString() string {
	return s.String()
}

func (s *CreateLoggingRequestLoggingDetails) SetLogstore(v string) *CreateLoggingRequestLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *CreateLoggingRequestLoggingDetails) SetType(v string) *CreateLoggingRequestLoggingDetails {
	s.Type = &v
	return s
}

type CreateLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoggingResponse) GoString() string {
	return s.String()
}

func (s *CreateLoggingResponse) SetHeaders(v map[string]*string) *CreateLoggingResponse {
	s.Headers = v
	return s
}

func (s *CreateLoggingResponse) SetStatusCode(v int32) *CreateLoggingResponse {
	s.StatusCode = &v
	return s
}

type CreateMachineGroupRequest struct {
	// 机器组属性。
	GroupAttribute *CreateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// 机器组名称。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 机器组类型，可选值，默认为空。
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 机器组标识种类，支持 ip 、userdefined 两种。
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// 机器列表。
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s CreateMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupRequest) SetGroupAttribute(v *CreateMachineGroupRequestGroupAttribute) *CreateMachineGroupRequest {
	s.GroupAttribute = v
	return s
}

func (s *CreateMachineGroupRequest) SetGroupName(v string) *CreateMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateMachineGroupRequest) SetGroupType(v string) *CreateMachineGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateMachineGroupRequest) SetMachineIdentifyType(v string) *CreateMachineGroupRequest {
	s.MachineIdentifyType = &v
	return s
}

func (s *CreateMachineGroupRequest) SetMachineList(v []*string) *CreateMachineGroupRequest {
	s.MachineList = v
	return s
}

type CreateMachineGroupRequestGroupAttribute struct {
	// 机器组所依赖的外部管理系统标识。
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// 机器组的日志主题。
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s CreateMachineGroupRequestGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupRequestGroupAttribute) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupRequestGroupAttribute) SetExternalName(v string) *CreateMachineGroupRequestGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *CreateMachineGroupRequestGroupAttribute) SetGroupTopic(v string) *CreateMachineGroupRequestGroupAttribute {
	s.GroupTopic = &v
	return s
}

type CreateMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMachineGroupResponse) SetHeaders(v map[string]*string) *CreateMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMachineGroupResponse) SetStatusCode(v int32) *CreateMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type CreateProjectRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

type CreateSavedSearchRequest struct {
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Logstore        *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	SearchQuery     *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic           *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s CreateSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *CreateSavedSearchRequest) SetDisplayName(v string) *CreateSavedSearchRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSavedSearchRequest) SetLogstore(v string) *CreateSavedSearchRequest {
	s.Logstore = &v
	return s
}

func (s *CreateSavedSearchRequest) SetSavedsearchName(v string) *CreateSavedSearchRequest {
	s.SavedsearchName = &v
	return s
}

func (s *CreateSavedSearchRequest) SetSearchQuery(v string) *CreateSavedSearchRequest {
	s.SearchQuery = &v
	return s
}

func (s *CreateSavedSearchRequest) SetTopic(v string) *CreateSavedSearchRequest {
	s.Topic = &v
	return s
}

type CreateSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CreateSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *CreateSavedSearchResponse) SetHeaders(v map[string]*string) *CreateSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *CreateSavedSearchResponse) SetStatusCode(v int32) *CreateSavedSearchResponse {
	s.StatusCode = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetStatusCode(v int32) *DeleteConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

type DeleteEtlMetaRequest struct {
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,255]之间。
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字。由数字、大小写字母、下划线_、连字符-组成，长度需要在[2,64]之间。
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// 此处固定为 "__all_etl_meta_tag_match__"
	EtlMetaTag *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
}

func (s DeleteEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *DeleteEtlMetaRequest) SetEtlMetaKey(v string) *DeleteEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *DeleteEtlMetaRequest) SetEtlMetaName(v string) *DeleteEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *DeleteEtlMetaRequest) SetEtlMetaTag(v string) *DeleteEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

type DeleteEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *DeleteEtlMetaResponse) SetHeaders(v map[string]*string) *DeleteEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *DeleteEtlMetaResponse) SetStatusCode(v int32) *DeleteEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type DeleteIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

type DeleteLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogStoreResponse) SetHeaders(v map[string]*string) *DeleteLogStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogStoreResponse) SetStatusCode(v int32) *DeleteLogStoreResponse {
	s.StatusCode = &v
	return s
}

type DeleteLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoggingResponse) SetHeaders(v map[string]*string) *DeleteLoggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoggingResponse) SetStatusCode(v int32) *DeleteLoggingResponse {
	s.StatusCode = &v
	return s
}

type DeleteMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteMachineGroupResponse) SetHeaders(v map[string]*string) *DeleteMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteMachineGroupResponse) SetStatusCode(v int32) *DeleteMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

type DeleteSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *DeleteSavedSearchResponse) SetHeaders(v map[string]*string) *DeleteSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *DeleteSavedSearchResponse) SetStatusCode(v int32) *DeleteSavedSearchResponse {
	s.StatusCode = &v
	return s
}

type GetAppliedConfigsResponseBody struct {
	// Logtail配置名称列表。
	Configs []*string `json:"configs,omitempty" xml:"configs,omitempty" type:"Repeated"`
	// Logtail配置数量。
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s GetAppliedConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppliedConfigsResponseBody) SetConfigs(v []*string) *GetAppliedConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *GetAppliedConfigsResponseBody) SetCount(v int32) *GetAppliedConfigsResponseBody {
	s.Count = &v
	return s
}

type GetAppliedConfigsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppliedConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppliedConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppliedConfigsResponse) GoString() string {
	return s.String()
}

func (s *GetAppliedConfigsResponse) SetHeaders(v map[string]*string) *GetAppliedConfigsResponse {
	s.Headers = v
	return s
}

func (s *GetAppliedConfigsResponse) SetStatusCode(v int32) *GetAppliedConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppliedConfigsResponse) SetBody(v *GetAppliedConfigsResponseBody) *GetAppliedConfigsResponse {
	s.Body = v
	return s
}

type GetCheckPointRequest struct {
	// Shard ID。
	// 如果指定的Shard不存在，则返回空列表。
	// 如果不指定Shard，则返回所有Shard的checkpoint。
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
}

func (s GetCheckPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointRequest) GoString() string {
	return s.String()
}

func (s *GetCheckPointRequest) SetShard(v int32) *GetCheckPointRequest {
	s.Shard = &v
	return s
}

type GetCheckPointResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*GetCheckPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetCheckPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointResponse) GoString() string {
	return s.String()
}

func (s *GetCheckPointResponse) SetHeaders(v map[string]*string) *GetCheckPointResponse {
	s.Headers = v
	return s
}

func (s *GetCheckPointResponse) SetStatusCode(v int32) *GetCheckPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckPointResponse) SetBody(v []*GetCheckPointResponseBody) *GetCheckPointResponse {
	s.Body = v
	return s
}

type GetCheckPointResponseBody struct {
	// shard id。
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
	// checkpoint 值。
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// checkpoint最后的更新时间。Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// 消费者。
	Consumer *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
}

func (s GetCheckPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCheckPointResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckPointResponseBody) SetShard(v int32) *GetCheckPointResponseBody {
	s.Shard = &v
	return s
}

func (s *GetCheckPointResponseBody) SetCheckpoint(v string) *GetCheckPointResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *GetCheckPointResponseBody) SetUpdateTime(v int64) *GetCheckPointResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetCheckPointResponseBody) SetConsumer(v string) *GetCheckPointResponseBody {
	s.Consumer = &v
	return s
}

type GetContextLogsRequest struct {
	// 指定起始日志往前（上文）的日志条数，取值范围为(0,100]。
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// 指定起始日志往后（下文）的日志条数，取值范围为(0,100]。
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// 起始日志所属的LogGroup的唯一身份标识。
	PackId *string `json:"pack_id,omitempty" xml:"pack_id,omitempty"`
	// 起始日志在对应LogGroup内的唯一上下文结构标识。
	PackMeta *string `json:"pack_meta,omitempty" xml:"pack_meta,omitempty"`
	// Logstore中数据的类型。该接口中该参数固定为context_log。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetContextLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsRequest) GoString() string {
	return s.String()
}

func (s *GetContextLogsRequest) SetBackLines(v int64) *GetContextLogsRequest {
	s.BackLines = &v
	return s
}

func (s *GetContextLogsRequest) SetForwardLines(v int64) *GetContextLogsRequest {
	s.ForwardLines = &v
	return s
}

func (s *GetContextLogsRequest) SetPackId(v string) *GetContextLogsRequest {
	s.PackId = &v
	return s
}

func (s *GetContextLogsRequest) SetPackMeta(v string) *GetContextLogsRequest {
	s.PackMeta = &v
	return s
}

func (s *GetContextLogsRequest) SetType(v string) *GetContextLogsRequest {
	s.Type = &v
	return s
}

type GetContextLogsResponseBody struct {
	// 向前查询到的日志条数。
	BackLines *int64 `json:"back_lines,omitempty" xml:"back_lines,omitempty"`
	// 向后查询到的日志条数。
	ForwardLines *int64 `json:"forward_lines,omitempty" xml:"forward_lines,omitempty"`
	// 获取到的日志，按上下文顺序排列。当根据指定起始日志查询不到上下文日志时，此参数为空。
	Logs []map[string]interface{} `json:"logs,omitempty" xml:"logs,omitempty" type:"Repeated"`
	// 查询的结果是否完整。
	// Complete：查询已经完成，返回结果为完整结果。
	// Incomplete：查询已经完成，返回结果为不完整结果，需要重复请求以获得完整结果。
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
	// 返回的总日志条数，包含请求参数中所指定的起始日志。
	TotalLines *int64 `json:"total_lines,omitempty" xml:"total_lines,omitempty"`
}

func (s GetContextLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextLogsResponseBody) SetBackLines(v int64) *GetContextLogsResponseBody {
	s.BackLines = &v
	return s
}

func (s *GetContextLogsResponseBody) SetForwardLines(v int64) *GetContextLogsResponseBody {
	s.ForwardLines = &v
	return s
}

func (s *GetContextLogsResponseBody) SetLogs(v []map[string]interface{}) *GetContextLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetContextLogsResponseBody) SetProgress(v string) *GetContextLogsResponseBody {
	s.Progress = &v
	return s
}

func (s *GetContextLogsResponseBody) SetTotalLines(v int64) *GetContextLogsResponseBody {
	s.TotalLines = &v
	return s
}

type GetContextLogsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetContextLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetContextLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetContextLogsResponse) GoString() string {
	return s.String()
}

func (s *GetContextLogsResponse) SetHeaders(v map[string]*string) *GetContextLogsResponse {
	s.Headers = v
	return s
}

func (s *GetContextLogsResponse) SetStatusCode(v int32) *GetContextLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextLogsResponse) SetBody(v *GetContextLogsResponseBody) *GetContextLogsResponse {
	s.Body = v
	return s
}

type GetCursorRequest struct {
	// 时间点（Unix时间戳）或者字符串begin、end。
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	// 这里固定为 cursor。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCursorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCursorRequest) GoString() string {
	return s.String()
}

func (s *GetCursorRequest) SetFrom(v string) *GetCursorRequest {
	s.From = &v
	return s
}

func (s *GetCursorRequest) SetType(v string) *GetCursorRequest {
	s.Type = &v
	return s
}

type GetCursorResponseBody struct {
	// 游标位置。
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s GetCursorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCursorResponseBody) GoString() string {
	return s.String()
}

func (s *GetCursorResponseBody) SetCursor(v string) *GetCursorResponseBody {
	s.Cursor = &v
	return s
}

type GetCursorResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCursorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCursorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCursorResponse) GoString() string {
	return s.String()
}

func (s *GetCursorResponse) SetHeaders(v map[string]*string) *GetCursorResponse {
	s.Headers = v
	return s
}

func (s *GetCursorResponse) SetStatusCode(v int32) *GetCursorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCursorResponse) SetBody(v *GetCursorResponseBody) *GetCursorResponse {
	s.Body = v
	return s
}

type GetCursorTimeRequest struct {
	// 游标。
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// 固定为 cursor_time 。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetCursorTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeRequest) GoString() string {
	return s.String()
}

func (s *GetCursorTimeRequest) SetCursor(v string) *GetCursorTimeRequest {
	s.Cursor = &v
	return s
}

func (s *GetCursorTimeRequest) SetType(v string) *GetCursorTimeRequest {
	s.Type = &v
	return s
}

type GetCursorTimeResponseBody struct {
	// Cursor的服务端时间。Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	CursorTime *string `json:"cursor_time,omitempty" xml:"cursor_time,omitempty"`
}

func (s GetCursorTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCursorTimeResponseBody) SetCursorTime(v string) *GetCursorTimeResponseBody {
	s.CursorTime = &v
	return s
}

type GetCursorTimeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCursorTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCursorTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCursorTimeResponse) GoString() string {
	return s.String()
}

func (s *GetCursorTimeResponse) SetHeaders(v map[string]*string) *GetCursorTimeResponse {
	s.Headers = v
	return s
}

func (s *GetCursorTimeResponse) SetStatusCode(v int32) *GetCursorTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCursorTimeResponse) SetBody(v *GetCursorTimeResponseBody) *GetCursorTimeResponse {
	s.Body = v
	return s
}

type GetEtlMetaRequest struct {
	ElMetaName *string `json:"elMetaName,omitempty" xml:"elMetaName,omitempty"`
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 此处固定为 "__all_etl_meta_tag_match__"。
	EtlMetaTag *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
}

func (s GetEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *GetEtlMetaRequest) SetElMetaName(v string) *GetEtlMetaRequest {
	s.ElMetaName = &v
	return s
}

func (s *GetEtlMetaRequest) SetEtlMetaKey(v string) *GetEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *GetEtlMetaRequest) SetEtlMetaTag(v string) *GetEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

type GetEtlMetaResponseBody struct {
	EtlMetaList []*EtlMeta `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
	Total       *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetEtlMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetEtlMetaResponseBody) SetEtlMetaList(v []*EtlMeta) *GetEtlMetaResponseBody {
	s.EtlMetaList = v
	return s
}

func (s *GetEtlMetaResponseBody) SetTotal(v int32) *GetEtlMetaResponseBody {
	s.Total = &v
	return s
}

type GetEtlMetaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEtlMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *GetEtlMetaResponse) SetHeaders(v map[string]*string) *GetEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *GetEtlMetaResponse) SetStatusCode(v int32) *GetEtlMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEtlMetaResponse) SetBody(v *GetEtlMetaResponseBody) *GetEtlMetaResponse {
	s.Body = v
	return s
}

type GetHistogramsRequest struct {
	// 查询开始时间点。UNIX时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	//
	// 时间区间遵循“左闭右开”原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 查询语句。仅支持查询语句，不支持分析语句。关于查询语句的详细语法，请参见查询语法。
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 查询结束时间点。UNIX时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	//
	// 时间区间遵循“左闭右开”原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// 日志主题。
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// Logstore中数据的类型。该接口中固定取值为histogram。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetHistogramsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsRequest) GoString() string {
	return s.String()
}

func (s *GetHistogramsRequest) SetFrom(v int64) *GetHistogramsRequest {
	s.From = &v
	return s
}

func (s *GetHistogramsRequest) SetQuery(v string) *GetHistogramsRequest {
	s.Query = &v
	return s
}

func (s *GetHistogramsRequest) SetTo(v int64) *GetHistogramsRequest {
	s.To = &v
	return s
}

func (s *GetHistogramsRequest) SetTopic(v string) *GetHistogramsRequest {
	s.Topic = &v
	return s
}

func (s *GetHistogramsRequest) SetType(v string) *GetHistogramsRequest {
	s.Type = &v
	return s
}

type GetHistogramsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*GetHistogramsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetHistogramsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponse) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponse) SetHeaders(v map[string]*string) *GetHistogramsResponse {
	s.Headers = v
	return s
}

func (s *GetHistogramsResponse) SetStatusCode(v int32) *GetHistogramsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistogramsResponse) SetBody(v []*GetHistogramsResponseBody) *GetHistogramsResponse {
	s.Body = v
	return s
}

type GetHistogramsResponseBody struct {
	// 子时间区间的开始时间点。UNIX时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	//
	// 时间区间遵循“左闭右开”原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 子时间区间的结束时间点。UNIX时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	//
	// 时间区间遵循“左闭右开”原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// 该子时间区间内查询到的日志条数。
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 当前查询结果在该子时间区间内的结果是否完整。
	//
	// Complete：查询已经完成，返回结果为完整结果。
	//
	// Incomplete：查询已经完成，返回结果为不完整结果，需要重复请求以获得完整结果。
	Progress *string `json:"progress,omitempty" xml:"progress,omitempty"`
}

func (s GetHistogramsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistogramsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistogramsResponseBody) SetFrom(v int64) *GetHistogramsResponseBody {
	s.From = &v
	return s
}

func (s *GetHistogramsResponseBody) SetTo(v int64) *GetHistogramsResponseBody {
	s.To = &v
	return s
}

func (s *GetHistogramsResponseBody) SetCount(v int64) *GetHistogramsResponseBody {
	s.Count = &v
	return s
}

func (s *GetHistogramsResponseBody) SetProgress(v string) *GetHistogramsResponseBody {
	s.Progress = &v
	return s
}

type GetIndexResponseBody struct {
	// 索引模式
	IndexMode *string `json:"index_mode,omitempty" xml:"index_mode,omitempty"`
	// 字段索引配置。key为字段名称，value为索引配置。
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// 上次修改时间
	LastModifyTime *int64 `json:"lastModifyTime,omitempty" xml:"lastModifyTime,omitempty"`
	// 配置全文索引。
	Line *GetIndexResponseBodyLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// 是否开启日志聚类.
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// 日志聚类的聚类字段过滤黑名单，仅当日志聚类开启时有效。
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// 日志聚类的聚类字段过滤白名单，仅当日志聚类开启时有效。
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// 日志服务默认字段值的最大长度为2048字节，即2 KB。如果您需要修改字段值的最大长度，可设置统计字段（text）最大长度，取值范围为64~16384字节。
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// 存储类型，目前固定取值为pg。
	Storage *string `json:"storage,omitempty" xml:"storage,omitempty"`
	// 索引文件生命周期，支持7天、30天、90天。
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s GetIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBody) SetIndexMode(v string) *GetIndexResponseBody {
	s.IndexMode = &v
	return s
}

func (s *GetIndexResponseBody) SetKeys(v map[string]*KeysValue) *GetIndexResponseBody {
	s.Keys = v
	return s
}

func (s *GetIndexResponseBody) SetLastModifyTime(v int64) *GetIndexResponseBody {
	s.LastModifyTime = &v
	return s
}

func (s *GetIndexResponseBody) SetLine(v *GetIndexResponseBodyLine) *GetIndexResponseBody {
	s.Line = v
	return s
}

func (s *GetIndexResponseBody) SetLogReduce(v bool) *GetIndexResponseBody {
	s.LogReduce = &v
	return s
}

func (s *GetIndexResponseBody) SetLogReduceBlackList(v []*string) *GetIndexResponseBody {
	s.LogReduceBlackList = v
	return s
}

func (s *GetIndexResponseBody) SetLogReduceWhiteList(v []*string) *GetIndexResponseBody {
	s.LogReduceWhiteList = v
	return s
}

func (s *GetIndexResponseBody) SetMaxTextLen(v int32) *GetIndexResponseBody {
	s.MaxTextLen = &v
	return s
}

func (s *GetIndexResponseBody) SetStorage(v string) *GetIndexResponseBody {
	s.Storage = &v
	return s
}

func (s *GetIndexResponseBody) SetTtl(v int32) *GetIndexResponseBody {
	s.Ttl = &v
	return s
}

type GetIndexResponseBodyLine struct {
	// 大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// 是否包含中文。
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// 排除的字段列表。
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// 包含的字段列表。
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// 分词符列表。
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s GetIndexResponseBodyLine) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponseBodyLine) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyLine) SetCaseSensitive(v bool) *GetIndexResponseBodyLine {
	s.CaseSensitive = &v
	return s
}

func (s *GetIndexResponseBodyLine) SetChn(v bool) *GetIndexResponseBodyLine {
	s.Chn = &v
	return s
}

func (s *GetIndexResponseBodyLine) SetExcludeKeys(v []*string) *GetIndexResponseBodyLine {
	s.ExcludeKeys = v
	return s
}

func (s *GetIndexResponseBodyLine) SetIncludeKeys(v []*string) *GetIndexResponseBodyLine {
	s.IncludeKeys = v
	return s
}

func (s *GetIndexResponseBodyLine) SetToken(v []*string) *GetIndexResponseBodyLine {
	s.Token = v
	return s
}

type GetIndexResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexResponse) GoString() string {
	return s.String()
}

func (s *GetIndexResponse) SetHeaders(v map[string]*string) *GetIndexResponse {
	s.Headers = v
	return s
}

func (s *GetIndexResponse) SetStatusCode(v int32) *GetIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexResponse) SetBody(v *GetIndexResponseBody) *GetIndexResponse {
	s.Body = v
	return s
}

type GetLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Logstore          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogStoreResponse) GoString() string {
	return s.String()
}

func (s *GetLogStoreResponse) SetHeaders(v map[string]*string) *GetLogStoreResponse {
	s.Headers = v
	return s
}

func (s *GetLogStoreResponse) SetStatusCode(v int32) *GetLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogStoreResponse) SetBody(v *Logstore) *GetLogStoreResponse {
	s.Body = v
	return s
}

type GetLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Logging           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoggingResponse) GoString() string {
	return s.String()
}

func (s *GetLoggingResponse) SetHeaders(v map[string]*string) *GetLoggingResponse {
	s.Headers = v
	return s
}

func (s *GetLoggingResponse) SetStatusCode(v int32) *GetLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoggingResponse) SetBody(v *Logging) *GetLoggingResponse {
	s.Body = v
	return s
}

type GetLogsRequest struct {
	// 查询开始时间点。该时间是指写入日志数据时指定的日志时间。
	//
	// 请求参数from和to定义的时间区间遵循左闭右开原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	// Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 仅当query参数为查询语句时，该参数有效，表示请求返回的最大日志条数。最小值为0，最大值为100，默认值为100。
	Line *int64 `json:"line,omitempty" xml:"line,omitempty"`
	// 仅当query参数为查询语句时，该参数有效，表示查询开始行。默认值为0。
	Offset *int64 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 用于指定返回结果是否按日志时间戳降序返回日志，精确到分钟级别。
	//
	// true：按照日志时间戳降序返回日志。
	// false（默认值）：按照日志时间戳升序返回日志。
	// 注意
	// 当query参数为查询语句时，参数reverse有效，用于指定返回日志排序方式。
	// 当query参数为查询和分析语句时，参数reverse无效，由SQL分析语句中order by语法指定排序方式。如果order by为asc（默认），则为升序；如果order by为desc，则为降序。
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// 查询语句或者分析语句。更多信息，请参见查询概述和分析概述。
	//
	// 在query参数的分析语句中加上set session parallel_sql=true;，表示使用SQL独享版。例如* | set session parallel_sql=true; select count(*) as pv 。
	//
	// 说明 当query参数中有分析语句（SQL语句）时，该接口的line参数和offset参数无效，建议设置为0，需通过SQL语句的LIMIT语法实现翻页。更多信息，请参见分页显示查询分析结果。
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 用于指定返回结果是否按日志时间戳降序返回日志，精确到分钟级别。
	//
	// true：按照日志时间戳降序返回日志。
	// false（默认值）：按照日志时间戳升序返回日志。
	// 注意
	// 当query参数为查询语句时，参数reverse有效，用于指定返回日志排序方式。
	// 当query参数为查询和分析语句时，参数reverse无效，由SQL分析语句中order by语法指定排序方式。如果order by为asc（默认），则为升序；如果order by为desc，则为降序。
	Reverse *bool `json:"reverse,omitempty" xml:"reverse,omitempty"`
	// 查询结束时间点。该时间是指写入日志数据时指定的日志时间。
	//
	// 请求参数from和to定义的时间区间遵循左闭右开原则，即该时间区间包括区间开始时间点，但不包括区间结束时间点。如果from和to的值相同，则为无效区间，函数直接返回错误。
	// Unix时间戳格式，表示从1970-1-1 00:00:00 UTC计算起的秒数。
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
	// status: 401 | SELECT remote_addr,COUNT(*) as pv GROUP by remote_addr ORDER by pv desc limit 5
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// 查询Logstore数据的类型。在该接口中固定取值为log。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLogsRequest) GoString() string {
	return s.String()
}

func (s *GetLogsRequest) SetFrom(v int64) *GetLogsRequest {
	s.From = &v
	return s
}

func (s *GetLogsRequest) SetLine(v int64) *GetLogsRequest {
	s.Line = &v
	return s
}

func (s *GetLogsRequest) SetOffset(v int64) *GetLogsRequest {
	s.Offset = &v
	return s
}

func (s *GetLogsRequest) SetPowerSql(v bool) *GetLogsRequest {
	s.PowerSql = &v
	return s
}

func (s *GetLogsRequest) SetQuery(v string) *GetLogsRequest {
	s.Query = &v
	return s
}

func (s *GetLogsRequest) SetReverse(v bool) *GetLogsRequest {
	s.Reverse = &v
	return s
}

func (s *GetLogsRequest) SetTo(v int64) *GetLogsRequest {
	s.To = &v
	return s
}

func (s *GetLogsRequest) SetTopic(v string) *GetLogsRequest {
	s.Topic = &v
	return s
}

func (s *GetLogsRequest) SetType(v string) *GetLogsRequest {
	s.Type = &v
	return s
}

type GetLogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLogsResponse) GoString() string {
	return s.String()
}

func (s *GetLogsResponse) SetHeaders(v map[string]*string) *GetLogsResponse {
	s.Headers = v
	return s
}

func (s *GetLogsResponse) SetStatusCode(v int32) *GetLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogsResponse) SetBody(v []map[string]interface{}) *GetLogsResponse {
	s.Body = v
	return s
}

type GetMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MachineGroup      `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *GetMachineGroupResponse) SetHeaders(v map[string]*string) *GetMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *GetMachineGroupResponse) SetStatusCode(v int32) *GetMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMachineGroupResponse) SetBody(v *MachineGroup) *GetMachineGroupResponse {
	s.Body = v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Project           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *Project) *GetProjectResponse {
	s.Body = v
	return s
}

type GetProjectLogsRequest struct {
	// 是否使用SQL独享版。更多信息，请参见开启SQL独享版。
	//
	// true：使用SQL独享版。
	// false（默认值）：使用SQL普通版。
	// 除通过powerSql参数配置SQL独享版外，您还可以使用query参数。
	PowerSql *bool `json:"powerSql,omitempty" xml:"powerSql,omitempty"`
	// 标准SQL语句。例如日志库名称为nginx-moni，查询时间区间在2022-03-01 10:41:40到2022-03-01 10:56:40之间的访问数量。
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetProjectLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectLogsRequest) GoString() string {
	return s.String()
}

func (s *GetProjectLogsRequest) SetPowerSql(v bool) *GetProjectLogsRequest {
	s.PowerSql = &v
	return s
}

func (s *GetProjectLogsRequest) SetQuery(v string) *GetProjectLogsRequest {
	s.Query = &v
	return s
}

type GetProjectLogsResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []map[string]*string `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s GetProjectLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectLogsResponse) GoString() string {
	return s.String()
}

func (s *GetProjectLogsResponse) SetHeaders(v map[string]*string) *GetProjectLogsResponse {
	s.Headers = v
	return s
}

func (s *GetProjectLogsResponse) SetStatusCode(v int32) *GetProjectLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectLogsResponse) SetBody(v []map[string]*string) *GetProjectLogsResponse {
	s.Body = v
	return s
}

type GetSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SavedSearch       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *GetSavedSearchResponse) SetHeaders(v map[string]*string) *GetSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *GetSavedSearchResponse) SetStatusCode(v int32) *GetSavedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavedSearchResponse) SetBody(v *SavedSearch) *GetSavedSearchResponse {
	s.Body = v
	return s
}

type ListConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*ConsumerGroup   `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupResponse) SetHeaders(v map[string]*string) *ListConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupResponse) SetStatusCode(v int32) *ListConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupResponse) SetBody(v []*ConsumerGroup) *ListConsumerGroupResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// 用于搜索匹配的自定义域名
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	Offset     *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	Size       *int32  `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetDomainName(v string) *ListDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *ListDomainsRequest) SetOffset(v int32) *ListDomainsRequest {
	s.Offset = &v
	return s
}

func (s *ListDomainsRequest) SetSize(v int32) *ListDomainsRequest {
	s.Size = &v
	return s
}

type ListDomainsResponseBody struct {
	Count   *int64    `json:"count,omitempty" xml:"count,omitempty"`
	Domains []*string `json:"domains,omitempty" xml:"domains,omitempty" type:"Repeated"`
	Total   *int64    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetCount(v int64) *ListDomainsResponseBody {
	s.Count = &v
	return s
}

func (s *ListDomainsResponseBody) SetDomains(v []*string) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetTotal(v int64) *ListDomainsResponseBody {
	s.Total = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListEtlMetaRequest struct {
	EtlMetaKey  *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	EtlMetaTag  *string `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	// 默认值 0。
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 默认值 200.
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *ListEtlMetaRequest) SetEtlMetaKey(v string) *ListEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *ListEtlMetaRequest) SetEtlMetaName(v string) *ListEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *ListEtlMetaRequest) SetEtlMetaTag(v string) *ListEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *ListEtlMetaRequest) SetOffset(v int32) *ListEtlMetaRequest {
	s.Offset = &v
	return s
}

func (s *ListEtlMetaRequest) SetSize(v int32) *ListEtlMetaRequest {
	s.Size = &v
	return s
}

type ListEtlMetaResponseBody struct {
	EtlMetaList []*EtlMeta `json:"etlMetaList,omitempty" xml:"etlMetaList,omitempty" type:"Repeated"`
	Total       *int32     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEtlMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListEtlMetaResponseBody) SetEtlMetaList(v []*EtlMeta) *ListEtlMetaResponseBody {
	s.EtlMetaList = v
	return s
}

func (s *ListEtlMetaResponseBody) SetTotal(v int32) *ListEtlMetaResponseBody {
	s.Total = &v
	return s
}

type ListEtlMetaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEtlMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *ListEtlMetaResponse) SetHeaders(v map[string]*string) *ListEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *ListEtlMetaResponse) SetStatusCode(v int32) *ListEtlMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEtlMetaResponse) SetBody(v *ListEtlMetaResponseBody) *ListEtlMetaResponse {
	s.Body = v
	return s
}

type ListEtlMetaNameRequest struct {
	// 默认值为 0。
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 默认值 200。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListEtlMetaNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameRequest) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameRequest) SetOffset(v int32) *ListEtlMetaNameRequest {
	s.Offset = &v
	return s
}

func (s *ListEtlMetaNameRequest) SetSize(v int32) *ListEtlMetaNameRequest {
	s.Size = &v
	return s
}

type ListEtlMetaNameResponseBody struct {
	Count           *int32    `json:"count,omitempty" xml:"count,omitempty"`
	EtlMetaNameList []*string `json:"etlMetaNameList,omitempty" xml:"etlMetaNameList,omitempty" type:"Repeated"`
	Total           *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEtlMetaNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameResponseBody) SetCount(v int32) *ListEtlMetaNameResponseBody {
	s.Count = &v
	return s
}

func (s *ListEtlMetaNameResponseBody) SetEtlMetaNameList(v []*string) *ListEtlMetaNameResponseBody {
	s.EtlMetaNameList = v
	return s
}

func (s *ListEtlMetaNameResponseBody) SetTotal(v int32) *ListEtlMetaNameResponseBody {
	s.Total = &v
	return s
}

type ListEtlMetaNameResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEtlMetaNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEtlMetaNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEtlMetaNameResponse) GoString() string {
	return s.String()
}

func (s *ListEtlMetaNameResponse) SetHeaders(v map[string]*string) *ListEtlMetaNameResponse {
	s.Headers = v
	return s
}

func (s *ListEtlMetaNameResponse) SetStatusCode(v int32) *ListEtlMetaNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEtlMetaNameResponse) SetBody(v *ListEtlMetaNameResponseBody) *ListEtlMetaNameResponse {
	s.Body = v
	return s
}

type ListLogStoresRequest struct {
	LogstoreName *string `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	Offset       *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	// 默认值为 500。
	Size          *int32  `json:"size,omitempty" xml:"size,omitempty"`
	TelemetryType *string `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
}

func (s ListLogStoresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresRequest) GoString() string {
	return s.String()
}

func (s *ListLogStoresRequest) SetLogstoreName(v string) *ListLogStoresRequest {
	s.LogstoreName = &v
	return s
}

func (s *ListLogStoresRequest) SetOffset(v int32) *ListLogStoresRequest {
	s.Offset = &v
	return s
}

func (s *ListLogStoresRequest) SetSize(v int32) *ListLogStoresRequest {
	s.Size = &v
	return s
}

func (s *ListLogStoresRequest) SetTelemetryType(v string) *ListLogStoresRequest {
	s.TelemetryType = &v
	return s
}

type ListLogStoresResponseBody struct {
	Logstores []*string `json:"logstores,omitempty" xml:"logstores,omitempty" type:"Repeated"`
	Total     *int32    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListLogStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponseBody) SetLogstores(v []*string) *ListLogStoresResponseBody {
	s.Logstores = v
	return s
}

func (s *ListLogStoresResponseBody) SetTotal(v int32) *ListLogStoresResponseBody {
	s.Total = &v
	return s
}

type ListLogStoresResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLogStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLogStoresResponse) GoString() string {
	return s.String()
}

func (s *ListLogStoresResponse) SetHeaders(v map[string]*string) *ListLogStoresResponse {
	s.Headers = v
	return s
}

func (s *ListLogStoresResponse) SetStatusCode(v int32) *ListLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogStoresResponse) SetBody(v *ListLogStoresResponseBody) *ListLogStoresResponse {
	s.Body = v
	return s
}

type ListMachineGroupRequest struct {
	// 可将 groupName 作为 pattern 匹配名称，只会返回匹配的机器组。例如 test 可以匹配机器组 test-group。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 分页请求的起始位置。默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 分页查询时，设置的每页行数。默认值为2000。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *ListMachineGroupRequest) SetGroupName(v string) *ListMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListMachineGroupRequest) SetOffset(v int32) *ListMachineGroupRequest {
	s.Offset = &v
	return s
}

func (s *ListMachineGroupRequest) SetSize(v int32) *ListMachineGroupRequest {
	s.Size = &v
	return s
}

type ListMachineGroupResponseBody struct {
	// 当前页返回的机器组数量。
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// 机器组名称列表。
	Machinegroups []*string `json:"machinegroups,omitempty" xml:"machinegroups,omitempty" type:"Repeated"`
	// 机器组总数量。
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMachineGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineGroupResponseBody) SetCount(v int32) *ListMachineGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListMachineGroupResponseBody) SetMachinegroups(v []*string) *ListMachineGroupResponseBody {
	s.Machinegroups = v
	return s
}

func (s *ListMachineGroupResponseBody) SetTotal(v int32) *ListMachineGroupResponseBody {
	s.Total = &v
	return s
}

type ListMachineGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMachineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *ListMachineGroupResponse) SetHeaders(v map[string]*string) *ListMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *ListMachineGroupResponse) SetStatusCode(v int32) *ListMachineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineGroupResponse) SetBody(v *ListMachineGroupResponseBody) *ListMachineGroupResponse {
	s.Body = v
	return s
}

type ListMachinesRequest struct {
	// 查询开始行。默认值为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 分页查询时，设置的每页行数。默认值为2000。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListMachinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesRequest) GoString() string {
	return s.String()
}

func (s *ListMachinesRequest) SetOffset(v int32) *ListMachinesRequest {
	s.Offset = &v
	return s
}

func (s *ListMachinesRequest) SetSize(v int32) *ListMachinesRequest {
	s.Size = &v
	return s
}

type ListMachinesResponseBody struct {
	// 当前页返回的机器数目。
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// 返回的机器信息列表。
	Machines []*Machine `json:"machines,omitempty" xml:"machines,omitempty" type:"Repeated"`
	// 机器总数。
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMachinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachinesResponseBody) SetCount(v int32) *ListMachinesResponseBody {
	s.Count = &v
	return s
}

func (s *ListMachinesResponseBody) SetMachines(v []*Machine) *ListMachinesResponseBody {
	s.Machines = v
	return s
}

func (s *ListMachinesResponseBody) SetTotal(v int32) *ListMachinesResponseBody {
	s.Total = &v
	return s
}

type ListMachinesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMachinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachinesResponse) GoString() string {
	return s.String()
}

func (s *ListMachinesResponse) SetHeaders(v map[string]*string) *ListMachinesResponse {
	s.Headers = v
	return s
}

func (s *ListMachinesResponse) SetStatusCode(v int32) *ListMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachinesResponse) SetBody(v *ListMachinesResponseBody) *ListMachinesResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	Offset      *int32  `json:"offset,omitempty" xml:"offset,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// 默认值为 500。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) SetOffset(v int32) *ListProjectRequest {
	s.Offset = &v
	return s
}

func (s *ListProjectRequest) SetProjectName(v string) *ListProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *ListProjectRequest) SetSize(v int32) *ListProjectRequest {
	s.Size = &v
	return s
}

type ListProjectResponseBody struct {
	Count    *int64     `json:"count,omitempty" xml:"count,omitempty"`
	Projects []*Project `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	Total    *int64     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetProjects(v []*Project) *ListProjectResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectResponseBody) SetTotal(v int64) *ListProjectResponseBody {
	s.Total = &v
	return s
}

type ListProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponse) GoString() string {
	return s.String()
}

func (s *ListProjectResponse) SetHeaders(v map[string]*string) *ListProjectResponse {
	s.Headers = v
	return s
}

func (s *ListProjectResponse) SetStatusCode(v int32) *ListProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

type ListSavedSearchRequest struct {
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// 默认值为 500。
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *ListSavedSearchRequest) SetOffset(v int32) *ListSavedSearchRequest {
	s.Offset = &v
	return s
}

func (s *ListSavedSearchRequest) SetSize(v int32) *ListSavedSearchRequest {
	s.Size = &v
	return s
}

type ListSavedSearchResponseBody struct {
	Count            *int32         `json:"count,omitempty" xml:"count,omitempty"`
	SavedsearchItems []*SavedSearch `json:"savedsearchItems,omitempty" xml:"savedsearchItems,omitempty" type:"Repeated"`
	Total            *int32         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSavedSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavedSearchResponseBody) SetCount(v int32) *ListSavedSearchResponseBody {
	s.Count = &v
	return s
}

func (s *ListSavedSearchResponseBody) SetSavedsearchItems(v []*SavedSearch) *ListSavedSearchResponseBody {
	s.SavedsearchItems = v
	return s
}

func (s *ListSavedSearchResponseBody) SetTotal(v int32) *ListSavedSearchResponseBody {
	s.Total = &v
	return s
}

type ListSavedSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSavedSearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *ListSavedSearchResponse) SetHeaders(v map[string]*string) *ListSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *ListSavedSearchResponse) SetStatusCode(v int32) *ListSavedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSavedSearchResponse) SetBody(v *ListSavedSearchResponseBody) *ListSavedSearchResponse {
	s.Body = v
	return s
}

type ListShardsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s ListShardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShardsResponse) GoString() string {
	return s.String()
}

func (s *ListShardsResponse) SetHeaders(v map[string]*string) *ListShardsResponse {
	s.Headers = v
	return s
}

func (s *ListShardsResponse) SetStatusCode(v int32) *ListShardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShardsResponse) SetBody(v []*Shard) *ListShardsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// 查询的资源的 id 列表。resource id 与 tags 应至少存在一个。
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// 资源类型。目前取值范围：project。
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 精确查找时过滤的标签键值对。resource id 与 tags 应至少存在一个。
	Tags []*ListTagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

type ListTagResourcesRequestTags struct {
	// 精确过滤的标签的键。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 精确过滤的标签的值。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	// 查询的资源的 id 列表。resource id 与 tags 应至少存在一个。
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 资源类型。目前取值范围：project。
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 精确查找时过滤的标签键值对。resource id 与 tags 应至少存在一个。
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// 下一个查询开始Token。
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 返回的标签列表。
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"tagResources,omitempty" xml:"tagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// 资源 id。
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// 资源类型。
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 标签的键。
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// 标签的值。
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type MergeShardsRequest struct {
	// 固定为 merge。
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s MergeShardsRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeShardsRequest) GoString() string {
	return s.String()
}

func (s *MergeShardsRequest) SetAction(v string) *MergeShardsRequest {
	s.Action = &v
	return s
}

type MergeShardsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s MergeShardsResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeShardsResponse) GoString() string {
	return s.String()
}

func (s *MergeShardsResponse) SetHeaders(v map[string]*string) *MergeShardsResponse {
	s.Headers = v
	return s
}

func (s *MergeShardsResponse) SetStatusCode(v int32) *MergeShardsResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeShardsResponse) SetBody(v []*Shard) *MergeShardsResponse {
	s.Body = v
	return s
}

type RemoveConfigFromMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s RemoveConfigFromMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveConfigFromMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveConfigFromMachineGroupResponse) SetHeaders(v map[string]*string) *RemoveConfigFromMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveConfigFromMachineGroupResponse) SetStatusCode(v int32) *RemoveConfigFromMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type SplitShardRequest struct {
	// 这里固定为 split。
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 分裂的位置。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 要分裂成的 shard 数量，默认为 2。
	ShardCount *int32 `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
}

func (s SplitShardRequest) String() string {
	return tea.Prettify(s)
}

func (s SplitShardRequest) GoString() string {
	return s.String()
}

func (s *SplitShardRequest) SetAction(v string) *SplitShardRequest {
	s.Action = &v
	return s
}

func (s *SplitShardRequest) SetKey(v string) *SplitShardRequest {
	s.Key = &v
	return s
}

func (s *SplitShardRequest) SetShardCount(v int32) *SplitShardRequest {
	s.ShardCount = &v
	return s
}

type SplitShardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*Shard           `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
}

func (s SplitShardResponse) String() string {
	return tea.Prettify(s)
}

func (s SplitShardResponse) GoString() string {
	return s.String()
}

func (s *SplitShardResponse) SetHeaders(v map[string]*string) *SplitShardResponse {
	s.Headers = v
	return s
}

func (s *SplitShardResponse) SetStatusCode(v int32) *SplitShardResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitShardResponse) SetBody(v []*Shard) *SplitShardResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// 资源的 id 列表，可以一次为多个同类型资源打上相同的标签。
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// 资源的类型。目前取值范围：project。
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 标签列表。
	Tags []*TagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	// 标签的 key。
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 标签的 value。
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UnTagResourcesRequest struct {
	// 是否删除所有标签，默认为 false，表示仅删除 tags 列表中的标签项。值为 true 时删除资源上绑定的所有标签。
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// 资源的 id 列表，可以一次为多个同类型资源删除相同的标签。当 all 为 false 时生效。
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// 资源的类型。目前取值范围 ： project。
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// 标签 key 列表。当 all 为 false 时，仅删除列表中的标签。
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTags(v []*string) *UnTagResourcesRequest {
	s.Tags = v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

type UpdateCheckPointRequest struct {
	// checkpoint值。
	Checkpoint *string `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// shard 的 id。
	Shard *int32 `json:"shard,omitempty" xml:"shard,omitempty"`
	// 消费者。
	Consumer *string `json:"consumer,omitempty" xml:"consumer,omitempty"`
	// 当不指定消费者时，必须指定forceSuccess为true才能更新checkpoint。
	ForceSuccess *bool `json:"forceSuccess,omitempty" xml:"forceSuccess,omitempty"`
	// 固定为 checkpoint。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateCheckPointRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckPointRequest) GoString() string {
	return s.String()
}

func (s *UpdateCheckPointRequest) SetCheckpoint(v string) *UpdateCheckPointRequest {
	s.Checkpoint = &v
	return s
}

func (s *UpdateCheckPointRequest) SetShard(v int32) *UpdateCheckPointRequest {
	s.Shard = &v
	return s
}

func (s *UpdateCheckPointRequest) SetConsumer(v string) *UpdateCheckPointRequest {
	s.Consumer = &v
	return s
}

func (s *UpdateCheckPointRequest) SetForceSuccess(v bool) *UpdateCheckPointRequest {
	s.ForceSuccess = &v
	return s
}

func (s *UpdateCheckPointRequest) SetType(v string) *UpdateCheckPointRequest {
	s.Type = &v
	return s
}

type UpdateCheckPointResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateCheckPointResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCheckPointResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckPointResponse) SetHeaders(v map[string]*string) *UpdateCheckPointResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckPointResponse) SetStatusCode(v int32) *UpdateCheckPointResponse {
	s.StatusCode = &v
	return s
}

type UpdateConsumerGroupRequest struct {
	Order   *bool  `json:"order,omitempty" xml:"order,omitempty"`
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) SetOrder(v bool) *UpdateConsumerGroupRequest {
	s.Order = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetTimeout(v int32) *UpdateConsumerGroupRequest {
	s.Timeout = &v
	return s
}

type UpdateConsumerGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponse) SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerGroupResponse) SetStatusCode(v int32) *UpdateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

type UpdateEtlMetaRequest struct {
	// 是否启用。etlMetaTag、etlMetaValue、enable 至少需要存在一个。
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,255]之间。
	EtlMetaKey *string `json:"etlMetaKey,omitempty" xml:"etlMetaKey,omitempty"`
	// 名字。由数字、大小写字母、下划线_、连字符-组成，长度需要在[2,64]之间。
	EtlMetaName *string `json:"etlMetaName,omitempty" xml:"etlMetaName,omitempty"`
	// key。由 ascii 可打印字符组成，包括数字、英文大小写字母、下划线、连字符、英文标点符号等组成，长度在[1,128]之间。
	EtlMetaTag   *string                `json:"etlMetaTag,omitempty" xml:"etlMetaTag,omitempty"`
	EtlMetaValue map[string]interface{} `json:"etlMetaValue,omitempty" xml:"etlMetaValue,omitempty"`
}

func (s UpdateEtlMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlMetaRequest) GoString() string {
	return s.String()
}

func (s *UpdateEtlMetaRequest) SetEnable(v bool) *UpdateEtlMetaRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaKey(v string) *UpdateEtlMetaRequest {
	s.EtlMetaKey = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaName(v string) *UpdateEtlMetaRequest {
	s.EtlMetaName = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaTag(v string) *UpdateEtlMetaRequest {
	s.EtlMetaTag = &v
	return s
}

func (s *UpdateEtlMetaRequest) SetEtlMetaValue(v map[string]interface{}) *UpdateEtlMetaRequest {
	s.EtlMetaValue = v
	return s
}

type UpdateEtlMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateEtlMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEtlMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateEtlMetaResponse) SetHeaders(v map[string]*string) *UpdateEtlMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateEtlMetaResponse) SetStatusCode(v int32) *UpdateEtlMetaResponse {
	s.StatusCode = &v
	return s
}

type UpdateIndexRequest struct {
	// 字段索引配置，key为字段名称，value为字段索引配置。
	Keys map[string]*KeysValue `json:"keys,omitempty" xml:"keys,omitempty"`
	// 配置全文索引。
	Line *UpdateIndexRequestLine `json:"line,omitempty" xml:"line,omitempty" type:"Struct"`
	// 开启日志聚类，开启后白名单与黑名单至多生效其中一个。
	LogReduce *bool `json:"log_reduce,omitempty" xml:"log_reduce,omitempty"`
	// 日志聚类的聚类字段黑名单
	LogReduceBlackList []*string `json:"log_reduce_black_list,omitempty" xml:"log_reduce_black_list,omitempty" type:"Repeated"`
	// 日志聚类的聚类字段白名单
	LogReduceWhiteList []*string `json:"log_reduce_white_list,omitempty" xml:"log_reduce_white_list,omitempty" type:"Repeated"`
	// 统计字段的最大长度
	MaxTextLen *int32 `json:"max_text_len,omitempty" xml:"max_text_len,omitempty"`
	// 保存时间，单位为天
	Ttl *int32 `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s UpdateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexRequest) GoString() string {
	return s.String()
}

func (s *UpdateIndexRequest) SetKeys(v map[string]*KeysValue) *UpdateIndexRequest {
	s.Keys = v
	return s
}

func (s *UpdateIndexRequest) SetLine(v *UpdateIndexRequestLine) *UpdateIndexRequest {
	s.Line = v
	return s
}

func (s *UpdateIndexRequest) SetLogReduce(v bool) *UpdateIndexRequest {
	s.LogReduce = &v
	return s
}

func (s *UpdateIndexRequest) SetLogReduceBlackList(v []*string) *UpdateIndexRequest {
	s.LogReduceBlackList = v
	return s
}

func (s *UpdateIndexRequest) SetLogReduceWhiteList(v []*string) *UpdateIndexRequest {
	s.LogReduceWhiteList = v
	return s
}

func (s *UpdateIndexRequest) SetMaxTextLen(v int32) *UpdateIndexRequest {
	s.MaxTextLen = &v
	return s
}

func (s *UpdateIndexRequest) SetTtl(v int32) *UpdateIndexRequest {
	s.Ttl = &v
	return s
}

type UpdateIndexRequestLine struct {
	// 大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// 包含中文
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// 排除的字段列表，不能与include_keys同时指定。
	ExcludeKeys []*string `json:"exclude_keys,omitempty" xml:"exclude_keys,omitempty" type:"Repeated"`
	// 包含的字段列表，不能与exclude_keys同时指定。
	IncludeKeys []*string `json:"include_keys,omitempty" xml:"include_keys,omitempty" type:"Repeated"`
	// 分词符列表。可以设置一个分词参数，指定这个字段按照哪一种方式分词。
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
}

func (s UpdateIndexRequestLine) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexRequestLine) GoString() string {
	return s.String()
}

func (s *UpdateIndexRequestLine) SetCaseSensitive(v bool) *UpdateIndexRequestLine {
	s.CaseSensitive = &v
	return s
}

func (s *UpdateIndexRequestLine) SetChn(v bool) *UpdateIndexRequestLine {
	s.Chn = &v
	return s
}

func (s *UpdateIndexRequestLine) SetExcludeKeys(v []*string) *UpdateIndexRequestLine {
	s.ExcludeKeys = v
	return s
}

func (s *UpdateIndexRequestLine) SetIncludeKeys(v []*string) *UpdateIndexRequestLine {
	s.IncludeKeys = v
	return s
}

func (s *UpdateIndexRequestLine) SetToken(v []*string) *UpdateIndexRequestLine {
	s.Token = v
	return s
}

type UpdateIndexResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIndexResponse) GoString() string {
	return s.String()
}

func (s *UpdateIndexResponse) SetHeaders(v map[string]*string) *UpdateIndexResponse {
	s.Headers = v
	return s
}

func (s *UpdateIndexResponse) SetStatusCode(v int32) *UpdateIndexResponse {
	s.StatusCode = &v
	return s
}

type UpdateLogStoreRequest struct {
	AppendMeta     *bool        `json:"appendMeta,omitempty" xml:"appendMeta,omitempty"`
	AutoSplit      *bool        `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	EnableTracking *bool        `json:"enable_tracking,omitempty" xml:"enable_tracking,omitempty"`
	EncryptConf    *EncryptConf `json:"encrypt_conf,omitempty" xml:"encrypt_conf,omitempty"`
	HotTtl         *int32       `json:"hot_ttl,omitempty" xml:"hot_ttl,omitempty"`
	LogstoreName   *string      `json:"logstoreName,omitempty" xml:"logstoreName,omitempty"`
	MaxSplitShard  *int32       `json:"maxSplitShard,omitempty" xml:"maxSplitShard,omitempty"`
	ShardCount     *int32       `json:"shardCount,omitempty" xml:"shardCount,omitempty"`
	TelemetryType  *string      `json:"telemetryType,omitempty" xml:"telemetryType,omitempty"`
	Ttl            *int32       `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s UpdateLogStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreRequest) SetAppendMeta(v bool) *UpdateLogStoreRequest {
	s.AppendMeta = &v
	return s
}

func (s *UpdateLogStoreRequest) SetAutoSplit(v bool) *UpdateLogStoreRequest {
	s.AutoSplit = &v
	return s
}

func (s *UpdateLogStoreRequest) SetEnableTracking(v bool) *UpdateLogStoreRequest {
	s.EnableTracking = &v
	return s
}

func (s *UpdateLogStoreRequest) SetEncryptConf(v *EncryptConf) *UpdateLogStoreRequest {
	s.EncryptConf = v
	return s
}

func (s *UpdateLogStoreRequest) SetHotTtl(v int32) *UpdateLogStoreRequest {
	s.HotTtl = &v
	return s
}

func (s *UpdateLogStoreRequest) SetLogstoreName(v string) *UpdateLogStoreRequest {
	s.LogstoreName = &v
	return s
}

func (s *UpdateLogStoreRequest) SetMaxSplitShard(v int32) *UpdateLogStoreRequest {
	s.MaxSplitShard = &v
	return s
}

func (s *UpdateLogStoreRequest) SetShardCount(v int32) *UpdateLogStoreRequest {
	s.ShardCount = &v
	return s
}

func (s *UpdateLogStoreRequest) SetTelemetryType(v string) *UpdateLogStoreRequest {
	s.TelemetryType = &v
	return s
}

func (s *UpdateLogStoreRequest) SetTtl(v int32) *UpdateLogStoreRequest {
	s.Ttl = &v
	return s
}

type UpdateLogStoreResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateLogStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreResponse) SetHeaders(v map[string]*string) *UpdateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogStoreResponse) SetStatusCode(v int32) *UpdateLogStoreResponse {
	s.StatusCode = &v
	return s
}

type UpdateLoggingRequest struct {
	// 服务日志配置列表。
	LoggingDetails []*UpdateLoggingRequestLoggingDetails `json:"loggingDetails,omitempty" xml:"loggingDetails,omitempty" type:"Repeated"`
	// 服务日志要保存到的 project 名称。
	LoggingProject *string `json:"loggingProject,omitempty" xml:"loggingProject,omitempty"`
}

func (s UpdateLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoggingRequest) SetLoggingDetails(v []*UpdateLoggingRequestLoggingDetails) *UpdateLoggingRequest {
	s.LoggingDetails = v
	return s
}

func (s *UpdateLoggingRequest) SetLoggingProject(v string) *UpdateLoggingRequest {
	s.LoggingProject = &v
	return s
}

type UpdateLoggingRequestLoggingDetails struct {
	// 该种类服务日志要保存到的 logstore 名称。
	Logstore *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	// 服务日志的种类。可选 "consumergroup_log"、 "logtail_alarm"、"operation_log"、"logtail_profile"、"metering"、"logtail_status"、"scheduled_sql_alert"、 "etl_alert" 等。
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateLoggingRequestLoggingDetails) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingRequestLoggingDetails) GoString() string {
	return s.String()
}

func (s *UpdateLoggingRequestLoggingDetails) SetLogstore(v string) *UpdateLoggingRequestLoggingDetails {
	s.Logstore = &v
	return s
}

func (s *UpdateLoggingRequestLoggingDetails) SetType(v string) *UpdateLoggingRequestLoggingDetails {
	s.Type = &v
	return s
}

type UpdateLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoggingResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoggingResponse) SetHeaders(v map[string]*string) *UpdateLoggingResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoggingResponse) SetStatusCode(v int32) *UpdateLoggingResponse {
	s.StatusCode = &v
	return s
}

type UpdateMachineGroupRequest struct {
	// 机器组属性。
	GroupAttribute *UpdateMachineGroupRequestGroupAttribute `json:"groupAttribute,omitempty" xml:"groupAttribute,omitempty" type:"Struct"`
	// 机器组名称。
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// 机器组类型，可选值，默认为空。
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// 机器组标识种类，支持 ip 、userdefined 两种。
	MachineIdentifyType *string `json:"machineIdentifyType,omitempty" xml:"machineIdentifyType,omitempty"`
	// 机器列表。
	MachineList []*string `json:"machineList,omitempty" xml:"machineList,omitempty" type:"Repeated"`
}

func (s UpdateMachineGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupRequest) SetGroupAttribute(v *UpdateMachineGroupRequestGroupAttribute) *UpdateMachineGroupRequest {
	s.GroupAttribute = v
	return s
}

func (s *UpdateMachineGroupRequest) SetGroupName(v string) *UpdateMachineGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetGroupType(v string) *UpdateMachineGroupRequest {
	s.GroupType = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetMachineIdentifyType(v string) *UpdateMachineGroupRequest {
	s.MachineIdentifyType = &v
	return s
}

func (s *UpdateMachineGroupRequest) SetMachineList(v []*string) *UpdateMachineGroupRequest {
	s.MachineList = v
	return s
}

type UpdateMachineGroupRequestGroupAttribute struct {
	// 机器组所依赖的外部管理系统标识。
	ExternalName *string `json:"externalName,omitempty" xml:"externalName,omitempty"`
	// 机器组的日志主题。
	GroupTopic *string `json:"groupTopic,omitempty" xml:"groupTopic,omitempty"`
}

func (s UpdateMachineGroupRequestGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupRequestGroupAttribute) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupRequestGroupAttribute) SetExternalName(v string) *UpdateMachineGroupRequestGroupAttribute {
	s.ExternalName = &v
	return s
}

func (s *UpdateMachineGroupRequestGroupAttribute) SetGroupTopic(v string) *UpdateMachineGroupRequestGroupAttribute {
	s.GroupTopic = &v
	return s
}

type UpdateMachineGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateMachineGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMachineGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateMachineGroupResponse) SetHeaders(v map[string]*string) *UpdateMachineGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateMachineGroupResponse) SetStatusCode(v int32) *UpdateMachineGroupResponse {
	s.StatusCode = &v
	return s
}

type UpdateProjectRequest struct {
	// Project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

type UpdateSavedSearchRequest struct {
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	Logstore        *string `json:"logstore,omitempty" xml:"logstore,omitempty"`
	SavedsearchName *string `json:"savedsearchName,omitempty" xml:"savedsearchName,omitempty"`
	SearchQuery     *string `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	Topic           *string `json:"topic,omitempty" xml:"topic,omitempty"`
}

func (s UpdateSavedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedSearchRequest) GoString() string {
	return s.String()
}

func (s *UpdateSavedSearchRequest) SetDisplayName(v string) *UpdateSavedSearchRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetLogstore(v string) *UpdateSavedSearchRequest {
	s.Logstore = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetSavedsearchName(v string) *UpdateSavedSearchRequest {
	s.SavedsearchName = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetSearchQuery(v string) *UpdateSavedSearchRequest {
	s.SearchQuery = &v
	return s
}

func (s *UpdateSavedSearchRequest) SetTopic(v string) *UpdateSavedSearchRequest {
	s.Topic = &v
	return s
}

type UpdateSavedSearchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UpdateSavedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSavedSearchResponse) GoString() string {
	return s.String()
}

func (s *UpdateSavedSearchResponse) SetHeaders(v map[string]*string) *UpdateSavedSearchResponse {
	s.Headers = v
	return s
}

func (s *UpdateSavedSearchResponse) SetStatusCode(v int32) *UpdateSavedSearchResponse {
	s.StatusCode = &v
	return s
}

type KeysValue struct {
	// 大小写敏感
	CaseSensitive *bool `json:"caseSensitive,omitempty" xml:"caseSensitive,omitempty"`
	// 包含中文
	Chn *bool `json:"chn,omitempty" xml:"chn,omitempty"`
	// 字段的索引类型
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 别名
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// 分词符列表。仅当type参数取值为text时，必须设置。
	Token []*string `json:"token,omitempty" xml:"token,omitempty" type:"Repeated"`
	// 开启统计
	DocValue *bool `json:"doc_value,omitempty" xml:"doc_value,omitempty"`
}

func (s KeysValue) String() string {
	return tea.Prettify(s)
}

func (s KeysValue) GoString() string {
	return s.String()
}

func (s *KeysValue) SetCaseSensitive(v bool) *KeysValue {
	s.CaseSensitive = &v
	return s
}

func (s *KeysValue) SetChn(v bool) *KeysValue {
	s.Chn = &v
	return s
}

func (s *KeysValue) SetType(v string) *KeysValue {
	s.Type = &v
	return s
}

func (s *KeysValue) SetAlias(v string) *KeysValue {
	s.Alias = &v
	return s
}

func (s *KeysValue) SetToken(v []*string) *KeysValue {
	s.Token = v
	return s
}

func (s *KeysValue) SetDocValue(v bool) *KeysValue {
	s.DocValue = &v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("central")
	return nil
}

func (client *Client) ApplyConfigToMachineGroup(project *string, machineGroup *string, configName *string) (_result *ApplyConfigToMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyConfigToMachineGroupResponse{}
	_body, _err := client.ApplyConfigToMachineGroupWithOptions(project, machineGroup, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyConfigToMachineGroupWithOptions(project *string, machineGroup *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyConfigToMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	configName = openapiutil.GetEncodeParam(configName)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyConfigToMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs/" + tea.StringValue(configName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ApplyConfigToMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreateEtlMeta(project *string, request *BatchCreateEtlMetaRequest) (_result *BatchCreateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreateEtlMetaResponse{}
	_body, _err := client.BatchCreateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateEtlMetaWithOptions(project *string, request *BatchCreateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchCreateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaList)) {
		body["etlMetaList"] = request.EtlMetaList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchCreateEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchModifyEtlMetaStatus(project *string, request *BatchModifyEtlMetaStatusRequest) (_result *BatchModifyEtlMetaStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchModifyEtlMetaStatusResponse{}
	_body, _err := client.BatchModifyEtlMetaStatusWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchModifyEtlMetaStatusWithOptions(project *string, request *BatchModifyEtlMetaStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchModifyEtlMetaStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKeyList)) {
		body["etlMetaKeyList"] = request.EtlMetaKeyList
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.Range)) {
		body["range"] = request.Range
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchModifyEtlMetaStatus"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchModifyEtlMetaStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateEtlMeta(project *string, request *BatchUpdateEtlMetaRequest) (_result *BatchUpdateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateEtlMetaResponse{}
	_body, _err := client.BatchUpdateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateEtlMetaWithOptions(project *string, request *BatchUpdateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchUpdateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EtlMetaList))) {
		body["etlMetaList"] = request.EtlMetaList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &BatchUpdateEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsumerGroup(project *string, logstore *string, request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerGroupWithOptions(project *string, logstore *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerGroup)) {
		body["consumerGroup"] = request.ConsumerGroup
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDomain(project *string, request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDomainWithOptions(project *string, request *CreateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domainName"] = request.DomainName
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEtlMeta(project *string, request *CreateEtlMetaRequest) (_result *CreateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEtlMetaResponse{}
	_body, _err := client.CreateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEtlMetaWithOptions(project *string, request *CreateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		body["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaValue)) {
		body["etlMetaValue"] = request.EtlMetaValue
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIndex(project *string, logstore *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIndexWithOptions(project *string, logstore *string, request *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Line))) {
		body["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduce)) {
		body["log_reduce"] = request.LogReduce
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceBlackList)) {
		body["log_reduce_black_list"] = request.LogReduceBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceWhiteList)) {
		body["log_reduce_white_list"] = request.LogReduceWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTextLen)) {
		body["max_text_len"] = request.MaxTextLen
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogStore(project *string, request *CreateLogStoreRequest) (_result *CreateLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLogStoreResponse{}
	_body, _err := client.CreateLogStoreWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogStoreWithOptions(project *string, request *CreateLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppendMeta)) {
		body["appendMeta"] = request.AppendMeta
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSplit)) {
		body["autoSplit"] = request.AutoSplit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTracking)) {
		body["enable_tracking"] = request.EnableTracking
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EncryptConf))) {
		body["encrypt_conf"] = request.EncryptConf
	}

	if !tea.BoolValue(util.IsUnset(request.HotTtl)) {
		body["hot_ttl"] = request.HotTtl
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		body["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSplitShard)) {
		body["maxSplitShard"] = request.MaxSplitShard
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		body["shardCount"] = request.ShardCount
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		body["telemetryType"] = request.TelemetryType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogging(project *string, request *CreateLoggingRequest) (_result *CreateLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLoggingResponse{}
	_body, _err := client.CreateLoggingWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoggingWithOptions(project *string, request *CreateLoggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoggingDetails)) {
		body["loggingDetails"] = request.LoggingDetails
	}

	if !tea.BoolValue(util.IsUnset(request.LoggingProject)) {
		body["loggingProject"] = request.LoggingProject
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMachineGroup(project *string, request *CreateMachineGroupRequest) (_result *CreateMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMachineGroupResponse{}
	_body, _err := client.CreateMachineGroupWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMachineGroupWithOptions(project *string, request *CreateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.GroupAttribute))) {
		body["groupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineIdentifyType)) {
		body["machineIdentifyType"] = request.MachineIdentifyType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineList)) {
		body["machineList"] = request.MachineList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["projectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSavedSearch(project *string, request *CreateSavedSearchRequest) (_result *CreateSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSavedSearchResponse{}
	_body, _err := client.CreateSavedSearchWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSavedSearchWithOptions(project *string, request *CreateSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		body["logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.SavedsearchName)) {
		body["savedsearchName"] = request.SavedsearchName
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &CreateSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerGroup(project *string, logstore *string, consumerGroup *string) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(project, logstore, consumerGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(project *string, logstore *string, consumerGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	consumerGroup = openapiutil.GetEncodeParam(consumerGroup)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(project *string, domainName *string) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(project, domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(project *string, domainName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	domainName = openapiutil.GetEncodeParam(domainName)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains/" + tea.StringValue(domainName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEtlMeta(project *string, request *DeleteEtlMetaRequest) (_result *DeleteEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEtlMetaResponse{}
	_body, _err := client.DeleteEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEtlMetaWithOptions(project *string, request *DeleteEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		query["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIndex(project *string, logstore *string) (_result *DeleteIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIndexWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogStore(project *string, logstore *string) (_result *DeleteLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.DeleteLogStoreWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogStoreWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLogStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogging(project *string) (_result *DeleteLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLoggingResponse{}
	_body, _err := client.DeleteLoggingWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoggingWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMachineGroup(project *string, machineGroup *string) (_result *DeleteMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.DeleteMachineGroupWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMachineGroupWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(project *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSavedSearch(project *string, savedsearchName *string) (_result *DeleteSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSavedSearchResponse{}
	_body, _err := client.DeleteSavedSearchWithOptions(project, savedsearchName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSavedSearchWithOptions(project *string, savedsearchName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSavedSearchResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	savedsearchName = openapiutil.GetEncodeParam(savedsearchName)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppliedConfigs(project *string, machineGroup *string) (_result *GetAppliedConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppliedConfigsResponse{}
	_body, _err := client.GetAppliedConfigsWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppliedConfigsWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppliedConfigsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppliedConfigs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppliedConfigsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCheckPoint(project *string, logstore *string, consumerGroup *string, request *GetCheckPointRequest) (_result *GetCheckPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCheckPointResponse{}
	_body, _err := client.GetCheckPointWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCheckPointWithOptions(project *string, logstore *string, consumerGroup *string, request *GetCheckPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCheckPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	consumerGroup = openapiutil.GetEncodeParam(consumerGroup)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Shard)) {
		query["shard"] = request.Shard
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCheckPoint"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetCheckPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetContextLogs(project *string, logstore *string, request *GetContextLogsRequest) (_result *GetContextLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetContextLogsResponse{}
	_body, _err := client.GetContextLogsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetContextLogsWithOptions(project *string, logstore *string, request *GetContextLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetContextLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackLines)) {
		query["back_lines"] = request.BackLines
	}

	if !tea.BoolValue(util.IsUnset(request.ForwardLines)) {
		query["forward_lines"] = request.ForwardLines
	}

	if !tea.BoolValue(util.IsUnset(request.PackId)) {
		query["pack_id"] = request.PackId
	}

	if !tea.BoolValue(util.IsUnset(request.PackMeta)) {
		query["pack_meta"] = request.PackMeta
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetContextLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetContextLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCursor(project *string, logstore *string, shardId *string, request *GetCursorRequest) (_result *GetCursorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCursorResponse{}
	_body, _err := client.GetCursorWithOptions(project, logstore, shardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCursorWithOptions(project *string, logstore *string, shardId *string, request *GetCursorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCursorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	shardId = openapiutil.GetEncodeParam(shardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCursor"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCursorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCursorTime(project *string, logstore *string, shardId *string, request *GetCursorTimeRequest) (_result *GetCursorTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCursorTimeResponse{}
	_body, _err := client.GetCursorTimeWithOptions(project, logstore, shardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCursorTimeWithOptions(project *string, logstore *string, shardId *string, request *GetCursorTimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCursorTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	shardId = openapiutil.GetEncodeParam(shardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		query["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCursorTime"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCursorTimeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEtlMeta(project *string, request *GetEtlMetaRequest) (_result *GetEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEtlMetaResponse{}
	_body, _err := client.GetEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEtlMetaWithOptions(project *string, request *GetEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ElMetaName)) {
		query["elMetaName"] = request.ElMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistograms(project *string, logstore *string, request *GetHistogramsRequest) (_result *GetHistogramsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHistogramsResponse{}
	_body, _err := client.GetHistogramsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistogramsWithOptions(project *string, logstore *string, request *GetHistogramsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHistogramsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistograms"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetHistogramsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIndex(project *string, logstore *string) (_result *GetIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexResponse{}
	_body, _err := client.GetIndexWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIndexWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogStore(project *string, logstore *string) (_result *GetLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogStoreResponse{}
	_body, _err := client.GetLogStoreWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogStoreWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogStoreResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogging(project *string) (_result *GetLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLoggingResponse{}
	_body, _err := client.GetLoggingWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoggingWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLogs(project *string, logstore *string, request *GetLogsRequest) (_result *GetLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLogsResponse{}
	_body, _err := client.GetLogsWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLogsWithOptions(project *string, logstore *string, request *GetLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Line)) {
		query["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.PowerSql)) {
		query["powerSql"] = request.PowerSql
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMachineGroup(project *string, machineGroup *string) (_result *GetMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMachineGroupResponse{}
	_body, _err := client.GetMachineGroupWithOptions(project, machineGroup, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMachineGroupWithOptions(project *string, machineGroup *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(project *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(project, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(project *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectLogs(project *string, request *GetProjectLogsRequest) (_result *GetProjectLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectLogsResponse{}
	_body, _err := client.GetProjectLogsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectLogsWithOptions(project *string, request *GetProjectLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PowerSql)) {
		query["powerSql"] = request.PowerSql
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectLogs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &GetProjectLogsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSavedSearch(project *string, savedsearchName *string) (_result *GetSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSavedSearchResponse{}
	_body, _err := client.GetSavedSearchWithOptions(project, savedsearchName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSavedSearchWithOptions(project *string, savedsearchName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSavedSearchResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	savedsearchName = openapiutil.GetEncodeParam(savedsearchName)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConsumerGroup(project *string, logstore *string) (_result *ListConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupResponse{}
	_body, _err := client.ListConsumerGroupWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConsumerGroupWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDomains(project *string, request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDomainsWithOptions(project *string, request *ListDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["domainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/domains"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEtlMeta(project *string, request *ListEtlMetaRequest) (_result *ListEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEtlMetaResponse{}
	_body, _err := client.ListEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEtlMetaWithOptions(project *string, request *ListEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		query["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		query["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		query["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEtlMetaName(project *string, request *ListEtlMetaNameRequest) (_result *ListEtlMetaNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEtlMetaNameResponse{}
	_body, _err := client.ListEtlMetaNameWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEtlMetaNameWithOptions(project *string, request *ListEtlMetaNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEtlMetaNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEtlMetaName"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetanames"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEtlMetaNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLogStores(project *string, request *ListLogStoresRequest) (_result *ListLogStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLogStoresResponse{}
	_body, _err := client.ListLogStoresWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLogStoresWithOptions(project *string, request *ListLogStoresRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLogStoresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		query["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		query["telemetryType"] = request.TelemetryType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLogStores"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLogStoresResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMachineGroup(project *string, request *ListMachineGroupRequest) (_result *ListMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMachineGroupResponse{}
	_body, _err := client.ListMachineGroupWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMachineGroupWithOptions(project *string, request *ListMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMachines(project *string, machineGroup *string, request *ListMachinesRequest) (_result *ListMachinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMachinesResponse{}
	_body, _err := client.ListMachinesWithOptions(project, machineGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMachinesWithOptions(project *string, machineGroup *string, request *ListMachinesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMachinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachines"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/machines"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachinesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectWithOptions(request *ListProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["projectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSavedSearch(project *string, request *ListSavedSearchRequest) (_result *ListSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSavedSearchResponse{}
	_body, _err := client.ListSavedSearchWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSavedSearchWithOptions(project *string, request *ListSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListShards(project *string, logstore *string) (_result *ListShardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShardsResponse{}
	_body, _err := client.ListShardsWithOptions(project, logstore, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListShardsWithOptions(project *string, logstore *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListShardsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListShards"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &ListShardsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("resourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeShards(project *string, logstore *string, shardID *string, request *MergeShardsRequest) (_result *MergeShardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MergeShardsResponse{}
	_body, _err := client.MergeShardsWithOptions(project, logstore, shardID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeShardsWithOptions(project *string, logstore *string, shardID *string, request *MergeShardsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MergeShardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	shardID = openapiutil.GetEncodeParam(shardID)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MergeShards"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardID)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &MergeShardsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveConfigFromMachineGroup(project *string, machineGroup *string, configName *string) (_result *RemoveConfigFromMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveConfigFromMachineGroupResponse{}
	_body, _err := client.RemoveConfigFromMachineGroupWithOptions(project, machineGroup, configName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveConfigFromMachineGroupWithOptions(project *string, machineGroup *string, configName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveConfigFromMachineGroupResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	machineGroup = openapiutil.GetEncodeParam(machineGroup)
	configName = openapiutil.GetEncodeParam(configName)
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveConfigFromMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(machineGroup) + "/configs/" + tea.StringValue(configName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &RemoveConfigFromMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SplitShard(project *string, logstore *string, shardID *string, request *SplitShardRequest) (_result *SplitShardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SplitShardResponse{}
	_body, _err := client.SplitShardWithOptions(project, logstore, shardID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SplitShardWithOptions(project *string, logstore *string, shardID *string, request *SplitShardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SplitShardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	shardID = openapiutil.GetEncodeParam(shardID)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Action)) {
		query["action"] = request.Action
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		query["shardCount"] = request.ShardCount
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SplitShard"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/shards/" + tea.StringValue(shardID)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &SplitShardResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/untag"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCheckPoint(project *string, logstore *string, consumerGroup *string, request *UpdateCheckPointRequest) (_result *UpdateCheckPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCheckPointResponse{}
	_body, _err := client.UpdateCheckPointWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCheckPointWithOptions(project *string, logstore *string, consumerGroup *string, request *UpdateCheckPointRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCheckPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	consumerGroup = openapiutil.GetEncodeParam(consumerGroup)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Consumer)) {
		query["consumer"] = request.Consumer
	}

	if !tea.BoolValue(util.IsUnset(request.ForceSuccess)) {
		query["forceSuccess"] = request.ForceSuccess
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Checkpoint)) {
		body["checkpoint"] = request.Checkpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Shard)) {
		body["shard"] = request.Shard
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCheckPoint"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateCheckPointResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConsumerGroup(project *string, logstore *string, consumerGroup *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.UpdateConsumerGroupWithOptions(project, logstore, consumerGroup, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConsumerGroupWithOptions(project *string, logstore *string, consumerGroup *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	consumerGroup = openapiutil.GetEncodeParam(consumerGroup)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConsumerGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/consumergroups/" + tea.StringValue(consumerGroup)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEtlMeta(project *string, request *UpdateEtlMetaRequest) (_result *UpdateEtlMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEtlMetaResponse{}
	_body, _err := client.UpdateEtlMetaWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEtlMetaWithOptions(project *string, request *UpdateEtlMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEtlMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaKey)) {
		body["etlMetaKey"] = request.EtlMetaKey
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaName)) {
		body["etlMetaName"] = request.EtlMetaName
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaTag)) {
		body["etlMetaTag"] = request.EtlMetaTag
	}

	if !tea.BoolValue(util.IsUnset(request.EtlMetaValue)) {
		body["etlMetaValue"] = request.EtlMetaValue
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEtlMeta"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/etlmetas"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateEtlMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIndex(project *string, logstore *string, request *UpdateIndexRequest) (_result *UpdateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIndexResponse{}
	_body, _err := client.UpdateIndexWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIndexWithOptions(project *string, logstore *string, request *UpdateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		body["keys"] = request.Keys
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Line))) {
		body["line"] = request.Line
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduce)) {
		body["log_reduce"] = request.LogReduce
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceBlackList)) {
		body["log_reduce_black_list"] = request.LogReduceBlackList
	}

	if !tea.BoolValue(util.IsUnset(request.LogReduceWhiteList)) {
		body["log_reduce_white_list"] = request.LogReduceWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.MaxTextLen)) {
		body["max_text_len"] = request.MaxTextLen
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIndex"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore) + "/index"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateIndexResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogStore(project *string, logstore *string, request *UpdateLogStoreRequest) (_result *UpdateLogStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLogStoreResponse{}
	_body, _err := client.UpdateLogStoreWithOptions(project, logstore, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLogStoreWithOptions(project *string, logstore *string, request *UpdateLogStoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLogStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	logstore = openapiutil.GetEncodeParam(logstore)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppendMeta)) {
		body["appendMeta"] = request.AppendMeta
	}

	if !tea.BoolValue(util.IsUnset(request.AutoSplit)) {
		body["autoSplit"] = request.AutoSplit
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTracking)) {
		body["enable_tracking"] = request.EnableTracking
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.EncryptConf))) {
		body["encrypt_conf"] = request.EncryptConf
	}

	if !tea.BoolValue(util.IsUnset(request.HotTtl)) {
		body["hot_ttl"] = request.HotTtl
	}

	if !tea.BoolValue(util.IsUnset(request.LogstoreName)) {
		body["logstoreName"] = request.LogstoreName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSplitShard)) {
		body["maxSplitShard"] = request.MaxSplitShard
	}

	if !tea.BoolValue(util.IsUnset(request.ShardCount)) {
		body["shardCount"] = request.ShardCount
	}

	if !tea.BoolValue(util.IsUnset(request.TelemetryType)) {
		body["telemetryType"] = request.TelemetryType
	}

	if !tea.BoolValue(util.IsUnset(request.Ttl)) {
		body["ttl"] = request.Ttl
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogStore"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logstores/" + tea.StringValue(logstore)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLogStoreResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLogging(project *string, request *UpdateLoggingRequest) (_result *UpdateLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLoggingResponse{}
	_body, _err := client.UpdateLoggingWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoggingWithOptions(project *string, request *UpdateLoggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LoggingDetails)) {
		body["loggingDetails"] = request.LoggingDetails
	}

	if !tea.BoolValue(util.IsUnset(request.LoggingProject)) {
		body["loggingProject"] = request.LoggingProject
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLogging"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/logging"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMachineGroup(project *string, groupName *string, request *UpdateMachineGroupRequest) (_result *UpdateMachineGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMachineGroupResponse{}
	_body, _err := client.UpdateMachineGroupWithOptions(project, groupName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMachineGroupWithOptions(project *string, groupName *string, request *UpdateMachineGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMachineGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	groupName = openapiutil.GetEncodeParam(groupName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.GroupAttribute))) {
		body["groupAttribute"] = request.GroupAttribute
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["groupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		body["groupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineIdentifyType)) {
		body["machineIdentifyType"] = request.MachineIdentifyType
	}

	if !tea.BoolValue(util.IsUnset(request.MachineList)) {
		body["machineList"] = request.MachineList
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMachineGroup"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/machinegroups/" + tea.StringValue(groupName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateMachineGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(project *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(project, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(project *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSavedSearch(project *string, savedsearchName *string, request *UpdateSavedSearchRequest) (_result *UpdateSavedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSavedSearchResponse{}
	_body, _err := client.UpdateSavedSearchWithOptions(project, savedsearchName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSavedSearchWithOptions(project *string, savedsearchName *string, request *UpdateSavedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSavedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["project"] = project
	savedsearchName = openapiutil.GetEncodeParam(savedsearchName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["displayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		body["logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.SavedsearchName)) {
		body["savedsearchName"] = request.SavedsearchName
	}

	if !tea.BoolValue(util.IsUnset(request.SearchQuery)) {
		body["searchQuery"] = request.SearchQuery
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		body["topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSavedSearch"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/savedsearches/" + tea.StringValue(savedsearchName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &UpdateSavedSearchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

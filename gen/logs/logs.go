// Package logs provides a client for Amazon CloudWatch Logs.
package logs

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
)

// Logs is a client for Amazon CloudWatch Logs.
type Logs struct {
	client *aws.JSONClient
}

// New returns a new Logs client.
func New(key, secret, region string, client *http.Client) *Logs {
	if client == nil {
		client = http.DefaultClient
	}

	return &Logs{
		client: &aws.JSONClient{
			Client:       client,
			Region:       region,
			Endpoint:     fmt.Sprintf("https://logs.%s.amazonaws.com", region),
			Prefix:       "logs",
			Key:          key,
			Secret:       secret,
			JSONVersion:  "1.1",
			TargetPrefix: "Logs_20140328",
		},
	}
}

// CreateLogGroup creates a new log group with the specified name. The name
// of the log group must be unique within a region for an AWS account. You
// can create up to 500 log groups per account. You must use the following
// guidelines when naming a log group: Log group names can be between 1 and
// 512 characters long. Allowed characters are a-z, 0-9, '_' (underscore),
// '-' (hyphen), '/' (forward slash), and '.' (period).
func (c *Logs) CreateLogGroup(req CreateLogGroupRequest) (err error) {
	// NRE
	err = c.client.Do("CreateLogGroup", "POST", "/", req, nil)
	return
}

// CreateLogStream creates a new log stream in the specified log group. The
// name of the log stream must be unique within the log group. There is no
// limit on the number of log streams that can exist in a log group. You
// must use the following guidelines when naming a log stream: Log stream
// names can be between 1 and 512 characters long. The ':' colon character
// is not allowed.
func (c *Logs) CreateLogStream(req CreateLogStreamRequest) (err error) {
	// NRE
	err = c.client.Do("CreateLogStream", "POST", "/", req, nil)
	return
}

// DeleteLogGroup deletes the log group with the specified name and
// permanently deletes all the archived log events associated with it.
func (c *Logs) DeleteLogGroup(req DeleteLogGroupRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteLogGroup", "POST", "/", req, nil)
	return
}

// DeleteLogStream deletes a log stream and permanently deletes all the
// archived log events associated with it.
func (c *Logs) DeleteLogStream(req DeleteLogStreamRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteLogStream", "POST", "/", req, nil)
	return
}

// DeleteMetricFilter deletes a metric filter associated with the specified
// log group.
func (c *Logs) DeleteMetricFilter(req DeleteMetricFilterRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteMetricFilter", "POST", "/", req, nil)
	return
}

// DeleteRetentionPolicy deletes the retention policy of the specified log
// group. Log events would not expire if they belong to log groups without
// a retention policy.
func (c *Logs) DeleteRetentionPolicy(req DeleteRetentionPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteRetentionPolicy", "POST", "/", req, nil)
	return
}

// DescribeLogGroups returns all the log groups that are associated with
// the AWS account making the request. The list returned in the response is
// ASCII-sorted by log group name. By default, this operation returns up to
// 50 log groups. If there are more log groups to list, the response would
// contain a nextToken value in the response body. You can also limit the
// number of log groups returned in the response by specifying the limit
// parameter in the request.
func (c *Logs) DescribeLogGroups(req DescribeLogGroupsRequest) (resp *DescribeLogGroupsResponse, err error) {
	resp = &DescribeLogGroupsResponse{}
	err = c.client.Do("DescribeLogGroups", "POST", "/", req, resp)
	return
}

// DescribeLogStreams returns all the log streams that are associated with
// the specified log group. The list returned in the response is
// ASCII-sorted by log stream name. By default, this operation returns up
// to 50 log streams. If there are more log streams to list, the response
// would contain a nextToken value in the response body. You can also limit
// the number of log streams returned in the response by specifying the
// limit parameter in the request.
func (c *Logs) DescribeLogStreams(req DescribeLogStreamsRequest) (resp *DescribeLogStreamsResponse, err error) {
	resp = &DescribeLogStreamsResponse{}
	err = c.client.Do("DescribeLogStreams", "POST", "/", req, resp)
	return
}

// DescribeMetricFilters returns all the metrics filters associated with
// the specified log group. The list returned in the response is
// ASCII-sorted by filter name. By default, this operation returns up to 50
// metric filters. If there are more metric filters to list, the response
// would contain a nextToken value in the response body. You can also limit
// the number of metric filters returned in the response by specifying the
// limit parameter in the request.
func (c *Logs) DescribeMetricFilters(req DescribeMetricFiltersRequest) (resp *DescribeMetricFiltersResponse, err error) {
	resp = &DescribeMetricFiltersResponse{}
	err = c.client.Do("DescribeMetricFilters", "POST", "/", req, resp)
	return
}

// GetLogEvents retrieves log events from the specified log stream. You can
// provide an optional time range to filter the results on the event
// timestamp . By default, this operation returns as much log events as can
// fit in a response size of 1MB, up to 10,000 log events. The response
// will always include a nextForwardToken and a nextBackwardToken in the
// response body. You can use any of these tokens in subsequent
// GetLogEvents requests to paginate through events in either forward or
// backward direction. You can also limit the number of log events returned
// in the response by specifying the limit parameter in the request.
func (c *Logs) GetLogEvents(req GetLogEventsRequest) (resp *GetLogEventsResponse, err error) {
	resp = &GetLogEventsResponse{}
	err = c.client.Do("GetLogEvents", "POST", "/", req, resp)
	return
}

// PutLogEvents uploads a batch of log events to the specified log stream.
// Every PutLogEvents request must include the sequenceToken obtained from
// the response of the previous request. An upload in a newly created log
// stream does not require a sequenceToken . The batch of events must
// satisfy the following constraints: The maximum batch size is 32,768
// bytes, and this size is calculated as the sum of all event messages in
// UTF-8, plus 26 bytes for each log event. None of the log events in the
// batch can be more than 2 hours in the future. None of the log events in
// the batch can be older than 14 days or the retention period of the log
// group. The log events in the batch must be in chronological ordered by
// their timestamp The maximum number of log events in a batch is 1,000.
func (c *Logs) PutLogEvents(req PutLogEventsRequest) (resp *PutLogEventsResponse, err error) {
	resp = &PutLogEventsResponse{}
	err = c.client.Do("PutLogEvents", "POST", "/", req, resp)
	return
}

// PutMetricFilter creates or updates a metric filter and associates it
// with the specified log group. Metric filters allow you to configure
// rules to extract metric data from log events ingested through
// PutLogEvents requests.
func (c *Logs) PutMetricFilter(req PutMetricFilterRequest) (err error) {
	// NRE
	err = c.client.Do("PutMetricFilter", "POST", "/", req, nil)
	return
}

// PutRetentionPolicy sets the retention of the specified log group. A
// retention policy allows you to configure the number of days you want to
// retain log events in the specified log group.
func (c *Logs) PutRetentionPolicy(req PutRetentionPolicyRequest) (err error) {
	// NRE
	err = c.client.Do("PutRetentionPolicy", "POST", "/", req, nil)
	return
}

// TestMetricFilter tests the filter pattern of a metric filter against a
// sample of log event messages. You can use this operation to validate the
// correctness of a metric filter pattern.
func (c *Logs) TestMetricFilter(req TestMetricFilterRequest) (resp *TestMetricFilterResponse, err error) {
	resp = &TestMetricFilterResponse{}
	err = c.client.Do("TestMetricFilter", "POST", "/", req, resp)
	return
}

type CreateLogGroupRequest struct {
	LogGroupName string `json:"logGroupName"`
}

type CreateLogStreamRequest struct {
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

type DeleteLogGroupRequest struct {
	LogGroupName string `json:"logGroupName"`
}

type DeleteLogStreamRequest struct {
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
}

type DeleteMetricFilterRequest struct {
	FilterName   string `json:"filterName"`
	LogGroupName string `json:"logGroupName"`
}

type DeleteRetentionPolicyRequest struct {
	LogGroupName string `json:"logGroupName"`
}

type DescribeLogGroupsRequest struct {
	Limit              int    `json:"limit,omitempty"`
	LogGroupNamePrefix string `json:"logGroupNamePrefix,omitempty"`
	NextToken          string `json:"nextToken,omitempty"`
}

type DescribeLogGroupsResponse struct {
	LogGroups []LogGroup `json:"logGroups,omitempty"`
	NextToken string     `json:"nextToken,omitempty"`
}

type DescribeLogStreamsRequest struct {
	Limit               int    `json:"limit,omitempty"`
	LogGroupName        string `json:"logGroupName"`
	LogStreamNamePrefix string `json:"logStreamNamePrefix,omitempty"`
	NextToken           string `json:"nextToken,omitempty"`
}

type DescribeLogStreamsResponse struct {
	LogStreams []LogStream `json:"logStreams,omitempty"`
	NextToken  string      `json:"nextToken,omitempty"`
}

type DescribeMetricFiltersRequest struct {
	FilterNamePrefix string `json:"filterNamePrefix,omitempty"`
	Limit            int    `json:"limit,omitempty"`
	LogGroupName     string `json:"logGroupName"`
	NextToken        string `json:"nextToken,omitempty"`
}

type DescribeMetricFiltersResponse struct {
	MetricFilters []MetricFilter `json:"metricFilters,omitempty"`
	NextToken     string         `json:"nextToken,omitempty"`
}

type GetLogEventsRequest struct {
	EndTime       int    `json:"endTime,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	LogGroupName  string `json:"logGroupName"`
	LogStreamName string `json:"logStreamName"`
	NextToken     string `json:"nextToken,omitempty"`
	StartFromHead bool   `json:"startFromHead,omitempty"`
	StartTime     int    `json:"startTime,omitempty"`
}

type GetLogEventsResponse struct {
	Events            []OutputLogEvent `json:"events,omitempty"`
	NextBackwardToken string           `json:"nextBackwardToken,omitempty"`
	NextForwardToken  string           `json:"nextForwardToken,omitempty"`
}

type InputLogEvent struct {
	Message   string `json:"message"`
	Timestamp int    `json:"timestamp"`
}

type LogGroup struct {
	ARN               string `json:"arn,omitempty"`
	CreationTime      int    `json:"creationTime,omitempty"`
	LogGroupName      string `json:"logGroupName,omitempty"`
	MetricFilterCount int    `json:"metricFilterCount,omitempty"`
	RetentionInDays   int    `json:"retentionInDays,omitempty"`
	StoredBytes       int    `json:"storedBytes,omitempty"`
}

type LogStream struct {
	ARN                 string `json:"arn,omitempty"`
	CreationTime        int    `json:"creationTime,omitempty"`
	FirstEventTimestamp int    `json:"firstEventTimestamp,omitempty"`
	LastEventTimestamp  int    `json:"lastEventTimestamp,omitempty"`
	LastIngestionTime   int    `json:"lastIngestionTime,omitempty"`
	LogStreamName       string `json:"logStreamName,omitempty"`
	StoredBytes         int    `json:"storedBytes,omitempty"`
	UploadSequenceToken string `json:"uploadSequenceToken,omitempty"`
}

type MetricFilter struct {
	CreationTime          int                    `json:"creationTime,omitempty"`
	FilterName            string                 `json:"filterName,omitempty"`
	FilterPattern         string                 `json:"filterPattern,omitempty"`
	MetricTransformations []MetricTransformation `json:"metricTransformations,omitempty"`
}

type MetricFilterMatchRecord struct {
	EventMessage    string            `json:"eventMessage,omitempty"`
	EventNumber     int               `json:"eventNumber,omitempty"`
	ExtractedValues map[string]string `json:"extractedValues,omitempty"`
}

type MetricTransformation struct {
	MetricName      string `json:"metricName"`
	MetricNamespace string `json:"metricNamespace"`
	MetricValue     string `json:"metricValue"`
}

type OutputLogEvent struct {
	IngestionTime int    `json:"ingestionTime,omitempty"`
	Message       string `json:"message,omitempty"`
	Timestamp     int    `json:"timestamp,omitempty"`
}

type PutLogEventsRequest struct {
	LogEvents     []InputLogEvent `json:"logEvents"`
	LogGroupName  string          `json:"logGroupName"`
	LogStreamName string          `json:"logStreamName"`
	SequenceToken string          `json:"sequenceToken,omitempty"`
}

type PutLogEventsResponse struct {
	NextSequenceToken string `json:"nextSequenceToken,omitempty"`
}

type PutMetricFilterRequest struct {
	FilterName            string                 `json:"filterName"`
	FilterPattern         string                 `json:"filterPattern"`
	LogGroupName          string                 `json:"logGroupName"`
	MetricTransformations []MetricTransformation `json:"metricTransformations"`
}

type PutRetentionPolicyRequest struct {
	LogGroupName    string `json:"logGroupName"`
	RetentionInDays int    `json:"retentionInDays"`
}

type TestMetricFilterRequest struct {
	FilterPattern    string   `json:"filterPattern"`
	LogEventMessages []string `json:"logEventMessages"`
}

type TestMetricFilterResponse struct {
	Matches []MetricFilterMatchRecord `json:"matches,omitempty"`
}

var _ time.Time // to avoid errors if the time package isn't referenced

package models

// Schema defines the Schema
type Schema struct {
	// Name of the schema
	Name string `json:"name"`
	// PrimaryKeys is a list of columns that are set as primary keys
	PrimaryKeys []string `json:"primaryKeys,omitempty"`
	// Dimensions is a list of fields that represents the dimensions in the schema
	// ref: https://docs.pinot.apache.org/basics/components/schema#categories
	Dimensions []*DimensionFieldSpec `json:"dimensions,omitempty"`
	// Metrics is a list of fields that represents the metrics in the schema
	// ref: https://docs.pinot.apache.org/basics/components/schema#categories
	Metrics []*MetricFieldSpec `json:"metrics,omitempty"`
	// DateTimes is a list of fields that represents the datetimes in the schema
	// ref: https://docs.pinot.apache.org/basics/components/schema#categories
	DateTimes []*DatetimeFieldSpec `json:"dateTimes,omitempty"`
	// TimeField represents the granularity
	TimeField *TimeFieldSpec `json:"timeField,omitempty"`
}

// DimensionFieldSpec is typically used in slice and dice operations for answering business queries
type DimensionFieldSpec struct {
	CommonColumnSpec
}

// MetricFieldSpec represents the quantitative data of the table. Such columns are used for aggregation.
// In data warehouse terminology, these can also be referred to as fact or measure columns
type MetricFieldSpec struct {
	CommonColumnSpec
}

// DatetimeFieldSpec represents time columns in the data. There can be multiple time columns in a table, but only one of them
// can be treated as primary. Primary time column is the one that is present in the segment config.
type DatetimeFieldSpec struct {
	Format      *string `json:"format,omitempty"`
	Granularity *string `json:"granularity,omitempty"`
	CommonColumnSpec
}

// TimeFieldSpec represents the granularity for both ingestion and query segments
type TimeFieldSpec struct {
	IncomingGranularity *TimeGranularitySpec `json:"incomingGranularity,omitempty"`
	OutgoingGranularity *TimeGranularitySpec `json:"outgoingGranularity,omitempty"`
	CommonColumnSpec
}

// TimeGranularitySpec represents the granularity object
type TimeGranularitySpec struct {
	// Name of the time granularity specification
	Name     *string `json:"name,omitempty"`
	DataType *string `json:"dataType,omitempty"`
	// TimeType is one of  TimeUnit enum values. e.g. HOURS , MINUTES etc. If your date is not in EPOCH format,
	// this value is not used and can be set to MILLISECONDS or any other unit.
	TimeType *string `json:"typeType,omitempty"`
	// TimeUnitSize is multiplied to the value present in the time column to get an actual timestamp.
	// eg: if timesize is 5 and value in time column is 4996308 minutes. The value that will be converted
	// to epoch timestamp will be 4996308 * 5 * 60 * 1000 = 1498892400000 milliseconds.
	// If your date is not in EPOCH format, this value is not used and can be set to 1 or any other integer.
	TimeUnitSize *int32 `json:"timeUnitSize,omitempty"`
	// TimeFormat can be either EPOCH or SIMPLE_DATE_FORMAT. If it is SIMPLE_DATE_FORMAT, the pattern string is also specified.
	//
	// Here are some sample date-time formats you can use in the schema:
	// 1:MILLISECONDS:EPOCH - used when timestamp is in the epoch milliseconds and stored in LONG format
	// 1:HOURS:EPOCH - used when timestamp is in the epoch hours and stored in LONG  or INT format
	// 1:DAYS:SIMPLE_DATE_FORMAT:yyyy-MM-dd - when date is in STRING format and has the pattern year-month-date. e.g. 2020-08-21
	// 1:HOURS:SIMPLE_DATE_FORMAT:EEE MMM dd HH:mm:ss ZZZ yyyy - when date is in STRING format. e.g. s Mon Aug 24 12:36:50 America/Los_Angeles 2019
	TimeFormat *string `json:"timeFormat,omitempty"`
}

// CommonColumnSpec represents the default values of such column
type CommonColumnSpec struct {
	Name                   *string     `json:"name,omitempty"`
	SingleValueField       *bool       `json:"singleValueField,omitempty"`
	MaxLength              *int32      `json:"maxLength,omitempty"`
	DataType               *string     `json:"dataType,omitempty"`
	DefaultNullValue       interface{} `json:"defaultNullValue,omitempty"`
	VirtualColumnProvider  *string     `json:"virtualColumnProvider,omitempty"`
	TransformFunction      *string     `json:"transformFunction,omitempty"`
	DefaultNullValueString *string     `json:"defaultNullValueString,omitempty"`
}

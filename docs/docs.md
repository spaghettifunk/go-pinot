# Pinot Controller API
APIs for accessing Pinot Controller information
  

## Informations

### Version

1.0

### Contact

https://github.com/apache/incubator-pinot  

## Tags

  ### <span id="tag-broker"></span>Broker

  ### <span id="tag-cluster"></span>Cluster

  ### <span id="tag-health"></span>Health

  ### <span id="tag-table"></span>Table

  ### <span id="tag-instance"></span>Instance

  ### <span id="tag-leader"></span>Leader

  ### <span id="tag-schema"></span>Schema

  ### <span id="tag-segment"></span>Segment

  ### <span id="tag-tenant"></span>Tenant

  ### <span id="tag-task"></span>Task

  ### <span id="tag-version"></span>Version

  ### <span id="tag-zookeeper"></span>Zookeeper

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json
  * multipart/form-data
  * text/plain

### Produces
  * application/octet-stream
  * application/json
  * text/plain

## All endpoints

###  broker

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /brokers/tables/{tableName} | [get brokers for table](#get-brokers-for-table) | List brokers for a given table |
| GET | /v2/brokers/tables/{tableName} | [get brokers for table v2](#get-brokers-for-table-v2) | List brokers for a given table |
| GET | /brokers/tenants/{tenantName} | [get brokers for tenant](#get-brokers-for-tenant) | List brokers for a given tenant |
| GET | /v2/brokers/tenants/{tenantName} | [get brokers for tenant v2](#get-brokers-for-tenant-v2) | List brokers for a given tenant |
| GET | /brokers/tables | [get tables to brokers mapping](#get-tables-to-brokers-mapping) | List tables to brokers mappings |
| GET | /v2/brokers/tables | [get tables to brokers mapping v2](#get-tables-to-brokers-mapping-v2) | List tables to brokers mappings |
| GET | /brokers/tenants | [get tenants to brokers mapping](#get-tenants-to-brokers-mapping) | List tenants to brokers mappings |
| GET | /v2/brokers/tenants | [get tenants to brokers mapping v2](#get-tenants-to-brokers-mapping-v2) | List tenants to brokers mappings |
| GET | /brokers | [list brokers mapping](#list-brokers-mapping) | List tenants and tables to brokers mappings |
| GET | /v2/brokers | [list brokers mapping v2](#list-brokers-mapping-v2) | List tenants and tables to brokers mappings |
| POST | /brokers/instances/{instanceName}/qps | [toggle query rate limiting](#toggle-query-rate-limiting) | Enable/disable the query rate limiting for a broker instance |
  


###  cluster

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /cluster/configs/{configName} | [delete cluster config](#delete-cluster-config) | Delete cluster configuration |
| GET | /cluster/info | [get cluster info](#get-cluster-info) | Get cluster Info |
| GET | /cluster/configs | [list cluster configs](#list-cluster-configs) | List cluster configurations |
| POST | /cluster/configs | [update cluster config](#update-cluster-config) | Update cluster configuration |
  


###  health

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /health | [check health](#check-health) | Check controller health |
| GET | /pinot-controller/admin | [check health legacy](#check-health-legacy) | Check controller health |
  


###  instance

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /instances | [add instance](#add-instance) | Create a new instance |
| DELETE | /instances/{instanceName} | [drop instance](#drop-instance) | Drop an instance |
| GET | /instances | [get all instances](#get-all-instances) | List all instances |
| GET | /instances/{instanceName} | [get instance](#get-instance) | Get instance information |
| POST | /instances/{instanceName}/state | [toggle instance state](#toggle-instance-state) | Enable/disable/drop an instance |
| PUT | /instances/{instanceName} | [update instance](#update-instance) | Update the specified instance |
| PUT | /instances/{instanceName}/updateTags | [update instance tags](#update-instance-tags) | Update the tags of the specified instance |
  


###  leader

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /leader/tables/{tableName} | [get leader for table](#get-leader-for-table) | Gets leader for a given table |
| GET | /leader/tables | [get leaders for all tables](#get-leaders-for-all-tables) | Gets leaders for all tables |
  


###  schema

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /schemas | [add schema](#add-schema) | Add a new schema |
| DELETE | /schemas/{schemaName} | [delete schema](#delete-schema) | Delete a schema |
| GET | /schemas/{schemaName} | [get schema](#get-schema) | Get a schema |
| GET | /tables/{tableName}/schema | [get table schema](#get-table-schema) | Get table schema |
| GET | /schemas | [list schema names](#list-schema-names) | List all schema names |
| PUT | /schemas/{schemaName} | [update schema](#update-schema) | Update a schema |
| POST | /schemas/validate | [validate schema](#validate-schema) | Validate schema |
  


###  segment

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /segments/{tableName} | [delete all segments](#delete-all-segments) | Delete all segments |
| DELETE | /segments/{tableName}/{segmentName} | [delete segment](#delete-segment) | Delete a segment |
| POST | /segments/{tableName}/delete | [delete segments](#delete-segments) | Delete the segments in the JSON array payload |
| GET | /segments/{tableName}/{segmentName} | [download segment](#download-segment) | Download a segment |
| POST | /segments/{tableName}/endReplaceSegments | [end replace segments](#end-replace-segments) | End to replace segments |
| GET | /tables/{realtimeTableName}/consumingSegmentsInfo | [get consuming segments info](#get-consuming-segments-info) | Returns state of consuming segments |
| GET | /segments/{tableName}/{segmentName}/metadata | [get segment metadata](#get-segment-metadata) | Get the metadata for a segment |
| GET | /tables/{tableName}/segments/{segmentName}/metadata | [get segment metadata deprecated1](#get-segment-metadata-deprecated1) | Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead) |
| GET | /tables/{tableName}/segments/{segmentName} | [get segment metadata deprecated2](#get-segment-metadata-deprecated2) | Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead) |
| GET | /segments/{tableName}/crc | [get segment to crc map](#get-segment-to-crc-map) | Get a map from segment to CRC of the segment (only apply to OFFLINE table) |
| GET | /tables/{tableName}/segments/crc | [get segment to crc map deprecated](#get-segment-to-crc-map-deprecated) | Get a map from segment to CRC of the segment (deprecated, use 'GET /segments/{tableName}/crc' instead) |
| GET | /segments/{tableName} | [get segments](#get-segments) | List all segments |
| GET | /segments/{tableName}/metadata | [get server metadata](#get-server-metadata) | Get the server metadata for all table segments |
| GET | /segments/{tableName}/servers | [get server to segments map](#get-server-to-segments-map) | Get a map from server to segments hosted by the server |
| GET | /tables/{tableName}/segments | [get server to segments map deprecated1](#get-server-to-segments-map-deprecated1) | Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead) |
| GET | /tables/{tableName}/segments/metadata | [get server to segments map deprecated2](#get-server-to-segments-map-deprecated2) | Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead) |
| POST | /segments/{tableName}/reload | [reload all segments](#reload-all-segments) | Reload all segments |
| POST | /tables/{tableName}/segments/reload | [reload all segments deprecated1](#reload-all-segments-deprecated1) | Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead) |
| GET | /tables/{tableName}/segments/reload | [reload all segments deprecated2](#reload-all-segments-deprecated2) | Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead) |
| POST | /segments/{tableName}/{segmentName}/reload | [reload segment](#reload-segment) | Reload a segment |
| POST | /tables/{tableName}/segments/{segmentName}/reload | [reload segment deprecated1](#reload-segment-deprecated1) | Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead) |
| GET | /tables/{tableName}/segments/{segmentName}/reload | [reload segment deprecated2](#reload-segment-deprecated2) | Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead) |
| POST | /segments/{tableNameWithType}/reset | [reset all segments](#reset-all-segments) | Resets all segments of the table, by first disabling them, waiting for external view to stabilize, and finally enabling the segments |
| POST | /segments/{tableNameWithType}/{segmentName}/reset | [reset segment](#reset-segment) | Resets a segment by first disabling it, waiting for external view to stabilize, and finally enabling it again |
| POST | /segments/{tableName}/startReplaceSegments | [start replace segments](#start-replace-segments) | Start to replace segments |
| POST | /segments | [upload segment as multi part](#upload-segment-as-multi-part) | Upload a segment |
| POST | /v2/segments | [upload segment as multi part v2](#upload-segment-as-multi-part-v2) | Upload a segment |
  


###  table

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /tables | [add table](#add-table) | Adds a table |
| GET | /tables/{tableName} | [alter table state or list table config](#alter-table-state-or-list-table-config) | Get/Enable/Disable/Drop a table |
| POST | /tables/{tableName}/assignInstances | [assign instances](#assign-instances) | Assign server instances to a table |
| POST | /tables/validate | [check table config](#check-table-config) | Validate table config for a table |
| DELETE | /tables/{tableName} | [delete table](#delete-table) | Deletes a table |
| GET | /tables/{tableName}/externalview | [get external view](#get-external-view) | Get table external view |
| GET | /tables/{tableName}/idealstate | [get ideal state](#get-ideal-state) | Get table ideal state |
| GET | /tables/{tableName}/instancePartitions | [get instance partitions](#get-instance-partitions) | Get the instance partitions |
| GET | /tables/{tableName}/instances | [get table instances](#get-table-instances) | List table instances |
| GET | /tables/{tableName}/size | [get table size](#get-table-size) | Read table sizes |
| GET | /tables/{tableName}/state | [get table state](#get-table-state) | Get current table state |
| GET | /tables/{tableName}/stats | [get table stats](#get-table-stats) | table stats |
| POST | /ingestFromFile | [ingest from file](#ingest-from-file) | Ingest a file |
| POST | /ingestFromURI | [ingest from URI](#ingest-from-uri) | Ingest from the given URI |
| GET | /tables | [list table configs](#list-table-configs) | Lists all tables in cluster |
| PUT | /tables/{tableName}/segmentConfigs | [put](#put) | Update segments configuration |
| POST | /tables/{tableName}/rebalance | [rebalance](#rebalance) | Rebalances a table (reassign instances and segments for a table) |
| POST | /tables/{tableName}/rebuildBrokerResourceFromHelixTags | [rebuild broker resource](#rebuild-broker-resource) | Rebuild broker resource for table |
| PUT | /tables/recommender | [recommend config](#recommend-config) | Recommend config |
| DELETE | /tables/{tableName}/instancePartitions | [remove instance partitions](#remove-instance-partitions) | Remove the instance partitions |
| POST | /tables/{tableName}/replaceInstance | [replace instance](#replace-instance) | Replace an instance in the instance partitions |
| PUT | /tables/{tableName}/instancePartitions | [set instance partitions](#set-instance-partitions) | Create/update the instance partitions |
| PUT | /tables/{tableName}/indexingConfigs | [update indexing config](#update-indexing-config) | Update table indexing configuration |
| PUT | /tables/{tableName} | [update table config](#update-table-config) | Updates table config for a table |
| PUT | /tables/{tableName}/metadataConfigs | [update table metadata](#update-table-metadata) | Update table metadata |
  


###  task

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| PUT | /tasks/{taskType}/cleanup | [clean up tasks](#clean-up-tasks) | Clean up finished tasks (COMPLETED, FAILED) for the given task type |
| PUT | /tasks/cleanuptasks/{taskType} | [clean up tasks deprecated](#clean-up-tasks-deprecated) | Clean up finished tasks (COMPLETED, FAILED) for the given task type (deprecated) |
| DELETE | /tasks/taskqueue/{taskType} | [delete task queue](#delete-task-queue) | Delete a task queue (deprecated) |
| DELETE | /tasks/{taskType} | [delete tasks](#delete-tasks) | Delete all tasks (as well as the task queue) for the given task type |
| GET | /tasks/scheduler/information | [get cron scheduler information](#get-cron-scheduler-information) | Fetch cron scheduler information |
| GET | /tasks/scheduler/jobDetails | [get cron scheduler job details](#get-cron-scheduler-job-details) | Fetch cron scheduler job keys |
| GET | /tasks/scheduler/jobKeys | [get cron scheduler job keys](#get-cron-scheduler-job-keys) | Fetch cron scheduler job keys |
| GET | /tasks/task/{taskName}/config | [get task configs](#get-task-configs) | Get the task config (a list of child task configs) for the given task |
| GET | /tasks/taskconfig/{taskName} | [get task configs deprecated](#get-task-configs-deprecated) | Get the task config (a list of child task configs) for the given task (deprecated) |
| GET | /tasks/{taskType}/state | [get task queue state](#get-task-queue-state) | Get the state (task queue state) for the given task type |
| GET | /tasks/taskqueuestate/{taskType} | [get task queue state deprecated](#get-task-queue-state-deprecated) | Get the state (task queue state) for the given task type (deprecated) |
| GET | /tasks/taskqueues | [get task queues](#get-task-queues) | List all task queues (deprecated) |
| GET | /tasks/task/{taskName}/state | [get task state](#get-task-state) | Get the task state for the given task |
| GET | /tasks/taskstate/{taskName} | [get task state deprecated](#get-task-state-deprecated) | Get the task state for the given task (deprecated) |
| GET | /tasks/{taskType}/taskstates | [get task states](#get-task-states) | Get a map from task to task state for the given task type |
| GET | /tasks/taskstates/{taskType} | [get task states deprecated](#get-task-states-deprecated) | Get a map from task to task state for the given task type (deprecated) |
| GET | /tasks/{taskType}/tasks | [get tasks](#get-tasks) | List all tasks for the given task type |
| GET | /tasks/tasks/{taskType} | [get tasks deprecated](#get-tasks-deprecated) | List all tasks for the given task type (deprecated) |
| GET | /tasks/tasktypes | [list task types](#list-task-types) | List all task types |
| PUT | /tasks/{taskType}/resume | [resume tasks](#resume-tasks) | Resume all stopped tasks (as well as the task queue) for the given task type |
| POST | /tasks/schedule | [schedule tasks](#schedule-tasks) | Schedule tasks and return a map from task type to task name scheduled |
| PUT | /tasks/scheduletasks | [schedule tasks deprecated](#schedule-tasks-deprecated) | Schedule tasks (deprecated) |
| PUT | /tasks/{taskType}/stop | [stop tasks](#stop-tasks) | Stop all running/pending tasks (as well as the task queue) for the given task type |
| PUT | /tasks/taskqueue/{taskType} | [toggle task queue state](#toggle-task-queue-state) | Stop/resume a task queue (deprecated) |
  


###  tenant

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /tenants/{tenantName}/metadata | [change tenant state](#change-tenant-state) | Change tenant state |
| POST | /tenants | [create tenant](#create-tenant) |  Create a tenant |
| DELETE | /tenants/{tenantName} | [delete tenant](#delete-tenant) | Delete a tenant |
| GET | /tenants | [get all tenants](#get-all-tenants) | List all tenants |
| GET | /tenants/{tenantName}/tables | [get tables on tenant](#get-tables-on-tenant) | List tables on a a server tenant |
| GET | /tenants/{tenantName}/metadata | [get tenant metadata](#get-tenant-metadata) | Get tenant information |
| GET | /tenants/{tenantName} | [list instance or toggle tenant state](#list-instance-or-toggle-tenant-state) | List instance for a tenant, or enable/disable/drop a tenant |
| PUT | /tenants | [update tenant](#update-tenant) | Update a tenant |
  


###  version

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /version | [get version number](#get-version-number) | Get version number of Pinot components |
  


###  zookeeper

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /zk/delete | [delete](#delete) | Delete the znode at this path |
| GET | /zk/get | [get data](#get-data) | Get content of the znode |
| GET | /zk/ls | [ls](#ls) | List the child znodes |
| GET | /zk/lsl | [lsl](#lsl) | List the child znodes along with Stats |
| PUT | /zk/put | [put data](#put-data) | Update the content of the node |
| GET | /zk/stat | [stat](#stat) | Get the stat |
  


## Paths

### <span id="add-instance"></span> Create a new instance (*addInstance*)

```
POST /instances
```

Creates a new instance with given instance config

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [AddInstanceBody](#add-instance-body) | `AddInstanceBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#add-instance-200) | OK | Success |  | [schema](#add-instance-200-schema) |
| [409](#add-instance-409) | Conflict | Instance already exists |  | [schema](#add-instance-409-schema) |
| [500](#add-instance-500) | Internal Server Error | Internal error |  | [schema](#add-instance-500-schema) |

#### Responses


##### <span id="add-instance-200"></span> 200 - Success
Status: OK

###### <span id="add-instance-200-schema"></span> Schema

##### <span id="add-instance-409"></span> 409 - Instance already exists
Status: Conflict

###### <span id="add-instance-409-schema"></span> Schema

##### <span id="add-instance-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="add-instance-500-schema"></span> Schema

###### Inlined models

**<span id="add-instance-body"></span> AddInstanceBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| grpcPort | int32 (formatted integer)| `int32` |  | |  |  |
| host | string| `string` | ✓ | |  |  |
| pools | map of int32 (formatted integer)| `map[string]int32` |  | |  |  |
| port | int32 (formatted integer)| `int32` | ✓ | |  |  |
| tags | []string| `[]string` |  | |  |  |
| type | string| `string` | ✓ | |  |  |



### <span id="add-schema"></span> Add a new schema (*addSchema*)

```
POST /schemas
```

Adds a new schema

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| override | `query` | boolean | `bool` |  |  | `true` | Whether to override the schema if the schema exists |
| body | `body` | [AddSchemaBody](#add-schema-body) | `AddSchemaBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#add-schema-200) | OK | Successfully created schema |  | [schema](#add-schema-200-schema) |
| [400](#add-schema-400) | Bad Request | Missing or invalid request body |  | [schema](#add-schema-400-schema) |
| [404](#add-schema-404) | Not Found | Schema not found |  | [schema](#add-schema-404-schema) |
| [500](#add-schema-500) | Internal Server Error | Internal error |  | [schema](#add-schema-500-schema) |

#### Responses


##### <span id="add-schema-200"></span> 200 - Successfully created schema
Status: OK

###### <span id="add-schema-200-schema"></span> Schema

##### <span id="add-schema-400"></span> 400 - Missing or invalid request body
Status: Bad Request

###### <span id="add-schema-400-schema"></span> Schema

##### <span id="add-schema-404"></span> 404 - Schema not found
Status: Not Found

###### <span id="add-schema-404-schema"></span> Schema

##### <span id="add-schema-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="add-schema-500-schema"></span> Schema

###### Inlined models

**<span id="add-schema-body"></span> AddSchemaBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dateTimeFieldSpecs | [][AddSchemaParamsBodyDateTimeFieldSpecsItems0](#add-schema-params-body-date-time-field-specs-items0)| `[]*AddSchemaParamsBodyDateTimeFieldSpecsItems0` |  | |  |  |
| dimensionFieldSpecs | [][AddSchemaParamsBodyDimensionFieldSpecsItems0](#add-schema-params-body-dimension-field-specs-items0)| `[]*AddSchemaParamsBodyDimensionFieldSpecsItems0` |  | |  |  |
| metricFieldSpecs | [][AddSchemaParamsBodyMetricFieldSpecsItems0](#add-schema-params-body-metric-field-specs-items0)| `[]*AddSchemaParamsBodyMetricFieldSpecsItems0` |  | |  |  |
| primaryKeyColumns | []string| `[]string` |  | |  |  |
| schemaName | string| `string` |  | |  |  |
| timeFieldSpec | [AddSchemaParamsBodyTimeFieldSpec](#add-schema-params-body-time-field-spec)| `AddSchemaParamsBodyTimeFieldSpec` |  | |  |  |



**<span id="add-schema-params-body-date-time-field-specs-items0"></span> AddSchemaParamsBodyDateTimeFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| format | string| `string` |  | |  |  |
| granularity | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="add-schema-params-body-dimension-field-specs-items0"></span> AddSchemaParamsBodyDimensionFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="add-schema-params-body-metric-field-specs-items0"></span> AddSchemaParamsBodyMetricFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="add-schema-params-body-time-field-spec"></span> AddSchemaParamsBodyTimeFieldSpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| incomingGranularitySpec | [AddSchemaParamsBodyTimeFieldSpecIncomingGranularitySpec](#add-schema-params-body-time-field-spec-incoming-granularity-spec)| `AddSchemaParamsBodyTimeFieldSpecIncomingGranularitySpec` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| outgoingGranularitySpec | [AddSchemaParamsBodyTimeFieldSpecOutgoingGranularitySpec](#add-schema-params-body-time-field-spec-outgoing-granularity-spec)| `AddSchemaParamsBodyTimeFieldSpecOutgoingGranularitySpec` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="add-schema-params-body-time-field-spec-incoming-granularity-spec"></span> AddSchemaParamsBodyTimeFieldSpecIncomingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



**<span id="add-schema-params-body-time-field-spec-outgoing-granularity-spec"></span> AddSchemaParamsBodyTimeFieldSpecOutgoingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="add-table"></span> Adds a table (*addTable*)

```
POST /tables
```

Adds a table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#add-table-200) | OK | successful operation |  | [schema](#add-table-200-schema) |

#### Responses


##### <span id="add-table-200"></span> 200 - successful operation
Status: OK

###### <span id="add-table-200-schema"></span> Schema
   
  

[AddTableOKBody](#add-table-o-k-body)

###### Inlined models

**<span id="add-table-o-k-body"></span> AddTableOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="alter-table-state-or-list-table-config"></span> Get/Enable/Disable/Drop a table (*alterTableStateOrListTableConfig*)

```
GET /tables/{tableName}
```

Get/Enable/Disable/Drop a table. If table name is the only parameter specified , the tableconfig will be printed

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | enable|disable|drop |
| type | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#alter-table-state-or-list-table-config-200) | OK | successful operation |  | [schema](#alter-table-state-or-list-table-config-200-schema) |

#### Responses


##### <span id="alter-table-state-or-list-table-config-200"></span> 200 - successful operation
Status: OK

###### <span id="alter-table-state-or-list-table-config-200-schema"></span> Schema
   
  



### <span id="assign-instances"></span> Assign server instances to a table (*assignInstances*)

```
POST /tables/{tableName}/assignInstances
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| dryRun | `query` | boolean | `bool` |  |  |  | Whether to do dry-run |
| type | `query` | string | `string` |  |  |  | OFFLINE|CONSUMING|COMPLETED |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#assign-instances-200) | OK | successful operation |  | [schema](#assign-instances-200-schema) |

#### Responses


##### <span id="assign-instances-200"></span> 200 - successful operation
Status: OK

###### <span id="assign-instances-200-schema"></span> Schema
   
  

map of [AssignInstancesOKBodyAnon](#assign-instances-o-k-body-anon)

###### Inlined models

**<span id="assign-instances-o-k-body-anon"></span> AssignInstancesOKBodyAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="change-tenant-state"></span> Change tenant state (*changeTenantState*)

```
POST /tenants/{tenantName}/metadata
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Tenant name |
| state | `query` | string | `string` |  | ✓ |  | state |
| type | `query` | string | `string` |  |  |  | tenant type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#change-tenant-state-200) | OK | Success |  | [schema](#change-tenant-state-200-schema) |
| [404](#change-tenant-state-404) | Not Found | Tenant not found |  | [schema](#change-tenant-state-404-schema) |
| [500](#change-tenant-state-500) | Internal Server Error | Server error reading tenant information |  | [schema](#change-tenant-state-500-schema) |

#### Responses


##### <span id="change-tenant-state-200"></span> 200 - Success
Status: OK

###### <span id="change-tenant-state-200-schema"></span> Schema
   
  



##### <span id="change-tenant-state-404"></span> 404 - Tenant not found
Status: Not Found

###### <span id="change-tenant-state-404-schema"></span> Schema

##### <span id="change-tenant-state-500"></span> 500 - Server error reading tenant information
Status: Internal Server Error

###### <span id="change-tenant-state-500-schema"></span> Schema

### <span id="check-health"></span> Check controller health (*checkHealth*)

```
GET /health
```

#### Produces
  * text/plain

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-health-200) | OK | Good |  | [schema](#check-health-200-schema) |

#### Responses


##### <span id="check-health-200"></span> 200 - Good
Status: OK

###### <span id="check-health-200-schema"></span> Schema

### <span id="check-health-legacy"></span> Check controller health (*checkHealthLegacy*)

```
GET /pinot-controller/admin
```

#### Produces
  * text/plain

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-health-legacy-200) | OK | Good |  | [schema](#check-health-legacy-200-schema) |

#### Responses


##### <span id="check-health-legacy-200"></span> 200 - Good
Status: OK

###### <span id="check-health-legacy-200-schema"></span> Schema

### <span id="check-table-config"></span> Validate table config for a table (*checkTableConfig*)

```
POST /tables/validate
```

This API returns the table config that matches the one you get from 'GET /tables/{tableName}'. This allows us to validate table config before apply.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-table-config-200) | OK | successful operation |  | [schema](#check-table-config-200-schema) |

#### Responses


##### <span id="check-table-config-200"></span> 200 - successful operation
Status: OK

###### <span id="check-table-config-200-schema"></span> Schema
   
  



### <span id="clean-up-tasks"></span> Clean up finished tasks (COMPLETED, FAILED) for the given task type (*cleanUpTasks*)

```
PUT /tasks/{taskType}/cleanup
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#clean-up-tasks-200) | OK | successful operation |  | [schema](#clean-up-tasks-200-schema) |

#### Responses


##### <span id="clean-up-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="clean-up-tasks-200-schema"></span> Schema
   
  

[CleanUpTasksOKBody](#clean-up-tasks-o-k-body)

###### Inlined models

**<span id="clean-up-tasks-o-k-body"></span> CleanUpTasksOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="clean-up-tasks-deprecated"></span> Clean up finished tasks (COMPLETED, FAILED) for the given task type (deprecated) (*cleanUpTasksDeprecated*)

```
PUT /tasks/cleanuptasks/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#clean-up-tasks-deprecated-200) | OK | successful operation |  | [schema](#clean-up-tasks-deprecated-200-schema) |

#### Responses


##### <span id="clean-up-tasks-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="clean-up-tasks-deprecated-200-schema"></span> Schema
   
  

[CleanUpTasksDeprecatedOKBody](#clean-up-tasks-deprecated-o-k-body)

###### Inlined models

**<span id="clean-up-tasks-deprecated-o-k-body"></span> CleanUpTasksDeprecatedOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="create-tenant"></span> Create a tenant (*createTenant*)

```
POST /tenants
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [CreateTenantBody](#create-tenant-body) | `CreateTenantBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-tenant-200) | OK | Success |  | [schema](#create-tenant-200-schema) |
| [500](#create-tenant-500) | Internal Server Error | Error creating tenant |  | [schema](#create-tenant-500-schema) |

#### Responses


##### <span id="create-tenant-200"></span> 200 - Success
Status: OK

###### <span id="create-tenant-200-schema"></span> Schema

##### <span id="create-tenant-500"></span> 500 - Error creating tenant
Status: Internal Server Error

###### <span id="create-tenant-500-schema"></span> Schema

###### Inlined models

**<span id="create-tenant-body"></span> CreateTenantBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| numberOfInstances | int32 (formatted integer)| `int32` |  | |  |  |
| offlineInstances | int32 (formatted integer)| `int32` |  | |  |  |
| realtimeInstances | int32 (formatted integer)| `int32` |  | |  |  |
| tenantName | string| `string` | ✓ | |  |  |
| tenantRole | string| `string` | ✓ | |  |  |



### <span id="delete"></span> Delete the znode at this path (*delete*)

```
DELETE /zk/delete
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-200) | OK | Success |  | [schema](#delete-200-schema) |
| [204](#delete-204) | No Content | No Content |  | [schema](#delete-204-schema) |
| [404](#delete-404) | Not Found | ZK Path not found |  | [schema](#delete-404-schema) |
| [500](#delete-500) | Internal Server Error | Internal server error |  | [schema](#delete-500-schema) |

#### Responses


##### <span id="delete-200"></span> 200 - Success
Status: OK

###### <span id="delete-200-schema"></span> Schema

##### <span id="delete-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-204-schema"></span> Schema

##### <span id="delete-404"></span> 404 - ZK Path not found
Status: Not Found

###### <span id="delete-404-schema"></span> Schema

##### <span id="delete-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="delete-500-schema"></span> Schema

### <span id="delete-all-segments"></span> Delete all segments (*deleteAllSegments*)

```
DELETE /segments/{tableName}
```

Delete all segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  | ✓ |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-all-segments-200) | OK | successful operation |  | [schema](#delete-all-segments-200-schema) |

#### Responses


##### <span id="delete-all-segments-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-all-segments-200-schema"></span> Schema
   
  

[DeleteAllSegmentsOKBody](#delete-all-segments-o-k-body)

###### Inlined models

**<span id="delete-all-segments-o-k-body"></span> DeleteAllSegmentsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-cluster-config"></span> Delete cluster configuration (*deleteClusterConfig*)

```
DELETE /cluster/configs/{configName}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| configName | `path` | string | `string` |  | ✓ |  | Name of the config to delete |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-cluster-config-200) | OK | Success |  | [schema](#delete-cluster-config-200-schema) |
| [500](#delete-cluster-config-500) | Internal Server Error | Server error deleting configuration |  | [schema](#delete-cluster-config-500-schema) |

#### Responses


##### <span id="delete-cluster-config-200"></span> 200 - Success
Status: OK

###### <span id="delete-cluster-config-200-schema"></span> Schema

##### <span id="delete-cluster-config-500"></span> 500 - Server error deleting configuration
Status: Internal Server Error

###### <span id="delete-cluster-config-500-schema"></span> Schema

### <span id="delete-schema"></span> Delete a schema (*deleteSchema*)

```
DELETE /schemas/{schemaName}
```

Deletes a schema by name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| schemaName | `path` | string | `string` |  | ✓ |  | Schema name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-schema-200) | OK | Successfully deleted schema |  | [schema](#delete-schema-200-schema) |
| [404](#delete-schema-404) | Not Found | Schema not found |  | [schema](#delete-schema-404-schema) |
| [409](#delete-schema-409) | Conflict | Schema is in use |  | [schema](#delete-schema-409-schema) |
| [500](#delete-schema-500) | Internal Server Error | Error deleting schema |  | [schema](#delete-schema-500-schema) |

#### Responses


##### <span id="delete-schema-200"></span> 200 - Successfully deleted schema
Status: OK

###### <span id="delete-schema-200-schema"></span> Schema

##### <span id="delete-schema-404"></span> 404 - Schema not found
Status: Not Found

###### <span id="delete-schema-404-schema"></span> Schema

##### <span id="delete-schema-409"></span> 409 - Schema is in use
Status: Conflict

###### <span id="delete-schema-409-schema"></span> Schema

##### <span id="delete-schema-500"></span> 500 - Error deleting schema
Status: Internal Server Error

###### <span id="delete-schema-500-schema"></span> Schema

### <span id="delete-segment"></span> Delete a segment (*deleteSegment*)

```
DELETE /segments/{tableName}/{segmentName}
```

Delete a segment

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-segment-200) | OK | successful operation |  | [schema](#delete-segment-200-schema) |

#### Responses


##### <span id="delete-segment-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-segment-200-schema"></span> Schema
   
  

[DeleteSegmentOKBody](#delete-segment-o-k-body)

###### Inlined models

**<span id="delete-segment-o-k-body"></span> DeleteSegmentOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-segments"></span> Delete the segments in the JSON array payload (*deleteSegments*)

```
POST /segments/{tableName}/delete
```

Delete the segments in the JSON array payload

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| body | `body` | []string | `[]string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-segments-200) | OK | successful operation |  | [schema](#delete-segments-200-schema) |

#### Responses


##### <span id="delete-segments-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-segments-200-schema"></span> Schema
   
  

[DeleteSegmentsOKBody](#delete-segments-o-k-body)

###### Inlined models

**<span id="delete-segments-o-k-body"></span> DeleteSegmentsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-table"></span> Deletes a table (*deleteTable*)

```
DELETE /tables/{tableName}
```

Deletes a table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table to delete |
| type | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-table-200) | OK | successful operation |  | [schema](#delete-table-200-schema) |

#### Responses


##### <span id="delete-table-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-table-200-schema"></span> Schema
   
  

[DeleteTableOKBody](#delete-table-o-k-body)

###### Inlined models

**<span id="delete-table-o-k-body"></span> DeleteTableOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-task-queue"></span> Delete a task queue (deprecated) (*deleteTaskQueue*)

```
DELETE /tasks/taskqueue/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |
| forceDelete | `query` | boolean | `bool` |  |  |  | Whether to force delete the task queue (expert only option, enable with cautious |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-task-queue-200) | OK | successful operation |  | [schema](#delete-task-queue-200-schema) |

#### Responses


##### <span id="delete-task-queue-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-task-queue-200-schema"></span> Schema
   
  

[DeleteTaskQueueOKBody](#delete-task-queue-o-k-body)

###### Inlined models

**<span id="delete-task-queue-o-k-body"></span> DeleteTaskQueueOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-tasks"></span> Delete all tasks (as well as the task queue) for the given task type (*deleteTasks*)

```
DELETE /tasks/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |
| forceDelete | `query` | boolean | `bool` |  |  |  | Whether to force deleting the tasks (expert only option, enable with cautious |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-tasks-200) | OK | successful operation |  | [schema](#delete-tasks-200-schema) |

#### Responses


##### <span id="delete-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="delete-tasks-200-schema"></span> Schema
   
  

[DeleteTasksOKBody](#delete-tasks-o-k-body)

###### Inlined models

**<span id="delete-tasks-o-k-body"></span> DeleteTasksOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="delete-tenant"></span> Delete a tenant (*deleteTenant*)

```
DELETE /tenants/{tenantName}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Tenant name |
| type | `query` | string | `string` |  | ✓ |  | Tenant type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#delete-tenant-200) | OK | Success |  | [schema](#delete-tenant-200-schema) |
| [400](#delete-tenant-400) | Bad Request | Tenant can not be deleted |  | [schema](#delete-tenant-400-schema) |
| [404](#delete-tenant-404) | Not Found | Tenant not found |  | [schema](#delete-tenant-404-schema) |
| [500](#delete-tenant-500) | Internal Server Error | Error deleting tenant |  | [schema](#delete-tenant-500-schema) |

#### Responses


##### <span id="delete-tenant-200"></span> 200 - Success
Status: OK

###### <span id="delete-tenant-200-schema"></span> Schema

##### <span id="delete-tenant-400"></span> 400 - Tenant can not be deleted
Status: Bad Request

###### <span id="delete-tenant-400-schema"></span> Schema

##### <span id="delete-tenant-404"></span> 404 - Tenant not found
Status: Not Found

###### <span id="delete-tenant-404-schema"></span> Schema

##### <span id="delete-tenant-500"></span> 500 - Error deleting tenant
Status: Internal Server Error

###### <span id="delete-tenant-500-schema"></span> Schema

### <span id="download-segment"></span> Download a segment (*downloadSegment*)

```
GET /segments/{tableName}/{segmentName}
```

Download a segment

#### Produces
  * application/octet-stream

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#download-segment-default) | | successful operation |  | [schema](#download-segment-default-schema) |

#### Responses


##### <span id="download-segment-default"></span> Default Response
successful operation

###### <span id="download-segment-default-schema"></span> Schema
empty schema

### <span id="drop-instance"></span> Drop an instance (*dropInstance*)

```
DELETE /instances/{instanceName}
```

Drop an instance

#### Consumes
  * text/plain

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Instance name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#drop-instance-200) | OK | Success |  | [schema](#drop-instance-200-schema) |
| [404](#drop-instance-404) | Not Found | Instance not found |  | [schema](#drop-instance-404-schema) |
| [409](#drop-instance-409) | Conflict | Instance cannot be dropped |  | [schema](#drop-instance-409-schema) |
| [500](#drop-instance-500) | Internal Server Error | Internal error |  | [schema](#drop-instance-500-schema) |

#### Responses


##### <span id="drop-instance-200"></span> 200 - Success
Status: OK

###### <span id="drop-instance-200-schema"></span> Schema

##### <span id="drop-instance-404"></span> 404 - Instance not found
Status: Not Found

###### <span id="drop-instance-404-schema"></span> Schema

##### <span id="drop-instance-409"></span> 409 - Instance cannot be dropped
Status: Conflict

###### <span id="drop-instance-409-schema"></span> Schema

##### <span id="drop-instance-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="drop-instance-500-schema"></span> Schema

### <span id="end-replace-segments"></span> End to replace segments (*endReplaceSegments*)

```
POST /segments/{tableName}/endReplaceSegments
```

End to replace segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| segmentLineageEntryId | `query` | string | `string` |  |  |  | Segment lineage entry id returned by startReplaceSegments API |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#end-replace-segments-default) | | successful operation |  | [schema](#end-replace-segments-default-schema) |

#### Responses


##### <span id="end-replace-segments-default"></span> Default Response
successful operation

###### <span id="end-replace-segments-default-schema"></span> Schema
empty schema

### <span id="get-all-instances"></span> List all instances (*getAllInstances*)

```
GET /instances
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-instances-200) | OK | Success |  | [schema](#get-all-instances-200-schema) |
| [500](#get-all-instances-500) | Internal Server Error | Internal error |  | [schema](#get-all-instances-500-schema) |

#### Responses


##### <span id="get-all-instances-200"></span> 200 - Success
Status: OK

###### <span id="get-all-instances-200-schema"></span> Schema

##### <span id="get-all-instances-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="get-all-instances-500-schema"></span> Schema

### <span id="get-all-tenants"></span> List all tenants (*getAllTenants*)

```
GET /tenants
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| type | `query` | string | `string` |  |  |  | Tenant type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-tenants-200) | OK | Success |  | [schema](#get-all-tenants-200-schema) |
| [500](#get-all-tenants-500) | Internal Server Error | Error reading tenants list |  | [schema](#get-all-tenants-500-schema) |

#### Responses


##### <span id="get-all-tenants-200"></span> 200 - Success
Status: OK

###### <span id="get-all-tenants-200-schema"></span> Schema

##### <span id="get-all-tenants-500"></span> 500 - Error reading tenants list
Status: Internal Server Error

###### <span id="get-all-tenants-500-schema"></span> Schema

### <span id="get-brokers-for-table"></span> List brokers for a given table (*getBrokersForTable*)

```
GET /brokers/tables/{tableName}
```

List brokers for a given table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-brokers-for-table-200) | OK | successful operation |  | [schema](#get-brokers-for-table-200-schema) |

#### Responses


##### <span id="get-brokers-for-table-200"></span> 200 - successful operation
Status: OK

###### <span id="get-brokers-for-table-200-schema"></span> Schema
   
  

[]string

### <span id="get-brokers-for-table-v2"></span> List brokers for a given table (*getBrokersForTableV2*)

```
GET /v2/brokers/tables/{tableName}
```

List brokers for a given table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-brokers-for-table-v2-200) | OK | successful operation |  | [schema](#get-brokers-for-table-v2-200-schema) |

#### Responses


##### <span id="get-brokers-for-table-v2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-brokers-for-table-v2-200-schema"></span> Schema
   
  

[][GetBrokersForTableV2OKBodyItems0](#get-brokers-for-table-v2-o-k-body-items0)

###### Inlined models

**<span id="get-brokers-for-table-v2-o-k-body-items0"></span> GetBrokersForTableV2OKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="get-brokers-for-tenant"></span> List brokers for a given tenant (*getBrokersForTenant*)

```
GET /brokers/tenants/{tenantName}
```

List brokers for a given tenant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Name of the tenant |
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-brokers-for-tenant-200) | OK | successful operation |  | [schema](#get-brokers-for-tenant-200-schema) |

#### Responses


##### <span id="get-brokers-for-tenant-200"></span> 200 - successful operation
Status: OK

###### <span id="get-brokers-for-tenant-200-schema"></span> Schema
   
  

[]string

### <span id="get-brokers-for-tenant-v2"></span> List brokers for a given tenant (*getBrokersForTenantV2*)

```
GET /v2/brokers/tenants/{tenantName}
```

List brokers for a given tenant

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Name of the tenant |
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-brokers-for-tenant-v2-200) | OK | successful operation |  | [schema](#get-brokers-for-tenant-v2-200-schema) |

#### Responses


##### <span id="get-brokers-for-tenant-v2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-brokers-for-tenant-v2-200-schema"></span> Schema
   
  

[][GetBrokersForTenantV2OKBodyItems0](#get-brokers-for-tenant-v2-o-k-body-items0)

###### Inlined models

**<span id="get-brokers-for-tenant-v2-o-k-body-items0"></span> GetBrokersForTenantV2OKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="get-cluster-info"></span> Get cluster Info (*getClusterInfo*)

```
GET /cluster/info
```

Get cluster Info

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cluster-info-200) | OK | Success |  | [schema](#get-cluster-info-200-schema) |
| [500](#get-cluster-info-500) | Internal Server Error | Internal server error |  | [schema](#get-cluster-info-500-schema) |

#### Responses


##### <span id="get-cluster-info-200"></span> 200 - Success
Status: OK

###### <span id="get-cluster-info-200-schema"></span> Schema

##### <span id="get-cluster-info-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="get-cluster-info-500-schema"></span> Schema

### <span id="get-consuming-segments-info"></span> Returns state of consuming segments (*getConsumingSegmentsInfo*)

```
GET /tables/{realtimeTableName}/consumingSegmentsInfo
```

Gets the status of consumers from all servers

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| realtimeTableName | `path` | string | `string` |  | ✓ |  | Realtime table name with or without type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-consuming-segments-info-200) | OK | Success |  | [schema](#get-consuming-segments-info-200-schema) |
| [404](#get-consuming-segments-info-404) | Not Found | Table not found |  | [schema](#get-consuming-segments-info-404-schema) |
| [500](#get-consuming-segments-info-500) | Internal Server Error | Internal server error |  | [schema](#get-consuming-segments-info-500-schema) |

#### Responses


##### <span id="get-consuming-segments-info-200"></span> 200 - Success
Status: OK

###### <span id="get-consuming-segments-info-200-schema"></span> Schema

##### <span id="get-consuming-segments-info-404"></span> 404 - Table not found
Status: Not Found

###### <span id="get-consuming-segments-info-404-schema"></span> Schema

##### <span id="get-consuming-segments-info-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="get-consuming-segments-info-500-schema"></span> Schema

### <span id="get-cron-scheduler-information"></span> Fetch cron scheduler information (*getCronSchedulerInformation*)

```
GET /tasks/scheduler/information
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cron-scheduler-information-200) | OK | successful operation |  | [schema](#get-cron-scheduler-information-200-schema) |

#### Responses


##### <span id="get-cron-scheduler-information-200"></span> 200 - successful operation
Status: OK

###### <span id="get-cron-scheduler-information-200-schema"></span> Schema
   
  

map of any 

### <span id="get-cron-scheduler-job-details"></span> Fetch cron scheduler job keys (*getCronSchedulerJobDetails*)

```
GET /tasks/scheduler/jobDetails
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `query` | string | `string` |  |  |  | Table name (with type suffix) |
| taskType | `query` | string | `string` |  |  |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cron-scheduler-job-details-200) | OK | successful operation |  | [schema](#get-cron-scheduler-job-details-200-schema) |

#### Responses


##### <span id="get-cron-scheduler-job-details-200"></span> 200 - successful operation
Status: OK

###### <span id="get-cron-scheduler-job-details-200-schema"></span> Schema
   
  

map of any 

### <span id="get-cron-scheduler-job-keys"></span> Fetch cron scheduler job keys (*getCronSchedulerJobKeys*)

```
GET /tasks/scheduler/jobKeys
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cron-scheduler-job-keys-200) | OK | successful operation |  | [schema](#get-cron-scheduler-job-keys-200-schema) |

#### Responses


##### <span id="get-cron-scheduler-job-keys-200"></span> 200 - successful operation
Status: OK

###### <span id="get-cron-scheduler-job-keys-200-schema"></span> Schema
   
  

[][GetCronSchedulerJobKeysOKBodyItems0](#get-cron-scheduler-job-keys-o-k-body-items0)

###### Inlined models

**<span id="get-cron-scheduler-job-keys-o-k-body-items0"></span> GetCronSchedulerJobKeysOKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| group | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="get-data"></span> Get content of the znode (*getData*)

```
GET /zk/get
```

#### Produces
  * text/plain

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-data-200) | OK | Success |  | [schema](#get-data-200-schema) |
| [204](#get-data-204) | No Content | No Content |  | [schema](#get-data-204-schema) |
| [404](#get-data-404) | Not Found | ZK Path not found |  | [schema](#get-data-404-schema) |
| [500](#get-data-500) | Internal Server Error | Internal server error |  | [schema](#get-data-500-schema) |

#### Responses


##### <span id="get-data-200"></span> 200 - Success
Status: OK

###### <span id="get-data-200-schema"></span> Schema

##### <span id="get-data-204"></span> 204 - No Content
Status: No Content

###### <span id="get-data-204-schema"></span> Schema

##### <span id="get-data-404"></span> 404 - ZK Path not found
Status: Not Found

###### <span id="get-data-404-schema"></span> Schema

##### <span id="get-data-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="get-data-500-schema"></span> Schema

### <span id="get-external-view"></span> Get table external view (*getExternalView*)

```
GET /tables/{tableName}/externalview
```

Get table external view

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| tableType | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-external-view-200) | OK | successful operation |  | [schema](#get-external-view-200-schema) |

#### Responses


##### <span id="get-external-view-200"></span> 200 - successful operation
Status: OK

###### <span id="get-external-view-200-schema"></span> Schema
   
  

[GetExternalViewOKBody](#get-external-view-o-k-body)

###### Inlined models

**<span id="get-external-view-o-k-body"></span> GetExternalViewOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| OFFLINE | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |
| REALTIME | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |



### <span id="get-ideal-state"></span> Get table ideal state (*getIdealState*)

```
GET /tables/{tableName}/idealstate
```

Get table ideal state

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| tableType | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-ideal-state-200) | OK | successful operation |  | [schema](#get-ideal-state-200-schema) |

#### Responses


##### <span id="get-ideal-state-200"></span> 200 - successful operation
Status: OK

###### <span id="get-ideal-state-200-schema"></span> Schema
   
  

[GetIdealStateOKBody](#get-ideal-state-o-k-body)

###### Inlined models

**<span id="get-ideal-state-o-k-body"></span> GetIdealStateOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| OFFLINE | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |
| REALTIME | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |



### <span id="get-instance"></span> Get instance information (*getInstance*)

```
GET /instances/{instanceName}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Instance name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-instance-200) | OK | Success |  | [schema](#get-instance-200-schema) |
| [404](#get-instance-404) | Not Found | Instance not found |  | [schema](#get-instance-404-schema) |
| [500](#get-instance-500) | Internal Server Error | Internal error |  | [schema](#get-instance-500-schema) |

#### Responses


##### <span id="get-instance-200"></span> 200 - Success
Status: OK

###### <span id="get-instance-200-schema"></span> Schema

##### <span id="get-instance-404"></span> 404 - Instance not found
Status: Not Found

###### <span id="get-instance-404-schema"></span> Schema

##### <span id="get-instance-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="get-instance-500-schema"></span> Schema

### <span id="get-instance-partitions"></span> Get the instance partitions (*getInstancePartitions*)

```
GET /tables/{tableName}/instancePartitions
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|CONSUMING|COMPLETED |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-instance-partitions-200) | OK | successful operation |  | [schema](#get-instance-partitions-200-schema) |

#### Responses


##### <span id="get-instance-partitions-200"></span> 200 - successful operation
Status: OK

###### <span id="get-instance-partitions-200-schema"></span> Schema
   
  

map of [GetInstancePartitionsOKBodyAnon](#get-instance-partitions-o-k-body-anon)

###### Inlined models

**<span id="get-instance-partitions-o-k-body-anon"></span> GetInstancePartitionsOKBodyAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="get-leader-for-table"></span> Gets leader for a given table (*getLeaderForTable*)

```
GET /leader/tables/{tableName}
```

Gets leader for a given table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-leader-for-table-200) | OK | successful operation |  | [schema](#get-leader-for-table-200-schema) |

#### Responses


##### <span id="get-leader-for-table-200"></span> 200 - successful operation
Status: OK

###### <span id="get-leader-for-table-200-schema"></span> Schema
   
  

[GetLeaderForTableOKBody](#get-leader-for-table-o-k-body)

###### Inlined models

**<span id="get-leader-for-table-o-k-body"></span> GetLeaderForTableOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerEntryMap | map of [GetLeaderForTableOKBodyLeadControllerEntryMapAnon](#get-leader-for-table-o-k-body-lead-controller-entry-map-anon)| `map[string]GetLeaderForTableOKBodyLeadControllerEntryMapAnon` |  | |  |  |
| leadControllerResourceEnabled | boolean| `bool` |  | |  |  |



**<span id="get-leader-for-table-o-k-body-lead-controller-entry-map-anon"></span> GetLeaderForTableOKBodyLeadControllerEntryMapAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerId | string| `string` |  | |  |  |
| tableNames | []string| `[]string` |  | |  |  |



### <span id="get-leaders-for-all-tables"></span> Gets leaders for all tables (*getLeadersForAllTables*)

```
GET /leader/tables
```

Gets leaders for all tables

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-leaders-for-all-tables-200) | OK | successful operation |  | [schema](#get-leaders-for-all-tables-200-schema) |

#### Responses


##### <span id="get-leaders-for-all-tables-200"></span> 200 - successful operation
Status: OK

###### <span id="get-leaders-for-all-tables-200-schema"></span> Schema
   
  

[GetLeadersForAllTablesOKBody](#get-leaders-for-all-tables-o-k-body)

###### Inlined models

**<span id="get-leaders-for-all-tables-o-k-body"></span> GetLeadersForAllTablesOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerEntryMap | map of [GetLeadersForAllTablesOKBodyLeadControllerEntryMapAnon](#get-leaders-for-all-tables-o-k-body-lead-controller-entry-map-anon)| `map[string]GetLeadersForAllTablesOKBodyLeadControllerEntryMapAnon` |  | |  |  |
| leadControllerResourceEnabled | boolean| `bool` |  | |  |  |



**<span id="get-leaders-for-all-tables-o-k-body-lead-controller-entry-map-anon"></span> GetLeadersForAllTablesOKBodyLeadControllerEntryMapAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerId | string| `string` |  | |  |  |
| tableNames | []string| `[]string` |  | |  |  |



### <span id="get-schema"></span> Get a schema (*getSchema*)

```
GET /schemas/{schemaName}
```

Gets a schema by name

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| schemaName | `path` | string | `string` |  | ✓ |  | Schema name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-schema-200) | OK | Success |  | [schema](#get-schema-200-schema) |
| [404](#get-schema-404) | Not Found | Schema not found |  | [schema](#get-schema-404-schema) |
| [500](#get-schema-500) | Internal Server Error | Internal error |  | [schema](#get-schema-500-schema) |

#### Responses


##### <span id="get-schema-200"></span> 200 - Success
Status: OK

###### <span id="get-schema-200-schema"></span> Schema

##### <span id="get-schema-404"></span> 404 - Schema not found
Status: Not Found

###### <span id="get-schema-404-schema"></span> Schema

##### <span id="get-schema-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="get-schema-500-schema"></span> Schema

### <span id="get-segment-metadata"></span> Get the metadata for a segment (*getSegmentMetadata*)

```
GET /segments/{tableName}/{segmentName}/metadata
```

Get the metadata for a segment

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-metadata-200) | OK | successful operation |  | [schema](#get-segment-metadata-200-schema) |

#### Responses


##### <span id="get-segment-metadata-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segment-metadata-200-schema"></span> Schema
   
  

map of string

### <span id="get-segment-metadata-deprecated1"></span> Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead) (*getSegmentMetadataDeprecated1*)

```
GET /tables/{tableName}/segments/{segmentName}/metadata
```

Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-metadata-deprecated1-200) | OK | successful operation |  | [schema](#get-segment-metadata-deprecated1-200-schema) |

#### Responses


##### <span id="get-segment-metadata-deprecated1-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segment-metadata-deprecated1-200-schema"></span> Schema
   
  

[][[]map[string]interface{}](#map-string-interface)

### <span id="get-segment-metadata-deprecated2"></span> Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead) (*getSegmentMetadataDeprecated2*)

```
GET /tables/{tableName}/segments/{segmentName}
```

Get the metadata for a segment (deprecated, use 'GET /segments/{tableName}/{segmentName}/metadata' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | MUST be null |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-metadata-deprecated2-200) | OK | successful operation |  | [schema](#get-segment-metadata-deprecated2-200-schema) |

#### Responses


##### <span id="get-segment-metadata-deprecated2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segment-metadata-deprecated2-200-schema"></span> Schema
   
  

[][[]map[string]interface{}](#map-string-interface)

### <span id="get-segment-to-crc-map"></span> Get a map from segment to CRC of the segment (only apply to OFFLINE table) (*getSegmentToCrcMap*)

```
GET /segments/{tableName}/crc
```

Get a map from segment to CRC of the segment (only apply to OFFLINE table)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-to-crc-map-200) | OK | successful operation |  | [schema](#get-segment-to-crc-map-200-schema) |

#### Responses


##### <span id="get-segment-to-crc-map-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segment-to-crc-map-200-schema"></span> Schema
   
  

map of string

### <span id="get-segment-to-crc-map-deprecated"></span> Get a map from segment to CRC of the segment (deprecated, use 'GET /segments/{tableName}/crc' instead) (*getSegmentToCrcMapDeprecated*)

```
GET /tables/{tableName}/segments/crc
```

Get a map from segment to CRC of the segment (deprecated, use 'GET /segments/{tableName}/crc' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-to-crc-map-deprecated-200) | OK | successful operation |  | [schema](#get-segment-to-crc-map-deprecated-200-schema) |

#### Responses


##### <span id="get-segment-to-crc-map-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segment-to-crc-map-deprecated-200-schema"></span> Schema
   
  

map of string

### <span id="get-segments"></span> List all segments (*getSegments*)

```
GET /segments/{tableName}
```

List all segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segments-200) | OK | successful operation |  | [schema](#get-segments-200-schema) |

#### Responses


##### <span id="get-segments-200"></span> 200 - successful operation
Status: OK

###### <span id="get-segments-200-schema"></span> Schema
   
  

[][map[string][]string](#map-string-string)

### <span id="get-server-metadata"></span> Get the server metadata for all table segments (*getServerMetadata*)

```
GET /segments/{tableName}/metadata
```

Get the server metadata for all table segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-server-metadata-200) | OK | successful operation |  | [schema](#get-server-metadata-200-schema) |

#### Responses


##### <span id="get-server-metadata-200"></span> 200 - successful operation
Status: OK

###### <span id="get-server-metadata-200-schema"></span> Schema
   
  



### <span id="get-server-to-segments-map"></span> Get a map from server to segments hosted by the server (*getServerToSegmentsMap*)

```
GET /segments/{tableName}/servers
```

Get a map from server to segments hosted by the server

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-server-to-segments-map-200) | OK | successful operation |  | [schema](#get-server-to-segments-map-200-schema) |

#### Responses


##### <span id="get-server-to-segments-map-200"></span> 200 - successful operation
Status: OK

###### <span id="get-server-to-segments-map-200-schema"></span> Schema
   
  

[][map[string]interface{}](#map-string-interface)

### <span id="get-server-to-segments-map-deprecated1"></span> Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead) (*getServerToSegmentsMapDeprecated1*)

```
GET /tables/{tableName}/segments
```

Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | MUST be null |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-server-to-segments-map-deprecated1-200) | OK | successful operation |  | [schema](#get-server-to-segments-map-deprecated1-200-schema) |

#### Responses


##### <span id="get-server-to-segments-map-deprecated1-200"></span> 200 - successful operation
Status: OK

###### <span id="get-server-to-segments-map-deprecated1-200-schema"></span> Schema
   
  

[][map[string]string](#map-string-string)

### <span id="get-server-to-segments-map-deprecated2"></span> Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead) (*getServerToSegmentsMapDeprecated2*)

```
GET /tables/{tableName}/segments/metadata
```

Get a map from server to segments hosted by the server (deprecated, use 'GET /segments/{tableName}/servers' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| state | `query` | string | `string` |  |  |  | MUST be null |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-server-to-segments-map-deprecated2-200) | OK | successful operation |  | [schema](#get-server-to-segments-map-deprecated2-200-schema) |

#### Responses


##### <span id="get-server-to-segments-map-deprecated2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-server-to-segments-map-deprecated2-200-schema"></span> Schema
   
  

[][map[string]string](#map-string-string)

### <span id="get-table-instances"></span> List table instances (*getTableInstances*)

```
GET /tables/{tableName}/instances
```

List instances of the given table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name without type |
| type | `query` | string | `string` |  |  |  | Instance type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-table-instances-200) | OK | Success |  | [schema](#get-table-instances-200-schema) |
| [404](#get-table-instances-404) | Not Found | Table not found |  | [schema](#get-table-instances-404-schema) |
| [500](#get-table-instances-500) | Internal Server Error | Internal server error |  | [schema](#get-table-instances-500-schema) |

#### Responses


##### <span id="get-table-instances-200"></span> 200 - Success
Status: OK

###### <span id="get-table-instances-200-schema"></span> Schema

##### <span id="get-table-instances-404"></span> 404 - Table not found
Status: Not Found

###### <span id="get-table-instances-404-schema"></span> Schema

##### <span id="get-table-instances-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="get-table-instances-500-schema"></span> Schema

### <span id="get-table-schema"></span> Get table schema (*getTableSchema*)

```
GET /tables/{tableName}/schema
```

Read table schema

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name (without type) |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-table-schema-200) | OK | Success |  | [schema](#get-table-schema-200-schema) |
| [404](#get-table-schema-404) | Not Found | Table not found |  | [schema](#get-table-schema-404-schema) |

#### Responses


##### <span id="get-table-schema-200"></span> 200 - Success
Status: OK

###### <span id="get-table-schema-200-schema"></span> Schema

##### <span id="get-table-schema-404"></span> 404 - Table not found
Status: Not Found

###### <span id="get-table-schema-404-schema"></span> Schema

### <span id="get-table-size"></span> Read table sizes (*getTableSize*)

```
GET /tables/{tableName}/size
```

Get table size details. Table size is the size of untarred segments including replication

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name without type |
| detailed | `query` | boolean | `bool` |  |  | `true` | Get detailed information |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-table-size-200) | OK | Success |  | [schema](#get-table-size-200-schema) |
| [404](#get-table-size-404) | Not Found | Table not found |  | [schema](#get-table-size-404-schema) |
| [500](#get-table-size-500) | Internal Server Error | Internal server error |  | [schema](#get-table-size-500-schema) |

#### Responses


##### <span id="get-table-size-200"></span> 200 - Success
Status: OK

###### <span id="get-table-size-200-schema"></span> Schema

##### <span id="get-table-size-404"></span> 404 - Table not found
Status: Not Found

###### <span id="get-table-size-404-schema"></span> Schema

##### <span id="get-table-size-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="get-table-size-500-schema"></span> Schema

### <span id="get-table-state"></span> Get current table state (*getTableState*)

```
GET /tables/{tableName}/state
```

Get current table state

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table to get its state |
| type | `query` | string | `string` |  | ✓ |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-table-state-200) | OK | successful operation |  | [schema](#get-table-state-200-schema) |

#### Responses


##### <span id="get-table-state-200"></span> 200 - successful operation
Status: OK

###### <span id="get-table-state-200-schema"></span> Schema
   
  



### <span id="get-table-stats"></span> table stats (*getTableStats*)

```
GET /tables/{tableName}/stats
```

Provides metadata info/stats about the table.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-table-stats-200) | OK | successful operation |  | [schema](#get-table-stats-200-schema) |

#### Responses


##### <span id="get-table-stats-200"></span> 200 - successful operation
Status: OK

###### <span id="get-table-stats-200-schema"></span> Schema
   
  



### <span id="get-tables-on-tenant"></span> List tables on a a server tenant (*getTablesOnTenant*)

```
GET /tenants/{tenantName}/tables
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Tenant name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tables-on-tenant-200) | OK | Success |  | [schema](#get-tables-on-tenant-200-schema) |
| [500](#get-tables-on-tenant-500) | Internal Server Error | Error reading list |  | [schema](#get-tables-on-tenant-500-schema) |

#### Responses


##### <span id="get-tables-on-tenant-200"></span> 200 - Success
Status: OK

###### <span id="get-tables-on-tenant-200-schema"></span> Schema

##### <span id="get-tables-on-tenant-500"></span> 500 - Error reading list
Status: Internal Server Error

###### <span id="get-tables-on-tenant-500-schema"></span> Schema

### <span id="get-tables-to-brokers-mapping"></span> List tables to brokers mappings (*getTablesToBrokersMapping*)

```
GET /brokers/tables
```

List tables to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tables-to-brokers-mapping-200) | OK | successful operation |  | [schema](#get-tables-to-brokers-mapping-200-schema) |

#### Responses


##### <span id="get-tables-to-brokers-mapping-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tables-to-brokers-mapping-200-schema"></span> Schema
   
  

map of [[]string](#string)

### <span id="get-tables-to-brokers-mapping-v2"></span> List tables to brokers mappings (*getTablesToBrokersMappingV2*)

```
GET /v2/brokers/tables
```

List tables to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tables-to-brokers-mapping-v2-200) | OK | successful operation |  | [schema](#get-tables-to-brokers-mapping-v2-200-schema) |

#### Responses


##### <span id="get-tables-to-brokers-mapping-v2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tables-to-brokers-mapping-v2-200-schema"></span> Schema
   
  

map of [[]*GetTablesToBrokersMappingV2OKBodyItems0](#get-tables-to-brokers-mapping-v2-o-k-body-items0)

###### Inlined models

**<span id="get-tables-to-brokers-mapping-v2-o-k-body-items0"></span> GetTablesToBrokersMappingV2OKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="get-task-configs"></span> Get the task config (a list of child task configs) for the given task (*getTaskConfigs*)

```
GET /tasks/task/{taskName}/config
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskName | `path` | string | `string` |  | ✓ |  | Task name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-configs-200) | OK | successful operation |  | [schema](#get-task-configs-200-schema) |

#### Responses


##### <span id="get-task-configs-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-configs-200-schema"></span> Schema
   
  

[][GetTaskConfigsOKBodyItems0](#get-task-configs-o-k-body-items0)

###### Inlined models

**<span id="get-task-configs-o-k-body-items0"></span> GetTaskConfigsOKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| configs | map of string| `map[string]string` |  | |  |  |
| taskType | string| `string` |  | |  |  |



### <span id="get-task-configs-deprecated"></span> Get the task config (a list of child task configs) for the given task (deprecated) (*getTaskConfigsDeprecated*)

```
GET /tasks/taskconfig/{taskName}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskName | `path` | string | `string` |  | ✓ |  | Task name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-configs-deprecated-200) | OK | successful operation |  | [schema](#get-task-configs-deprecated-200-schema) |

#### Responses


##### <span id="get-task-configs-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-configs-deprecated-200-schema"></span> Schema
   
  

[][GetTaskConfigsDeprecatedOKBodyItems0](#get-task-configs-deprecated-o-k-body-items0)

###### Inlined models

**<span id="get-task-configs-deprecated-o-k-body-items0"></span> GetTaskConfigsDeprecatedOKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| configs | map of string| `map[string]string` |  | |  |  |
| taskType | string| `string` |  | |  |  |



### <span id="get-task-queue-state"></span> Get the state (task queue state) for the given task type (*getTaskQueueState*)

```
GET /tasks/{taskType}/state
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-queue-state-200) | OK | successful operation |  | [schema](#get-task-queue-state-200-schema) |

#### Responses


##### <span id="get-task-queue-state-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-queue-state-200-schema"></span> Schema
   
  



### <span id="get-task-queue-state-deprecated"></span> Get the state (task queue state) for the given task type (deprecated) (*getTaskQueueStateDeprecated*)

```
GET /tasks/taskqueuestate/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-queue-state-deprecated-200) | OK | successful operation |  | [schema](#get-task-queue-state-deprecated-200-schema) |

#### Responses


##### <span id="get-task-queue-state-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-queue-state-deprecated-200-schema"></span> Schema
   
  

[GetTaskQueueStateDeprecatedOKBody](#get-task-queue-state-deprecated-o-k-body)

###### Inlined models

**<span id="get-task-queue-state-deprecated-o-k-body"></span> GetTaskQueueStateDeprecatedOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` |  | |  |  |



### <span id="get-task-queues"></span> List all task queues (deprecated) (*getTaskQueues*)

```
GET /tasks/taskqueues
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-queues-200) | OK | successful operation |  | [schema](#get-task-queues-200-schema) |

#### Responses


##### <span id="get-task-queues-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-queues-200-schema"></span> Schema
   
  

[]string

### <span id="get-task-state"></span> Get the task state for the given task (*getTaskState*)

```
GET /tasks/task/{taskName}/state
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskName | `path` | string | `string` |  | ✓ |  | Task name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-state-200) | OK | successful operation |  | [schema](#get-task-state-200-schema) |

#### Responses


##### <span id="get-task-state-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-state-200-schema"></span> Schema
   
  



### <span id="get-task-state-deprecated"></span> Get the task state for the given task (deprecated) (*getTaskStateDeprecated*)

```
GET /tasks/taskstate/{taskName}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskName | `path` | string | `string` |  | ✓ |  | Task name |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-state-deprecated-200) | OK | successful operation |  | [schema](#get-task-state-deprecated-200-schema) |

#### Responses


##### <span id="get-task-state-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-state-deprecated-200-schema"></span> Schema
   
  

[GetTaskStateDeprecatedOKBody](#get-task-state-deprecated-o-k-body)

###### Inlined models

**<span id="get-task-state-deprecated-o-k-body"></span> GetTaskStateDeprecatedOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` |  | |  |  |



### <span id="get-task-states"></span> Get a map from task to task state for the given task type (*getTaskStates*)

```
GET /tasks/{taskType}/taskstates
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-states-200) | OK | successful operation |  | [schema](#get-task-states-200-schema) |

#### Responses


##### <span id="get-task-states-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-states-200-schema"></span> Schema
   
  

map of string

### <span id="get-task-states-deprecated"></span> Get a map from task to task state for the given task type (deprecated) (*getTaskStatesDeprecated*)

```
GET /tasks/taskstates/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-task-states-deprecated-200) | OK | successful operation |  | [schema](#get-task-states-deprecated-200-schema) |

#### Responses


##### <span id="get-task-states-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-task-states-deprecated-200-schema"></span> Schema
   
  

map of string

### <span id="get-tasks"></span> List all tasks for the given task type (*getTasks*)

```
GET /tasks/{taskType}/tasks
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tasks-200) | OK | successful operation |  | [schema](#get-tasks-200-schema) |

#### Responses


##### <span id="get-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tasks-200-schema"></span> Schema
   
  

[]string

### <span id="get-tasks-deprecated"></span> List all tasks for the given task type (deprecated) (*getTasksDeprecated*)

```
GET /tasks/tasks/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tasks-deprecated-200) | OK | successful operation |  | [schema](#get-tasks-deprecated-200-schema) |

#### Responses


##### <span id="get-tasks-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tasks-deprecated-200-schema"></span> Schema
   
  

[]string

### <span id="get-tenant-metadata"></span> Get tenant information (*getTenantMetadata*)

```
GET /tenants/{tenantName}/metadata
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Tenant name |
| type | `query` | string | `string` |  |  |  | tenant type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tenant-metadata-200) | OK | Success |  | [schema](#get-tenant-metadata-200-schema) |
| [404](#get-tenant-metadata-404) | Not Found | Tenant not found |  | [schema](#get-tenant-metadata-404-schema) |
| [500](#get-tenant-metadata-500) | Internal Server Error | Server error reading tenant information |  | [schema](#get-tenant-metadata-500-schema) |

#### Responses


##### <span id="get-tenant-metadata-200"></span> 200 - Success
Status: OK

###### <span id="get-tenant-metadata-200-schema"></span> Schema
   
  

[GetTenantMetadataOKBody](#get-tenant-metadata-o-k-body)

##### <span id="get-tenant-metadata-404"></span> 404 - Tenant not found
Status: Not Found

###### <span id="get-tenant-metadata-404-schema"></span> Schema

##### <span id="get-tenant-metadata-500"></span> 500 - Server error reading tenant information
Status: Internal Server Error

###### <span id="get-tenant-metadata-500-schema"></span> Schema

###### Inlined models

**<span id="get-tenant-metadata-o-k-body"></span> GetTenantMetadataOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| BrokerInstances | []string| `[]string` |  | |  |  |
| ServerInstances | []string| `[]string` |  | |  |  |
| tenantName | string| `string` |  | |  |  |



### <span id="get-tenants-to-brokers-mapping"></span> List tenants to brokers mappings (*getTenantsToBrokersMapping*)

```
GET /brokers/tenants
```

List tenants to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tenants-to-brokers-mapping-200) | OK | successful operation |  | [schema](#get-tenants-to-brokers-mapping-200-schema) |

#### Responses


##### <span id="get-tenants-to-brokers-mapping-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tenants-to-brokers-mapping-200-schema"></span> Schema
   
  

map of [[]string](#string)

### <span id="get-tenants-to-brokers-mapping-v2"></span> List tenants to brokers mappings (*getTenantsToBrokersMappingV2*)

```
GET /v2/brokers/tenants
```

List tenants to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-tenants-to-brokers-mapping-v2-200) | OK | successful operation |  | [schema](#get-tenants-to-brokers-mapping-v2-200-schema) |

#### Responses


##### <span id="get-tenants-to-brokers-mapping-v2-200"></span> 200 - successful operation
Status: OK

###### <span id="get-tenants-to-brokers-mapping-v2-200-schema"></span> Schema
   
  

map of [[]*GetTenantsToBrokersMappingV2OKBodyItems0](#get-tenants-to-brokers-mapping-v2-o-k-body-items0)

###### Inlined models

**<span id="get-tenants-to-brokers-mapping-v2-o-k-body-items0"></span> GetTenantsToBrokersMappingV2OKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="get-version-number"></span> Get version number of Pinot components (*getVersionNumber*)

```
GET /version
```

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-version-number-200) | OK | Success |  | [schema](#get-version-number-200-schema) |

#### Responses


##### <span id="get-version-number-200"></span> 200 - Success
Status: OK

###### <span id="get-version-number-200-schema"></span> Schema

### <span id="ingest-from-file"></span> Ingest a file (*ingestFromFile*)

```
POST /ingestFromFile
```

Creates a segment using given file and pushes it to Pinot. 
 All steps happen on the controller. This API is NOT meant for production environments/large input files. 
 Example usage (query params need encoding):
```
curl -X POST -F file=@data.json -H "Content-Type: multipart/form-data" "http://localhost:9000/ingestFromFile?tableNameWithType=foo_OFFLINE&
batchConfigMapStr={
  "inputFormat":"csv",
  "recordReader.prop.delimiter":"|"
}" 
```

#### Consumes
  * multipart/form-data

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| batchConfigMapStr | `query` | string | `string` |  | ✓ |  | Batch config Map as json string. Must pass inputFormat, and optionally record reader properties. e.g. {"inputFormat":"json"} |
| tableNameWithType | `query` | string | `string` |  | ✓ |  | Name of the table to upload the file to |
| body | `body` | [IngestFromFileBody](#ingest-from-file-body) | `IngestFromFileBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#ingest-from-file-default) | | successful operation |  | [schema](#ingest-from-file-default-schema) |

#### Responses


##### <span id="ingest-from-file-default"></span> Default Response
successful operation

###### <span id="ingest-from-file-default-schema"></span> Schema
empty schema

###### Inlined models

**<span id="ingest-from-file-body"></span> IngestFromFileBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][IngestFromFileParamsBodyBodyPartsItems0](#ingest-from-file-params-body-body-parts-items0)| `[]*IngestFromFileParamsBodyBodyPartsItems0` |  | |  |  |
| contentDisposition | [IngestFromFileParamsBodyContentDisposition](#ingest-from-file-params-body-content-disposition)| `IngestFromFileParamsBodyContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*IngestFromFileParamsBodyFieldsItems0](#ingest-from-file-params-body-fields-items0)| `map[string][]IngestFromFileParamsBodyFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [IngestFromFileParamsBodyMediaType](#ingest-from-file-params-body-media-type)| `IngestFromFileParamsBodyMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*IngestFromFileParamsBodyParameterizedHeadersItems0](#ingest-from-file-params-body-parameterized-headers-items0)| `map[string][]IngestFromFileParamsBodyParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0"></span> IngestFromFileParamsBodyBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [IngestFromFileParamsBodyBodyPartsItems0ContentDisposition](#ingest-from-file-params-body-body-parts-items0-content-disposition)| `IngestFromFileParamsBodyBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [IngestFromFileParamsBodyBodyPartsItems0MediaType](#ingest-from-file-params-body-body-parts-items0-media-type)| `IngestFromFileParamsBodyBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*IngestFromFileParamsBodyBodyPartsItems0ParameterizedHeadersItems0](#ingest-from-file-params-body-body-parts-items0-parameterized-headers-items0)| `map[string][]IngestFromFileParamsBodyBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [IngestFromFileParamsBodyBodyPartsItems0Parent](#ingest-from-file-params-body-body-parts-items0-parent)| `IngestFromFileParamsBodyBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-content-disposition"></span> IngestFromFileParamsBodyBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-media-type"></span> IngestFromFileParamsBodyBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parameterized-headers-items0"></span> IngestFromFileParamsBodyBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent"></span> IngestFromFileParamsBodyBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0](#ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0)| `[]*IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [IngestFromFileParamsBodyBodyPartsItems0ParentContentDisposition](#ingest-from-file-params-body-body-parts-items0-parent-content-disposition)| `IngestFromFileParamsBodyBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [IngestFromFileParamsBodyBodyPartsItems0ParentMediaType](#ingest-from-file-params-body-body-parts-items0-parent-media-type)| `IngestFromFileParamsBodyBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*IngestFromFileParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0](#ingest-from-file-params-body-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]IngestFromFileParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0"></span> IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-content-disposition)| `IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType](#ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-media-type)| `IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-content-disposition"></span> IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-media-type"></span> IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> IngestFromFileParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-content-disposition"></span> IngestFromFileParamsBodyBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-media-type"></span> IngestFromFileParamsBodyBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="ingest-from-file-params-body-body-parts-items0-parent-parameterized-headers-items0"></span> IngestFromFileParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-content-disposition"></span> IngestFromFileParamsBodyContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-fields-items0"></span> IngestFromFileParamsBodyFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [IngestFromFileParamsBodyFieldsItems0ContentDisposition](#ingest-from-file-params-body-fields-items0-content-disposition)| `IngestFromFileParamsBodyFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [IngestFromFileParamsBodyFieldsItems0FormDataContentDisposition](#ingest-from-file-params-body-fields-items0-form-data-content-disposition)| `IngestFromFileParamsBodyFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [IngestFromFileParamsBodyFieldsItems0MediaType](#ingest-from-file-params-body-fields-items0-media-type)| `IngestFromFileParamsBodyFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*IngestFromFileParamsBodyFieldsItems0ParameterizedHeadersItems0](#ingest-from-file-params-body-fields-items0-parameterized-headers-items0)| `map[string][]IngestFromFileParamsBodyFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-fields-items0-content-disposition"></span> IngestFromFileParamsBodyFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-fields-items0-form-data-content-disposition"></span> IngestFromFileParamsBodyFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-fields-items0-media-type"></span> IngestFromFileParamsBodyFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="ingest-from-file-params-body-fields-items0-parameterized-headers-items0"></span> IngestFromFileParamsBodyFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="ingest-from-file-params-body-media-type"></span> IngestFromFileParamsBodyMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="ingest-from-file-params-body-parameterized-headers-items0"></span> IngestFromFileParamsBodyParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="ingest-from-uri"></span> Ingest from the given URI (*ingestFromURI*)

```
POST /ingestFromURI
```

Creates a segment using file at the given URI and pushes it to Pinot. 
 All steps happen on the controller. This API is NOT meant for production environments/large input files. 
Example usage (query params need encoding):
```
curl -X POST "http://localhost:9000/ingestFromURI?tableNameWithType=foo_OFFLINE
&batchConfigMapStr={
  "inputFormat":"json",
  "input.fs.className":"org.apache.pinot.plugin.filesystem.S3PinotFS",
  "input.fs.prop.region":"us-central",
  "input.fs.prop.accessKey":"foo",
  "input.fs.prop.secretKey":"bar"
}
&sourceURIStr=s3://test.bucket/path/to/json/data/data.json"
```

#### Consumes
  * multipart/form-data

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| batchConfigMapStr | `query` | string | `string` |  | ✓ |  | Batch config Map as json string. Must pass inputFormat, and optionally input FS properties. e.g. {"inputFormat":"json"} |
| sourceURIStr | `query` | string | `string` |  | ✓ |  | URI of file to upload |
| tableNameWithType | `query` | string | `string` |  | ✓ |  | Name of the table to upload the file to |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#ingest-from-uri-default) | | successful operation |  | [schema](#ingest-from-uri-default-schema) |

#### Responses


##### <span id="ingest-from-uri-default"></span> Default Response
successful operation

###### <span id="ingest-from-uri-default-schema"></span> Schema
empty schema

### <span id="list-brokers-mapping"></span> List tenants and tables to brokers mappings (*listBrokersMapping*)

```
GET /brokers
```

List tenants and tables to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-brokers-mapping-200) | OK | successful operation |  | [schema](#list-brokers-mapping-200-schema) |

#### Responses


##### <span id="list-brokers-mapping-200"></span> 200 - successful operation
Status: OK

###### <span id="list-brokers-mapping-200-schema"></span> Schema
   
  

map of [map[string][]string](#map-string-string)

### <span id="list-brokers-mapping-v2"></span> List tenants and tables to brokers mappings (*listBrokersMappingV2*)

```
GET /v2/brokers
```

List tenants and tables to brokers mappings

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| state | `query` | string | `string` |  |  |  | ONLINE|OFFLINE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-brokers-mapping-v2-200) | OK | successful operation |  | [schema](#list-brokers-mapping-v2-200-schema) |

#### Responses


##### <span id="list-brokers-mapping-v2-200"></span> 200 - successful operation
Status: OK

###### <span id="list-brokers-mapping-v2-200-schema"></span> Schema
   
  

map of [map[string][]*ListBrokersMappingV2OKBodyItems0](#map-string-list-brokers-mapping-v2-o-k-body-items0)

###### Inlined models

**<span id="list-brokers-mapping-v2-o-k-body-items0"></span> ListBrokersMappingV2OKBodyItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="list-cluster-configs"></span> List cluster configurations (*listClusterConfigs*)

```
GET /cluster/configs
```

List cluster level configurations

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-cluster-configs-200) | OK | Success |  | [schema](#list-cluster-configs-200-schema) |
| [500](#list-cluster-configs-500) | Internal Server Error | Internal server error |  | [schema](#list-cluster-configs-500-schema) |

#### Responses


##### <span id="list-cluster-configs-200"></span> 200 - Success
Status: OK

###### <span id="list-cluster-configs-200-schema"></span> Schema

##### <span id="list-cluster-configs-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="list-cluster-configs-500-schema"></span> Schema

### <span id="list-instance-or-toggle-tenant-state"></span> List instance for a tenant, or enable/disable/drop a tenant (*listInstanceOrToggleTenantState*)

```
GET /tenants/{tenantName}
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tenantName | `path` | string | `string` |  | ✓ |  | Tenant name |
| state | `query` | string | `string` |  |  |  | state |
| type | `query` | string | `string` |  |  |  | Tenant type (server|broker) |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-instance-or-toggle-tenant-state-200) | OK | Success |  | [schema](#list-instance-or-toggle-tenant-state-200-schema) |
| [500](#list-instance-or-toggle-tenant-state-500) | Internal Server Error | Error reading tenants list |  | [schema](#list-instance-or-toggle-tenant-state-500-schema) |

#### Responses


##### <span id="list-instance-or-toggle-tenant-state-200"></span> 200 - Success
Status: OK

###### <span id="list-instance-or-toggle-tenant-state-200-schema"></span> Schema

##### <span id="list-instance-or-toggle-tenant-state-500"></span> 500 - Error reading tenants list
Status: Internal Server Error

###### <span id="list-instance-or-toggle-tenant-state-500-schema"></span> Schema

### <span id="list-schema-names"></span> List all schema names (*listSchemaNames*)

```
GET /schemas
```

Lists all schema names

#### Produces
  * application/json

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-schema-names-200) | OK | successful operation |  | [schema](#list-schema-names-200-schema) |

#### Responses


##### <span id="list-schema-names-200"></span> 200 - successful operation
Status: OK

###### <span id="list-schema-names-200-schema"></span> Schema
   
  



### <span id="list-table-configs"></span> Lists all tables in cluster (*listTableConfigs*)

```
GET /tables
```

Lists all tables in cluster

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| type | `query` | string | `string` |  |  |  | realtime|offline |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-table-configs-200) | OK | successful operation |  | [schema](#list-table-configs-200-schema) |

#### Responses


##### <span id="list-table-configs-200"></span> 200 - successful operation
Status: OK

###### <span id="list-table-configs-200-schema"></span> Schema
   
  



### <span id="list-task-types"></span> List all task types (*listTaskTypes*)

```
GET /tasks/tasktypes
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-task-types-200) | OK | successful operation |  | [schema](#list-task-types-200-schema) |

#### Responses


##### <span id="list-task-types-200"></span> 200 - successful operation
Status: OK

###### <span id="list-task-types-200-schema"></span> Schema
   
  

[]string

### <span id="ls"></span> List the child znodes (*ls*)

```
GET /zk/ls
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#ls-200) | OK | Success |  | [schema](#ls-200-schema) |
| [404](#ls-404) | Not Found | ZK Path not found |  | [schema](#ls-404-schema) |
| [500](#ls-500) | Internal Server Error | Internal server error |  | [schema](#ls-500-schema) |

#### Responses


##### <span id="ls-200"></span> 200 - Success
Status: OK

###### <span id="ls-200-schema"></span> Schema

##### <span id="ls-404"></span> 404 - ZK Path not found
Status: Not Found

###### <span id="ls-404-schema"></span> Schema

##### <span id="ls-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="ls-500-schema"></span> Schema

### <span id="lsl"></span> List the child znodes along with Stats (*lsl*)

```
GET /zk/lsl
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#lsl-200) | OK | Success |  | [schema](#lsl-200-schema) |
| [404](#lsl-404) | Not Found | ZK Path not found |  | [schema](#lsl-404-schema) |
| [500](#lsl-500) | Internal Server Error | Internal server error |  | [schema](#lsl-500-schema) |

#### Responses


##### <span id="lsl-200"></span> 200 - Success
Status: OK

###### <span id="lsl-200-schema"></span> Schema

##### <span id="lsl-404"></span> 404 - ZK Path not found
Status: Not Found

###### <span id="lsl-404-schema"></span> Schema

##### <span id="lsl-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="lsl-500-schema"></span> Schema

### <span id="put"></span> Update segments configuration (*put*)

```
PUT /tables/{tableName}/segmentConfigs
```

Updates segmentsConfig section (validation and retention) of a table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-200) | OK | Success |  | [schema](#put-200-schema) |
| [404](#put-404) | Not Found | Table not found |  | [schema](#put-404-schema) |
| [500](#put-500) | Internal Server Error | Internal server error |  | [schema](#put-500-schema) |

#### Responses


##### <span id="put-200"></span> 200 - Success
Status: OK

###### <span id="put-200-schema"></span> Schema

##### <span id="put-404"></span> 404 - Table not found
Status: Not Found

###### <span id="put-404-schema"></span> Schema

##### <span id="put-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="put-500-schema"></span> Schema

### <span id="put-data"></span> Update the content of the node (*putData*)

```
PUT /zk/put
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| accessOption | `query` | string | `string` |  | ✓ | `"1"` | accessOption |
| data | `query` | string | `string` |  | ✓ |  | Content |
| expectedVersion | `query` | string | `string` |  | ✓ | `"-1"` | expectedVersion |
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#put-data-200) | OK | Success |  | [schema](#put-data-200-schema) |
| [204](#put-data-204) | No Content | No Content |  | [schema](#put-data-204-schema) |
| [404](#put-data-404) | Not Found | ZK Path not found |  | [schema](#put-data-404-schema) |
| [500](#put-data-500) | Internal Server Error | Internal server error |  | [schema](#put-data-500-schema) |

#### Responses


##### <span id="put-data-200"></span> 200 - Success
Status: OK

###### <span id="put-data-200-schema"></span> Schema

##### <span id="put-data-204"></span> 204 - No Content
Status: No Content

###### <span id="put-data-204-schema"></span> Schema

##### <span id="put-data-404"></span> 404 - ZK Path not found
Status: Not Found

###### <span id="put-data-404-schema"></span> Schema

##### <span id="put-data-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="put-data-500-schema"></span> Schema

### <span id="rebalance"></span> Rebalances a table (reassign instances and segments for a table) (*rebalance*)

```
POST /tables/{tableName}/rebalance
```

Rebalances a table (reassign instances and segments for a table)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table to rebalance |
| bestEfforts | `query` | boolean | `bool` |  |  |  | Whether to use best-efforts to rebalance (not fail the rebalance when the no-downtime contract cannot be achieved) |
| bootstrap | `query` | boolean | `bool` |  |  |  | Whether to rebalance table in bootstrap mode (regardless of minimum segment movement, reassign all segments in a round-robin fashion as if adding new segments to an empty table) |
| downtime | `query` | boolean | `bool` |  |  |  | Whether to allow downtime for the rebalance |
| dryRun | `query` | boolean | `bool` |  |  |  | Whether to rebalance table in dry-run mode |
| includeConsuming | `query` | boolean | `bool` |  |  |  | Whether to reassign CONSUMING segments for real-time table |
| minAvailableReplicas | `query` | int32 (formatted integer) | `int32` |  |  | `1` | For no-downtime rebalance, minimum number of replicas to keep alive during rebalance, or maximum number of replicas allowed to be unavailable if value is negative |
| reassignInstances | `query` | boolean | `bool` |  |  |  | Whether to reassign instances before reassigning segments |
| type | `query` | string | `string` |  | ✓ |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#rebalance-200) | OK | successful operation |  | [schema](#rebalance-200-schema) |

#### Responses


##### <span id="rebalance-200"></span> 200 - successful operation
Status: OK

###### <span id="rebalance-200-schema"></span> Schema
   
  

[RebalanceOKBody](#rebalance-o-k-body)

###### Inlined models

**<span id="rebalance-o-k-body"></span> RebalanceOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | |  |  |
| instanceAssignment | map of [RebalanceOKBodyInstanceAssignmentAnon](#rebalance-o-k-body-instance-assignment-anon)| `map[string]RebalanceOKBodyInstanceAssignmentAnon` |  | |  |  |
| segmentAssignment | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |
| status | string| `string` |  | |  |  |



**<span id="rebalance-o-k-body-instance-assignment-anon"></span> RebalanceOKBodyInstanceAssignmentAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="rebuild-broker-resource"></span> Rebuild broker resource for table (*rebuildBrokerResource*)

```
POST /tables/{tableName}/rebuildBrokerResourceFromHelixTags
```

when new brokers are added

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name (with type) |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#rebuild-broker-resource-200) | OK | Success |  | [schema](#rebuild-broker-resource-200-schema) |
| [400](#rebuild-broker-resource-400) | Bad Request | Bad request: table name has to be with table type |  | [schema](#rebuild-broker-resource-400-schema) |
| [500](#rebuild-broker-resource-500) | Internal Server Error | Internal error rebuilding broker resource or serializing response |  | [schema](#rebuild-broker-resource-500-schema) |

#### Responses


##### <span id="rebuild-broker-resource-200"></span> 200 - Success
Status: OK

###### <span id="rebuild-broker-resource-200-schema"></span> Schema

##### <span id="rebuild-broker-resource-400"></span> 400 - Bad request: table name has to be with table type
Status: Bad Request

###### <span id="rebuild-broker-resource-400-schema"></span> Schema

##### <span id="rebuild-broker-resource-500"></span> 500 - Internal error rebuilding broker resource or serializing response
Status: Internal Server Error

###### <span id="rebuild-broker-resource-500-schema"></span> Schema

### <span id="recommend-config"></span> Recommend config (*recommendConfig*)

```
PUT /tables/recommender
```

Recommend a config with input json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#recommend-config-200) | OK | successful operation |  | [schema](#recommend-config-200-schema) |

#### Responses


##### <span id="recommend-config-200"></span> 200 - successful operation
Status: OK

###### <span id="recommend-config-200-schema"></span> Schema
   
  



### <span id="reload-all-segments"></span> Reload all segments (*reloadAllSegments*)

```
POST /segments/{tableName}/reload
```

Reload all segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-all-segments-200) | OK | successful operation |  | [schema](#reload-all-segments-200-schema) |

#### Responses


##### <span id="reload-all-segments-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-all-segments-200-schema"></span> Schema
   
  

[ReloadAllSegmentsOKBody](#reload-all-segments-o-k-body)

###### Inlined models

**<span id="reload-all-segments-o-k-body"></span> ReloadAllSegmentsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reload-all-segments-deprecated1"></span> Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead) (*reloadAllSegmentsDeprecated1*)

```
POST /tables/{tableName}/segments/reload
```

Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-all-segments-deprecated1-200) | OK | successful operation |  | [schema](#reload-all-segments-deprecated1-200-schema) |

#### Responses


##### <span id="reload-all-segments-deprecated1-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-all-segments-deprecated1-200-schema"></span> Schema
   
  

[ReloadAllSegmentsDeprecated1OKBody](#reload-all-segments-deprecated1-o-k-body)

###### Inlined models

**<span id="reload-all-segments-deprecated1-o-k-body"></span> ReloadAllSegmentsDeprecated1OKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reload-all-segments-deprecated2"></span> Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead) (*reloadAllSegmentsDeprecated2*)

```
GET /tables/{tableName}/segments/reload
```

Reload all segments (deprecated, use 'POST /segments/{tableName}/reload' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-all-segments-deprecated2-200) | OK | successful operation |  | [schema](#reload-all-segments-deprecated2-200-schema) |

#### Responses


##### <span id="reload-all-segments-deprecated2-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-all-segments-deprecated2-200-schema"></span> Schema
   
  

[ReloadAllSegmentsDeprecated2OKBody](#reload-all-segments-deprecated2-o-k-body)

###### Inlined models

**<span id="reload-all-segments-deprecated2-o-k-body"></span> ReloadAllSegmentsDeprecated2OKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reload-segment"></span> Reload a segment (*reloadSegment*)

```
POST /segments/{tableName}/{segmentName}/reload
```

Reload a segment

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-segment-200) | OK | successful operation |  | [schema](#reload-segment-200-schema) |

#### Responses


##### <span id="reload-segment-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-segment-200-schema"></span> Schema
   
  

[ReloadSegmentOKBody](#reload-segment-o-k-body)

###### Inlined models

**<span id="reload-segment-o-k-body"></span> ReloadSegmentOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reload-segment-deprecated1"></span> Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead) (*reloadSegmentDeprecated1*)

```
POST /tables/{tableName}/segments/{segmentName}/reload
```

Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-segment-deprecated1-200) | OK | successful operation |  | [schema](#reload-segment-deprecated1-200-schema) |

#### Responses


##### <span id="reload-segment-deprecated1-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-segment-deprecated1-200-schema"></span> Schema
   
  

[ReloadSegmentDeprecated1OKBody](#reload-segment-deprecated1-o-k-body)

###### Inlined models

**<span id="reload-segment-deprecated1-o-k-body"></span> ReloadSegmentDeprecated1OKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reload-segment-deprecated2"></span> Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead) (*reloadSegmentDeprecated2*)

```
GET /tables/{tableName}/segments/{segmentName}/reload
```

Reload a segment (deprecated, use 'POST /segments/{tableName}/{segmentName}/reload' instead)

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reload-segment-deprecated2-200) | OK | successful operation |  | [schema](#reload-segment-deprecated2-200-schema) |

#### Responses


##### <span id="reload-segment-deprecated2-200"></span> 200 - successful operation
Status: OK

###### <span id="reload-segment-deprecated2-200-schema"></span> Schema
   
  

[ReloadSegmentDeprecated2OKBody](#reload-segment-deprecated2-o-k-body)

###### Inlined models

**<span id="reload-segment-deprecated2-o-k-body"></span> ReloadSegmentDeprecated2OKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="remove-instance-partitions"></span> Remove the instance partitions (*removeInstancePartitions*)

```
DELETE /tables/{tableName}/instancePartitions
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|CONSUMING|COMPLETED |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#remove-instance-partitions-200) | OK | successful operation |  | [schema](#remove-instance-partitions-200-schema) |

#### Responses


##### <span id="remove-instance-partitions-200"></span> 200 - successful operation
Status: OK

###### <span id="remove-instance-partitions-200-schema"></span> Schema
   
  

[RemoveInstancePartitionsOKBody](#remove-instance-partitions-o-k-body)

###### Inlined models

**<span id="remove-instance-partitions-o-k-body"></span> RemoveInstancePartitionsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="replace-instance"></span> Replace an instance in the instance partitions (*replaceInstance*)

```
POST /tables/{tableName}/replaceInstance
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| newInstanceId | `query` | string | `string` |  | ✓ |  | New instance to replace with |
| oldInstanceId | `query` | string | `string` |  | ✓ |  | Old instance to be replaced |
| type | `query` | string | `string` |  |  |  | OFFLINE|CONSUMING|COMPLETED |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#replace-instance-200) | OK | successful operation |  | [schema](#replace-instance-200-schema) |

#### Responses


##### <span id="replace-instance-200"></span> 200 - successful operation
Status: OK

###### <span id="replace-instance-200-schema"></span> Schema
   
  

map of [ReplaceInstanceOKBodyAnon](#replace-instance-o-k-body-anon)

###### Inlined models

**<span id="replace-instance-o-k-body-anon"></span> ReplaceInstanceOKBodyAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="reset-all-segments"></span> Resets all segments of the table, by first disabling them, waiting for external view to stabilize, and finally enabling the segments (*resetAllSegments*)

```
POST /segments/{tableNameWithType}/reset
```

Resets a segment by disabling and then enabling a segment

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableNameWithType | `path` | string | `string` |  | ✓ |  | Name of the table with type |
| maxWaitTimeMs | `query` | int64 (formatted integer) | `int64` |  |  |  | Maximum time in milliseconds to wait for reset to be completed. By default, uses serverAdminRequestTimeout |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reset-all-segments-200) | OK | successful operation |  | [schema](#reset-all-segments-200-schema) |

#### Responses


##### <span id="reset-all-segments-200"></span> 200 - successful operation
Status: OK

###### <span id="reset-all-segments-200-schema"></span> Schema
   
  

[ResetAllSegmentsOKBody](#reset-all-segments-o-k-body)

###### Inlined models

**<span id="reset-all-segments-o-k-body"></span> ResetAllSegmentsOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="reset-segment"></span> Resets a segment by first disabling it, waiting for external view to stabilize, and finally enabling it again (*resetSegment*)

```
POST /segments/{tableNameWithType}/{segmentName}/reset
```

Resets a segment by disabling and then enabling the segment

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| segmentName | `path` | string | `string` |  | ✓ |  | Name of the segment |
| tableNameWithType | `path` | string | `string` |  | ✓ |  | Name of the table with type |
| maxWaitTimeMs | `query` | int64 (formatted integer) | `int64` |  |  |  | Maximum time in milliseconds to wait for reset to be completed. By default, uses serverAdminRequestTimeout |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#reset-segment-200) | OK | successful operation |  | [schema](#reset-segment-200-schema) |

#### Responses


##### <span id="reset-segment-200"></span> 200 - successful operation
Status: OK

###### <span id="reset-segment-200-schema"></span> Schema
   
  

[ResetSegmentOKBody](#reset-segment-o-k-body)

###### Inlined models

**<span id="reset-segment-o-k-body"></span> ResetSegmentOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="resume-tasks"></span> Resume all stopped tasks (as well as the task queue) for the given task type (*resumeTasks*)

```
PUT /tasks/{taskType}/resume
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#resume-tasks-200) | OK | successful operation |  | [schema](#resume-tasks-200-schema) |

#### Responses


##### <span id="resume-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="resume-tasks-200-schema"></span> Schema
   
  

[ResumeTasksOKBody](#resume-tasks-o-k-body)

###### Inlined models

**<span id="resume-tasks-o-k-body"></span> ResumeTasksOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="schedule-tasks"></span> Schedule tasks and return a map from task type to task name scheduled (*scheduleTasks*)

```
POST /tasks/schedule
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `query` | string | `string` |  |  |  | Table name (with type suffix) |
| taskType | `query` | string | `string` |  |  |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#schedule-tasks-200) | OK | successful operation |  | [schema](#schedule-tasks-200-schema) |

#### Responses


##### <span id="schedule-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="schedule-tasks-200-schema"></span> Schema
   
  

map of string

### <span id="schedule-tasks-deprecated"></span> Schedule tasks (deprecated) (*scheduleTasksDeprecated*)

```
PUT /tasks/scheduletasks
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#schedule-tasks-deprecated-200) | OK | successful operation |  | [schema](#schedule-tasks-deprecated-200-schema) |

#### Responses


##### <span id="schedule-tasks-deprecated-200"></span> 200 - successful operation
Status: OK

###### <span id="schedule-tasks-deprecated-200-schema"></span> Schema
   
  

map of string

### <span id="set-instance-partitions"></span> Create/update the instance partitions (*setInstancePartitions*)

```
PUT /tables/{tableName}/instancePartitions
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#set-instance-partitions-200) | OK | successful operation |  | [schema](#set-instance-partitions-200-schema) |

#### Responses


##### <span id="set-instance-partitions-200"></span> 200 - successful operation
Status: OK

###### <span id="set-instance-partitions-200-schema"></span> Schema
   
  

map of [SetInstancePartitionsOKBodyAnon](#set-instance-partitions-o-k-body-anon)

###### Inlined models

**<span id="set-instance-partitions-o-k-body-anon"></span> SetInstancePartitionsOKBodyAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="start-replace-segments"></span> Start to replace segments (*startReplaceSegments*)

```
POST /segments/{tableName}/startReplaceSegments
```

Start to replace segments

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table |
| type | `query` | string | `string` |  |  |  | OFFLINE|REALTIME |
| body | `body` | [StartReplaceSegmentsBody](#start-replace-segments-body) | `StartReplaceSegmentsBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#start-replace-segments-default) | | successful operation |  | [schema](#start-replace-segments-default-schema) |

#### Responses


##### <span id="start-replace-segments-default"></span> Default Response
successful operation

###### <span id="start-replace-segments-default-schema"></span> Schema
empty schema

###### Inlined models

**<span id="start-replace-segments-body"></span> StartReplaceSegmentsBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| segmentsFrom | []string| `[]string` |  | |  |  |
| segmentsTo | []string| `[]string` |  | |  |  |



### <span id="stat"></span> Get the stat (*stat*)

```
GET /zk/stat
```

 Use this api to fetch additional details of a znode such as creation time, modified time, numChildren etc 

#### Produces
  * text/plain

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| path | `query` | string | `string` |  | ✓ | `"/"` | Zookeeper Path, must start with / |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#stat-200) | OK | Success |  | [schema](#stat-200-schema) |
| [404](#stat-404) | Not Found | Table not found |  | [schema](#stat-404-schema) |
| [500](#stat-500) | Internal Server Error | Internal server error |  | [schema](#stat-500-schema) |

#### Responses


##### <span id="stat-200"></span> 200 - Success
Status: OK

###### <span id="stat-200-schema"></span> Schema

##### <span id="stat-404"></span> 404 - Table not found
Status: Not Found

###### <span id="stat-404-schema"></span> Schema

##### <span id="stat-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="stat-500-schema"></span> Schema

### <span id="stop-tasks"></span> Stop all running/pending tasks (as well as the task queue) for the given task type (*stopTasks*)

```
PUT /tasks/{taskType}/stop
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#stop-tasks-200) | OK | successful operation |  | [schema](#stop-tasks-200-schema) |

#### Responses


##### <span id="stop-tasks-200"></span> 200 - successful operation
Status: OK

###### <span id="stop-tasks-200-schema"></span> Schema
   
  

[StopTasksOKBody](#stop-tasks-o-k-body)

###### Inlined models

**<span id="stop-tasks-o-k-body"></span> StopTasksOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="toggle-instance-state"></span> Enable/disable/drop an instance (*toggleInstanceState*)

```
POST /instances/{instanceName}/state
```

Enable/disable/drop an instance

#### Consumes
  * text/plain

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Instance name |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#toggle-instance-state-200) | OK | Success |  | [schema](#toggle-instance-state-200-schema) |
| [400](#toggle-instance-state-400) | Bad Request | Bad Request |  | [schema](#toggle-instance-state-400-schema) |
| [404](#toggle-instance-state-404) | Not Found | Instance not found |  | [schema](#toggle-instance-state-404-schema) |
| [409](#toggle-instance-state-409) | Conflict | Instance cannot be dropped |  | [schema](#toggle-instance-state-409-schema) |
| [500](#toggle-instance-state-500) | Internal Server Error | Internal error |  | [schema](#toggle-instance-state-500-schema) |

#### Responses


##### <span id="toggle-instance-state-200"></span> 200 - Success
Status: OK

###### <span id="toggle-instance-state-200-schema"></span> Schema

##### <span id="toggle-instance-state-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="toggle-instance-state-400-schema"></span> Schema

##### <span id="toggle-instance-state-404"></span> 404 - Instance not found
Status: Not Found

###### <span id="toggle-instance-state-404-schema"></span> Schema

##### <span id="toggle-instance-state-409"></span> 409 - Instance cannot be dropped
Status: Conflict

###### <span id="toggle-instance-state-409-schema"></span> Schema

##### <span id="toggle-instance-state-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="toggle-instance-state-500-schema"></span> Schema

### <span id="toggle-query-rate-limiting"></span> Enable/disable the query rate limiting for a broker instance (*toggleQueryRateLimiting*)

```
POST /brokers/instances/{instanceName}/qps
```

Enable/disable the query rate limiting for a broker instance

#### Consumes
  * text/plain

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Broker instance name |
| state | `query` | string | `string` |  | ✓ |  | ENABLE|DISABLE |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#toggle-query-rate-limiting-200) | OK | Success |  | [schema](#toggle-query-rate-limiting-200-schema) |
| [400](#toggle-query-rate-limiting-400) | Bad Request | Bad Request |  | [schema](#toggle-query-rate-limiting-400-schema) |
| [404](#toggle-query-rate-limiting-404) | Not Found | Instance not found |  | [schema](#toggle-query-rate-limiting-404-schema) |
| [500](#toggle-query-rate-limiting-500) | Internal Server Error | Internal error |  | [schema](#toggle-query-rate-limiting-500-schema) |

#### Responses


##### <span id="toggle-query-rate-limiting-200"></span> 200 - Success
Status: OK

###### <span id="toggle-query-rate-limiting-200-schema"></span> Schema

##### <span id="toggle-query-rate-limiting-400"></span> 400 - Bad Request
Status: Bad Request

###### <span id="toggle-query-rate-limiting-400-schema"></span> Schema

##### <span id="toggle-query-rate-limiting-404"></span> 404 - Instance not found
Status: Not Found

###### <span id="toggle-query-rate-limiting-404-schema"></span> Schema

##### <span id="toggle-query-rate-limiting-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="toggle-query-rate-limiting-500-schema"></span> Schema

### <span id="toggle-task-queue-state"></span> Stop/resume a task queue (deprecated) (*toggleTaskQueueState*)

```
PUT /tasks/taskqueue/{taskType}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| taskType | `path` | string | `string` |  | ✓ |  | Task type |
| state | `query` | string | `string` |  | ✓ |  | state |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#toggle-task-queue-state-200) | OK | successful operation |  | [schema](#toggle-task-queue-state-200-schema) |

#### Responses


##### <span id="toggle-task-queue-state-200"></span> 200 - successful operation
Status: OK

###### <span id="toggle-task-queue-state-200-schema"></span> Schema
   
  

[ToggleTaskQueueStateOKBody](#toggle-task-queue-state-o-k-body)

###### Inlined models

**<span id="toggle-task-queue-state-o-k-body"></span> ToggleTaskQueueStateOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="update-cluster-config"></span> Update cluster configuration (*updateClusterConfig*)

```
POST /cluster/configs
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-cluster-config-200) | OK | Success |  | [schema](#update-cluster-config-200-schema) |
| [500](#update-cluster-config-500) | Internal Server Error | Server error updating configuration |  | [schema](#update-cluster-config-500-schema) |

#### Responses


##### <span id="update-cluster-config-200"></span> 200 - Success
Status: OK

###### <span id="update-cluster-config-200-schema"></span> Schema

##### <span id="update-cluster-config-500"></span> 500 - Server error updating configuration
Status: Internal Server Error

###### <span id="update-cluster-config-500-schema"></span> Schema

### <span id="update-indexing-config"></span> Update table indexing configuration (*updateIndexingConfig*)

```
PUT /tables/{tableName}/indexingConfigs
```

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Table name (without type) |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-indexing-config-200) | OK | Success |  | [schema](#update-indexing-config-200-schema) |
| [404](#update-indexing-config-404) | Not Found | Table not found |  | [schema](#update-indexing-config-404-schema) |
| [500](#update-indexing-config-500) | Internal Server Error | Server error updating configuration |  | [schema](#update-indexing-config-500-schema) |

#### Responses


##### <span id="update-indexing-config-200"></span> 200 - Success
Status: OK

###### <span id="update-indexing-config-200-schema"></span> Schema

##### <span id="update-indexing-config-404"></span> 404 - Table not found
Status: Not Found

###### <span id="update-indexing-config-404-schema"></span> Schema

##### <span id="update-indexing-config-500"></span> 500 - Server error updating configuration
Status: Internal Server Error

###### <span id="update-indexing-config-500-schema"></span> Schema

### <span id="update-instance"></span> Update the specified instance (*updateInstance*)

```
PUT /instances/{instanceName}
```

Update specified instance with given instance config

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Instance name |
| body | `body` | [UpdateInstanceBody](#update-instance-body) | `UpdateInstanceBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-instance-200) | OK | Success |  | [schema](#update-instance-200-schema) |
| [500](#update-instance-500) | Internal Server Error | Internal error |  | [schema](#update-instance-500-schema) |

#### Responses


##### <span id="update-instance-200"></span> 200 - Success
Status: OK

###### <span id="update-instance-200-schema"></span> Schema

##### <span id="update-instance-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="update-instance-500-schema"></span> Schema

###### Inlined models

**<span id="update-instance-body"></span> UpdateInstanceBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| grpcPort | int32 (formatted integer)| `int32` |  | |  |  |
| host | string| `string` | ✓ | |  |  |
| pools | map of int32 (formatted integer)| `map[string]int32` |  | |  |  |
| port | int32 (formatted integer)| `int32` | ✓ | |  |  |
| tags | []string| `[]string` |  | |  |  |
| type | string| `string` | ✓ | |  |  |



### <span id="update-instance-tags"></span> Update the tags of the specified instance (*updateInstanceTags*)

```
PUT /instances/{instanceName}/updateTags
```

Update the tags of the specified instance

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| instanceName | `path` | string | `string` |  | ✓ |  | Instance name |
| tags | `query` | string | `string` |  | ✓ |  | Comma separated tags list |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-instance-tags-200) | OK | Success |  | [schema](#update-instance-tags-200-schema) |
| [500](#update-instance-tags-500) | Internal Server Error | Internal error |  | [schema](#update-instance-tags-500-schema) |

#### Responses


##### <span id="update-instance-tags-200"></span> 200 - Success
Status: OK

###### <span id="update-instance-tags-200-schema"></span> Schema

##### <span id="update-instance-tags-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="update-instance-tags-500-schema"></span> Schema

### <span id="update-schema"></span> Update a schema (*updateSchema*)

```
PUT /schemas/{schemaName}
```

Updates a schema

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| schemaName | `path` | string | `string` |  | ✓ |  | Name of the schema |
| reload | `query` | boolean | `bool` |  |  |  | Whether to reload the table if the new schema is backward compatible |
| body | `body` | [UpdateSchemaBody](#update-schema-body) | `UpdateSchemaBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-schema-200) | OK | Successfully updated schema |  | [schema](#update-schema-200-schema) |
| [400](#update-schema-400) | Bad Request | Missing or invalid request body |  | [schema](#update-schema-400-schema) |
| [404](#update-schema-404) | Not Found | Schema not found |  | [schema](#update-schema-404-schema) |
| [500](#update-schema-500) | Internal Server Error | Internal error |  | [schema](#update-schema-500-schema) |

#### Responses


##### <span id="update-schema-200"></span> 200 - Successfully updated schema
Status: OK

###### <span id="update-schema-200-schema"></span> Schema

##### <span id="update-schema-400"></span> 400 - Missing or invalid request body
Status: Bad Request

###### <span id="update-schema-400-schema"></span> Schema

##### <span id="update-schema-404"></span> 404 - Schema not found
Status: Not Found

###### <span id="update-schema-404-schema"></span> Schema

##### <span id="update-schema-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="update-schema-500-schema"></span> Schema

###### Inlined models

**<span id="update-schema-body"></span> UpdateSchemaBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UpdateSchemaParamsBodyBodyPartsItems0](#update-schema-params-body-body-parts-items0)| `[]*UpdateSchemaParamsBodyBodyPartsItems0` |  | |  |  |
| contentDisposition | [UpdateSchemaParamsBodyContentDisposition](#update-schema-params-body-content-disposition)| `UpdateSchemaParamsBodyContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*UpdateSchemaParamsBodyFieldsItems0](#update-schema-params-body-fields-items0)| `map[string][]UpdateSchemaParamsBodyFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UpdateSchemaParamsBodyMediaType](#update-schema-params-body-media-type)| `UpdateSchemaParamsBodyMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UpdateSchemaParamsBodyParameterizedHeadersItems0](#update-schema-params-body-parameterized-headers-items0)| `map[string][]UpdateSchemaParamsBodyParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0"></span> UpdateSchemaParamsBodyBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UpdateSchemaParamsBodyBodyPartsItems0ContentDisposition](#update-schema-params-body-body-parts-items0-content-disposition)| `UpdateSchemaParamsBodyBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UpdateSchemaParamsBodyBodyPartsItems0MediaType](#update-schema-params-body-body-parts-items0-media-type)| `UpdateSchemaParamsBodyBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UpdateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0](#update-schema-params-body-body-parts-items0-parameterized-headers-items0)| `map[string][]UpdateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [UpdateSchemaParamsBodyBodyPartsItems0Parent](#update-schema-params-body-body-parts-items0-parent)| `UpdateSchemaParamsBodyBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-content-disposition"></span> UpdateSchemaParamsBodyBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-media-type"></span> UpdateSchemaParamsBodyBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parameterized-headers-items0"></span> UpdateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent"></span> UpdateSchemaParamsBodyBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0](#update-schema-params-body-body-parts-items0-parent-body-parts-items0)| `[]*UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [UpdateSchemaParamsBodyBodyPartsItems0ParentContentDisposition](#update-schema-params-body-body-parts-items0-parent-content-disposition)| `UpdateSchemaParamsBodyBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UpdateSchemaParamsBodyBodyPartsItems0ParentMediaType](#update-schema-params-body-body-parts-items0-parent-media-type)| `UpdateSchemaParamsBodyBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UpdateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0](#update-schema-params-body-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]UpdateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-body-parts-items0"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#update-schema-params-body-body-parts-items0-parent-body-parts-items0-content-disposition)| `UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType](#update-schema-params-body-body-parts-items0-parent-body-parts-items0-media-type)| `UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#update-schema-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-body-parts-items0-content-disposition"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-body-parts-items0-media-type"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-content-disposition"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-media-type"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="update-schema-params-body-body-parts-items0-parent-parameterized-headers-items0"></span> UpdateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="update-schema-params-body-content-disposition"></span> UpdateSchemaParamsBodyContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-fields-items0"></span> UpdateSchemaParamsBodyFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UpdateSchemaParamsBodyFieldsItems0ContentDisposition](#update-schema-params-body-fields-items0-content-disposition)| `UpdateSchemaParamsBodyFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [UpdateSchemaParamsBodyFieldsItems0FormDataContentDisposition](#update-schema-params-body-fields-items0-form-data-content-disposition)| `UpdateSchemaParamsBodyFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UpdateSchemaParamsBodyFieldsItems0MediaType](#update-schema-params-body-fields-items0-media-type)| `UpdateSchemaParamsBodyFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*UpdateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0](#update-schema-params-body-fields-items0-parameterized-headers-items0)| `map[string][]UpdateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="update-schema-params-body-fields-items0-content-disposition"></span> UpdateSchemaParamsBodyFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-fields-items0-form-data-content-disposition"></span> UpdateSchemaParamsBodyFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="update-schema-params-body-fields-items0-media-type"></span> UpdateSchemaParamsBodyFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="update-schema-params-body-fields-items0-parameterized-headers-items0"></span> UpdateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="update-schema-params-body-media-type"></span> UpdateSchemaParamsBodyMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="update-schema-params-body-parameterized-headers-items0"></span> UpdateSchemaParamsBodyParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="update-table-config"></span> Updates table config for a table (*updateTableConfig*)

```
PUT /tables/{tableName}
```

Updates table config for a table

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  | Name of the table to update |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-table-config-200) | OK | successful operation |  | [schema](#update-table-config-200-schema) |

#### Responses


##### <span id="update-table-config-200"></span> 200 - successful operation
Status: OK

###### <span id="update-table-config-200-schema"></span> Schema
   
  

[UpdateTableConfigOKBody](#update-table-config-o-k-body)

###### Inlined models

**<span id="update-table-config-o-k-body"></span> UpdateTableConfigOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="update-table-metadata"></span> Update table metadata (*updateTableMetadata*)

```
PUT /tables/{tableName}/metadataConfigs
```

Updates table configuration

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tableName | `path` | string | `string` |  | ✓ |  |  |
| body | `body` | string | `string` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-table-metadata-200) | OK | Success |  | [schema](#update-table-metadata-200-schema) |
| [404](#update-table-metadata-404) | Not Found | Table not found |  | [schema](#update-table-metadata-404-schema) |
| [500](#update-table-metadata-500) | Internal Server Error | Internal server error |  | [schema](#update-table-metadata-500-schema) |

#### Responses


##### <span id="update-table-metadata-200"></span> 200 - Success
Status: OK

###### <span id="update-table-metadata-200-schema"></span> Schema

##### <span id="update-table-metadata-404"></span> 404 - Table not found
Status: Not Found

###### <span id="update-table-metadata-404-schema"></span> Schema

##### <span id="update-table-metadata-500"></span> 500 - Internal server error
Status: Internal Server Error

###### <span id="update-table-metadata-500-schema"></span> Schema

### <span id="update-tenant"></span> Update a tenant (*updateTenant*)

```
PUT /tenants
```

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [UpdateTenantBody](#update-tenant-body) | `UpdateTenantBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#update-tenant-200) | OK | Success |  | [schema](#update-tenant-200-schema) |
| [500](#update-tenant-500) | Internal Server Error | Failed to update the tenant |  | [schema](#update-tenant-500-schema) |

#### Responses


##### <span id="update-tenant-200"></span> 200 - Success
Status: OK

###### <span id="update-tenant-200-schema"></span> Schema

##### <span id="update-tenant-500"></span> 500 - Failed to update the tenant
Status: Internal Server Error

###### <span id="update-tenant-500-schema"></span> Schema

###### Inlined models

**<span id="update-tenant-body"></span> UpdateTenantBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| numberOfInstances | int32 (formatted integer)| `int32` |  | |  |  |
| offlineInstances | int32 (formatted integer)| `int32` |  | |  |  |
| realtimeInstances | int32 (formatted integer)| `int32` |  | |  |  |
| tenantName | string| `string` | ✓ | |  |  |
| tenantRole | string| `string` | ✓ | |  |  |



### <span id="upload-segment-as-multi-part"></span> Upload a segment (*uploadSegmentAsMultiPart*)

```
POST /segments
```

Upload a segment as binary

#### Consumes
  * multipart/form-data

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| enableParallelPushProtection | `query` | boolean | `bool` |  |  |  | Whether to enable parallel push protection |
| tableName | `query` | string | `string` |  |  |  | Name of the table |
| body | `body` | [UploadSegmentAsMultiPartBody](#upload-segment-as-multi-part-body) | `UploadSegmentAsMultiPartBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#upload-segment-as-multi-part-default) | | successful operation |  | [schema](#upload-segment-as-multi-part-default-schema) |

#### Responses


##### <span id="upload-segment-as-multi-part-default"></span> Default Response
successful operation

###### <span id="upload-segment-as-multi-part-default-schema"></span> Schema
empty schema

###### Inlined models

**<span id="upload-segment-as-multi-part-body"></span> UploadSegmentAsMultiPartBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UploadSegmentAsMultiPartParamsBodyBodyPartsItems0](#upload-segment-as-multi-part-params-body-body-parts-items0)| `[]*UploadSegmentAsMultiPartParamsBodyBodyPartsItems0` |  | |  |  |
| contentDisposition | [UploadSegmentAsMultiPartParamsBodyContentDisposition](#upload-segment-as-multi-part-params-body-content-disposition)| `UploadSegmentAsMultiPartParamsBodyContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*UploadSegmentAsMultiPartParamsBodyFieldsItems0](#upload-segment-as-multi-part-params-body-fields-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartParamsBodyMediaType](#upload-segment-as-multi-part-params-body-media-type)| `UploadSegmentAsMultiPartParamsBodyMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartParamsBodyParameterizedHeadersItems0](#upload-segment-as-multi-part-params-body-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ContentDisposition](#upload-segment-as-multi-part-params-body-body-parts-items0-content-disposition)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0MediaType](#upload-segment-as-multi-part-params-body-body-parts-items0-media-type)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-params-body-body-parts-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0Parent](#upload-segment-as-multi-part-params-body-body-parts-items0-parent)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-media-type"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0)| `[]*UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentContentDisposition](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-content-disposition)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentMediaType](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-media-type)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-content-disposition)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-media-type)| `UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-media-type"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-media-type"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-body-parts-items0-parent-parameterized-headers-items0"></span> UploadSegmentAsMultiPartParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-fields-items0"></span> UploadSegmentAsMultiPartParamsBodyFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartParamsBodyFieldsItems0ContentDisposition](#upload-segment-as-multi-part-params-body-fields-items0-content-disposition)| `UploadSegmentAsMultiPartParamsBodyFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [UploadSegmentAsMultiPartParamsBodyFieldsItems0FormDataContentDisposition](#upload-segment-as-multi-part-params-body-fields-items0-form-data-content-disposition)| `UploadSegmentAsMultiPartParamsBodyFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartParamsBodyFieldsItems0MediaType](#upload-segment-as-multi-part-params-body-fields-items0-media-type)| `UploadSegmentAsMultiPartParamsBodyFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartParamsBodyFieldsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-params-body-fields-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartParamsBodyFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-fields-items0-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-fields-items0-form-data-content-disposition"></span> UploadSegmentAsMultiPartParamsBodyFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-fields-items0-media-type"></span> UploadSegmentAsMultiPartParamsBodyFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-fields-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartParamsBodyFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-media-type"></span> UploadSegmentAsMultiPartParamsBodyMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-params-body-parameterized-headers-items0"></span> UploadSegmentAsMultiPartParamsBodyParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="upload-segment-as-multi-part-v2"></span> Upload a segment (*uploadSegmentAsMultiPartV2*)

```
POST /v2/segments
```

Upload a segment as binary

#### Consumes
  * multipart/form-data

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| enableParallelPushProtection | `query` | boolean | `bool` |  |  |  | Whether to enable parallel push protection |
| tableName | `query` | string | `string` |  |  |  | Name of the table |
| body | `body` | [UploadSegmentAsMultiPartV2Body](#upload-segment-as-multi-part-v2-body) | `UploadSegmentAsMultiPartV2Body` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [default](#upload-segment-as-multi-part-v2-default) | | successful operation |  | [schema](#upload-segment-as-multi-part-v2-default-schema) |

#### Responses


##### <span id="upload-segment-as-multi-part-v2-default"></span> Default Response
successful operation

###### <span id="upload-segment-as-multi-part-v2-default-schema"></span> Schema
empty schema

###### Inlined models

**<span id="upload-segment-as-multi-part-v2-body"></span> UploadSegmentAsMultiPartV2Body**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0](#upload-segment-as-multi-part-v2-params-body-body-parts-items0)| `[]*UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0` |  | |  |  |
| contentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyContentDisposition](#upload-segment-as-multi-part-v2-params-body-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0](#upload-segment-as-multi-part-v2-params-body-fields-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartV2ParamsBodyMediaType](#upload-segment-as-multi-part-v2-params-body-media-type)| `UploadSegmentAsMultiPartV2ParamsBodyMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyParameterizedHeadersItems0](#upload-segment-as-multi-part-v2-params-body-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ContentDisposition](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0MediaType](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-media-type)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0Parent](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-media-type"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0)| `[]*UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentContentDisposition](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentMediaType](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-media-type)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-media-type)| `UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-media-type"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-media-type"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-body-parts-items0-parent-parameterized-headers-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-fields-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ContentDisposition](#upload-segment-as-multi-part-v2-params-body-fields-items0-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0FormDataContentDisposition](#upload-segment-as-multi-part-v2-params-body-fields-items0-form-data-content-disposition)| `UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0MediaType](#upload-segment-as-multi-part-v2-params-body-fields-items0-media-type)| `UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ParameterizedHeadersItems0](#upload-segment-as-multi-part-v2-params-body-fields-items0-parameterized-headers-items0)| `map[string][]UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-fields-items0-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-fields-items0-form-data-content-disposition"></span> UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-fields-items0-media-type"></span> UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-fields-items0-parameterized-headers-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-media-type"></span> UploadSegmentAsMultiPartV2ParamsBodyMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="upload-segment-as-multi-part-v2-params-body-parameterized-headers-items0"></span> UploadSegmentAsMultiPartV2ParamsBodyParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="validate-schema"></span> Validate schema (*validateSchema*)

```
POST /schemas/validate
```

This API returns the schema that matches the one you get from 'GET /schema/{schemaName}'. This allows us to validate schema before apply.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| body | `body` | [ValidateSchemaBody](#validate-schema-body) | `ValidateSchemaBody` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#validate-schema-200) | OK | Successfully validated schema |  | [schema](#validate-schema-200-schema) |
| [400](#validate-schema-400) | Bad Request | Missing or invalid request body |  | [schema](#validate-schema-400-schema) |
| [500](#validate-schema-500) | Internal Server Error | Internal error |  | [schema](#validate-schema-500-schema) |

#### Responses


##### <span id="validate-schema-200"></span> 200 - Successfully validated schema
Status: OK

###### <span id="validate-schema-200-schema"></span> Schema

##### <span id="validate-schema-400"></span> 400 - Missing or invalid request body
Status: Bad Request

###### <span id="validate-schema-400-schema"></span> Schema

##### <span id="validate-schema-500"></span> 500 - Internal error
Status: Internal Server Error

###### <span id="validate-schema-500-schema"></span> Schema

###### Inlined models

**<span id="validate-schema-body"></span> ValidateSchemaBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][ValidateSchemaParamsBodyBodyPartsItems0](#validate-schema-params-body-body-parts-items0)| `[]*ValidateSchemaParamsBodyBodyPartsItems0` |  | |  |  |
| contentDisposition | [ValidateSchemaParamsBodyContentDisposition](#validate-schema-params-body-content-disposition)| `ValidateSchemaParamsBodyContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*ValidateSchemaParamsBodyFieldsItems0](#validate-schema-params-body-fields-items0)| `map[string][]ValidateSchemaParamsBodyFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [ValidateSchemaParamsBodyMediaType](#validate-schema-params-body-media-type)| `ValidateSchemaParamsBodyMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*ValidateSchemaParamsBodyParameterizedHeadersItems0](#validate-schema-params-body-parameterized-headers-items0)| `map[string][]ValidateSchemaParamsBodyParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0"></span> ValidateSchemaParamsBodyBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [ValidateSchemaParamsBodyBodyPartsItems0ContentDisposition](#validate-schema-params-body-body-parts-items0-content-disposition)| `ValidateSchemaParamsBodyBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [ValidateSchemaParamsBodyBodyPartsItems0MediaType](#validate-schema-params-body-body-parts-items0-media-type)| `ValidateSchemaParamsBodyBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*ValidateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0](#validate-schema-params-body-body-parts-items0-parameterized-headers-items0)| `map[string][]ValidateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [ValidateSchemaParamsBodyBodyPartsItems0Parent](#validate-schema-params-body-body-parts-items0-parent)| `ValidateSchemaParamsBodyBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-content-disposition"></span> ValidateSchemaParamsBodyBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-media-type"></span> ValidateSchemaParamsBodyBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parameterized-headers-items0"></span> ValidateSchemaParamsBodyBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent"></span> ValidateSchemaParamsBodyBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0](#validate-schema-params-body-body-parts-items0-parent-body-parts-items0)| `[]*ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [ValidateSchemaParamsBodyBodyPartsItems0ParentContentDisposition](#validate-schema-params-body-body-parts-items0-parent-content-disposition)| `ValidateSchemaParamsBodyBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [ValidateSchemaParamsBodyBodyPartsItems0ParentMediaType](#validate-schema-params-body-body-parts-items0-parent-media-type)| `ValidateSchemaParamsBodyBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*ValidateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0](#validate-schema-params-body-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]ValidateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-body-parts-items0"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#validate-schema-params-body-body-parts-items0-parent-body-parts-items0-content-disposition)| `ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType](#validate-schema-params-body-body-parts-items0-parent-body-parts-items0-media-type)| `ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#validate-schema-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-body-parts-items0-content-disposition"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-body-parts-items0-media-type"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-content-disposition"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-media-type"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="validate-schema-params-body-body-parts-items0-parent-parameterized-headers-items0"></span> ValidateSchemaParamsBodyBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-content-disposition"></span> ValidateSchemaParamsBodyContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-fields-items0"></span> ValidateSchemaParamsBodyFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [ValidateSchemaParamsBodyFieldsItems0ContentDisposition](#validate-schema-params-body-fields-items0-content-disposition)| `ValidateSchemaParamsBodyFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [ValidateSchemaParamsBodyFieldsItems0FormDataContentDisposition](#validate-schema-params-body-fields-items0-form-data-content-disposition)| `ValidateSchemaParamsBodyFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [ValidateSchemaParamsBodyFieldsItems0MediaType](#validate-schema-params-body-fields-items0-media-type)| `ValidateSchemaParamsBodyFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*ValidateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0](#validate-schema-params-body-fields-items0-parameterized-headers-items0)| `map[string][]ValidateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `models.MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-fields-items0-content-disposition"></span> ValidateSchemaParamsBodyFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-fields-items0-form-data-content-disposition"></span> ValidateSchemaParamsBodyFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-fields-items0-media-type"></span> ValidateSchemaParamsBodyFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="validate-schema-params-body-fields-items0-parameterized-headers-items0"></span> ValidateSchemaParamsBodyFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="validate-schema-params-body-media-type"></span> ValidateSchemaParamsBodyMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="validate-schema-params-body-parameterized-headers-items0"></span> ValidateSchemaParamsBodyParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



## Models

### <span id="body-part"></span> BodyPart


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [BodyPartContentDisposition](#body-part-content-disposition)| `BodyPartContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [BodyPartMediaType](#body-part-media-type)| `BodyPartMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*BodyPartParameterizedHeadersItems0](#body-part-parameterized-headers-items0)| `map[string][]BodyPartParameterizedHeadersItems0` |  | |  |  |
| parent | [BodyPartParent](#body-part-parent)| `BodyPartParent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



#### Inlined models

**<span id="body-part-content-disposition"></span> BodyPartContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="body-part-media-type"></span> BodyPartMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="body-part-parameterized-headers-items0"></span> BodyPartParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="body-part-parent"></span> BodyPartParent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][BodyPartParentBodyPartsItems0](#body-part-parent-body-parts-items0)| `[]*BodyPartParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [BodyPartParentContentDisposition](#body-part-parent-content-disposition)| `BodyPartParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [BodyPartParentMediaType](#body-part-parent-media-type)| `BodyPartParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*BodyPartParentParameterizedHeadersItems0](#body-part-parent-parameterized-headers-items0)| `map[string][]BodyPartParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="body-part-parent-body-parts-items0"></span> BodyPartParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [BodyPartParentBodyPartsItems0ContentDisposition](#body-part-parent-body-parts-items0-content-disposition)| `BodyPartParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [BodyPartParentBodyPartsItems0MediaType](#body-part-parent-body-parts-items0-media-type)| `BodyPartParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*BodyPartParentBodyPartsItems0ParameterizedHeadersItems0](#body-part-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]BodyPartParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="body-part-parent-body-parts-items0-content-disposition"></span> BodyPartParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="body-part-parent-body-parts-items0-media-type"></span> BodyPartParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="body-part-parent-body-parts-items0-parameterized-headers-items0"></span> BodyPartParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="body-part-parent-content-disposition"></span> BodyPartParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="body-part-parent-media-type"></span> BodyPartParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="body-part-parent-parameterized-headers-items0"></span> BodyPartParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="consuming-segment-info"></span> ConsumingSegmentInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| _consumerState | string| `string` |  | |  |  |
| _lastConsumedTimestamp | int64 (formatted integer)| `int64` |  | |  |  |
| _partitionToOffsetMap | map of string| `map[string]string` |  | |  |  |
| _serverName | string| `string` |  | |  |  |



### <span id="consuming-segments-info-map"></span> ConsumingSegmentsInfoMap


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| _segmentToConsumingInfoMap | map of [[]*ConsumingSegmentsInfoMapSegmentToConsumingInfoMapItems0](#consuming-segments-info-map-segment-to-consuming-info-map-items0)| `map[string][]ConsumingSegmentsInfoMapSegmentToConsumingInfoMapItems0` |  | |  |  |



#### Inlined models

**<span id="consuming-segments-info-map-segment-to-consuming-info-map-items0"></span> ConsumingSegmentsInfoMapSegmentToConsumingInfoMapItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| _consumerState | string| `string` |  | |  |  |
| _lastConsumedTimestamp | int64 (formatted integer)| `int64` |  | |  |  |
| _partitionToOffsetMap | map of string| `map[string]string` |  | |  |  |
| _serverName | string| `string` |  | |  |  |



### <span id="content-disposition"></span> ContentDisposition


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="date-time-field-spec"></span> DateTimeFieldSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| format | string| `string` |  | |  |  |
| granularity | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



### <span id="dimension-field-spec"></span> DimensionFieldSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



### <span id="form-data-body-part"></span> FormDataBodyPart


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [FormDataBodyPartContentDisposition](#form-data-body-part-content-disposition)| `FormDataBodyPartContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [FormDataBodyPartFormDataContentDisposition](#form-data-body-part-form-data-content-disposition)| `FormDataBodyPartFormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataBodyPartMediaType](#form-data-body-part-media-type)| `FormDataBodyPartMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataBodyPartParameterizedHeadersItems0](#form-data-body-part-parameterized-headers-items0)| `map[string][]FormDataBodyPartParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



#### Inlined models

**<span id="form-data-body-part-content-disposition"></span> FormDataBodyPartContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-body-part-form-data-content-disposition"></span> FormDataBodyPartFormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-body-part-media-type"></span> FormDataBodyPartMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-body-part-parameterized-headers-items0"></span> FormDataBodyPartParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="form-data-content-disposition"></span> FormDataContentDisposition


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



### <span id="form-data-multi-part"></span> FormDataMultiPart


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][FormDataMultiPartBodyPartsItems0](#form-data-multi-part-body-parts-items0)| `[]*FormDataMultiPartBodyPartsItems0` |  | |  |  |
| contentDisposition | [FormDataMultiPartContentDisposition](#form-data-multi-part-content-disposition)| `FormDataMultiPartContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| fields | map of [[]*FormDataMultiPartFieldsItems0](#form-data-multi-part-fields-items0)| `map[string][]FormDataMultiPartFieldsItems0` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataMultiPartMediaType](#form-data-multi-part-media-type)| `FormDataMultiPartMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataMultiPartParameterizedHeadersItems0](#form-data-multi-part-parameterized-headers-items0)| `map[string][]FormDataMultiPartParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



#### Inlined models

**<span id="form-data-multi-part-body-parts-items0"></span> FormDataMultiPartBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [FormDataMultiPartBodyPartsItems0ContentDisposition](#form-data-multi-part-body-parts-items0-content-disposition)| `FormDataMultiPartBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataMultiPartBodyPartsItems0MediaType](#form-data-multi-part-body-parts-items0-media-type)| `FormDataMultiPartBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataMultiPartBodyPartsItems0ParameterizedHeadersItems0](#form-data-multi-part-body-parts-items0-parameterized-headers-items0)| `map[string][]FormDataMultiPartBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [FormDataMultiPartBodyPartsItems0Parent](#form-data-multi-part-body-parts-items0-parent)| `FormDataMultiPartBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-content-disposition"></span> FormDataMultiPartBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-media-type"></span> FormDataMultiPartBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parameterized-headers-items0"></span> FormDataMultiPartBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent"></span> FormDataMultiPartBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0](#form-data-multi-part-body-parts-items0-parent-body-parts-items0)| `[]*FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [FormDataMultiPartBodyPartsItems0ParentContentDisposition](#form-data-multi-part-body-parts-items0-parent-content-disposition)| `FormDataMultiPartBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataMultiPartBodyPartsItems0ParentMediaType](#form-data-multi-part-body-parts-items0-parent-media-type)| `FormDataMultiPartBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataMultiPartBodyPartsItems0ParentParameterizedHeadersItems0](#form-data-multi-part-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]FormDataMultiPartBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-body-parts-items0"></span> FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#form-data-multi-part-body-parts-items0-parent-body-parts-items0-content-disposition)| `FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0MediaType](#form-data-multi-part-body-parts-items0-parent-body-parts-items0-media-type)| `FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#form-data-multi-part-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-body-parts-items0-content-disposition"></span> FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-body-parts-items0-media-type"></span> FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> FormDataMultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-content-disposition"></span> FormDataMultiPartBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-media-type"></span> FormDataMultiPartBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-multi-part-body-parts-items0-parent-parameterized-headers-items0"></span> FormDataMultiPartBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="form-data-multi-part-content-disposition"></span> FormDataMultiPartContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-fields-items0"></span> FormDataMultiPartFieldsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [FormDataMultiPartFieldsItems0ContentDisposition](#form-data-multi-part-fields-items0-content-disposition)| `FormDataMultiPartFieldsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| formDataContentDisposition | [FormDataMultiPartFieldsItems0FormDataContentDisposition](#form-data-multi-part-fields-items0-form-data-content-disposition)| `FormDataMultiPartFieldsItems0FormDataContentDisposition` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [FormDataMultiPartFieldsItems0MediaType](#form-data-multi-part-fields-items0-media-type)| `FormDataMultiPartFieldsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameterizedHeaders | map of [[]*FormDataMultiPartFieldsItems0ParameterizedHeadersItems0](#form-data-multi-part-fields-items0-parameterized-headers-items0)| `map[string][]FormDataMultiPartFieldsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |
| simple | boolean| `bool` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="form-data-multi-part-fields-items0-content-disposition"></span> FormDataMultiPartFieldsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-fields-items0-form-data-content-disposition"></span> FormDataMultiPartFieldsItems0FormDataContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| name | string| `string` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="form-data-multi-part-fields-items0-media-type"></span> FormDataMultiPartFieldsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-multi-part-fields-items0-parameterized-headers-items0"></span> FormDataMultiPartFieldsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="form-data-multi-part-media-type"></span> FormDataMultiPartMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="form-data-multi-part-parameterized-headers-items0"></span> FormDataMultiPartParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="instance"></span> Instance


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| grpcPort | int32 (formatted integer)| `int32` |  | |  |  |
| host | string| `string` | ✓ | |  |  |
| pools | map of int32 (formatted integer)| `map[string]int32` |  | |  |  |
| port | int32 (formatted integer)| `int32` | ✓ | |  |  |
| tags | []string| `[]string` |  | |  |  |
| type | string| `string` | ✓ | |  |  |



### <span id="instance-info"></span> InstanceInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| host | string| `string` |  | |  |  |
| instanceName | string| `string` |  | |  |  |
| port | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="instance-partitions"></span> InstancePartitions


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="instances"></span> Instances


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instances | []string| `[]string` |  | |  |  |



### <span id="job-key"></span> JobKey


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| group | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="lead-controller-entry"></span> LeadControllerEntry


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerId | string| `string` |  | |  |  |
| tableNames | []string| `[]string` |  | |  |  |



### <span id="lead-controller-response"></span> LeadControllerResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerEntryMap | map of [LeadControllerResponseLeadControllerEntryMapAnon](#lead-controller-response-lead-controller-entry-map-anon)| `map[string]LeadControllerResponseLeadControllerEntryMapAnon` |  | |  |  |
| leadControllerResourceEnabled | boolean| `bool` |  | |  |  |



#### Inlined models

**<span id="lead-controller-response-lead-controller-entry-map-anon"></span> LeadControllerResponseLeadControllerEntryMapAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| leadControllerId | string| `string` |  | |  |  |
| tableNames | []string| `[]string` |  | |  |  |



### <span id="media-type"></span> MediaType


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



### <span id="message-body-workers"></span> MessageBodyWorkers


  

[interface{}](#interface)

### <span id="metric-field-spec"></span> MetricFieldSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



### <span id="multi-part"></span> MultiPart


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][MultiPartBodyPartsItems0](#multi-part-body-parts-items0)| `[]*MultiPartBodyPartsItems0` |  | |  |  |
| contentDisposition | [MultiPartContentDisposition](#multi-part-content-disposition)| `MultiPartContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [MultiPartMediaType](#multi-part-media-type)| `MultiPartMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*MultiPartParameterizedHeadersItems0](#multi-part-parameterized-headers-items0)| `map[string][]MultiPartParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



#### Inlined models

**<span id="multi-part-body-parts-items0"></span> MultiPartBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [MultiPartBodyPartsItems0ContentDisposition](#multi-part-body-parts-items0-content-disposition)| `MultiPartBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [MultiPartBodyPartsItems0MediaType](#multi-part-body-parts-items0-media-type)| `MultiPartBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*MultiPartBodyPartsItems0ParameterizedHeadersItems0](#multi-part-body-parts-items0-parameterized-headers-items0)| `map[string][]MultiPartBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPartBodyPartsItems0Parent](#multi-part-body-parts-items0-parent)| `MultiPartBodyPartsItems0Parent` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="multi-part-body-parts-items0-content-disposition"></span> MultiPartBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="multi-part-body-parts-items0-media-type"></span> MultiPartBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="multi-part-body-parts-items0-parameterized-headers-items0"></span> MultiPartBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent"></span> MultiPartBodyPartsItems0Parent**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bodyParts | [][MultiPartBodyPartsItems0ParentBodyPartsItems0](#multi-part-body-parts-items0-parent-body-parts-items0)| `[]*MultiPartBodyPartsItems0ParentBodyPartsItems0` |  | |  |  |
| contentDisposition | [MultiPartBodyPartsItems0ParentContentDisposition](#multi-part-body-parts-items0-parent-content-disposition)| `MultiPartBodyPartsItems0ParentContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [MultiPartBodyPartsItems0ParentMediaType](#multi-part-body-parts-items0-parent-media-type)| `MultiPartBodyPartsItems0ParentMediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*MultiPartBodyPartsItems0ParentParameterizedHeadersItems0](#multi-part-body-parts-items0-parent-parameterized-headers-items0)| `map[string][]MultiPartBodyPartsItems0ParentParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-body-parts-items0"></span> MultiPartBodyPartsItems0ParentBodyPartsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| contentDisposition | [MultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition](#multi-part-body-parts-items0-parent-body-parts-items0-content-disposition)| `MultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition` |  | |  |  |
| entity | [interface{}](#interface)| `interface{}` |  | |  |  |
| headers | map of [[]string](#string)| `map[string][]string` |  | |  |  |
| mediaType | [MultiPartBodyPartsItems0ParentBodyPartsItems0MediaType](#multi-part-body-parts-items0-parent-body-parts-items0-media-type)| `MultiPartBodyPartsItems0ParentBodyPartsItems0MediaType` |  | |  |  |
| messageBodyWorkers | [interface{}](#interface)| `interface{}` |  | |  |  |
| parameterizedHeaders | map of [[]*MultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0](#multi-part-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0)| `map[string][]MultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0` |  | |  |  |
| parent | [MultiPart](#multi-part)| `MultiPart` |  | |  |  |
| providers | [interface{}](#interface)| `interface{}` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-body-parts-items0-content-disposition"></span> MultiPartBodyPartsItems0ParentBodyPartsItems0ContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-body-parts-items0-media-type"></span> MultiPartBodyPartsItems0ParentBodyPartsItems0MediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-body-parts-items0-parameterized-headers-items0"></span> MultiPartBodyPartsItems0ParentBodyPartsItems0ParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-content-disposition"></span> MultiPartBodyPartsItems0ParentContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-media-type"></span> MultiPartBodyPartsItems0ParentMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="multi-part-body-parts-items0-parent-parameterized-headers-items0"></span> MultiPartBodyPartsItems0ParentParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



**<span id="multi-part-content-disposition"></span> MultiPartContentDisposition**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| fileName | string| `string` |  | |  |  |
| modificationDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| parameters | map of string| `map[string]string` |  | |  |  |
| readDate | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| size | int64 (formatted integer)| `int64` |  | |  |  |
| type | string| `string` |  | |  |  |



**<span id="multi-part-media-type"></span> MultiPartMediaType**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| subtype | string| `string` |  | |  |  |
| type | string| `string` |  | |  |  |
| wildcardSubtype | boolean| `bool` |  | |  |  |
| wildcardType | boolean| `bool` |  | |  |  |



**<span id="multi-part-parameterized-headers-items0"></span> MultiPartParameterizedHeadersItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="parameterized-header"></span> ParameterizedHeader


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| parameters | map of string| `map[string]string` |  | |  |  |
| value | string| `string` |  | |  |  |



### <span id="pinot-task-config"></span> PinotTaskConfig


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| configs | map of string| `map[string]string` |  | |  |  |
| taskType | string| `string` |  | |  |  |



### <span id="providers"></span> Providers


  

[interface{}](#interface)

### <span id="rebalance-result"></span> RebalanceResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| description | string| `string` |  | |  |  |
| instanceAssignment | map of [RebalanceResultInstanceAssignmentAnon](#rebalance-result-instance-assignment-anon)| `map[string]RebalanceResultInstanceAssignmentAnon` |  | |  |  |
| segmentAssignment | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |
| status | string| `string` |  | |  |  |



#### Inlined models

**<span id="rebalance-result-instance-assignment-anon"></span> RebalanceResultInstanceAssignmentAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| instancePartitionsName | string| `string` |  | |  |  |
| partitionToInstancesMap | map of [[]string](#string)| `map[string][]string` |  | |  |  |



### <span id="schema"></span> Schema


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dateTimeFieldSpecs | [][SchemaDateTimeFieldSpecsItems0](#schema-date-time-field-specs-items0)| `[]*SchemaDateTimeFieldSpecsItems0` |  | |  |  |
| dimensionFieldSpecs | [][SchemaDimensionFieldSpecsItems0](#schema-dimension-field-specs-items0)| `[]*SchemaDimensionFieldSpecsItems0` |  | |  |  |
| metricFieldSpecs | [][SchemaMetricFieldSpecsItems0](#schema-metric-field-specs-items0)| `[]*SchemaMetricFieldSpecsItems0` |  | |  |  |
| primaryKeyColumns | []string| `[]string` |  | |  |  |
| schemaName | string| `string` |  | |  |  |
| timeFieldSpec | [SchemaTimeFieldSpec](#schema-time-field-spec)| `SchemaTimeFieldSpec` |  | |  |  |



#### Inlined models

**<span id="schema-date-time-field-specs-items0"></span> SchemaDateTimeFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| format | string| `string` |  | |  |  |
| granularity | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="schema-dimension-field-specs-items0"></span> SchemaDimensionFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="schema-metric-field-specs-items0"></span> SchemaMetricFieldSpecsItems0**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="schema-time-field-spec"></span> SchemaTimeFieldSpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| incomingGranularitySpec | [SchemaTimeFieldSpecIncomingGranularitySpec](#schema-time-field-spec-incoming-granularity-spec)| `SchemaTimeFieldSpecIncomingGranularitySpec` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| outgoingGranularitySpec | [SchemaTimeFieldSpecOutgoingGranularitySpec](#schema-time-field-spec-outgoing-granularity-spec)| `SchemaTimeFieldSpecOutgoingGranularitySpec` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



**<span id="schema-time-field-spec-incoming-granularity-spec"></span> SchemaTimeFieldSpecIncomingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



**<span id="schema-time-field-spec-outgoing-granularity-spec"></span> SchemaTimeFieldSpecOutgoingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="segment-size-details"></span> SegmentSizeDetails


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| serverInfo | map of [SegmentSizeDetailsServerInfoAnon](#segment-size-details-server-info-anon)| `map[string]SegmentSizeDetailsServerInfoAnon` |  | |  |  |



#### Inlined models

**<span id="segment-size-details-server-info-anon"></span> SegmentSizeDetailsServerInfoAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| diskSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segmentName | string| `string` |  | |  |  |



### <span id="segment-size-info"></span> SegmentSizeInfo


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| diskSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segmentName | string| `string` |  | |  |  |



### <span id="start-replace-segments-request"></span> StartReplaceSegmentsRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| segmentsFrom | []string| `[]string` |  | |  |  |
| segmentsTo | []string| `[]string` |  | |  |  |



### <span id="string-result-response"></span> StringResultResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` |  | |  |  |



### <span id="success-response"></span> SuccessResponse


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| status | string| `string` |  | |  |  |



### <span id="table-size-details"></span> TableSizeDetails


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| offlineSegments | [TableSizeDetailsOfflineSegments](#table-size-details-offline-segments)| `TableSizeDetailsOfflineSegments` |  | |  |  |
| realtimeSegments | [TableSizeDetailsRealtimeSegments](#table-size-details-realtime-segments)| `TableSizeDetailsRealtimeSegments` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| tableName | string| `string` |  | |  |  |



#### Inlined models

**<span id="table-size-details-offline-segments"></span> TableSizeDetailsOfflineSegments**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| missingSegments | int32 (formatted integer)| `int32` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segments | map of [TableSizeDetailsOfflineSegmentsSegmentsAnon](#table-size-details-offline-segments-segments-anon)| `map[string]TableSizeDetailsOfflineSegmentsSegmentsAnon` |  | |  |  |



**<span id="table-size-details-offline-segments-segments-anon"></span> TableSizeDetailsOfflineSegmentsSegmentsAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| serverInfo | map of [TableSizeDetailsOfflineSegmentsSegmentsAnonServerInfoAnon](#table-size-details-offline-segments-segments-anon-server-info-anon)| `map[string]TableSizeDetailsOfflineSegmentsSegmentsAnonServerInfoAnon` |  | |  |  |



**<span id="table-size-details-offline-segments-segments-anon-server-info-anon"></span> TableSizeDetailsOfflineSegmentsSegmentsAnonServerInfoAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| diskSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segmentName | string| `string` |  | |  |  |



**<span id="table-size-details-realtime-segments"></span> TableSizeDetailsRealtimeSegments**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| missingSegments | int32 (formatted integer)| `int32` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segments | map of [TableSizeDetailsRealtimeSegmentsSegmentsAnon](#table-size-details-realtime-segments-segments-anon)| `map[string]TableSizeDetailsRealtimeSegmentsSegmentsAnon` |  | |  |  |



**<span id="table-size-details-realtime-segments-segments-anon"></span> TableSizeDetailsRealtimeSegmentsSegmentsAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| serverInfo | map of [TableSizeDetailsRealtimeSegmentsSegmentsAnonServerInfoAnon](#table-size-details-realtime-segments-segments-anon-server-info-anon)| `map[string]TableSizeDetailsRealtimeSegmentsSegmentsAnonServerInfoAnon` |  | |  |  |



**<span id="table-size-details-realtime-segments-segments-anon-server-info-anon"></span> TableSizeDetailsRealtimeSegmentsSegmentsAnonServerInfoAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| diskSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segmentName | string| `string` |  | |  |  |



### <span id="table-sub-type-size-details"></span> TableSubTypeSizeDetails


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| missingSegments | int32 (formatted integer)| `int32` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segments | map of [TableSubTypeSizeDetailsSegmentsAnon](#table-sub-type-size-details-segments-anon)| `map[string]TableSubTypeSizeDetailsSegmentsAnon` |  | |  |  |



#### Inlined models

**<span id="table-sub-type-size-details-segments-anon"></span> TableSubTypeSizeDetailsSegmentsAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| estimatedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| reportedSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| serverInfo | map of [TableSubTypeSizeDetailsSegmentsAnonServerInfoAnon](#table-sub-type-size-details-segments-anon-server-info-anon)| `map[string]TableSubTypeSizeDetailsSegmentsAnonServerInfoAnon` |  | |  |  |



**<span id="table-sub-type-size-details-segments-anon-server-info-anon"></span> TableSubTypeSizeDetailsSegmentsAnonServerInfoAnon**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| diskSizeInBytes | int64 (formatted integer)| `int64` |  | |  |  |
| segmentName | string| `string` |  | |  |  |



### <span id="table-view"></span> TableView


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| OFFLINE | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |
| REALTIME | map of [map[string]string](#map-string-string)| `map[string]map[string]string` |  | |  |  |



### <span id="tenant"></span> Tenant


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| numberOfInstances | int32 (formatted integer)| `int32` |  | |  |  |
| offlineInstances | int32 (formatted integer)| `int32` |  | |  |  |
| realtimeInstances | int32 (formatted integer)| `int32` |  | |  |  |
| tenantName | string| `string` | ✓ | |  |  |
| tenantRole | string| `string` | ✓ | |  |  |



### <span id="tenant-metadata"></span> TenantMetadata


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| BrokerInstances | []string| `[]string` |  | |  |  |
| ServerInstances | []string| `[]string` |  | |  |  |
| tenantName | string| `string` |  | |  |  |



### <span id="tenants-list"></span> TenantsList


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| BROKER_TENANTS | []string| `[]string` |  | |  |  |
| SERVER_TENANTS | []string| `[]string` |  | |  |  |



### <span id="time-field-spec"></span> TimeFieldSpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| defaultNullValue | [interface{}](#interface)| `interface{}` |  | |  |  |
| defaultNullValueString | string| `string` |  | |  |  |
| incomingGranularitySpec | [TimeFieldSpecIncomingGranularitySpec](#time-field-spec-incoming-granularity-spec)| `TimeFieldSpecIncomingGranularitySpec` |  | |  |  |
| maxLength | int32 (formatted integer)| `int32` |  | |  |  |
| name | string| `string` |  | |  |  |
| outgoingGranularitySpec | [TimeFieldSpecOutgoingGranularitySpec](#time-field-spec-outgoing-granularity-spec)| `TimeFieldSpecOutgoingGranularitySpec` |  | |  |  |
| singleValueField | boolean| `bool` |  | |  |  |
| transformFunction | string| `string` |  | |  |  |
| virtualColumnProvider | string| `string` |  | |  |  |



#### Inlined models

**<span id="time-field-spec-incoming-granularity-spec"></span> TimeFieldSpecIncomingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



**<span id="time-field-spec-outgoing-granularity-spec"></span> TimeFieldSpecOutgoingGranularitySpec**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



### <span id="time-granularity-spec"></span> TimeGranularitySpec


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dataType | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| timeFormat | string| `string` |  | |  |  |
| timeType | string| `string` |  | |  |  |
| timeUnitSize | int32 (formatted integer)| `int32` |  | |  |  |



/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// TaskrouterV1TaskChannel struct for TaskrouterV1TaskChannel
type TaskrouterV1TaskChannel struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Task Channel resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that we created to identify the Task Channel resource.
	Sid *string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the Task Channel, such as `voice` or `sms`.
	UniqueName *string `json:"unique_name,omitempty"`
	// The SID of the Workspace that contains the Task Channel.
	WorkspaceSid *string `json:"workspace_sid,omitempty"`
	// Whether the Task Channel will prioritize Workers that have been idle. When `true`, Workers that have been idle the longest are prioritized.
	ChannelOptimizedRouting *bool `json:"channel_optimized_routing,omitempty"`
	// The absolute URL of the Task Channel resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
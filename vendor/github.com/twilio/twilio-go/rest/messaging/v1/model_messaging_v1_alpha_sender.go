/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// MessagingV1AlphaSender struct for MessagingV1AlphaSender
type MessagingV1AlphaSender struct {
	// The unique string that we created to identify the AlphaSender resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AlphaSender resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The Alphanumeric Sender ID string.
	AlphaSender *string `json:"alpha_sender,omitempty"`
	// An array of values that describe whether the number can receive calls or messages. Can be: `SMS`.
	Capabilities *[]string `json:"capabilities,omitempty"`
	// The absolute URL of the AlphaSender resource.
	Url *string `json:"url,omitempty"`
}

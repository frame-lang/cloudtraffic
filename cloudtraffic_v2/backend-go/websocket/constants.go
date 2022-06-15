// Contains the IDs used in Cloud PubSub

package websocket

var PROJECT_ID string = "oceanic-oxide-303322"

/**
 * TL_TOPIC_ID: Send events from Utils service to TL service
 * A cloud function is added as a subscriber to this topic
 * So, when we publish events to this topic, as a result cloud function will get trigeered
 */
var TL_TOPIC_ID string = "cloudtraffic-trafficlight-service-topic"

/**
 * UTILS_TOPIC_ID: This topic is added inside the TL service cloud function
 * Send events from TL service to Utils service
 * A pull subscriber (cloudtraffic-utils-service-sub) is added here in Utils service code
 * So, when we publish events to this topic from cloud function, as a result the pull subscriber is get called
 * We need this variable to this to create a new subscription in this topic, only if already not existed
 */
var UTILS_TOPIC_ID string = "cloudtraffic-utils-service-topic"

var SUBSCRIPTION_ID string = "cloudtraffic-utils-service-sub"
package chapter4

// The first step to start consuming records is to create a KafkaConsumer instance
// We need to use the three mandatory properties: bootstrap.servers, key.deserializer, and value.deserializer
// There is a fourth property, which is not strictly mandatory but very commonyly used. The property is group.id and
// is is specifies the consumer group the KafkaConsumer instance belong to.

// Subscribing to Topics
// Once we create a consumer, the next step is to subscribe to one or more topics.

// The Poll Loop

// Thread Safety
/*
 */

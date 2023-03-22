# 

There has been a huge gap in our infrastructure for data. Traditionally, data management was all about storage - the file stores and databases that keep our data safe and let us look up the right bit at the right time. A modern company is an incredibly complex system built out of hundreds or even thousands of custom applications, microservices, databases, SaaS layers and analytics platform.

## Publish/Subscribe Messaging

Apache Kafka is often described as a "distributed commit log" or more recently as a "distributed streaming platform". Similary, data within Kafka is stored durably, in order, and can be read deterministically. In addition, the data can be distributed within the system to provide additional protections against failures.

**Messages and Batches**

The unit of data within Kafka is called a message. Message in kafka does not have a specific format or meaning to Kafka. A message can have an optional piece of metadata, which is referred to as a key. The key is also a byte array and, as with the message, has no specific meaning to Kafka. Keys are used when messages are to be written to partitoins in a more controlled manner. The simplest such scheme is to generate a consistent hash of the key and then select the partition number for that message by taking the result of the hash modulo the total number of partitions in the topic. This ensures that messages with the same key are always written to the same partition.

For efficiency, messages are written into Kafka is batches. A batch is just a collection or messages, all of with are being produced to the same topic and partition. An individual round trip accross the network for each message would result in excessive over head, and collecting messages together into a batch reduces this. Of course, this is a trade-off between latency and throughput: the larger the batches, the more messages that can be handled per unit of time, but the longer it takes an individual message to propagate.

**Schemas**

While messages are opaque byte arrays to Kafka itself, it is recommended that additional structure, or schema, be imposed on the message content so that it can be understood.(JSON, XML)

**Topics and Partitions**

Messages in Kafka are categoried into topics. Topics are additionally broken down into a number of partitions. A topic typically has multiple partitions. Partitions are also the way that Kafka provides redundancy and scalability. Each partition can be hosted on a different server, which means that a single topic can be scaled horizontally across multiple servers to provide performance for beyond the ability of a single server. Partitions can be replicated, such that different servers will store a copy of the same partition in case one server fails.

**Producers and Consumers**

A message will be produced to a specific topic. By default, the producer will balance messages over all partitions of a topic evenly. In some cases, the producer will direct messages to specific partitions. This is typically done using the message key and a partitioner that will generate a hash of the key and map it to a specific partition. 

The consumer subscribes to one or more topics and reads the messages in the order in which the were produced to each partition. The consumer keep tracks of which messages it has already consumed by keeping track of the offset of messages. By storing the next posible offset for each partition, typically in Kafka itself, a consumer can stop and restart without losing its place.

Consumers work as part of consumer group, which is one or more consumers that work together to consume a topic. The group ensures that each partition is only consumed by one member. The mapping of a consumer to a partition is often called ownership of the partition by the customer.

**Brokers and Clusters**

A single kafka server is called a broker. 

Kafka brokers are designed to operate as part of cluster. Within a cluster of brokers, one broker will also function as the cluster controller (elected automatically from the live members of the cluster). The controller is responsible for administrative operations, including assign partitions to brokers and monitoring for broker failures. A partition is owned by a single broker in the cluster, and that broker is called the leader of the partition. A replicated partition is assigned to additional brokers, called followers of the partition.
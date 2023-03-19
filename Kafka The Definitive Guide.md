# 

There has been a huge gap in our infrastructure for data. Traditionally, data management was all about storage - the file stores and databases that keep our data safe and let us look up the right bit at the right time. A modern company is an incredibly complex system built out of hundreds or even thousands of custom applications, microservices, databases, SaaS layers and analytics platform.

## Publish/Subscribe Messaging

Apache Kafka is often described as a "distributed commit log" or more recently as a "distributed streaming platform". Similary, data within Kafka is stored durably, in order, and can be read deterministically. In addition, the data can be distributed within the system to provide additional protections against failures.

**Messages and Batches**

The unit of data within Kafka is called a message. Message in kafka does not have a specific format or meaning to Kafka. A message can have an optional piece of metadata, which is referred to as a key. The key is also a byte array and, as with the message, has no specific meaning to Kafka. Keys are used when messages are to be written to partitoins in a more controlled manner. The simplest such scheme is to generate a consistent hash of the key and then select the partition number for that message by taking the result of the hash modulo the total number of partitions in the topic. This ensures that messages with the same key are always written to the same partition.

For efficiency, messages are written into Kafka is batches. A batch is just a collection or messages, all of with are being produced to the same topic and partition. An individual round trip accross the network for each message would result in excessive over head, and collecting messages together into a batch reduces this. Of course, this is a trade-off between latency and throughput: the larger the batches, the more messages that can be handled per unit of time, but the longer it takes an individual message to propagate.

**Schemas**

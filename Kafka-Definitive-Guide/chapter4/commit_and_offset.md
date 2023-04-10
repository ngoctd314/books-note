# Commits and Offsets

Whenever we call poll(), it returns records written to Kafka that consumers in our group have not read yet. This means that we have a way of tracking which records were read by a consumer of the group.

We call the action of updating the current position in the partition an offset commit. Unlike traditional message queues, Kafka does not commit records individually. Instead, consumers commit the last message they've successfully processed from a partition and implicitly assume that every message before the last was also successfully processed.

Problem, if a comsumer crashes or a new consumer joins the consumer group, this will trigger a rebalance. After a rebalance, each consumer may be assigned a new set of partitions than the one it processed before. In order to know where to pick up the work, the consumer will read the latest committed offset of each partition and continue from there.

If the commited offset is smaller than the offset of the last message the client processed, the messages between the last processed offset and the commited offset will be processed twice.

## Automatic Commit

Allow the consumer to do it. By config enable.auto.commit = true, then every 5s the consumer will commit the lastest offset that your client received from poll().

Automatic commits are convenient, but they don't give devlopers enough control to avoid duplicate messages.

## Commit Current Offset

Auto commit base on timer, so we need a way to get more control that makes sense to the application developer rather than based on a timer.   

By setting enable.auto.commit = false, offsets will only be committed when the application explicity chooses to do so.

It is important to remember that commitSync() will commit the lastest offset returned by poll(), so if you call commitSync() before you done processing all the records in the collection, you risk missing the messages that were committed but not processed, in case the application crashes.

**Synchronous commit**
**Asynchronous Commit**
**Combining Synchronous and Asynchronous Commits**
**Commiting a Specified Offset**

## Consuming Records with Specific Offsets
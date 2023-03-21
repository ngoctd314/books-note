# Redis in action

## Getting to know Redis

Redis is an in-memory remote database that offers high performance, replication and a unique data model to produce a platform for solving problem. Redis is a very fast non-relational database that stores a mapping of keys to five different types of values. Redis supports in-memory persistent storage on disk, replication to scale read performance, and client-side sharding to scale write performance.

|Name|Type|Data storage options|Query types|Additional features|
|-|-|-|-|-|
|Redis|In-memory, non-relational, database|Strings, lists, sets, hashes, sorted sets|Commands for each data type for common access patterns, with bulk operations, and partial transaction support|Publish/Subscribe master/slave replication, disk persistence scripting (stored procedures)|

When using an in-memory database like Redis, one of the first questions that's asked is "What happens when my server gets turned off". Redis has two different forms of persistence available for writing in-memory data to disk in a compact format. The other method uses an append-only file that writes every command that alters data in Redis to disk as it happens (sync, sync once per second, or sync at the completion of every operation).

|Structure type|What it contains|Structure read/write ability|
|-|-|-|
|String|Strings, integers, or floating-point values|Operate on the whole string, parts, increment, decrement the integers and floats|
|List|Linked list of string|Push or pop items from both ends, trim based on offsets, read individual or multiple items, find or remove items by value|
|Set|Unordered collection of unique strings|Add, fetch, or remove individual items, check membership, intersect, union, difference, fetch random items|
|Hash|Unordered hash table of keys to values|Add, fetch, or remove individual items, fetch the whole hash|
|ZSet|Ordered mapping of string members to floating-point scores, ordered by score|Add, fetch, or remove individual values, fetch items based on score ranges or member value|

**Strings in Redis**
|Command|What it does|
|-|-|
|GET|Fetches the data stored at the given key|
|SET|Sets the value stored at the given key|
|DEL|Deletes the value stored at the given key (works for all types)|

**List in Redis**
|Command|What it does|
|-|-|
|RPUSH|Pushes the value onto the right end of the list|
|LRANGE|Fetches a range of values from the list|
|LINDEX|Fetches an item at a given position in the list|
|LPOP|Pops the value from the left end of the list and returns it|

**Sets in Redis**
|Command|What it does|
|-|-|
|SADD|Adds the item to the set|
|SMEMBERS|Returns the entire set of items|
|SISMEMBER|Checks if an item is in the set|
|SREM|Removes the item from the set, if it exists|


**Hashes in Redis**
|Command|What it does|
|-|-|
|HSET|Stores the value at the key in the hash|
|HGET|Fetches the value at the given hash key|
|HGETALL|Fetches the entire hash|
|HDEL|Removes a key from the hash, if it exists|

**Sorted sets in Redis**
|Command|What it does|
|-|-|
|ZADD|Adds member with the given store to the ZSET|
|ZRANGE|Fetchs the items in the ZSET from their positions in sorted order|
|ZRANGEBYSCORE|Fetchs items in the ZSET based on a range of scores|
|ZREM|Removes the item from ZSET, if it exists|

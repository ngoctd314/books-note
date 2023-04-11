# Introduction to Transactions

A transaction is a logical unit of processing in a database that includes one or more database operations, which can be read or write operations. There are situations where your application may require reads and writes to multiple documents (in one or more collection) as part of this logical unit of processing.

## A definition of ACID

ACID is an acronym for Atomicity, Consistency, Isolation, and Durability. 

**Atomicity** ensures that all operations inside a transaction will either be applied or nothing will be applied.
**Consistency** ensures that if a transaction succeeds, the database will move from one consistent state to the next consistent state.
**Isolation** is the property that permits multiple transactions to run at the same time in your database.  It guarantees that a transaction will not view the partial results of any other transaction, which means multiple parallel transactions will have the same results as running each of the transactions sequentially.
**Durability** ensures that when a transaction is committed all data will persist even in the case of a system failure.

A database is said to be ACID-compliant when it ensures that all these properties are met and that only successful transactions can be processed.

MongoDB is a distributed database with ACID compliant transactions across replica sets and/or across shards. The network layer adds additional level of complexity. 

## How to Use Transactions

Two apis to use transactions. The first is a similar syntax to relational databases (start_transaction and commit_transaction) called the core API and the second is called the callback API, which is the recommend approach to using transaction.


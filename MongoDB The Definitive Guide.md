# MongoDB The Definitive Guide

## Part 1: Introduction to MongoDB

### Chapter 1: Introduction

MongoDB is a powerful, flexible, and scalable general-purpose database. It combines the ability to scale out with features such as secondary indexes, range queries, sorting aggregations, and geospatial indexes. This chapter covers the major design decisions that made MongoDB what it is.

#### Ease of Use

MongoDB is a document-oriented database, not a relational one. The document-oriented database replaces the concept of a "row" with a more flexible model, the "document". By allowing embedded documents and arrays, the document-oriented approach makes it possible to represent complex hierarchical relationships with a single record.
 
#### Designed to scale
#### Rich with Features
#### Without Sacrificing
#### The Philosophy

### Chapter 2: Getting Started

#### Documents

At the heart of MongoDB is the document: an ordered set of keys with associated values. The keys in a document are strings. Any UTF-8 character is allowed in a key, with a few notable exceptions:

- Key must not contain the character \0 (the null character). This character is used to signify the end of key.
- The . and $ characters have some special properties and should be used in certain circumstances. 

MongoDB is type-sensitive and case-sensitive.

#### Collections

A collection is a group of documents.

#### Dynamic Schemas
Collections have dynamic schemas. This means that the documents within a single collection can have any number of different shapes. For example, both of the following documents could be stored in a single collection:

```json
{"greeting": "Hello, world", "views": 3}
{"signoff": "Good night, and gook luck"}
```

But we dont do that :))

#### Databases

In addition to grouping documents by collection, MongoDB groups collections into databases. A single instance of MongoDB can host several databases. A good rule of thumb is to store all data for a single application in the same database.

There are also some reserved database names, which you can access but which have special semantics. There are as follows: 
- admin: authentication, authorization. Access to this database is required for some administrative operations.
- local: stores data specific to a single server.
- config: shared MongoDB clusters use the config database to store information about each shard.

#### Data Types

**Basic Data Types**

Documents in MongoDB can be thought of as "JSON-like" in that they are conceptual similar to objects in Javascript, but with more type support

```json
// Null
{"x": null}

// Boolean
{"x": true}

// Number
{"x": 3.14}
{"x": 3}
{"x": NumberInt("3")} // 32 bit
{"x": NumberLong("3")} // 64 bit

// String
{"x": "foobar"}

// Date
// MongoDB stores dates as 64-bit integers representing milliseconds since the Unix epoch
{"x": new Date()}

// Regular expression
{"x": /foobar/i}

// Array
{"x": ["a", "b", "c"]}

// Embedded document
// Documents can contain entire documents embedded as values in a parent document
{"x": {"foo": "bar"}}

// ObjectID
// An object ID is a 12 byte ID for documents
{"x": ObjectId()}

// Binary data
Binary data is a string of arbitrary bytes. It cannot be manipulated from the shell.

// Code
MongoDB also makes it possible to store arbitrary JavaScript in queries and documents:
{"x": function() {}}
```

**Dates**

In JavaScript, the Date class is used for MongoDB's data type. When creating a new Date object, always call new Date(), not just Date(). Calling the constructor as a function returns a string representation of the date, not an actual Date object. This is not MongoDB's choice; it is how JS works. If you are not careful to always use the Date constructor, you can end up with a mishmash of strings and dates.


### Chapter 3: Creating, Updating and Deleting Documents

### Chapter 4: Querying

## Part 2: Designing Your Application

## Part 3: Replication

## Part 4: Sharding
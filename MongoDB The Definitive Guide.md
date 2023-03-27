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

### Chapter 3: Creating, Updating and Deleting Documents

### Chapter 4: Querying

## Part 2: Designing Your Application

## Part 3: Replication

## Part 4: Sharding
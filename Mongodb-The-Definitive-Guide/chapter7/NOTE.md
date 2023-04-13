# Chapter 7

Many applications require data analysis of one form or another. MongoDB provides powerful support for running analytics natively using the aggregation framework.

## Pipelines, Stages and Tunables

The aggregation framework is a set of analytics tools within MongoDB that allow you to do analytics do documents in one or more collections.

**Aggregation**
Collection and summary of data

**Stages**
One of the built-in methods that can be completed on the data, but does not permanently alter it

```bash
$match: Filters for data that matches criteria
$group: Groups documents based on criteria
$sort: Puts the documents in a specified order
```
**Aggregation pipeline**
A series completed on the data in order

**Quiz note**
Which tasks can be completed with an aggregation pipeline?
- You can filter for relevant pieces of data by using aggregation, but you can change the documents in the database
- You can group documents together using aggregation, but you cann't change those documents in the database
- You can calculate totals from a group of documents by using aggregation

## Lesson 2: Using $match and $group stages in MongoDB aggregation pipeline

**$match**
Filter for documents matching criteria.
Place as early as possible in the pipeline so it can use indexes.
Reduces the number of documents.

**$group**
Create a single document for each distinct value.

## Lesson 3: Using $sort and $limit stages in a MongoDB aggregation pipeline

**$sort**
Sorts all input documents and passes them through pipeline in sorted order
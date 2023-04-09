# Indexes

Indexes enable you to perform queries effciently.
They're an important part of application development and are even required for certain types of queries.
- What indexes are and why you'd want to use them
- How to choose which fields to index
- How to enforce and evaluate index usage
- Administrative details on creating and removing indexes.

Choosing the right indexes for your collections is critial to performance.

## Introduction to indexes

A database index is similar to a book's index. Instead of looking through the whole book, the database takes a shortcut and just looks at an ordered list with references to the content.

A query that does not use an index is called a collection scan, which means that the server has to "look through the whole book" to find a query's results.

- totalKeysExamined: This is how many kes within the index MongoDB walked through in order to generate the result set.
- we can compare "totalKeysExamined" to "nReturnd" to get a sense for how much of the index MongoDB
had to traverse in order to find just the documents matching the query.

## Recap

To recap, when designing a compound index:
- Keys for equality filters should appear first
- Keys used for sorting should appear before multivalue fields
- Keys for multivalue filters should appear last
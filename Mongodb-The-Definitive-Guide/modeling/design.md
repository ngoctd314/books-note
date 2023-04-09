# Design mongodb schema

When you are designing your MongoDB schema design, the only thing that matters is that you design a schema that will work well for your application. Two different apps that use the same exact data might have very different schemas if the applications are used differently. When designing a schema, we want to take into consideration the following:  

- Store the data.
- Provide good query performance.
- Require reasonable amount of hardware.

Let's take a look at how we might model the relational User model in MongoDB

```json
{
    "first_name": "Paul",
    "surname": "Miller",
    "cell": "1234",
    "city": "London",
    "location": [45.123, 47.232],
    "profession": ["banking", "finance", "trader"],
    "cars": [
        {
            "model": "Bentley",
            "year": 1973
        },
        {
            "model": "Rolls Royce",
            "year": 1965
        }
    ]
}
```

You can see that instead of splitting our data up into separate collections or documents, we take advantage of MongoDB's document based design to embed data into arrays and objects within the User object. Now we can make one simple query to pull all that data together for our application.


## Embedding vs. Referencing

MongoDB schema design actually comes down to only two choices for every piece of data. You can either embed that data directly or reference another piece of data using the $lookup operator.

## Embedding

**Advantages**
- You can retrieve all relevant information in a single query.
- Avoid implementing joins in application code or using $lookup.
- Update related information as a single atomic operation.
- By default, all CRUD operations on a single document are ACID compliant.
- However, if you need to a transaction across multiple operations, you can use the transaction operator.
- Though transactions are available starting 4.0, however, i should add that it's anti-pattern to be overly reliant on using them in your application.

**Limitations**
- Large documents mean more overhead if most fields are not relevant. You can increase query performance by limiting the size of the documents that you are sending over the wire for each query. 
- There is a 16-MB documentation size limit in MongoDB. If you are embedding too much data inside a single document, you could potentially his this limit.

## Referencing

The other option for designing our schema is referencing another document using a documents unique object ID and connecting them together using the $lookup operator. Referencing works similarly as the JOIN operator in an SQL query. 

**Advantages**
- By splitting up data, you will have smaller documents
- Less likely to reach 16-MB-per document limit
- Infrequently accessed information not needed on every query.

**Limitations**

- In order to retrieve all the data in the referenced documents, a minium of two queries
or $lookup required to retrieve all the information

## Type of Relationships

**One-to-One**
- Prefer key-value pair embedded in the document.
- For example, an employee can work in one and only one department.

**One-to-Few**

Rule 1: Favor embedding unless there is a compelling reason not to.
Generally speaking, my default action is to embed data within a document. I pull it our reference it only if
i need to access it on its own, it's too big, i rarely need it, or any other reason.

**One-to-Many**

Rule 2: Needing to access an object on its own is a compelling reason not to embed it.
Rule 3: Avoid join/lookups if possible, but don't be afraid they can provide a better schema design.

**One-to-Squillions**
What if we have a schema where there could be potentially millions of subdocuments, or more?

```json
{
	"_id": ObjectID("AAAB"),
	"name": "goofy.example.com",
	"ipaddr": "127.66.66.66"
}

{
	"message": "cpu is on fire!",
	"host": ObjectID("AAAB")
}
```

**Many-to-Many**

```json
// Each user has a sub-array of linked tasks, and each task has a sub-array of owners for each item in our to-do app
{
    "_id": ObjectID("AAF1"),
    "name": "Kate Moster",
    "tasks": [ObjectID("ADF9"), ObjectID("AE02"), ObjectID("AE73")]
}

{
    "_id": ObjectID("ADF9"),
    "description": "Write blog post about MongoDB schema design",
    "due_date": ISODate("2014-04-01"),
    "owners": [ObjectID("AAF1"), ObjectID("BB3G")]
}
```

Rule 5: You want to structure your data to match the ways that your application queries and updates it.

Summary:
One-to-one: Prefer key value pairs within the document
One-to-few: Prefer embedding
One-to-many: Prefer embedding
One-to-squillions: Prefer referencing
Many-to-many: Prefer referencing

Rule 1: Favor embedding unless these is a compelling reason not to.
Rule 2: Needing to access an object on its own is a compelling reason not to embed it.
Rule 3: Avoid joins and lookups if possible, but don't afraid if they can provide a better schema design.
Rule 4: Arrays should not grow without bound. If there are more than a couple of hundred documents on the many side, don't embed them; if there are more than a few thousand documents on the many side, don't use an array of ObjectID references.
Rule 5: As always, with MongoDB, how you model your data depends entirely on your particular application's data access patterns. You want to structure your data to match the ways that your application queries and updates it.


## One-to-many: embed, in the "one" side

- The documents from the "many" side are embedded.
- Most common representation for simle applications, few documents to embeded
- Need to process main object and the N realted documents together
- Indexing is done on the array

## One-to-many: embed, in the "one" side

- Array of references
- Allows for large documents and a high count of these
- List of references available when retrieving the main object
- Cascade deletes are not supported by MongoDB and must be managed by the application

## Recap for the One-to-Many Relationships

- There are a lot of choices: embed or reference, and choose the side between "one" and "many"
- Duplication may occur when embedding on the "many" side. However, it may be OK, or even preferable.
- Prefer embedding over referencing for simplicity, or when there is a small number of referenced documents as all related information is kept together.
- Embed on the side of the most queried collection
- Prefer referencing when the associated documents are not always needed with the most often queried documents.

## Many-to-many Representations

1. Embed
a. Array of subdocuments in the "many" side
b. Array of subdocuments in the other "many" side

Usually, only the most queries side is considered

2. Reference
a. Array of references in one "many" side
b. Array of references in the other "many" side

## Recap for the Many-to-Many Relationships:

- Ensure it is a "many-to-many" relationship that should not be simplified.
- A "many-to-many" relationshiop can be replaced by two "one-to-many" relationships but does not have to with the document model.
- Prefer embedding on the most queried side
- Prefer embedding for information that is primarily static over time and may profit from duplication.
- Prefer referencing over embedding to avoid managing duplication.

## One-to-one relationship
1. Embed
a. Fields at same level
b. Grouping in sub-documents

2. Reference
a. Same identifier in both documents
b. in the main "one" side
c. in the secondary "one" side

## Recap for the One-to-One relationships
- Prefere embedding over referencing for simplicity
- Use subdocuments to organize the fields
- Use a reference for optimization purposes

## Questions about relationships

1. Why did we introduce the one-to-zillion relationship in our modeling notation?
A. To highlight the fact that huge cardinalities may impact design choices
B. To graphically represent a relationship that has a higher order of manitude than commonly seen in a "one-to-many" relationship

2. Consider a one-to-many relationship observed between a country and the cities in that country
Which of the following are valid ways to represent this one-to-many relationship with the document model in MongoDB?
A. Embed the entities for the cities as an array of sub-documents in the corresponding county document.
B. Have a collection for the countries and a collection for the cities with each city document having a field to reference the document of its country 
C. Embed all the fields for a city as a subdocument in the corresponding country document.
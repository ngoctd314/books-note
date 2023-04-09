## Attribute pattern

```json
{
    "description": "Cherry Coke 6-pack",
    "manufacturer": "Coca-Cola",
    "brand": "Coke",
    "sub_brand": "Cherry Coke",
    "price": 5.99,
    "color": "red",
    "size": "12 ounces",
    "container": "can",
    "sweetener": "sugar"
}
```

```json
{
    "descripton": "Evian 500ml",
    "manufacturer": "Danone",
    "brand": "Evian",
    "price": 1.99,
    "size": "500ml",
    "container": "plastic bottele"
}
```
To search effectively on one of those fields

```go
db.products.find({"capacity": {"$gt": 4000}})
// index on "capacity"
db.products.find({"output": "5v"})
// index on "output"

// May need a lot of indexes
```

To use the attribute pattern you start by identifying the list of fields you want to transpose. Here we transpose the fields input, output and capacity.

```json
// Define fields we need to transpose
{
    "input": "5V/1300 mA",
    "output": "5V/1A",
    "capacity": "4200 mAh"
}
```
Then for each field in associated value, we create that pair. The name of the keys for those pairs do not matter. Only for consistency, let's use K for key and V for value.

```json
{
    "add_specs": [
        {"k": "input", "v": "5V/1300 mA"},
        {"k": "output", "v": "5V/1A"},
        {"k": "capacity", "v": 4200}
    ]
}
```

The attribute pattern addesses the problem of having a lot of similar fields in a document. Often, those fields have similar value types. Or there's a need to search across many of those fields. Fields present in only a small subset of documents.

### Benefits and Trade-Offs
- Easier to index
- Allow for non-deterministic field names
- Ability to qualify the relationship of the original field and value

## Extended Reference Pattern


# Design patterns trade-off

Applying Patterns may lead to
1. Duplication
Let's start with the concern of duplication
**Why?**
Why do we have duplication?
It is usually the result of embedding information in a given document for faster access.

**Concern**
The concern is that it makes handling changes to duplicated information a challenge for correctness and consistency, where multiple document across different collections may need to be updated. There are some miss conception that duplication should not exist. In some cases, duplication is better than no duplication. However, not all pieces of information are affected in the same way by duplication.

Example about duplication is better than not doing it
```json
// order
{
    "customer_id": "12345", // order reference customer
    "items": [
        {
            "item_id": "B000HC2LI0"
        }
    ]
}

// customer
{
    "_id": "12345", 
    // customer reference address
    "address": {
        "street": "TK",
        "city": "HN",
        "country": "VIE"
    }
}
```

Với thiết kế này, việc cập nhật địa chỉ gây ảnh hương đến tính nhất quán dữ liệu trong hệ thống. Ví dụ trong db đã có order.

```json
// order
{
    "_id": 1,
    "cutomer_id": "12345",
    "items": [
        {
            "item_id": "B000HC2LI0"
        }
    ]
}
```
Đơn hàng này đang được giao đến TK, HN. Việc cập nhật lại địa chỉ, khiến sai lệnh dữ liệu trong hệ thống (data inconsistency)

Với thiết kế chuẩn thì address sẽ được embeded trong order
```json
// order
{
    "_id": 1,
    "items": [
        {
            "item_id": "B000HC2LI0"
        }
    ],
    "shipping_address": {
        "street": "TK",
        "city": "HN",
        "country": "VIE"
    }
}
```

Ví dụ khác, trường hợp copy data không bị thay đổi. Let's say we want to model movies and actors. Movies have many actors and actors play in many movies (many-to-many relationship.)

```json
// movie
{
    "_id": "tt1",    
    "title": "Star wars",
    "cast": [
        "nm1", "nm2"
    ]
}

// actor
{
    "_id": "nm1",
    "name": "Mark Hamill",
    "filmography": [
        "tt1": "Luke Skywalker"
    ]
}
```

Ở đây đang thiết kế theo kiểu reference. Tuy nhiên thông tin về diễn viên trong một bộ phim khi phim đã ra rạp thì nó sẽ không bao giờ thay đổi (never update). Vậy ta có thể embed thông tin về diễn viên trong movie collection.

```go
{
    "_id": "tt1",    
    "title": "Star wars",
    "cast": [
        "nm1", "nm2"
    ]
}
```

2. Data staleness
accepting staleness in some pieces of data

Staleness in about facing a piece of data to a user that may have been out of date.
We now in live in a world has more staleness than a few years ago.

**Why?**
- New events come along at such a rate that updating some data constantly cause performance issues.

**Concern?**
- Data quality and reliability

3. Data integrity issues
writing extra application side logic to ensure referential integrity

**Why?**
- Linking information between documents or tables 
- No support for cascading deletes

**Concern?**
- Challenge for correctness and consistency


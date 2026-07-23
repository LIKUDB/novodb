# NQL Reference (NovoDB Query Language)

This document summarizes the content of the interactive `HELP` command (`internal/novodb/help.go`), organized by category.

## Databases

| Command | Description |
|---|---|
| `CREATE DB <name>` | Create a database |
| `DROP DB <name>` | Drop a database |
| `RENAME DB <old> TO <new>` | Rename a database |
| `USE <name>` | Switch active database |
| `SHOW DBS` | List databases |
| `INFO DB <name>` / `DESCRIBE DB <name>` | Detail / schema |
| `STATS DB [<name>]` / `SIZE DB [<name>]` | Statistics / size |
| `COMPACT <db>` | Garbage collection |
| `ANALYZE DB` / `OPTIMIZE DB` | Analysis and optimization |
| `BACKUP <db> TO <file>` / `RESTORE <db> FROM <file>` | Backups |

## Blocks (Collections)

| Command | Description |
|---|---|
| `CREATE BLOCK [<db>] <name>` | Create a block |
| `DROP BLOCK [<db>] <name>` | Drop a block |
| `RENAME BLOCK [<db>] <old> TO <new>` | Rename |
| `SHOW BLOCKS [<db>]` | List blocks |
| `EMPTY BLOCK [<db>] <name>` / `CLEAR [<db>] <name>` | Empty block |
| `REBUILD BLOCK` / `CHECK BLOCK` / `REPAIR BLOCK` | Index maintenance / integrity |

## Documents — INSERT

```sql
INSERT users {"name": "John", "age": 30}
INSERT users name: "John", age: 30
INSERT users {"name": "John"}; {"name": "Jane"}
INSERT users [{"name": "John"}, {"name": "Jane"}]
INSERT users FROM "file.json"
```

## Documents — FIND

```sql
FIND users WHERE age > 18 AND status = "active"
FIND users SELECT name, age WHERE age > 18
FIND users ORDER name, age:DESC WHERE age > 18
FIND users WHERE age > 18 LIMIT 50 OFFSET 100
GET users <id>
```

## Documents — SEARCH (full-text)

```sql
SEARCH articles "search text"
SEARCH articles "exact phrase" EXACT
SEARCH articles "~fuzzy" FUZZY
SEARCH articles "text" WITH SCORE WITH MATCHES
```

## Documents — UPDATE / DELETE

```sql
UPDATE users WHERE _id = "abc" SET name = "John", age = 30
UPDATE users WHERE _id = "abc" INC views = 1
UPDATE users WHERE _id = "abc" PUSH tags = "newtag"
UPDATE ALL users SET status = "archived"

DELETE users WHERE age < 18
DELETE ALL users
```

## Aggregations and GROUP BY

```sql
COUNT users WHERE active = true
SUM orders amount WHERE status = "completed"
AVG products price
GROUP orders BY status SUM amount
```

Available functions: `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`, `MEDIAN`, `MODE`, `STDDEV`.

## ACID Transactions

```sql
BEGIN mydb users
INSERT users {"name": "John"}
COMMIT
-- or ROLLBACK / ABORT

TX STATUS
TX LIST
TX ISOLATION
```

Isolation levels: `read_committed`, `repeatable_read` (default), `serializable`.

## JOINs and Views

```sql
JOIN orders WITH customers ON orders.customer_id = customers._id

VIEW CREATE <name> AS FIND <block> WHERE ...
VIEW DROP <name>
VIEW SHOW
```

## Export / Import

```sql
EXPORT users TO "file.json"
EXPORT users WHERE age > 18 TO "file.csv"
IMPORT users FROM "file.json"
```

## Users, Shards, and Cluster

```sql
CREATE USER alice PASSWORD "..." ROLE readwrite
SHOW USERS

SHARD STATUS
SHARD REBALANCE
SHARD SCALE mydb 8

CLUSTER STATUS
```

## Navigation / System

`PWD`, `LS`, `LS <db>`, `CD <db>`, `TREE`, `STATUS`, `HEALTH`, `VERSION`, `PING`, `HELP`, `EXIT` / `QUIT`.

## Filter Operators

`=`/`==`, `!=`/`<>`, `>`, `<`, `>=`, `<=`, `LIKE`, `CONTAINS`, `EXISTS`, `IN`, `NOT IN`, `BETWEEN`, `STARTS WITH`, `ENDS WITH`, `AND`, `OR`.

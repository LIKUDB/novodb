# NovoDB NQL - Complete Syntax Reference

## INDEX

1. [Databases](#1-databases)
2. [Blocks (Collections)](#2-blocks-collections)
3. [Documents - INSERT](#3-documents---insert)
4. [Documents - FIND](#4-documents---find)
5. [Documents - SEARCH (Full-Text)](#5-documents---search-full-text)
6. [Documents - UPDATE](#6-documents---update)
7. [Documents - DELETE](#7-documents---delete)
8. [Aggregations](#8-aggregations)
9. [GROUP BY](#9-group-by)
10. [JOIN](#10-join)
11. [Views](#11-views)
12. [ACID Transactions](#12-acid-transactions)
13. [EXPORT / IMPORT](#13-export--import)
14. [Users](#14-users)
15. [SHARD](#15-shard)
16. [CLUSTER](#16-cluster)
17. [Navigation / System](#17-navigation--system)
18. [Complete Example](#18-complete-example)

---

## 1. Databases

### CREATE DB
```sql
CREATE DB mydb
```
Creates a new database named `mydb`.

### DROP DB
```sql
DROP DB mydb
```
Deletes the database `mydb` and all its data.

### RENAME DB
```sql
RENAME DB old_name TO new_name
```
Renames a database.

### USE
```sql
USE mydb
```
Switches to the active database.

### SHOW DBS
```sql
SHOW DBS
```
Lists all databases.

### INFO DB
```sql
INFO DB mydb
```
Shows database details.

### DESCRIBE DB
```sql
DESCRIBE DB mydb
```
Shows database schema.

### STATS DB
```sql
STATS DB mydb
```
Shows database statistics.

### SIZE DB
```sql
SIZE DB mydb
```
Shows database size.

### COMPACT
```sql
COMPACT mydb
```
Runs garbage collection.

### ANALYZE DB
```sql
ANALYZE DB mydb
```
Analyzes database performance.

### OPTIMIZE DB
```sql
OPTIMIZE DB mydb
```
Optimizes the database.

### BACKUP
```sql
BACKUP mydb TO backup.json
```
Creates a database backup.

### RESTORE
```sql
RESTORE mydb FROM backup.json
```
Restores a database from backup.

---

## 2. Blocks (Collections)

### CREATE BLOCK
```sql
CREATE BLOCK users
CREATE BLOCK mydb users
```
Creates a block (collection) in the active database or the specified one.

### DROP BLOCK
```sql
DROP BLOCK users
DROP BLOCK mydb users
```
Deletes a block.

### RENAME BLOCK
```sql
RENAME BLOCK old TO new
RENAME BLOCK mydb old TO new
```
Renames a block.

### SHOW BLOCKS
```sql
SHOW BLOCKS
SHOW BLOCKS mydb
```
Lists all blocks in the active database or the specified one.

### EMPTY BLOCK / CLEAR
```sql
EMPTY BLOCK users
CLEAR users
```
Deletes all documents from a block.

### ANALYZE BLOCK
```sql
ANALYZE BLOCK users
```
Analyzes block performance.

### OPTIMIZE BLOCK
```sql
OPTIMIZE BLOCK users
```
Optimizes a block (rebuilds indexes, compacts).

### REBUILD BLOCK
```sql
REBUILD BLOCK users
```
Rebuilds all indexes for the block.

### CHECK BLOCK
```sql
CHECK BLOCK users
```
Verifies block integrity.

### REPAIR BLOCK
```sql
REPAIR BLOCK users
```
Repairs a corrupted block.

### SIZE BLOCK
```sql
SIZE BLOCK users
```
Shows block size.

### INFO BLOCK
```sql
INFO BLOCK users
```
Shows block details.

### DESCRIBE BLOCK
```sql
DESCRIBE BLOCK users
```
Shows block schema.

---

## 3. Documents - INSERT

### INSERT - JSON Format
```sql
INSERT users {"name": "John", "age": 30, "active": true}
INSERT users {"user": {"name": "Maria", "email": "maria@email.com"}}
```

### INSERT - key:value Format
```sql
INSERT users name: "John", age: 30, active: true
```

### INSERT - key=value Format
```sql
INSERT users name = "John", age = 30, active = true
```

### INSERT - Multiple Documents (;)
```sql
INSERT users {"name": "John"}; {"name": "Maria"}; {"name": "Peter"}
```

### INSERT - Batch (JSON Array)
```sql
INSERT users [{"name": "John"}, {"name": "Maria"}, {"name": "Peter"}]
```

### INSERT - With Custom ID
```sql
INSERT users user123 {"name": "John", "age": 30}
INSERT users id = "user456" name: "Maria", age: 25
```

### INSERT - From File
```sql
INSERT users FROM "users.json"
INSERT users FROM "users.csv"
```

---

## 4. Documents - FIND

### FIND - Basic
```sql
FIND users
```
Finds all documents in the block (max 100 by default).

### FIND - By ID
```sql
FIND users abc123
```
Finds a document by its ID.

### FIND - With Filters
```sql
FIND users WHERE age > 18
FIND users WHERE age > 18 AND active = true
FIND users WHERE name LIKE "%John%"
FIND users WHERE email CONTAINS "@gmail.com"
FIND users WHERE age BETWEEN 18 AND 65
FIND users WHERE status IN ("active", "pending")
FIND users WHERE tags IN ["go", "database", "nosql"]
FIND users WHERE EXISTS email
FIND users WHERE email IS NOT NULL
```

### FIND - With Projection
```sql
FIND users SELECT name, email, age WHERE age > 18
FIND users EXCLUDE password, token WHERE age > 18
```

### FIND - With Sorting
```sql
FIND users ORDER name, age:DESC WHERE age > 18
```

### FIND - With Pagination
```sql
FIND users WHERE age > 18 LIMIT 50 OFFSET 100
```

### FIND - Table Output (Human Readable)
```sql
FIND users WHERE age > 18 --type:table
FIND users WHERE age > 18 --format:table
```

### GET - By ID
```sql
GET users abc123
GET users @ abc123
```

---

## 5. Documents - SEARCH (Full-Text)

### SEARCH - Basic
```sql
SEARCH articles "data engine"
```

### SEARCH - Exact Phrase
```sql
SEARCH articles "distributed data engine" EXACT
SEARCH articles "distributed data engine" PHRASE
```

### SEARCH - Fuzzy Search
```sql
SEARCH articles "~data engine" FUZZY
SEARCH articles "~data engine" SIMILAR
```

### SEARCH - With Relevance Score
```sql
SEARCH articles "data engine" WITH SCORE
SEARCH articles "data engine" SHOW SCORE
SEARCH articles "data engine" WITH RELEVANCE
```

### SEARCH - With Matches
```sql
SEARCH articles "data engine" WITH MATCHES
SEARCH articles "data engine" SHOW MATCHES
```

### SEARCH - With Operators (+ -)
```sql
SEARCH articles "+engine -closed data"
```

### SEARCH - With Filters and Pagination
```sql
SEARCH articles "engine" WHERE author = "John" LIMIT 50 ORDER date:DESC
```

### SEARCH - Combined Options
```sql
SEARCH articles "database" WITH SCORE WITH MATCHES LIMIT 20
```

---

## 6. Documents - UPDATE

### UPDATE - SET Fields
```sql
UPDATE users WHERE _id = "abc" SET name = "John", age = 30
UPDATE users WHERE id = "abc" SET name = "John", age = 30
UPDATE users WHERE name = "Maria" SET active = false
UPDATE users WHERE email = "john@email.com" SET name = "John Doe"
```

### UPDATE - INC (Increment)
```sql
UPDATE users WHERE _id = "abc" INC visits = 1
UPDATE users WHERE age > 18 INC points = 10
UPDATE users WHERE _id = "abc" INC views = 1, likes = 5
```

### UPDATE - DEC (Decrement)
```sql
UPDATE users WHERE _id = "abc" DEC stock = 5
UPDATE users WHERE _id = "abc" DEC balance = 100
```

### UPDATE - PUSH (to Array)
```sql
UPDATE users WHERE _id = "abc" PUSH tags = "new"
UPDATE users WHERE _id = "abc" PUSH tags = "go", "database"
```

### UPDATE - PULL (from Array)
```sql
UPDATE users WHERE _id = "abc" PULL tags = "old"
UPDATE users WHERE _id = "abc" PULL tags = "obsolete", "deprecated"
```

### UPDATE - ALL (All Documents)
```sql
UPDATE ALL users SET status = "archived"
UPDATE ALL users SET updated_at = "2024-01-01"
```

### UPDATE - Combined Operations
```sql
UPDATE users WHERE status = "draft" SET status = "published", published_at = "2024-01-01"
UPDATE users WHERE _id = "abc" SET name = "John" INC age = 1 PUSH tags = "active"
```

---

## 7. Documents - DELETE

### DELETE - With Filters
```sql
DELETE users WHERE _id = "abc"
DELETE users WHERE age < 18
DELETE users WHERE status = "inactive" OR age > 65
DELETE users WHERE email CONTAINS "@test.com"
```

### DELETE - ALL (with Confirmation)
```sql
DELETE ALL users
```
Deletes ALL documents from the block (asks for confirmation in interactive mode).

### EMPTY BLOCK / CLEAR
```sql
EMPTY BLOCK users
CLEAR users
```
Alias for DELETE ALL.

---

## 8. Aggregations

### COUNT
```sql
COUNT users
COUNT users WHERE active = true
COUNT users WHERE age > 18 AND status = "active"
```

### SUM
```sql
SUM orders total
SUM orders total WHERE status = "completed"
SUM products price WHERE category = "electronics"
```

### AVG
```sql
AVG products price
AVG products price WHERE category = "electronics"
AVG salaries amount WHERE department = "engineering"
```

### MIN
```sql
MIN products price
MIN products price WHERE category = "electronics"
MIN salaries amount
```

### MAX
```sql
MAX products price
MAX products price WHERE category = "electronics"
MAX salaries amount
```

### MEDIAN
```sql
MEDIAN salaries amount
MEDIAN salaries amount WHERE department = "engineering"
MEDIAN scores value
```

### MODE
```sql
MODE products category
MODE logs level
MODE users city
```

### STDDEV
```sql
STDDEV scores value
STDDEV scores value WHERE test = "final"
STDDEV measurements value
```

---

## 9. GROUP BY

### GROUP - Basic
```sql
GROUP users BY city COUNT
```

### GROUP - With SUM
```sql
GROUP orders BY status SUM total
GROUP products BY category SUM price
```

### GROUP - With AVG
```sql
GROUP products BY category AVG price
GROUP salaries BY department AVG amount
```

### GROUP - With Filters
```sql
GROUP users BY city COUNT WHERE age > 18
GROUP orders BY status SUM total WHERE date > "2024-01-01"
GROUP products BY category AVG price WHERE price > 10
```

### GROUP - With MIN/MAX
```sql
GROUP products BY category MIN price
GROUP products BY category MAX price
GROUP salaries BY department MIN amount
```

### GROUP - With Multiple Aggregations
```sql
GROUP products BY category COUNT, SUM price, AVG price
```

---

## 10. JOIN

### JOIN - Basic
```sql
JOIN orders WITH customers ON orders.customer_id = customers._id
```

### JOIN - With Filters
```sql
JOIN articles WITH authors ON articles.author_id = authors._id
JOIN orders WITH customers ON orders.customer_id = customers._id WHERE orders.total > 100
```

### JOIN - With Different Field Names
```sql
JOIN books WITH authors ON books.author_id = authors.id
JOIN posts WITH users ON posts.author = users.username
```

---

## 11. Views

### VIEW CREATE
```sql
VIEW CREATE active_users AS FIND users WHERE active = true
VIEW CREATE adults AS FIND users WHERE age > 18
VIEW CREATE active_adults AS FIND users WHERE active = true AND age > 18
```

### VIEW DROP
```sql
VIEW DROP active_users
```

### VIEW SHOW
```sql
VIEW SHOW
```
Lists all views.

### VIEW INFO
```sql
VIEW INFO active_users
```
Shows view definition details.

### VIEW - Execute
```sql
active_users
adults
active_adults
```

---

## 12. ACID Transactions

### BEGIN - Start Transaction
```sql
BEGIN
BEGIN mydb
BEGIN mydb users
```

### Transaction Operations
```sql
BEGIN
  INSERT users {"name": "John", "age": 30}
  INSERT users {"name": "Maria", "age": 25}
  UPDATE users WHERE _id = "abc" SET balance = 500
  DELETE users WHERE _id = "xyz"
COMMIT
```

### ROLLBACK
```sql
BEGIN
  INSERT users {"name": "John", "age": 30}
  INSERT users {"name": "Maria", "age": 25}
ROLLBACK
```

### ABORT (Alias for ROLLBACK)
```sql
BEGIN
  INSERT users {"name": "John", "age": 30}
ABORT
```

### TX STATUS
```sql
TX STATUS
```
Shows current transaction status.

### TX LIST
```sql
TX LIST
```
Lists all active transactions.

### TX ISOLATION
```sql
TX ISOLATION
```
Shows current isolation level (`read_committed`, `repeatable_read`, `serializable`).

---

## 13. EXPORT / IMPORT

### EXPORT - JSON
```sql
EXPORT users TO "users.json"
EXPORT users WHERE age > 18 TO "adults.json"
EXPORT users WHERE active = true TO "active_users.json"
```

### EXPORT - CSV
```sql
EXPORT users TO "users.csv"
EXPORT users WHERE age > 18 TO "adults.csv"
```

### IMPORT - JSON
```sql
IMPORT users FROM "users.json"
IMPORT users FROM "users_backup.json"
```

### IMPORT - CSV
```sql
IMPORT users FROM "users.csv"
IMPORT users FROM "data/users.csv"
```

---

## 14. Users

### CREATE USER
```sql
CREATE USER alice PASSWORD "secret123"
CREATE USER bob PASSWORD "secret456" ROLE admin
CREATE USER carlos PASSWORD "secret789" ROLE readonly
CREATE USER diana PASSWORD "secret000" ROLE readwrite
```

**Roles:**
- `admin` - Full administrative access
- `readwrite` - Read and write access
- `readonly` - Read-only access

### DROP USER
```sql
DROP USER alice
DROP USER bob
```

### SHOW USERS
```sql
SHOW USERS
```
Lists all users with their roles and status.

---

## 15. SHARD

### SHARD STATUS
```sql
SHARD STATUS
```
Shows detailed shard status including document count, size, primary node, and nodes.

### SHARD REBALANCE
```sql
SHARD REBALANCE
```
Triggers automatic shard rebalancing.

### SHARD SCALE
```sql
SHARD SCALE mydb 32
```
Scales a database to a specific number of shards (minimum 4).

---

## 16. CLUSTER

### CLUSTER STATUS
```sql
CLUSTER STATUS
```
Shows cluster status including node ID, state, leader, peers, and max nodes.

---

## 17. Navigation / System

### PWD
```sql
PWD
```
Shows current working directory path.

### LS
```sql
LS
LS mydb
```
Lists databases or blocks.

### CD
```sql
CD mydb
```
Changes to a database.

### TREE
```sql
TREE
```
Shows the complete directory tree with all databases and blocks.

### STATUS
```sql
STATUS
```
Shows comprehensive system status including:
- Version and build
- Node ID and data root
- Current database
- Shards and replicas
- Operations count
- Memory stats
- Cache stats
- Metrics
- Uptime
- Cluster info (if enabled)
- WAL stats (if enabled)
- Flex-Column stats

### HEALTH
```sql
HEALTH
```
Health check with system status and uptime.

### VERSION
```sql
VERSION
```
Shows NovoDB version.

### PING
```sql
PING
```
Connection test with timestamp.

### HELP
```sql
HELP
```
Shows complete help with all commands and syntax.

### EXIT / QUIT
```sql
EXIT
QUIT
\Q
```
Exits the shell.

---

## 18. Complete Example

```sql
-- ============================================================
-- COMPLETE EXAMPLE - E-Commerce System
-- ============================================================

-- 1. CREATE DATABASE
CREATE DB ecommerce

-- 2. USE THE DATABASE
USE ecommerce

-- 3. CREATE BLOCKS (COLLECTIONS)
CREATE BLOCK products
CREATE BLOCK categories
CREATE BLOCK customers
CREATE BLOCK orders
CREATE BLOCK order_items
CREATE BLOCK reviews

-- 4. INSERT DATA

-- Categories
INSERT categories {"_id": "cat1", "name": "Electronics", "description": "Electronic devices and gadgets"}
INSERT categories {"_id": "cat2", "name": "Books", "description": "Books and publications"}
INSERT categories {"_id": "cat3", "name": "Clothing", "description": "Apparel and fashion"}
INSERT categories {"_id": "cat4", "name": "Home", "description": "Home and kitchen items"}

-- Products
INSERT products {"_id": "p1", "name": "Laptop", "category_id": "cat1", "price": 999.99, "stock": 50, "rating": 4.5}
INSERT products {"_id": "p2", "name": "Smartphone", "category_id": "cat1", "price": 699.99, "stock": 100, "rating": 4.8}
INSERT products {"_id": "p3", "name": "Programming Book", "category_id": "cat2", "price": 49.99, "stock": 200, "rating": 4.7}
INSERT products {"_id": "p4", "name": "T-Shirt", "category_id": "cat3", "price": 19.99, "stock": 500, "rating": 4.2}
INSERT products {"_id": "p5", "name": "Coffee Maker", "category_id": "cat4", "price": 89.99, "stock": 75, "rating": 4.3}
INSERT products {"_id": "p6", "name": "Headphones", "category_id": "cat1", "price": 149.99, "stock": 150, "rating": 4.6}
INSERT products {"_id": "p7", "name": "Cookbook", "category_id": "cat2", "price": 29.99, "stock": 300, "rating": 4.4}

-- Customers
INSERT customers {"_id": "c1", "name": "Alice Johnson", "email": "alice@email.com", "city": "New York", "since": "2023-01-15"}
INSERT customers {"_id": "c2", "name": "Bob Smith", "email": "bob@email.com", "city": "Los Angeles", "since": "2023-03-20"}
INSERT customers {"_id": "c3", "name": "Carol Davis", "email": "carol@email.com", "city": "Chicago", "since": "2023-06-10"}
INSERT customers {"_id": "c4", "name": "David Wilson", "email": "david@email.com", "city": "New York", "since": "2023-08-05"}

-- Orders
INSERT orders {"_id": "o1", "customer_id": "c1", "date": "2024-01-15", "status": "delivered", "total": 1049.98}
INSERT orders {"_id": "o2", "customer_id": "c2", "date": "2024-02-20", "status": "shipped", "total": 699.99}
INSERT orders {"_id": "o3", "customer_id": "c3", "date": "2024-03-10", "status": "delivered", "total": 139.98}
INSERT orders {"_id": "o4", "customer_id": "c4", "date": "2024-03-15", "status": "pending", "total": 89.99}
INSERT orders {"_id": "o5", "customer_id": "c1", "date": "2024-04-01", "status": "shipped", "total": 199.98}

-- Order Items
INSERT order_items {"_id": "oi1", "order_id": "o1", "product_id": "p1", "quantity": 1, "price": 999.99}
INSERT order_items {"_id": "oi2", "order_id": "o1", "product_id": "p5", "quantity": 1, "price": 89.99}
INSERT order_items {"_id": "oi3", "order_id": "o2", "product_id": "p2", "quantity": 1, "price": 699.99}
INSERT order_items {"_id": "oi4", "order_id": "o3", "product_id": "p3", "quantity": 2, "price": 49.99}
INSERT order_items {"_id": "oi5", "order_id": "o4", "product_id": "p5", "quantity": 1, "price": 89.99}
INSERT order_items {"_id": "oi6", "order_id": "o5", "product_id": "p3", "quantity": 1, "price": 49.99}
INSERT order_items {"_id": "oi7", "order_id": "o5", "product_id": "p6", "quantity": 1, "price": 149.99}

-- Reviews
INSERT reviews {"_id": "r1", "product_id": "p1", "customer_id": "c1", "rating": 5, "comment": "Excellent laptop!", "date": "2024-01-20"}
INSERT reviews {"_id": "r2", "product_id": "p2", "customer_id": "c2", "rating": 4, "comment": "Good phone", "date": "2024-02-25"}
INSERT reviews {"_id": "r3", "product_id": "p3", "customer_id": "c3", "rating": 5, "comment": "Great book", "date": "2024-03-15"}
INSERT reviews {"_id": "r4", "product_id": "p4", "customer_id": "c4", "rating": 4, "comment": "Nice t-shirt", "date": "2024-03-20"}

-- 5. QUERIES

-- All products
FIND products

-- Products by category
FIND products WHERE category_id = "cat1"

-- Products in stock
FIND products WHERE stock > 0

-- Products with price between 50 and 100
FIND products WHERE price BETWEEN 50 AND 100

-- 6. SEARCH
SEARCH products "laptop"
SEARCH products "book" WITH SCORE
SEARCH products "coffee" WITH MATCHES

-- 7. JOINS
-- Products with categories
JOIN products WITH categories ON products.category_id = categories._id

-- Orders with customers
JOIN orders WITH customers ON orders.customer_id = customers._id

-- Orders with products (through order_items)
JOIN order_items WITH orders ON order_items.order_id = orders._id

-- 8. AGGREGATIONS
-- Total products
COUNT products

-- Average price
AVG products price

-- Total revenue (sum of order totals)
SUM orders total

-- Average order value
AVG orders total

-- 9. GROUP BY
-- Products per category
GROUP products BY category_id COUNT

-- Orders by status
GROUP orders BY status COUNT

-- Average price by category
GROUP products BY category_id AVG price

-- Total sales by category
GROUP products BY category_id SUM price

-- 10. COMPLEX QUERIES

-- Products with high rating and in stock
FIND products WHERE rating >= 4.5 AND stock > 0

-- Orders from New York customers
JOIN orders WITH customers ON orders.customer_id = customers._id WHERE customers.city = "New York"

-- Customers with orders > $500
JOIN orders WITH customers ON orders.customer_id = customers._id WHERE orders.total > 500

-- Top 5 most expensive products
FIND products ORDER price:DESC LIMIT 5

-- 11. UPDATE
-- Update product price
UPDATE products WHERE _id = "p1" SET price = 899.99

-- Reduce stock after order
UPDATE products WHERE _id = "p1" DEC stock = 1

-- Update order status
UPDATE orders WHERE _id = "o4" SET status = "shipped"

-- 12. DELETE
-- Delete a review
DELETE reviews WHERE _id = "r4"

-- 13. VIEWS
VIEW CREATE active_orders AS FIND orders WHERE status IN ("pending", "shipped")
VIEW CREATE expensive_products AS FIND products WHERE price > 500
VIEW CREATE ny_customers AS FIND customers WHERE city = "New York"

-- Execute views
active_orders
expensive_products
ny_customers

-- 14. TRANSACTION
BEGIN
  -- Create new order
  INSERT orders {"customer_id": "c2", "date": "2024-04-10", "status": "pending", "total": 149.99}
  -- Add order items
  INSERT order_items {"order_id": "o6", "product_id": "p6", "quantity": 1, "price": 149.99}
  -- Reduce stock
  UPDATE products WHERE _id = "p6" DEC stock = 1
COMMIT

-- 15. EXPORT
EXPORT products TO "products_export.json"
EXPORT orders WHERE status = "delivered" TO "delivered_orders.json"

-- 16. STATISTICS
STATS DB ecommerce
STATS DB ecommerce

-- 17. SYSTEM
STATUS
HEALTH
PING
VERSION

-- 18. VIEW STRUCTURE
TREE
SHOW DBS
SHOW BLOCKS

-- 19. EXIT
EXIT
```

---

## Filter Operators Reference

### Comparison Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `=`, `==` | Equal | `age = 18` |
| `!=`, `<>` | Not equal | `status != "inactive"` |
| `>` | Greater than | `age > 18` |
| `<` | Less than | `age < 65` |
| `>=` | Greater than or equal | `price >= 100` |
| `<=` | Less than or equal | `price <= 1000` |

### Text Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `LIKE` | Pattern with % | `name LIKE "%John%"` |
| `NOT LIKE` | Negative pattern | `name NOT LIKE "%John%"` |
| `CONTAINS` | Contains substring | `email CONTAINS "@gmail.com"` |
| `NOT CONTAINS` | Does not contain | `email NOT CONTAINS "@gmail.com"` |
| `STARTS WITH` | Starts with | `name STARTS WITH "A"` |
| `ENDS WITH` | Ends with | `email ENDS WITH ".com"` |

### Set Operators
| Operator | Description | Example |
|----------|-------------|---------|
| `IN` | Value in list | `status IN ("active", "pending")` |
| `NOT IN` | Value not in list | `status NOT IN ("inactive", "deleted")` |
| `BETWEEN` | Range | `age BETWEEN 18 AND 65` |
| `EXISTS` | Field exists | `EXISTS email` |
| `IS NULL` | Is null | `email IS NULL` |
| `IS NOT NULL` | Is not null | `email IS NOT NULL` |

### Logical Operators
| Operator | Description |
|----------|-------------|
| `AND` | And (default) |
| `OR` | Or |

---

## Environment Variables

### Data and Storage
```bash
NOVODB_DATA=./data                    # Data directory
NOVODB_CACHE=2048                     # L1 Cache in MB
NOVODB_L2_CACHE=4096                  # L2 Cache (indexes)
```

### Network
```bash
NOVODB_QUERY_PORT=1555                # HTTP Query Port
NOVODB_ADMIN_PORT=1556                # Admin API Port
NOVODB_METRICS=9090                   # Metrics Port
NOVODB_RAFT_PORT=2335                 # Raft Port
NOVODB_RAFT_BIND=0.0.0.0              # Raft Bind
```

### Sharding
```bash
NOVODB_SHARD_COUNT=16                 # Number of shards
NOVODB_REPLICA_COUNT=3                # Replication factor
```

### Performance
```bash
NOVODB_FAST_STARTUP=true              # Fast startup
NOVODB_COMPRESS_LARGE=true            # Compress large docs
NOVODB_MAX_DOC_SIZE=10485760          # Max size (10MB)
NOVODB_WORKERS=0                      # Workers (auto)
```

### Security
```bash
NOVODB_JWT_SECRET=your-secret         # JWT Secret
NOVODB_QUERY_USER=admin               # Default user
NOVODB_QUERY_PASSWORD=password        # Default password
NOVODB_TLS_ENABLED=false              # TLS enabled
```

### Logging
```bash
NOVODB_LOG_LEVEL=info                 # debug|info|warn|error
```

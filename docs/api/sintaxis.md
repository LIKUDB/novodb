-- ============================================================
-- COMPLETE EXAMPLE - Library Management System
-- ============================================================

-- 1. CREATE DATABASE
CREATE DB library

-- 2. USE THE DATABASE
USE library

-- 3. CREATE BLOCKS (COLLECTIONS)
CREATE BLOCK books
CREATE BLOCK authors
CREATE BLOCK loans
CREATE BLOCK users

-- 4. INSERT DOCUMENTS
-- Authors
INSERT authors {"_id": "a1", "name": "Gabriel García Márquez", "nationality": "Colombian", "birth": 1927}
INSERT authors {"_id": "a2", "name": "Isabel Allende", "nationality": "Chilean", "birth": 1942}
INSERT authors {"_id": "a3", "name": "Jorge Luis Borges", "nationality": "Argentine", "birth": 1899}

-- Books
INSERT books {"_id": "b1", "title": "One Hundred Years of Solitude", "author_id": "a1", "year": 1967, "genre": "Magical Realism", "copies": 5}
INSERT books {"_id": "b2", "title": "The House of the Spirits", "author_id": "a2", "year": 1982, "genre": "Novel", "copies": 3}
INSERT books {"_id": "b3", "title": "The Aleph", "author_id": "a3", "year": 1949, "genre": "Short Story", "copies": 2}
INSERT books {"_id": "b4", "title": "Love in the Time of Cholera", "author_id": "a1", "year": 1985, "genre": "Romantic Novel", "copies": 4}
INSERT books {"_id": "b5", "title": "Ficciones", "author_id": "a3", "year": 1944, "genre": "Short Story", "copies": 3}

-- Users
INSERT users {"_id": "u1", "name": "Ana Martinez", "email": "ana@email.com", "register_date": "2024-01-15"}
INSERT users {"_id": "u2", "name": "Carlos Lopez", "email": "carlos@email.com", "register_date": "2024-02-20"}
INSERT users {"_id": "u3", "name": "Elena Ruiz", "email": "elena@email.com", "register_date": "2024-03-10"}

-- Loans
INSERT loans {"_id": "l1", "book_id": "b1", "user_id": "u1", "loan_date": "2024-06-01", "return_date": "2024-06-15", "returned": false}
INSERT loans {"_id": "l2", "book_id": "b2", "user_id": "u2", "loan_date": "2024-06-05", "return_date": "2024-06-20", "returned": false}
INSERT loans {"_id": "l3", "book_id": "b3", "user_id": "u3", "loan_date": "2024-05-20", "return_date": "2024-06-04", "returned": true}

-- 5. QUERY DATA
-- View all books
FIND books

-- Search books by genre
FIND books WHERE genre = "Magical Realism"

-- Search books with more than 3 copies
FIND books WHERE copies > 3

-- Search Colombian authors
FIND authors WHERE nationality = "Colombian"

-- 6. SEARCH - Full-Text Search
SEARCH books "solitude"
SEARCH books "love" WITH SCORE
SEARCH books "house" WITH MATCHES

-- 7. GROUP BY
-- Count books by genre
GROUP books BY genre COUNT

-- Count books by author
GROUP books BY author_id COUNT

-- 8. JOINS
-- Join books with authors
JOIN books WITH authors ON books.author_id = authors._id

-- 9. AGGREGATIONS
-- Count total books
COUNT books

-- Average copies per book
AVG books copies

-- Maximum copies
MAX books copies

-- 10. UPDATE
-- Increment copies of a book
UPDATE books WHERE _id = "b1" INC copies = 1

-- Update loan status
UPDATE loans WHERE _id = "l1" SET returned = true

-- 11. ACID TRANSACTION
BEGIN
  -- Register a new loan
  INSERT loans {"book_id": "b4", "user_id": "u1", "loan_date": "2024-06-10", "return_date": "2024-06-25", "returned": false}
  -- Reduce available copies
  UPDATE books WHERE _id = "b4" DEC copies = 1
COMMIT

-- 12. VIEW
-- Create view of active loans
VIEW CREATE active_loans AS FIND loans WHERE returned = false

-- View the view
active_loans

-- 13. EXPORT
EXPORT books TO "books_backup.json"
EXPORT authors WHERE nationality = "Colombian" TO "colombian_authors.json"

-- 14. STATISTICS
STATS DB library
STATS DB library

-- 15. VIEW FULL STRUCTURE
TREE

-- 16. SYSTEM STATUS
STATUS
HEALTH
PING
VERSION

-- 17. VIEW DATABASES AND BLOCKS
SHOW DBS
SHOW BLOCKS

-- 18. CLEANUP (optional)
-- DELETE ALL loans
-- EMPTY BLOCK loans
-- DROP BLOCK loans
-- DROP DB library

-- 19. EXIT
EXIT

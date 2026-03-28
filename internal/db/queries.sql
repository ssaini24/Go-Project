-- User queries

-- Missing index on WHERE clause, SELECT *
SELECT * FROM users WHERE email = 'test@example.com';

-- N+1 risk: fetching orders per user in a loop
SELECT * FROM orders WHERE user_id = 1;

-- No LIMIT on large table scan
SELECT * FROM audit_logs WHERE created_at > '2024-01-01';

-- Implicit type cast, missing index hint
SELECT id, name FROM products WHERE CAST(price AS TEXT) = '99.99';

-- Cartesian join (missing JOIN condition)
SELECT u.name, o.total FROM users u, orders o;

-- Non-SARGable: function on indexed column
SELECT * FROM sessions WHERE DATE(created_at) = '2024-03-01';

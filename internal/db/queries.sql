-- Rule: select_star
select_star_query = "SELECT * FROM users WHERE id = 1";

-- Rule: missing_where_clause
missing_where_clause_query = "DELETE FROM sessions";

-- Rule: function_on_indexed_column
function_on_indexed_column_query = "SELECT id FROM products WHERE CAST(price AS TEXT) = '99.99'";

-- Rule: join_without_condition
join_without_condition_query = "SELECT u.name, o.total FROM users u, orders o";

-- Rule: n_plus_one_pattern
n_plus_one_pattern_query = "SELECT id, (SELECT COUNT(*) FROM orders o WHERE o.user_id = u.id) AS order_count FROM users u";

-- Rule: destructive_ddl
destructive_ddl_query = "DROP TABLE audit_logs";

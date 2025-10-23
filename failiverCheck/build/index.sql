EXPLAIN ANALYZE SELECT * FROM components WHERE id BETWEEN 10 AND 10000;
SELECT * from pg_indexes WHERE tablename = 'components';
CREATE INDEX IF NOT EXISTS components_title_idx ON components(title);
DROP INDEX IF EXISTS components_title_idx;
ANALYSE;
EXPLAIN ANALYZE SELECT * FROM components WHERE title IN ('Component 100000', 'Component 20000');
SELECT COUNT(*) FROM components;


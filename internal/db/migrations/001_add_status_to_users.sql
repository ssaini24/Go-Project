-- Migration: add status column to users table

-- up
ALTER TABLE users ADD COLUMN status VARCHAR(20) NOT NULL DEFAULT 'active';

-- down
ALTER TABLE users DROP COLUMN status;

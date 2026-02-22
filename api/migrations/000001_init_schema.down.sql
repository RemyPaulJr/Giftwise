-- Need to drop tables in reverse order they were created
DROP TABLE IF EXISTS feedback;
DROP TABLE IF EXISTS suggestions;
DROP TABLE IF EXISTS gift_lists;
DROP TABLE IF EXISTS recipients;
DROP TABLE IF EXISTS users;

-- Need to drop enum created after the recipients table is dropped
DROP TYPE IF EXISTS budget;
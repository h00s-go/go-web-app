package db

const sqlCreateSchema = `
CREATE TABLE IF NOT EXISTS schema (
	version integer
)
`

const sqlGetSchema = `
SELECT * FROM schema
`

const sqlInsertSchema = `
INSERT INTO schema (version) VALUES (1)
`

const sqlCreatePosts = `
CREATE TABLE IF NOT EXISTS posts (
	id bigserial PRIMARY KEY,
	title text UNIQUE NOT NULL,
	body text UNIQUE NOT NULL
)
`

// create index for faster link post throttling
const sqlCreatePostsTitleIndex = `
CREATE INDEX IF NOT EXISTS posts_title_idx ON posts (title)
`

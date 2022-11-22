-- Users Table
CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	uuid VARCHAR(36) UNIQUE NOT NULL,
	name VARCHAR(100) NOT NULL,
  github VARCHAR(50),
	twitter VARCHAR(50),
	mastodon VARCHAR(50),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE INDEX users_deleted_at_idx ON users (deleted_at);

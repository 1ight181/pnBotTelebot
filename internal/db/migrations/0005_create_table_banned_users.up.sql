CREATE TABLE
    IF NOT EXISTS banned_users (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        reason TEXT,
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        expires_at TIMESTAMP,
        created_by TEXT
    );
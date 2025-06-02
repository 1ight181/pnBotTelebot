-- 1. Таблица users
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        chat_id VARCHAR(50) UNIQUE NOT NULL,
        username VARCHAR(100),
        created_at TIMESTAMP NOT NULL DEFAULT NOW (),
        is_active BOOLEAN NOT NULL DEFAULT TRUE
    );

-- 2. Таблица sources
CREATE TABLE
    sources (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL UNIQUE,
        url TEXT,
        logo_url TEXT
    );

-- 3. Таблица categories
CREATE TABLE
    categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL UNIQUE
    );

-- 4. Таблица promotions
CREATE TABLE
    promotions (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        start_date DATE NOT NULL,
        end_date DATE NOT NULL,
        source_id INTEGER NOT NULL REFERENCES sources (id) ON DELETE CASCADE,
        category_id INTEGER NOT NULL REFERENCES categories (id) ON DELETE CASCADE,
        image_url TEXT,
        created_at TIMESTAMP NOT NULL DEFAULT NOW ()
    );

-- 5. Таблица subscriptions
CREATE TABLE
    subscriptions (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        category_id INTEGER NOT NULL REFERENCES categories (id) ON DELETE CASCADE,
        source_id INTEGER REFERENCES sources (id) ON DELETE CASCADE,
        UNIQUE (user_id, category_id, source_id)
    );

-- 6. Таблица notifications_log
CREATE TABLE
    notifications_log (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        promo_id INTEGER NOT NULL REFERENCES promotions (id) ON DELETE CASCADE,
        sent_at TIMESTAMP NOT NULL DEFAULT NOW (),
        UNIQUE (user_id, promo_id)
    );
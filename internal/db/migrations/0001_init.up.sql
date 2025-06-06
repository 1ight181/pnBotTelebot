-- Тип enum для статуса оффера
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'offer_status') THEN
        CREATE TYPE offer_status AS ENUM ('active', 'paused', 'archived');
    END IF;
END$$;

-- Таблица категорий
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

-- Таблица партнёров
CREATE TABLE IF NOT EXISTS partners (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    logo_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица офферов
CREATE TABLE IF NOT EXISTS offers (
    id SERIAL PRIMARY KEY,
    partner_internal_offer_id TEXT,
    description TEXT,
    title TEXT,
    status offer_status,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    partner_id INTEGER REFERENCES partners(id) ON DELETE SET NULL,
    tracking_link TEXT,
    payout DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица креативов
CREATE TABLE IF NOT EXISTS creatives (
    id SERIAL PRIMARY KEY,
    partner_internal_creative_id TEXT,
    offer_id INTEGER REFERENCES offers(id) ON DELETE CASCADE,
    type TEXT,
    resource_url TEXT,
    width INTEGER,
    height INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    tg_id BIGINT UNIQUE NOT NULL,
    chat_id BIGINT UNIQUE NOT NULL,
    username TEXT,
    fullname TEXT,
    is_subscribed BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица для many-to-many связи: пользователи - категории
CREATE TABLE IF NOT EXISTS user_categories (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE
);

-- Таблица логов статистики
CREATE TABLE IF NOT EXISTS statistics_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    offer_id INTEGER REFERENCES offers(id) ON DELETE SET NULL,
    clicked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address TEXT
);

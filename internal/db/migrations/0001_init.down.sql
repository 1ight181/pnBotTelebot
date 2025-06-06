-- Удаляем таблицы в порядке зависимостей

DROP TABLE IF EXISTS statistics_logs;
DROP TABLE IF EXISTS user_categories;
DROP TABLE IF EXISTS creatives;
DROP TABLE IF EXISTS offers;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS partners;
DROP TABLE IF EXISTS categories;

-- Удаляем enum тип
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'offer_status') THEN
        DROP TYPE offer_status;
    END IF;
END$$;

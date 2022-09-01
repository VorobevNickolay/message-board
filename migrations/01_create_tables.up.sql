CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE TABLE IF NOT EXISTS messages (
                                     id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     userId TEXT NOT NULL UNIQUE,
                                     text TEXT NOT NULL,
                                     created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
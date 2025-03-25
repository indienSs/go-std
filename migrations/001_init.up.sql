CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT,
    email TEXT,
    password TEXT,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title TEXT,
    author TEXT,
    genres TEXT[],
    publicationDate TIMESTAMP WITH TIME ZONE
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
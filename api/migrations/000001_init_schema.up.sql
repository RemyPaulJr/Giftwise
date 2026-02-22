-- Create ENUM for budget to set budget choices for end user
CREATE TYPE budget AS ENUM ('under_25', '25_50', '50_100', '100_plus');

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,  
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS recipients (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    recipient_name VARCHAR(255) NOT NULL,
    hobbies TEXT NOT NULL,
    interests TEXT NOT NULL,
    budget budget NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS gift_lists (
    id SERIAL PRIMARY KEY NOT NULL,
    recipient_id INTEGER NOT NULL REFERENCES recipients(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS suggestions (
    id SERIAL PRIMARY KEY NOT NULL,
    gift_list_id INTEGER NOT NULL REFERENCES gift_lists(id),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()  
);

CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY NOT NULL,
    suggestion_id INTEGER NOT NULL references suggestions(id),
    user_id UUID NOT NULL references users(id),
    rating SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_feedback_per_user UNIQUE (suggestion_id, user_id)
);



CREATE TABLE users_shell (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL,
    password VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
);
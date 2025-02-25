

CREATE TABLE commands_history (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text VARCHAR(300) NOT NULL,
	created_by BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users_shell(id)
);

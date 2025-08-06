CREATE TABLE IF NOT EXISTS refresh_jti (
    jti VARCHAR(36) PRIMARY KEY,
    issued_by INT NOT NULL,
    expires_at BIGINT NOT NULL
);

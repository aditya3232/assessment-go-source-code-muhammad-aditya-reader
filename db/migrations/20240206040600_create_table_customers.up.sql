CREATE TABLE customers (
    id              VARCHAR(255) NOT NULL PRIMARY KEY,
    national_id     BIGINT       NOT NULL,
    name            VARCHAR(255) NOT NULL,
    detail_address  TEXT         NOT NULL,
    created_at      BIGINT       NOT NULL,
    updated_at      BIGINT       NOT NULL
);

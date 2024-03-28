CREATE TABLE sellers 
(
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    national_id     BIGINT       NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    detail_address  TEXT         NOT NULL,
    created_at      BIGINT       NOT NULL,
    updated_at      BIGINT       NOT NULL
) ENGINE=InnoDB;
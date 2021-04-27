CREATE TABLE IF NOT EXISTS todos
(
    id         CHAR(36)     NOT NULL,
    title      VARCHAR(255) NOT NULL,
    dead_line  DATETIME(6)  NOT NULL,
    is_done    BOOL         NOT NULL,
    created_at DATETIME(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6)  NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    deleted_at DATETIME(6)           DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (title)
) ENGINE = InnoDB
    COMMENT ='TODOタスク';

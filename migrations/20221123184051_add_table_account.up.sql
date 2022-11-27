CREATE TABLE IF NOT EXISTS "public".account
(
    id          UUID PRIMARY KEY                     DEFAULT gen_random_uuid(),
    user_id     UUID REFERENCES "user" (id) NOT NULL,
    name        VARCHAR(255)                NOT NULL,
    description VARCHAR(255)                NOT NULL DEFAULT '',
    url         VARCHAR(255)                NOT NULL,
    username    VARCHAR(255)                NOT NULL,
    password    VARCHAR(255)                NOT NULL
);
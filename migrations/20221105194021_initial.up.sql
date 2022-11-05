CREATE TABLE public."user"
(
    id            UUID                   NOT NULL DEFAULT gen_random_uuid(),
    username      CHARACTER VARYING(40)  NOT NULL,
    email         CHARACTER VARYING(56)  NOT NULL,
    password_hash CHARACTER VARYING(255) NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_username_key UNIQUE (username),
    CONSTRAINT user_email_key UNIQUE (email)
);

ALTER TABLE IF EXISTS public."user"
    OWNER TO knvovk;

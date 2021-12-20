BEGIN;

CREATE TABLE IF NOT EXISTS public.messages
(
    id serial NOT NULL,
    user_id INT,
    message text NOT NULL,
    created_at TIMESTAMPTZ DEFAULT Now(),
    PRIMARY KEY (id),
    CONSTRAINT fk_message_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);

ALTER TABLE IF EXISTS public.messages
    OWNER to postgres;

COMMIT;
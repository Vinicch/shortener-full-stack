-- migrate:up
CREATE TABLE public."url" (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    alias VARCHAR NOT NULL,
    original VARCHAR NOT NULL,
    shortened VARCHAR NOT NULL,
    visits bigint NOT NULL
);

CREATE UNIQUE INDEX url_alias_key ON public."url"(alias);
CREATE UNIQUE INDEX url_shortened_key ON public."url"(shortened);

-- migrate:down
DROP TABLE public."url";

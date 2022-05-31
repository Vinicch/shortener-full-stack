-- migrate:up
CREATE TABLE public.url_alias (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    alias VARCHAR(6) NOT NULL,
    "url" VARCHAR NOT NULL
);

CREATE UNIQUE INDEX alias_alias_key ON public.url_alias(alias);

-- migrate:down
DROP TABLE public.url_alias;

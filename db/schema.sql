SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: __schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.__schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: url_alias; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.url_alias (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    alias character varying(6) NOT NULL,
    url character varying NOT NULL
);


--
-- Name: __schema_migrations __schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.__schema_migrations
    ADD CONSTRAINT __schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: url_alias url_alias_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.url_alias
    ADD CONSTRAINT url_alias_pkey PRIMARY KEY (id);


--
-- Name: alias_alias_key; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX alias_alias_key ON public.url_alias USING btree (alias);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.__schema_migrations (version) VALUES
    ('20220531005240');

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
-- Name: url; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.url (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    alias character varying(6) NOT NULL,
    original character varying NOT NULL,
    shortened character varying NOT NULL,
    visits bigint NOT NULL
);


--
-- Name: __schema_migrations __schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.__schema_migrations
    ADD CONSTRAINT __schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: url url_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.url
    ADD CONSTRAINT url_pkey PRIMARY KEY (id);


--
-- Name: url_alias_key; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX url_alias_key ON public.url USING btree (alias);


--
-- Name: url_shortened_key; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX url_shortened_key ON public.url USING btree (shortened);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.__schema_migrations (version) VALUES
    ('20220531005240');

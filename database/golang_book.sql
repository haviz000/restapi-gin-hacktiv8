-- Database: golang_book

-- DROP DATABASE IF EXISTS golang_book;

CREATE DATABASE golang_book
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_Indonesia.1252'
    LC_CTYPE = 'English_Indonesia.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


-- Table: public.books

-- DROP TABLE IF EXISTS public.books;

CREATE TABLE IF NOT EXISTS public.books
(
    id bigint NOT NULL DEFAULT 'nextval('books_id_seq'::regclass)',
    title character varying(300) COLLATE pg_catalog."default",
    author character varying(300) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    CONSTRAINT books_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.books
    OWNER to postgres;
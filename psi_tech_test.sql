--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

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
-- Name: auth_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.auth_users (
    id character varying(50) NOT NULL,
    username text,
    password text,
    created_at timestamp with time zone
);


ALTER TABLE public.auth_users OWNER TO postgres;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.companies (
    id character varying(50) NOT NULL,
    user_id character varying(50),
    company_code text,
    company_name text
);


ALTER TABLE public.companies OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(50) NOT NULL,
    nama text,
    email text,
    telp text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: vouchers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vouchers (
    id character varying(50) NOT NULL,
    code text,
    discount numeric,
    is_active boolean
);


ALTER TABLE public.vouchers OWNER TO postgres;

--
-- Data for Name: auth_users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.auth_users (id, username, password, created_at) FROM stdin;
301eebd8-f278-48b8-b7d6-9f8a2dc8274d	testing	$2a$10$grtM.NOMJwcwFCpVTdZ/7eaKxD5lkV1aac9I4nooD5h1TSjlCX5Oq	2026-04-07 16:55:39.511411+07
0027acd5-b288-4abd-9b79-662492373a0c	testing2	$2a$10$Z/1KSwW.Q7GFXBFP5AKzm.MNBg15XKYI4KYBCUeHh/NQtHXp9CEQ6	2026-04-07 17:21:46.914549+07
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.companies (id, user_id, company_code, company_name) FROM stdin;
trew098	12qwer	SPI	
poiuyt1234	321rewq	PIC	Samudera
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, nama, email, telp) FROM stdin;
12qwer	Imron		081234567890
321rewq	Juli	sammy@mail.com	087654321
\.


--
-- Data for Name: vouchers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vouchers (id, code, discount, is_active) FROM stdin;
v001	DISKON50	50	t
\.


--
-- Name: auth_users auth_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_users
    ADD CONSTRAINT auth_users_pkey PRIMARY KEY (id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: auth_users uni_auth_users_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.auth_users
    ADD CONSTRAINT uni_auth_users_username UNIQUE (username);


--
-- Name: vouchers uni_vouchers_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT uni_vouchers_code UNIQUE (code);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: vouchers vouchers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vouchers
    ADD CONSTRAINT vouchers_pkey PRIMARY KEY (id);


--
-- Name: companies fk_companies_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT fk_companies_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--


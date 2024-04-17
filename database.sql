--
-- PostgreSQL database dump
--

-- Dumped from database version 13.6
-- Dumped by pg_dump version 13.6

-- Started on 2024-04-17 19:03:13

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

DROP DATABASE qrmenu;
--
-- TOC entry 3079 (class 1262 OID 49406)
-- Name: qrmenu; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE qrmenu WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Russian_Russia.1251';


ALTER DATABASE qrmenu OWNER TO postgres;

\connect qrmenu

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

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3080 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 209 (class 1259 OID 49468)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    name character varying NOT NULL,
    background character varying NOT NULL,
    menu_id bigint NOT NULL,
    is_visible boolean NOT NULL,
    order_param bigint NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 208 (class 1259 OID 49466)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.categories_id_seq OWNER TO postgres;

--
-- TOC entry 3081 (class 0 OID 0)
-- Dependencies: 208
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 207 (class 1259 OID 49442)
-- Name: currencies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.currencies (
    id bigint NOT NULL,
    name character varying NOT NULL,
    code character varying NOT NULL
);


ALTER TABLE public.currencies OWNER TO postgres;

--
-- TOC entry 206 (class 1259 OID 49440)
-- Name: currencies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.currencies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.currencies_id_seq OWNER TO postgres;

--
-- TOC entry 3082 (class 0 OID 0)
-- Dependencies: 206
-- Name: currencies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.currencies_id_seq OWNED BY public.currencies.id;


--
-- TOC entry 203 (class 1259 OID 49420)
-- Name: establishments; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.establishments (
    id bigint NOT NULL,
    name character varying NOT NULL,
    phone character varying DEFAULT ''::character varying NOT NULL,
    logo character varying DEFAULT ''::character varying NOT NULL,
    background character varying DEFAULT ''::character varying NOT NULL,
    wifi_password character varying DEFAULT ''::character varying NOT NULL,
    can_make_orders boolean DEFAULT true NOT NULL,
    country character varying DEFAULT ''::character varying NOT NULL,
    city character varying DEFAULT ''::character varying NOT NULL,
    address character varying DEFAULT ''::character varying NOT NULL,
    text text DEFAULT ''::text NOT NULL,
    profile_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    link character varying NOT NULL
);


ALTER TABLE public.establishments OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 49418)
-- Name: establishments_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.establishments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.establishments_id_seq OWNER TO postgres;

--
-- TOC entry 3083 (class 0 OID 0)
-- Dependencies: 202
-- Name: establishments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.establishments_id_seq OWNED BY public.establishments.id;


--
-- TOC entry 211 (class 1259 OID 49484)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id bigint NOT NULL,
    name character varying NOT NULL,
    weight character varying NOT NULL,
    additional_info text NOT NULL,
    is_visible boolean NOT NULL,
    is_available boolean NOT NULL,
    image character varying NOT NULL,
    category_id bigint NOT NULL,
    order_param bigint NOT NULL,
    price double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 210 (class 1259 OID 49482)
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO postgres;

--
-- TOC entry 3084 (class 0 OID 0)
-- Dependencies: 210
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- TOC entry 205 (class 1259 OID 49431)
-- Name: menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.menus (
    id bigint NOT NULL,
    name character varying NOT NULL,
    is_visible boolean NOT NULL,
    establishment_id bigint NOT NULL,
    order_param bigint NOT NULL
);


ALTER TABLE public.menus OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 49429)
-- Name: menus_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menus_id_seq OWNER TO postgres;

--
-- TOC entry 3085 (class 0 OID 0)
-- Dependencies: 204
-- Name: menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.menus_id_seq OWNED BY public.menus.id;


--
-- TOC entry 201 (class 1259 OID 49409)
-- Name: profiles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.profiles (
    id bigint NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    verified boolean DEFAULT false NOT NULL,
    verification_code character varying DEFAULT ''::character varying NOT NULL
);


ALTER TABLE public.profiles OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 49407)
-- Name: profiles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.profiles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.profiles_id_seq OWNER TO postgres;

--
-- TOC entry 3086 (class 0 OID 0)
-- Dependencies: 200
-- Name: profiles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.profiles_id_seq OWNED BY public.profiles.id;


--
-- TOC entry 212 (class 1259 OID 49498)
-- Name: variants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.variants (
    id bigint NOT NULL,
    name character varying NOT NULL,
    price double precision NOT NULL,
    item_id bigint NOT NULL
);


ALTER TABLE public.variants OWNER TO postgres;

--
-- TOC entry 2906 (class 2604 OID 49471)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 2905 (class 2604 OID 49445)
-- Name: currencies id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currencies ALTER COLUMN id SET DEFAULT nextval('public.currencies_id_seq'::regclass);


--
-- TOC entry 2894 (class 2604 OID 49423)
-- Name: establishments id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.establishments ALTER COLUMN id SET DEFAULT nextval('public.establishments_id_seq'::regclass);


--
-- TOC entry 2907 (class 2604 OID 49487)
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- TOC entry 2904 (class 2604 OID 49434)
-- Name: menus id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menus ALTER COLUMN id SET DEFAULT nextval('public.menus_id_seq'::regclass);


--
-- TOC entry 2891 (class 2604 OID 49412)
-- Name: profiles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles ALTER COLUMN id SET DEFAULT nextval('public.profiles_id_seq'::regclass);


--
-- TOC entry 3070 (class 0 OID 49468)
-- Dependencies: 209
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, background, menu_id, is_visible, order_param) FROM stdin;
\.


--
-- TOC entry 3068 (class 0 OID 49442)
-- Dependencies: 207
-- Data for Name: currencies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currencies (id, name, code) FROM stdin;
2	Американский доллар	USD
3	Российский рубль	RUB
4	Евро	EUR
1	Узбекский сум	UZS
5	Казахстанский тенге	KZT
\.


--
-- TOC entry 3064 (class 0 OID 49420)
-- Dependencies: 203
-- Data for Name: establishments; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.establishments (id, name, phone, logo, background, wifi_password, can_make_orders, country, city, address, text, profile_id, currency_id, link) FROM stdin;
\.


--
-- TOC entry 3072 (class 0 OID 49484)
-- Dependencies: 211
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, name, weight, additional_info, is_visible, is_available, image, category_id, order_param, price) FROM stdin;
\.


--
-- TOC entry 3066 (class 0 OID 49431)
-- Dependencies: 205
-- Data for Name: menus; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.menus (id, name, is_visible, establishment_id, order_param) FROM stdin;
\.


--
-- TOC entry 3062 (class 0 OID 49409)
-- Dependencies: 201
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.profiles (id, name, email, password, verified, verification_code) FROM stdin;
\.


--
-- TOC entry 3073 (class 0 OID 49498)
-- Dependencies: 212
-- Data for Name: variants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.variants (id, name, price, item_id) FROM stdin;
\.


--
-- TOC entry 3087 (class 0 OID 0)
-- Dependencies: 208
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 12, true);


--
-- TOC entry 3088 (class 0 OID 0)
-- Dependencies: 206
-- Name: currencies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.currencies_id_seq', 5, true);


--
-- TOC entry 3089 (class 0 OID 0)
-- Dependencies: 202
-- Name: establishments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.establishments_id_seq', 22, true);


--
-- TOC entry 3090 (class 0 OID 0)
-- Dependencies: 210
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 4, true);


--
-- TOC entry 3091 (class 0 OID 0)
-- Dependencies: 204
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.menus_id_seq', 17, true);


--
-- TOC entry 3092 (class 0 OID 0)
-- Dependencies: 200
-- Name: profiles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.profiles_id_seq', 18, true);


--
-- TOC entry 2920 (class 2606 OID 49476)
-- Name: categories categories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pk PRIMARY KEY (id);


--
-- TOC entry 2918 (class 2606 OID 49450)
-- Name: currencies currencies_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currencies
    ADD CONSTRAINT currencies_pk PRIMARY KEY (id);


--
-- TOC entry 2912 (class 2606 OID 49428)
-- Name: establishments establishments_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_pk PRIMARY KEY (id);


--
-- TOC entry 2914 (class 2606 OID 49523)
-- Name: establishments establishments_un; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_un UNIQUE (link);


--
-- TOC entry 2922 (class 2606 OID 49492)
-- Name: items items_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pk PRIMARY KEY (id);


--
-- TOC entry 2916 (class 2606 OID 49439)
-- Name: menus menus_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_pk PRIMARY KEY (id);


--
-- TOC entry 2910 (class 2606 OID 49417)
-- Name: profiles profiles_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pk PRIMARY KEY (id);


--
-- TOC entry 2924 (class 2606 OID 49505)
-- Name: variants variants_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.variants
    ADD CONSTRAINT variants_pk PRIMARY KEY (id);


--
-- TOC entry 2928 (class 2606 OID 49529)
-- Name: categories categories_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_fk FOREIGN KEY (menu_id) REFERENCES public.menus(id) ON DELETE CASCADE;


--
-- TOC entry 2926 (class 2606 OID 49544)
-- Name: establishments establishments_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_fk FOREIGN KEY (profile_id) REFERENCES public.profiles(id) ON DELETE CASCADE;


--
-- TOC entry 2925 (class 2606 OID 49456)
-- Name: establishments establishments_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.establishments
    ADD CONSTRAINT establishments_fk_1 FOREIGN KEY (currency_id) REFERENCES public.currencies(id);


--
-- TOC entry 2929 (class 2606 OID 49534)
-- Name: items items_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_fk FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE;


--
-- TOC entry 2927 (class 2606 OID 49524)
-- Name: menus menus_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_fk FOREIGN KEY (establishment_id) REFERENCES public.establishments(id) ON DELETE CASCADE;


--
-- TOC entry 2930 (class 2606 OID 49539)
-- Name: variants variants_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.variants
    ADD CONSTRAINT variants_fk FOREIGN KEY (item_id) REFERENCES public.items(id) ON DELETE CASCADE;


-- Completed on 2024-04-17 19:03:14

--
-- PostgreSQL database dump complete
--


--
-- PostgreSQL database dump
--

-- Dumped from database version 11.10
-- Dumped by pg_dump version 11.12 (Ubuntu 11.12-1.pgdg18.04+1)

-- Started on 2021-07-21 10:01:16 UTC

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

SET default_with_oids = false;

--
-- TOC entry 227 (class 1259 OID 16846)
-- Name: buildpacks; Type: TABLE; Schema: public; Owner: admin_user_group
--

CREATE TABLE public.buildpacks (
                                   id integer NOT NULL,
                                   guid text NOT NULL,
                                   created_at timestamp without time zone DEFAULT now() NOT NULL,
                                   updated_at timestamp without time zone,
                                   name text NOT NULL,
                                   key text,
                                   "position" integer NOT NULL,
                                   enabled boolean DEFAULT true,
                                   locked boolean DEFAULT false,
                                   filename text,
                                   sha256_checksum text,
                                   stack character varying(255)
);


ALTER TABLE public.buildpacks OWNER TO admin_user_group;

--
-- TOC entry 228 (class 1259 OID 16855)
-- Name: buildpacks_id_seq; Type: SEQUENCE; Schema: public; Owner: admin_user_group
--

CREATE SEQUENCE public.buildpacks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buildpacks_id_seq OWNER TO admin_user_group;

--
-- TOC entry 4646 (class 0 OID 0)
-- Dependencies: 228
-- Name: buildpacks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin_user_group
--

ALTER SEQUENCE public.buildpacks_id_seq OWNED BY public.buildpacks.id;


--
-- TOC entry 4508 (class 2604 OID 17813)
-- Name: buildpacks id; Type: DEFAULT; Schema: public; Owner: admin_user_group
--

ALTER TABLE ONLY public.buildpacks ALTER COLUMN id SET DEFAULT nextval('public.buildpacks_id_seq'::regclass);


--
-- TOC entry 4638 (class 0 OID 16846)
-- Dependencies: 227
-- Data for Name: buildpacks; Type: TABLE DATA; Schema: public; Owner: admin_user_group
--

COPY public.buildpacks (id, guid, created_at, updated_at, name, key, "position", enabled, locked, filename, sha256_checksum, stack) FROM stdin;
13	aeb7c510-2bec-4508-8309-fda7b4a67a1d	2020-02-12 20:52:52.803138	2021-06-14 17:51:44.61865	java_buildpack	aeb7c510-2bec-4508-8309-fda7b4a67a1d_0364796c0d5444ba6ea907ba6c88cb8951ac45b9f0e46c766250a306bb537cc9	2	t	f	java_buildpack-cached-cflinuxfs3-v4.39.zip	0364796c0d5444ba6ea907ba6c88cb8951ac45b9f0e46c766250a306bb537cc9	cflinuxfs3
26140	256020df-0aa0-4c29-a962-33b9c30b9400	2021-07-14 10:34:28.503763	2021-07-14 11:37:00.229995	uas_dataflow_server_buildpack-ad	256020df-0aa0-4c29-a962-33b9c30b9400_0d28251bcd9a5a2619bd408d33cde6983b28af942394f9020c3728b6f0e7c1ba	19	t	f	uas_dataflow_server_buildpack.zip	0d28251bcd9a5a2619bd408d33cde6983b28af942394f9020c3728b6f0e7c1ba	\N
12	38e43106-5177-46da-8da9-279a99000c0a	2020-02-12 20:52:11.688725	2021-06-28 15:37:25.656963	go_buildpack	38e43106-5177-46da-8da9-279a99000c0a_e859df7651e9f5765be137ae94c126d67a4bc74060a4788cf10007b3d98ab621	6	t	f	go_buildpack-cached-cflinuxfs3-v1.9.33.zip	e859df7651e9f5765be137ae94c126d67a4bc74060a4788cf10007b3d98ab621	cflinuxfs3
4	2eb534f1-dbfb-4af3-9680-8a266d91edca	2020-02-12 19:21:34.347254	2021-06-18 22:50:05.880184	sap_java_buildpack	2eb534f1-dbfb-4af3-9680-8a266d91edca_958318cb4d45f3a8a3882b485d9d926880478e1b344edbdb4639e4f468bd7a42	12	t	f	sap_java_buildpack-v1.36.0.zip	958318cb4d45f3a8a3882b485d9d926880478e1b344edbdb4639e4f468bd7a42	\N
1443	2c8f290f-0043-426f-9f05-e7667e8bda5a	2020-03-15 18:57:45.365631	2020-03-15 18:58:31.905208	nodejs_perl_buildpack	\N	16	t	f	\N	\N	\N
1435	5264813f-cb92-4768-b4fa-f9d5085768cb	2020-03-15 14:22:18.641306	2020-03-15 14:22:24.646175	nodejs_perl_buildpack	5264813f-cb92-4768-b4fa-f9d5085768cb_023d2fd89b7785878c15943a8b3a84c0cb5df8baf98981cad26a0728f7296984	17	t	f	nodejs_buildpack-cflinuxfs3-v1.7.14.zip	023d2fd89b7785878c15943a8b3a84c0cb5df8baf98981cad26a0728f7296984	cflinuxfs3
17211	7d6ae5f0-0086-4816-b249-dda1b1b2f792	2021-01-26 06:49:03.593505	2021-01-26 06:49:04.441532	uas_poc_buildpack	7d6ae5f0-0086-4816-b249-dda1b1b2f792_30309b6748bb0bf71616c6dee56970098ca013aa8666095407b91d9ebb8a356c	18	t	f	chen-buildpack.zip	30309b6748bb0bf71616c6dee56970098ca013aa8666095407b91d9ebb8a356c	\N
21802	42dee934-0c1e-4b04-a3bd-cd31e0b4962c	2021-04-23 15:05:29.75938	2021-05-10 16:27:38.996757	sap_java_buildpack_1_34	42dee934-0c1e-4b04-a3bd-cd31e0b4962c_a2bae492bce1d07fce84f2d7a74189f6e505a80707b9bcfa8c87732d2c802cc0	15	t	f	sap_java_buildpack-v1.34.1.zip	a2bae492bce1d07fce84f2d7a74189f6e505a80707b9bcfa8c87732d2c802cc0	\N
10	e14cc734-3bbf-4ae0-a0dd-cfac81dd8a72	2020-02-12 20:51:33.779403	2021-06-28 15:40:19.052789	binary_buildpack	e14cc734-3bbf-4ae0-a0dd-cfac81dd8a72_01950c1ae446f52cd0cdc80db416e8cc7582689e1e0f9219a99457ac54eac046	11	t	f	binary_buildpack-cached-cflinuxfs3-v1.0.39.zip	01950c1ae446f52cd0cdc80db416e8cc7582689e1e0f9219a99457ac54eac046	cflinuxfs3
24771	3bfbcf83-f685-4e09-b469-482547ff651d	2021-06-18 22:50:09.929526	2021-06-18 22:50:21.127921	sap_java_buildpack_1_36	3bfbcf83-f685-4e09-b469-482547ff651d_958318cb4d45f3a8a3882b485d9d926880478e1b344edbdb4639e4f468bd7a42	13	t	f	sap_java_buildpack-v1.36.0.zip	958318cb4d45f3a8a3882b485d9d926880478e1b344edbdb4639e4f468bd7a42	\N
21	0316c9a0-384b-4212-ae12-3ce441124933	2020-02-12 20:56:34.422105	2021-07-09 12:50:00.432449	staticfile_buildpack	0316c9a0-384b-4212-ae12-3ce441124933_36f8c9b875277ccb90a687e745356143377c8fc1e801f1da8d9ee6c36b631db4	1	t	f	staticfile_buildpack-cached-cflinuxfs3-v1.5.23.zip	36f8c9b875277ccb90a687e745356143377c8fc1e801f1da8d9ee6c36b631db4	cflinuxfs3
11	351908ec-72af-4a13-8fd6-2e35a6f5f9bc	2020-02-12 20:51:38.007603	2021-07-09 12:50:52.911782	dotnet_core_buildpack	351908ec-72af-4a13-8fd6-2e35a6f5f9bc_7a6dc70dbee0d334a3f6df5d5611e05b17da68d678fa60a927cc316073a6f342	4	t	f	dotnet-core_buildpack-cached-cflinuxfs3-v2.3.31.zip	7a6dc70dbee0d334a3f6df5d5611e05b17da68d678fa60a927cc316073a6f342	cflinuxfs3
17	8747fdf2-d72a-4fdd-95ae-f2c4bd2427ff	2020-02-12 20:54:50.133573	2021-07-09 12:51:44.031417	php_buildpack	8747fdf2-d72a-4fdd-95ae-f2c4bd2427ff_2c2154945dbeb0ce54a7936fb691a531da06da8d1368ebf409eb7a6526447227	8	t	f	php_buildpack-cached-cflinuxfs3-v4.4.43.zip	2c2154945dbeb0ce54a7936fb691a531da06da8d1368ebf409eb7a6526447227	cflinuxfs3
14	07f83b4a-da4c-4884-b0c1-6fd1ed1080cc	2020-02-12 20:53:46.058283	2021-07-09 12:51:49.464557	nginx_buildpack	07f83b4a-da4c-4884-b0c1-6fd1ed1080cc_8c4f3e5d6f9af250888e486b10f6adef127fff819294ea2a1966dc5174afbe5b	9	t	f	nginx_buildpack-cached-cflinuxfs3-v1.1.29.zip	8c4f3e5d6f9af250888e486b10f6adef127fff819294ea2a1966dc5174afbe5b	cflinuxfs3
26149	f42f4439-ab39-4d0e-8dd2-c233f2fb68f7	2021-07-14 13:09:44.789286	2021-07-19 13:32:51.419227	uas_dataflow_server_buildpack-adm	f42f4439-ab39-4d0e-8dd2-c233f2fb68f7_d2825f4129a86c1adc8c1dca4c5a45cab0c5d754765780eb8dd46cfd820f73cf	21	t	f	uas_dataflow_server_buildpack.zip	d2825f4129a86c1adc8c1dca4c5a45cab0c5d754765780eb8dd46cfd820f73cf	\N
18164	3e392712-6ac2-40ed-beeb-b8acd401722d	2021-02-14 13:08:15.202085	2021-07-14 14:09:24.413672	uas_dataflow_server_buildpack	3e392712-6ac2-40ed-beeb-b8acd401722d_c6bee5820fce8136d1a60c74c1a6d438447c9489efeae28f5cbb2181fa4e26a1	20	t	f	uas_dataflow_server_buildpack.zip	c6bee5820fce8136d1a60c74c1a6d438447c9489efeae28f5cbb2181fa4e26a1	\N
24118	6c5ba5b4-0386-4e0e-8ccb-fb8b4cc938ac	2021-06-07 13:32:06.39552	2021-06-07 13:32:12.806875	sap_java_buildpack_1_35	6c5ba5b4-0386-4e0e-8ccb-fb8b4cc938ac_19726bb542453501b38cc6b6820e4785a852b8622b788d358bf541a3a69b4189	14	t	f	sap_java_buildpack-v1.35.0.zip	19726bb542453501b38cc6b6820e4785a852b8622b788d358bf541a3a69b4189	\N
18	15f5cf7b-8c1e-44a6-bee9-fe45d4587b98	2020-02-12 20:55:34.707994	2021-06-28 15:37:48.187921	python_buildpack	15f5cf7b-8c1e-44a6-bee9-fe45d4587b98_75acfff3ddb6633d20668b1ea0f1a8b9311d24945481df07756d3ff62ff8b036	7	t	f	python_buildpack-cached-cflinuxfs3-v1.7.42.zip	75acfff3ddb6633d20668b1ea0f1a8b9311d24945481df07756d3ff62ff8b036	cflinuxfs3
20	eff2f8fa-a640-425d-8deb-a4f17ba969b4	2020-02-12 20:56:14.630658	2021-07-09 12:50:19.981563	ruby_buildpack	eff2f8fa-a640-425d-8deb-a4f17ba969b4_82973be4dd4b94e5c7f2da682218b2865d3e670f5de1a28a775aef6d6da44e2a	3	t	f	ruby_buildpack-cached-cflinuxfs3-v1.8.42.zip	82973be4dd4b94e5c7f2da682218b2865d3e670f5de1a28a775aef6d6da44e2a	cflinuxfs3
15	af6e73b0-0f2e-4dbd-bd4f-d86c889d68fe	2020-02-12 20:53:52.902633	2021-07-09 12:52:34.34631	r_buildpack	af6e73b0-0f2e-4dbd-bd4f-d86c889d68fe_30985a9d95a6aaba6ff1715e0ce2f3239afeb142cc6f6f2223a0eea4446ade47	10	t	f	r_buildpack-cached-cflinuxfs3-v1.1.19.zip	30985a9d95a6aaba6ff1715e0ce2f3239afeb142cc6f6f2223a0eea4446ade47	cflinuxfs3
16	9aec774a-4155-4849-af75-965bad1c3484	2020-02-12 20:54:35.002372	2021-07-14 18:31:47.373121	nodejs_buildpack	9aec774a-4155-4849-af75-965bad1c3484_6a9ae0f3eb28d9074c7509f8b24534f74fb3624f170379dba7365c0d8c08b69a	5	t	f	nodejs_buildpack-cached-cflinuxfs3-v1.7.52.zip	6a9ae0f3eb28d9074c7509f8b24534f74fb3624f170379dba7365c0d8c08b69a	cflinuxfs3
\.


--
-- TOC entry 4648 (class 0 OID 0)
-- Dependencies: 228
-- Name: buildpacks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin_user_group
--

SELECT pg_catalog.setval('public.buildpacks_id_seq', 26532, true);


--
-- TOC entry 4513 (class 2606 OID 26204)
-- Name: buildpacks buildpacks_pkey; Type: CONSTRAINT; Schema: public; Owner: admin_user_group
--

ALTER TABLE ONLY public.buildpacks
    ADD CONSTRAINT buildpacks_pkey PRIMARY KEY (id);


--
-- TOC entry 4509 (class 1259 OID 26608)
-- Name: buildpacks_created_at_index; Type: INDEX; Schema: public; Owner: admin_user_group
--

CREATE INDEX buildpacks_created_at_index ON public.buildpacks USING btree (created_at);


--
-- TOC entry 4510 (class 1259 OID 26609)
-- Name: buildpacks_guid_index; Type: INDEX; Schema: public; Owner: admin_user_group
--

CREATE UNIQUE INDEX buildpacks_guid_index ON public.buildpacks USING btree (guid);


--
-- TOC entry 4511 (class 1259 OID 26610)
-- Name: buildpacks_key_index; Type: INDEX; Schema: public; Owner: admin_user_group
--

CREATE INDEX buildpacks_key_index ON public.buildpacks USING btree (key);


--
-- TOC entry 4514 (class 1259 OID 26611)
-- Name: buildpacks_updated_at_index; Type: INDEX; Schema: public; Owner: admin_user_group
--

CREATE INDEX buildpacks_updated_at_index ON public.buildpacks USING btree (updated_at);


--
-- TOC entry 4515 (class 1259 OID 27055)
-- Name: unique_name_and_stack; Type: INDEX; Schema: public; Owner: admin_user_group
--

CREATE UNIQUE INDEX unique_name_and_stack ON public.buildpacks USING btree (name, stack);


--
-- TOC entry 4645 (class 0 OID 0)
-- Dependencies: 227
-- Name: TABLE buildpacks; Type: ACL; Schema: public; Owner: admin_user_group
--

GRANT SELECT ON TABLE public.buildpacks TO ccdbread;


--
-- TOC entry 4647 (class 0 OID 0)
-- Dependencies: 228
-- Name: SEQUENCE buildpacks_id_seq; Type: ACL; Schema: public; Owner: admin_user_group
--

GRANT SELECT ON SEQUENCE public.buildpacks_id_seq TO ccdbread;


-- Completed on 2021-07-21 10:01:26 UTC

--
-- PostgreSQL database dump complete
--
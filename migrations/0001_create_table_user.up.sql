DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;
-- CREATE EXTENSION IF NOT EXISTS postgis;
-- CREATE EXTENSION IF NOT EXISTS postgis_topology;


CREATE TABLE users
(
    id      INTEGER  PRIMARY KEY ,
    first_name   VARCHAR(32)                 NOT NULL CHECK ( first_name <> '' ),
    last_name    VARCHAR(32)                 NOT NULL CHECK ( last_name <> '' ),
    email        VARCHAR(64) UNIQUE          NOT NULL CHECK ( email <> '' ),
    password     VARCHAR(250)                NOT NULL CHECK ( octet_length(password) <> 0 ),
    role         VARCHAR(10)                 NOT NULL DEFAULT 'user',
    about        VARCHAR(1024)                        DEFAULT '',
    avatar       VARCHAR(512),
    phone_number VARCHAR(20),
    address      VARCHAR(250),
    city         VARCHAR(30),
    country      VARCHAR(30),
    gender       VARCHAR(20)                 NOT NULL DEFAULT 'male',
    postcode     INTEGER,
    birthday     DATE                                 DEFAULT NULL,
    created_at   TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE     NOT NULL        DEFAULT CURRENT_TIMESTAMP

);




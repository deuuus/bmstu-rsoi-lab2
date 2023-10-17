CREATE DATABASE payments;
--GRANT ALL PRIVILEGES ON DATABASE payments TO program;

\c payments;
CREATE TABLE payment
(
    id          SERIAL PRIMARY KEY,
    payment_uid uuid        NOT NULL,
    status      VARCHAR(20) NOT NULL
        CHECK (status IN ('PAID', 'CANCELED')),
    price       INT         NOT NULL
);

CREATE DATABASE reservations;
--GRANT ALL PRIVILEGES ON DATABASE reservations TO program;

\c reservations;
CREATE TABLE hotels
(
    id        SERIAL PRIMARY KEY,
    hotel_uid uuid         NOT NULL UNIQUE,
    name      VARCHAR(255) NOT NULL,
    country   VARCHAR(80)  NOT NULL,
    city      VARCHAR(80)  NOT NULL,
    address   VARCHAR(255) NOT NULL,
    stars     INT,
    price     INT          NOT NULL
);

INSERT INTO hotels VALUES (1, '049161bb-badd-4fa8-9d90-87c9a82b0668', 'Ararat Park Hyatt Moscow', 'Россия', 'Москва', 'Неглинная ул., 4', 5, 10000);
INSERT INTO hotels VALUES (2, gen_random_uuid(), 'Chester', 'Россия', 'Владимир', 'Большая нижегородская ул., 27', 4, 5000);
INSERT INTO hotels VALUES (3, gen_random_uuid(), 'Bronza', 'Россия', 'Санкт-Петербург', 'Невский пр-кт., 4', 3, 3000);

CREATE TABLE reservation
(
    id              SERIAL PRIMARY KEY,
    reservation_uid uuid UNIQUE NOT NULL,
    username        VARCHAR(80) NOT NULL,
    payment_uid     uuid        NOT NULL,
    hotel_id        INT REFERENCES hotels (id),
    status          VARCHAR(20) NOT NULL
        CHECK (status IN ('PAID', 'CANCELED')),
    start_date      TIMESTAMP WITH TIME ZONE,
    end_data        TIMESTAMP WITH TIME ZONE
);

CREATE DATABASE loyalties;
--GRANT ALL PRIVILEGES ON DATABASE loyalties TO program;

\c loyalties;
CREATE TABLE loyalty
(
    id                SERIAL PRIMARY KEY,
    username          VARCHAR(80) NOT NULL UNIQUE,
    reservation_count INT         NOT NULL DEFAULT 0,
    status            VARCHAR(80) NOT NULL DEFAULT 'BRONZE'
        CHECK (status IN ('BRONZE', 'SILVER', 'GOLD')),
    discount          INT         NOT NULL
);

INSERT INTO loyalty VALUES (1, 'Test Max', 25, 'GOLD', 10);
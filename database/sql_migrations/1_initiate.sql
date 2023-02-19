-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category(
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE book(
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255),
    description VARCHAR(255),
    image_url VARCHAR(255),
    release_year INT NOT NULL,
    price VARCHAR(255),
    total_page INT NOT NULL,
    thickness VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    category_id  INT not null,
    foreign key (category_id) references category (id)
);

-- +migrate StatementEnd

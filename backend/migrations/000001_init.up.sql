CREATE TABLE IF NOT EXISTS admins (
    id                    BIGSERIAL PRIMARY KEY,
    email                 TEXT UNIQUE NOT NULL,
    password_hash         TEXT NOT NULL,

    first_name            TEXT NOT NULL DEFAULT '',
    last_name             TEXT NOT NULL DEFAULT '',
    father_name           TEXT NOT NULL DEFAULT '',
    city                  TEXT NOT NULL DEFAULT '',
    telegram              TEXT NOT NULL DEFAULT '',
    image_id              BIGINT NOT NULL DEFAULT 0,

    email_notification    BOOLEAN NOT NULL DEFAULT false,
    telegram_notification BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS websites (
    id       SERIAL PRIMARY KEY,
    admin_id BIGINT NOT NULL,
    alias    TEXT UNIQUE NOT NULL,
    active   BOOLEAN NOT NULL DEFAULT false,

    FOREIGN KEY (admin_id) REFERENCES admins (id)
);

CREATE TABLE IF NOT EXISTS sections (
    id            BIGSERIAL PRIMARY KEY,
    uuid          TEXT UNIQUE,
    website_alias TEXT NOT NULL,
    width         SMALLINT NOT NULL,
    full_width    BOOLEAN NOT NULL,
    height        SMALLINT NOT NULL,
    full_height   BOOLEAN NOT NULL,

    FOREIGN KEY (website_alias) REFERENCES websites (alias)
);

CREATE TABLE IF NOT EXISTS blocks (
    id            SERIAL PRIMARY KEY,
    section_uuid  TEXT NOT NULL,
    website_alias TEXT NOT NULL,
    text          TEXT NOT NULL DEFAULT '',

    FOREIGN KEY (section_uuid) REFERENCES sections (uuid),
    FOREIGN KEY (website_alias) REFERENCES websites (alias)
);

CREATE TABLE IF NOT EXISTS customers (
    id                    BIGSERIAL PRIMARY KEY,
    website_alias         TEXT NOT NULL,
    email                 TEXT NOT NULL,
    password_hash         TEXT NOT NULL,

    first_name            TEXT NOT NULL DEFAULT '',
    last_name             TEXT NOT NULL DEFAULT '',
    father_name           TEXT NOT NULL DEFAULT '',
    phone                 TEXT NOT NULL DEFAULT '',
    telegram              TEXT NOT NULL DEFAULT '',
    delivery_type         TEXT NOT NULL DEFAULT '',
    payment_type          TEXT NOT NULL DEFAULT '',

    email_notification    BOOLEAN NOT NULL DEFAULT false,
    telegram_notification BOOLEAN NOT NULL DEFAULT false,

    UNIQUE (website_alias, email),
    FOREIGN KEY (website_alias) REFERENCES websites (alias)
);

CREATE TABLE IF NOT EXISTS products (
    id             BIGSERIAL PRIMARY KEY,
    website_alias  TEXT      NOT NULL,
    name           TEXT      NOT NULL,
    description    TEXT      NOT NULL,
    price          INTEGER   NOT NULL,
    image_ids      TEXT[]    NOT NULL,
    active         BOOLEAN   NOT NULL,
    tags           TEXT[]    NOT NULL,

    FOREIGN KEY (website_alias) REFERENCES websites (alias)
);

CREATE TABLE IF NOT EXISTS carts (
    id BIGINT PRIMARY KEY,

    FOREIGN KEY (id) REFERENCES customers (id)
);

CREATE TABLE IF NOT EXISTS cart_items (
    id         BIGSERIAL PRIMARY KEY,
    cart_id    BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    count      INTEGER NOT NULL,

    UNIQUE (cart_id, product_id),

    FOREIGN KEY (cart_id) REFERENCES carts (id),
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS saved_products (
    id             BIGSERIAL PRIMARY KEY,
    website_alias  TEXT      NOT NULL,
    name           TEXT      NOT NULL,
    description    TEXT      NOT NULL,
    price          INTEGER   NOT NULL,
    image_ids      TEXT[]    NOT NULL,
    active         BOOLEAN   NOT NULL,
    tags           TEXT[]    NOT NULL,

    FOREIGN KEY (website_alias) REFERENCES websites (alias)
);

CREATE TABLE IF NOT EXISTS orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT    NOT NULL,
    total_sum   INTEGER   NOT NULL,
    date_time   TIMESTAMP NOT NULL,
    status      SMALLINT  NOT NULL,
    comment     TEXT      NOT NULL,

    FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id               BIGSERIAL PRIMARY KEY,
    order_id         BIGINT    NOT NULL,
    saved_product_id BIGINT    NOT NULL,
    count            INTEGER   NOT NULL ,

    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (saved_product_id) REFERENCES saved_products (id)
);

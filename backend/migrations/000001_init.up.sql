CREATE TABLE IF NOT EXISTS admins (
    id                    BIGSERIAL PRIMARY KEY,
    email                 TEXT UNIQUE NOT NULL ,
    password_hash         TEXT NOT NULL ,

    first_name            TEXT DEFAULT '',
    last_name             TEXT DEFAULT '',
    father_name           TEXT DEFAULT '',
    city                  TEXT DEFAULT '',
    telegram              TEXT DEFAULT '',
    image_id              BIGINT DEFAULT 0,

    email_notification    BOOLEAN DEFAULT false,
    telegram_notification BOOLEAN DEFAULT false
);

CREATE TABLE IF NOT EXISTS websites (
    id                SERIAL PRIMARY KEY,
    admin_id          BIGINT NOT NULL ,
    alias             TEXT UNIQUE NOT NULL ,

    background_color  TEXT DEFAULT 'white',
    text_color        TEXT DEFAULT 'black',
    font              TEXT DEFAULT 'Arial',

    main_one          TEXT DEFAULT '',
    main_two          TEXT DEFAULT '',

    about_one         TEXT DEFAULT '',
    about_two         TEXT DEFAULT '',
    about_three       TEXT DEFAULT '',
    about_four        TEXT DEFAULT '',
    about_five        TEXT DEFAULT '',
    about_six         TEXT DEFAULT '',
    about_image_one   INTEGER DEFAULT 0,
    about_image_two   INTEGER DEFAULT 0,
    about_image_three INTEGER DEFAULT 0,
    about_image_four  INTEGER DEFAULT 0,

    new_product_one   TEXT DEFAULT '',
    product_one       TEXT DEFAULT '',

    contact_one       TEXT DEFAULT '',
    contact_two       TEXT DEFAULT '',
    contact_three     TEXT DEFAULT '',
    contact_four      TEXT DEFAULT '',
    contact_five      TEXT DEFAULT '',

    FOREIGN KEY (admin_id) REFERENCES admins (id)
);

CREATE TABLE IF NOT EXISTS customers (
    id                    BIGSERIAL PRIMARY KEY,
    website_id            INTEGER NOT NULL,
    email                 TEXT NOT NULL,
    password_hash         TEXT NOT NULL,

    first_name            TEXT DEFAULT '',
    last_name             TEXT DEFAULT '',
    father_name           TEXT DEFAULT '',
    phone                 TEXT DEFAULT '',
    telegram              TEXT DEFAULT '',
    delivery_type         TEXT DEFAULT '',
    payment_type          TEXT DEFAULT '',

    email_notification    INTEGER DEFAULT 0,
    telegram_notification INTEGER DEFAULT 0,

    FOREIGN KEY (website_id) REFERENCES websites (id)
);

CREATE TABLE IF NOT EXISTS images (
    id   BIGSERIAL PRIMARY KEY,
    path TEXT
);

CREATE TABLE IF NOT EXISTS products (
    id          BIGSERIAL PRIMARY KEY,
    website_id  BIGINT NOT NULL,
    name        TEXT,
    description TEXT,
    price       INTEGER,
    images_id   TEXT,
    active      BOOLEAN,
    tags        TEXT DEFAULT '',

    FOREIGN KEY (website_id) REFERENCES websites (id)
);

CREATE TABLE IF NOT EXISTS carts (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT UNIQUE NOT NULL,

    FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE IF NOT EXISTS cart_items (
    id         BIGSERIAL PRIMARY KEY,
    cart_id    BIGINT NOT NULL ,
    product_id BIGINT NOT NULL ,
    count      INTEGER,

    FOREIGN KEY (cart_id) REFERENCES carts (id),
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS saved_products (
    id          BIGSERIAL PRIMARY KEY,
    website_id  INTEGER NOT NULL,
    name        TEXT,
    description TEXT,
    price       INTEGER,
    images_id   TEXT,
    active      INTEGER,
    tags        TEXT DEFAULT '',

    FOREIGN KEY (website_id) REFERENCES websites (id)
);

CREATE TABLE IF NOT EXISTS orders (
    id          BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL ,
    date_time   TEXT,
    status      INTEGER,
    comment     TEXT,

    FOREIGN KEY (customer_id) REFERENCES customers (id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id               BIGSERIAL PRIMARY KEY,
    order_id         BIGINT NOT NULL,
    saved_product_id BIGINT NOT NULL,
    count            INTEGER,

    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (saved_product_id) REFERENCES saved_products (id)
);

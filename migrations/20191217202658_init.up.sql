CREATE TABLE album
(
    id         VARCHAR PRIMARY KEY,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product
(
    id              SERIAL PRIMARY KEY,     -- uint32
    name            VARCHAR NOT NULL,
    supplier_id     SERIAL NOT NULL,        -- uint32
    category_id     SERIAL NOT NULL,        -- uint32
    units_in_stock  SERIAL NOT NULL,        -- uint32
    unit_price      FLOAT NOT NULL,
    discontinued    BOOLEAN NOT NULL
);
 
-- TABLE "products" --
CREATE TABLE IF NOT EXISTS products (
  "url"           TEXT      UNIQUE PRIMARY KEY,
  "name"          TEXT      NOT NULL,
  "old_price"     INTEGER   DEFAULT 0,
  "current_price" INTEGER   NOT NULL,
  "quantity"      INTEGER   NOT NULL,
  "image_url"     TEXT,
  "category"      TEXT []
);
-- TABLE "products" --
CREATE TABLE IF NOT EXISTS products (
  "product_id"    TEXT      UNIQUE PRIMARY KEY,
  "url"           TEXT      UNIQUE,
  "name"          TEXT      NOT NULL,
  "old_price"     INTEGER   DEFAULT 0,
  "current_price" INTEGER   NOT NULL,
  "image_url"     TEXT,
  "category"      TEXT []
);
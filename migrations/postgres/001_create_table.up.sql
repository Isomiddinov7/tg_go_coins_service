CREATE TABLE IF NOT EXISTS "coins"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "coin_icon" TEXT,
    "coin_buy_price" VARCHAR NOT NULL,
    "coin_sale_price" VARCHAR NOT NULL,
    "address" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "half_coins_price"(
    "coin_id" UUID NOT NULL REFERENCES "coins"("id"),
    "halfCoinAmount" VARCHAR NOT NULL,
    "halfCoinPrice" VARCHAR NOT NULL
);
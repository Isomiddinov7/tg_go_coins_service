CREATE TYPE StatusUser AS ENUM('active', 'inactive');
CREATE TYPE BuyOrSell AS ENUM('buy', 'sell');
CREATE TYPE MessageStatus AS ENUM('user', 'admin');
CREATE TYPE MessageReadStatus AS ENUM('false', 'true');
CREATE TYPE TransactionStatus AS ENUM('pending', 'success', 'error');

CREATE TABLE IF NOT EXISTS "telegram_user"(
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR NOT NULL,
    "telegram_id" VARCHAR NOT NULL
);


CREATE TABLE IF NOT EXISTS "coins"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "coin_icon" VARCHAR NOT NULL,
    "coin_buy_price" VARCHAR NOT NULL,
    "coin_sell_price" VARCHAR NOT NULL,
    "address" VARCHAR,
    "card_number" VARCHAR,
    "status" BOOLEAN DEFAULT false,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "half_coins_price"(
    "coin_id" UUID NOT NULL REFERENCES "coins"("id"),
    "halfCoinAmount" VARCHAR NOT NULL,
    "halfCoinPrice" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "users"(
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR,
    "username" VARCHAR,
    "status" StatusUser DEFAULT 'active',
    "telegram_id" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "admin"(
    "id" UUID NOT NULL PRIMARY KEY,
    "login" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "admin_address"(
    "admin_id" UUID NOT NULL REFERENCES "admin"("id"),
    "coin_id" UUID NOT NULL REFERENCES "coins"("id"),
    "address" VARCHAR NOT NULL,
    "card_number" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "user_transaction"(
    "id" UUID NOT NULL PRIMARY KEY,
    "coin_id" UUID NOT NULL REFERENCES "coins"("id"),
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "user_confirmation_img" VARCHAR NOT NULL,
    "coin_price" VARCHAR NOT NULL,
    "coin_amount" VARCHAR NOT NULL,
    "all_price" VARCHAR NOT NULL,
    "status" BuyOrSell NOT NULL,
    "user_address" VARCHAR,
    "card_name" VARCHAR,
    "payment_card" VARCHAR,
    "message" TEXT,
    "transaction_status" TransactionStatus DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "messages"(
    "id" UUID NOT NULL PRIMARY KEY,
    "status" MessageStatus NOT NULL,
    "message" TEXT NOT NULL,
    "read" MessageReadStatus NOT NULL,
    "admin_id" UUID NOT NULL REFERENCES "admin"("id"),
    "user_id" UUID NOT NULL REFERENCES "users"("id"), 
    "file" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "premium"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "card_number" VARCHAR NOT NULL,
    "img" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "premium_price_month"(
    "id" UUID NOT NULL PRIMARY KEY,
    "month" VARCHAR NOT NULL,
    "price" VARCHAR NOT NULL,
    "premium_id" UUID NOT NULL REFERENCES "premium"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);


CREATE TABLE IF NOT EXISTS "premium_transaction"(
    "id" UUID NOT NULL PRIMARY KEY,
    "phone_number" VARCHAR NOT NULL,
    "telegram_username" VARCHAR NOT NULL,
    "user_id" UUID NOT NULL REFERENCES "users"("id"),
    "price_id" UUID NOT NULL REFERENCES "premium_price_month"("id"),
    "payment_img" VARCHAR NOT NULL,
    "transaction_status" TransactionStatus DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "stars"(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "count" VARCHAR NOT NULL,
    "price" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "starts_img"(
    "id" UUID NOT NULL PRIMARY KEY,
    "stars_img" VARCHAR NOT NULL,
    "stars_id" NOT NULL REFERENCES "stars"("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
)

CREATE TABLE IF NOT EXISTS "stars_transaction"(
    "id" UUID NOT NULL PRIMARY KEY,
    "stars_id" NOT NULL REFERENCES "stars"("id"),
    "transfer_img" VARCHAR NOT NULL,
    "stars_count" VARCHAR NOT NULL,
    "stars_price" VARCHAR NOT NULL,
    "user_name" VARCHAR NOT NULL,
    "telegram_id" VARCHAR NOT NULL,
    "status" TransactionStatus DEFAULT 'pending',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

-- INSERT INTO "admin"("id", "login", "password") VALUES('dbecf401-64b3-4b9b-829a-c8b061431286', 'bahodir2809', '123456789');
-- INSERT INTO "super_admin"("id","login","password") VALUES('690d15b1-b3bf-416f-83e1-02b183ccb2f2', 'azam1222', '938791222');
-- INSERT INTO "admin_address"("admin_id", "coin_id", "address") VALUES('dbecf401-64b3-4b9b-829a-c8b061431286', 'ecd98c25-4cd3-41f7-8526-5efe021533f7', 'addres$$TON');
-- [
--       {"HalfCoinAmount": "0.5", "HalfCoinPrice": "650000"},
--       {"HalfCoinAmount": "0.8", "HalfCoinPrice": "80000"}
-- ]
-- CREATE TABLE IF NOT EXISTS "sell_coin"(
--     "user_id" UUID NOT NULL REFERENCES "users"("id"),
--     "coin_id" UUID NOT NULL REFERENCES "coins"("id"),
--     "address" VARCHAR NOT NULL,
--     "coin_amount" VARCHAR NOT NULL,
--     "number_of_card" VARCHAR NOT NULL,
--     "check_img" TEXT,
--     "price" VARCHAR NOT NULL,
--     "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     "updated_at" TIMESTAMP
-- );


    -- login password 
    --     success
    --     accsess token
        




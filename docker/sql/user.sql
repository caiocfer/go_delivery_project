
CREATE DATABASE "DeliveryProject";

CREATE TABLE "users" ( 
    "user_id" INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR UNIQUE NOT NULL,
    "password_hash" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE "restaurants" (
    "restaurant_id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "owner_name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password_hash" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "restaurant_name" VARCHAR NOT NULL,
    "cuisine_type" VARCHAR,
    "description" TEXT,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE "menu_items" (
    "restaurant_id" INT NOT NULL,
    "item_id" INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "description" TEXT,
    "category" VARCHAR,
    "image_url" varchar,
    "price" DECIMAL(10, 2) NOT NULL,
    "is_avaliable" BOOLEAN DEFAULT TRUE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_restaurant FOREIGN KEY ("restaurant_id") REFERENCES "restaurants" ("restaurant_id") ON DELETE CASCADE
);

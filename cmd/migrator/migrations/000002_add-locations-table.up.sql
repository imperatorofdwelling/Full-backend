CREATE TABLE IF NOT EXISTS locations (
                                         "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                         "city" VARCHAR(255) NOT NULL,
                                         "federal_district" VARCHAR(255),
                                         "fias_id" VARCHAR(255),
                                         "kladr_id" VARCHAR(255),
                                         "lat" VARCHAR(255),
                                         "lon" VARCHAR(255),
                                         "okato" VARCHAR(255),
                                         "oktmo" VARCHAR(255),
                                         "population" FLOAT,
                                         "region_iso_code" VARCHAR(255),
                                         "region_name" VARCHAR(255) NOT NULL,
                                         "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
                                         "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
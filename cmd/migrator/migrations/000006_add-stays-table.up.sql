CREATE TYPE STAYS_TYPE AS ENUM ('apartment', 'house', 'hotel');

CREATE TABLE IF NOT EXISTS stays
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    location_id UUID NOT NULL,
    name VARCHAR(255) DEFAULT '',
    type STAYS_TYPE NOT NULL DEFAULT 'apartment',
    number_of_bedrooms INTEGER DEFAULT 0,
    number_of_beds INTEGER DEFAULT 0,
    number_of_bathrooms INTEGER DEFAULT 0,
    guests INTEGER DEFAULT 1,
    rating FLOAT DEFAULT 0.0,
    amenities JSONB DEFAULT JSONB_BUILD_OBJECT(
                                                  'Wi-fi', false,
                                                  'Air conditioner', false,
                                                  'Pets allowed', false,
                                                  'Breakfast included', false,
                                                  'Vacuum cleaner', false,
                                                  'Working area', false,
                                                  'Washing machine', false,
                                                  'TV', false,
                                                  'Home light control', false,
                                                  'Smart door locks', false,
                                                  'Voice assistant', false,
                                                  'Touch control panels', false
                                              ),
    is_smoking_prohibited BOOLEAN DEFAULT false,
    square FLOAT NOT NULL,
    street VARCHAR(255) DEFAULT '',
    house VARCHAR(255) DEFAULT '',
    entrance VARCHAR(255) DEFAULT '',
    floor VARCHAR(255) DEFAULT '',
    room VARCHAR(255) DEFAULT '',
    price DECIMAL(15, 2) DEFAULT 0.0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE NO ACTION,
    FOREIGN KEY (location_id) REFERENCES locations(id) ON DELETE NO ACTION
);

/*
Пример для запросов amenities
    SELECT *
    FROM stays
    WHERE amenities ->> 'Wi-fi' = 'true';
*/

CREATE INDEX stays_amenities_idx ON stays USING GIN (amenities jsonb_ops);
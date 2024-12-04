DO $$
    DECLARE
        location_id_moscow UUID;
        location_id_sp UUID;
    BEGIN
        SELECT id INTO location_id_moscow FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1;
        SELECT id INTO location_id_sp FROM locations WHERE LOWER(city) = LOWER('Санкт-Петербург') LIMIT 1;

        IF location_id_moscow IS NULL THEN
            RAISE EXCEPTION 'Location not found for city: Москва';
        END IF;

        IF location_id_sp IS NULL THEN
            RAISE EXCEPTION 'Location not found for city: Санкт-Петербург';
        END IF;

            INSERT INTO stays (
                id,
                user_id,
                location_id,
                name,
                type,
                number_of_bedrooms,
                number_of_beds,
                number_of_bathrooms,
                guests,
                rating,
                is_smoking_prohibited,
                square,
                street,
                house,
                entrance,
                floor,
                room,
                price,
                created_at,
                updated_at
            ) VALUES
                 ('e9b1f840-1a75-4a3d-9e92-fb762c40d8be', '550e8400-e29b-41d4-a716-446655440000', location_id_moscow, 'Cozy Apartment', 'apartment', 2, 3, 1, 4, 4.5, TRUE, 50.0, 'Main St', '10A', 'A', '1', '101', 100.0, NOW(), NOW()),
                 ('36d1d6f3-de13-4e66-9b12-1d5fba95a5ec', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Luxury Condo', 'apartment', 3, 4, 2, 6, 5.0, TRUE, 80.0, 'Second St', '20B', 'B', '2', '202', 200.0, NOW(), NOW()),
                 ('7e1d8c8f-bc0e-4f54-b5d6-cc188d60e31f', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Beach House', 'house', 4, 5, 3, 8, 4.8, FALSE, 120.0, 'Ocean Dr', '30C', NULL, NULL, NULL, 300.0, NOW(), NOW()),
                 ('5372b0f8-fc01-4c5d-9a6e-90e034fbc8b9', '550e8400-e29b-41d4-a716-446655440001', location_id_sp, 'Mountain Cabin', 'apartment', 2, 2, 1, 4, 4.0, TRUE, 70.0, 'Hill Rd', '40D', 'D', '1', '404', 150.0, NOW(), NOW()),
                 ('dbb8e576-5b64-4c5e-b0e1-6124eef86b5e', '550e8400-e29b-41d4-a716-446655440000', location_id_moscow, 'City Studio', 'apartment', 1, 1, 1, 2, 4.2, FALSE, 25.0, 'Downtown', '50E', 'E', '5', '505', 80.0, NOW(), NOW()),
                 ('8f8e9858-bdb6-498e-8447-58c5c5c8ae35', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Family Home', 'apartment', 3, 4, 2, 5, 4.6, TRUE, 90.0, 'Park Ln', '60F', 'F', '2', '606', 120.0, NOW(), NOW()),
                 ('a5fcb5c4-0b3a-4084-b785-4cf4e08d8a19', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Charming Cottage', 'house', 2, 3, 1, 3, 4.3, TRUE, 55.0, 'Old St', '70G', NULL, NULL, NULL, 110.0, NOW(), NOW()),
                 ('0ea0e4b7-2c07-4263-a57f-537fa82f6d8c', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Penthouse Suite', 'apartment', 5, 6, 3, 10, 4.9, TRUE, 150.0, 'Skyline Ave', '80H', 'H', '8', '808', 400.0, NOW(), NOW()),
                 ('9b758493-0bca-4142-8d30-ec536f68b6f5', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Suburban Bungalow', 'hotel', 3, 3, 2, 5, 4.7, FALSE, 65.0, 'Elm St', '90I', 'I', '1', '909', 140.0, NOW(), NOW()),
                 ('1fe5c282-a9ef-4d11-b0c2-9370f5e24a99', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Loft Apartment', 'house', 1, 1, 1, 1, 4.1, TRUE, 30.0, 'Creative St', '100J', NULL, NULL, NULL, 75.0, NOW(), NOW());

END $$;
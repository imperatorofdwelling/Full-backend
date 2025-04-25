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
    guests,
    amenities,
    house,
    entrance,
    address,
    rooms_count,
    beds_count,
    price,
    period,
    owners_rules,
    cancellation_policy,
    describe_property,
    created_at,
    updated_at
) VALUES
      ('e9b1f840-1a75-4a3d-9e92-fb762c40d8be', '550e8400-e29b-41d4-a716-446655440000', location_id_moscow, 'Cozy Apartment', 'apartment', 4, '{"Wi-fi": true, "Air conditioner": true}', '10A', 'A', 'Main St', '2', '3', '100.0', 'day', 'No smoking', 'Free cancellation 24h before check-in', 'A cozy apartment in the heart of Moscow', NOW(), NOW()),
      ('36d1d6f3-de13-4e66-9b12-1d5fba95a5ec', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Luxury Condo', 'apartment', 6, '{"Wi-fi": true, "Air conditioner": true, "Kitchen": true}', '20B', 'B', 'Second St', '3', '4', '200.0', 'day', 'No parties, no pets', 'Moderate - 50% refund up to 5 days before check-in', 'Luxurious condo with amazing city views', NOW(), NOW()),
      ('7e1d8c8f-bc0e-4f54-b5d6-cc188d60e31f', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Beach House', 'house', 8, '{"Wi-fi": true, "Air conditioner": false, "Pool": true}', '30C', '', 'Ocean Dr', '4', '5', '300.0', 'day', 'Quiet hours after 10PM', 'Strict - No refunds', 'Beautiful beach house with ocean views', NOW(), NOW()),
      ('5372b0f8-fc01-4c5d-9a6e-90e034fbc8b9', '550e8400-e29b-41d4-a716-446655440001', location_id_sp, 'Mountain Cabin', 'apartment', 4, '{"Wi-fi": false, "Fireplace": true}', '40D', 'D', 'Hill Rd', '2', '2', '150.0', 'day', 'No smoking inside', 'Flexible - Full refund 1 day prior to arrival', 'Cozy mountain cabin perfect for a getaway', NOW(), NOW()),
      ('dbb8e576-5b64-4c5e-b0e1-6124eef86b5e', '550e8400-e29b-41d4-a716-446655440000', location_id_moscow, 'City Studio', 'apartment', 2, '{"Wi-fi": true}', '50E', 'E', 'Downtown', '1', '1', '80.0', 'day', 'Keep noise to a minimum', 'Moderate', 'Modern studio in downtown Moscow', NOW(), NOW()),
      ('8f8e9858-bdb6-498e-8447-58c5c5c8ae35', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Family Home', 'apartment', 5, '{"Wi-fi": true, "Air conditioner": true, "Washer": true}', '60F', 'F', 'Park Ln', '3', '4', '120.0', 'day', 'Family-friendly, no parties', 'Flexible', 'Spacious family home near the park', NOW(), NOW()),
      ('a5fcb5c4-0b3a-4084-b785-4cf4e08d8a19', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Charming Cottage', 'house', 3, '{"Wi-fi": true, "Garden": true}', '70G', '', 'Old St', '2', '3', '110.0', 'day', 'Respect the neighbors', 'Moderate', 'Charming cottage with beautiful garden', NOW(), NOW()),
      ('0ea0e4b7-2c07-4263-a57f-537fa82f6d8c', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Penthouse Suite', 'apartment', 10, '{"Wi-fi": true, "Air conditioner": true, "Pool": true, "Gym": true}', '80H', 'H', 'Skyline Ave', '5', '6', '400.0', 'day', 'No smoking, no pets', 'Strict', 'Luxurious penthouse with panoramic city views', NOW(), NOW()),
      ('9b758493-0bca-4142-8d30-ec536f68b6f5', '550e8400-e29b-41d4-a716-446655440000', location_id_sp, 'Suburban Bungalow', 'hotel', 5, '{"Wi-fi": true, "Air conditioner": true, "Parking": true}', '90I', 'I', 'Elm St', '3', '3', '140.0', 'day', 'Check-out by 11 AM', 'Flexible', 'Comfortable bungalow in quiet suburban area', NOW(), NOW()),
      ('1fe5c282-a9ef-4d11-b0c2-9370f5e24a99', '550e8400-e29b-41d4-a716-446655440001', location_id_moscow, 'Loft Apartment', 'house', 1, '{"Wi-fi": true, "Air conditioner": false}', '100J', '', 'Creative St', '1', '1', '75.0', 'day', 'Art-friendly space', 'Moderate', 'Creative loft space in artistic district', NOW(), NOW());
END $$;
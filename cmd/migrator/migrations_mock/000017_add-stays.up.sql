CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

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
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440000', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'Cozy Apartment', 'type1', 2, 3, 1, 4, 4.5, TRUE, 50.0, 'Main St', '10A', 'A', '1', '101', 100.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440001', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'Luxury Condo', 'type2', 3, 4, 2, 6, 5.0, TRUE, 80.0, 'Second St', '20B', 'B', '2', '202', 200.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440000', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Санкт-Петербург') LIMIT 1), 'Beach House', 'type3', 4, 5, 3, 8, 4.8, FALSE, 120.0, 'Ocean Dr', '30C', NULL, NULL, NULL, 300.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440001', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Санкт-Петербург') LIMIT 1), 'Mountain Cabin', 'type4', 2, 2, 1, 4, 4.0, TRUE, 70.0, 'Hill Rd', '40D', 'D', '1', '404', 150.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440000', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'City Studio', 'type1', 1, 1, 1, 2, 4.2, FALSE, 25.0, 'Downtown', '50E', 'E', '5', '505', 80.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440001', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'Family Home', 'type2', 3, 4, 2, 5, 4.6, TRUE, 90.0, 'Park Ln', '60F', 'F', '2', '606', 120.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440000', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Санкт-Петербург') LIMIT 1), 'Charming Cottage', 'type3', 2, 3, 1, 3, 4.3, TRUE, 55.0, 'Old St', '70G', NULL, NULL, NULL, 110.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440001', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'Penthouse Suite', 'type4', 5, 6, 3, 10, 4.9, TRUE, 150.0, 'Skyline Ave', '80H', 'H', '8', '808', 400.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440000', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Санкт-Петербург') LIMIT 1), 'Suburban Bungalow', 'type1', 3, 3, 2, 5, 4.7, FALSE, 65.0, 'Elm St', '90I', 'I', '1', '909', 140.0, NOW(), NOW()),
      (uuid_generate_v4(), '550e8400-e29b-41d4-a716-446655440001', (SELECT id FROM locations WHERE LOWER(city) = LOWER('Москва') LIMIT 1), 'Loft Apartment', 'type2', 1, 1, 1, 1, 4.1, TRUE, 30.0, 'Creative St', '100J', NULL, NULL, NULL, 75.0, NOW(), NOW());
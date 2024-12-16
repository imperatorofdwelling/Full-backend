INSERT INTO users (
    id,
    name,
    password,
    email,
    phone,
    avatar,
    birth_date,
    national,
    gender,
    country,
    city,
    created_at,
    updated_at
) VALUES
      ('550e8400-e29b-41d4-a716-446655440000', 'John Doe', '12345678', 'john.doe@example.com', '+1234567890', NULL, '1990-05-20', 'American', 'Male', 'USA', 'New York', NOW(), NOW()),
      ('550e8400-e29b-41d4-a716-446655440001', 'Jane Smith', '12345678', 'jane.smith@example.com', '+0987654321', NULL, '1985-12-15', 'Canadian', 'Female', 'Canada', 'Toronto', NOW(), NOW());
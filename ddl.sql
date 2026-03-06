CREATE TABLE project_users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone VARCHAR(20) UNIQUE,
    role VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE project_rooms (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    room_number VARCHAR(20) NOT NULL,
    room_type VARCHAR(50),
    price DECIMAL(10,2) NOT NULL,
    capacity INT NOT NULL,
    status VARCHAR(20) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE TABLE project_bookings (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    room_id INT NOT NULL,
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    booking_status VARCHAR(20) DEFAULT 'confirmed',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id)
    REFERENCES project_users(id)
    ON DELETE CASCADE,

    FOREIGN KEY (room_id)
    REFERENCES project_rooms(id)
    ON DELETE CASCADE
);

CREATE INDEX idx_user_email ON project_users(email);
CREATE INDEX idx_booking_user_id ON project_bookings(user_id);
CREATE INDEX idx_booking_room_id ON project_bookings(room_id);

-- data seeding for rooms
INSERT INTO project_rooms (room_number, room_type, price, capacity, status)
VALUES
('101', 'Deluxe', 200.00, 2, 'available'),
('102', 'Deluxe', 200.00, 2, 'unavailable'),
('103', 'Deluxe', 200.00, 2, 'available'),

('201', 'Suites', 350.00, 4, 'available'),
('202', 'Suites', 350.00, 4, 'available'),
('203', 'Suites', 350.00, 4, 'unavailable'),

('301', 'Presidential', 800.00, 6, 'available'),
('302', 'Presidential', 800.00, 6, 'available'),
('303', 'Presidential', 800.00, 6, 'unavailable');
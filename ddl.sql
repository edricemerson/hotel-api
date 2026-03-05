CREATE TABLE project_users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE project_hotels (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    location VARCHAR(200) NOT NULL,
    description TEXT,
    rating DECIMAL(2,1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE project_rooms (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    hotel_id INT NOT NULL,
    room_number VARCHAR(20) NOT NULL,
    room_type VARCHAR(50),
    price DECIMAL(10,2) NOT NULL,
    capacity INT NOT NULL,
    status VARCHAR(20) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (hotel_id)
    REFERENCES project_hotels(id)
    ON DELETE CASCADE
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

CREATE TABLE project_booking_payments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    booking_id INT NOT NULL,
    payment_method VARCHAR(50),
    payment_status VARCHAR(20) DEFAULT 'pending',
    amount DECIMAL(10,2) NOT NULL,
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (booking_id)
    REFERENCES project_bookings(id)
    ON DELETE CASCADE
);

CREATE VIEW project_booking_report AS
SELECT
    b.id AS booking_id,
    u.name AS user_name,
    u.email,
    h.name AS hotel_name,
    r.room_number,
    r.room_type,
    b.check_in,
    b.check_out,
    b.total_price,
    b.booking_status,
    b.created_at
FROM project_bookings b
JOIN project_users u ON b.user_id = u.id
JOIN project_rooms r ON b.room_id = r.id
JOIN project_hotels h ON r.hotel_id = h.id;

CREATE INDEX idx_user_email ON project_users(email);
CREATE INDEX idx_rooms_hotel_id ON project_rooms(hotel_id);
CREATE INDEX idx_booking_user_id ON project_bookings(user_id);
CREATE INDEX idx_booking_room_id ON project_bookings(room_id);
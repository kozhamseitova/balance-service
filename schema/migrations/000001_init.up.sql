CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    balance DECIMAL(10, 2) NOT NULL DEFAULT 0
);

CREATE TABLE Reservations (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) NOT NULL,
    service_id INT NOT NULL,
    order_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    reserved_at TIMESTAMPTZ DEFAULT current_timestamp,
    recognized_at TIMESTAMPTZ
);

CREATE TABLE RevenueReports (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES Users(id) NOT NULL,
    service_id INT NOT NULL,
    order_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    recognition_date TIMESTAMPTZ DEFAULT current_timestamp
);

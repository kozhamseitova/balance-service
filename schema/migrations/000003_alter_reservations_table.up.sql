ALTER TABLE reservations
ADD CONSTRAINT unique_user_service UNIQUE (user_id, service_id);

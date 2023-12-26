-- Create the role_type enum
CREATE TYPE role_type AS ENUM ('admin', 'user');

-- Create the reservation_type enum
CREATE TYPE reservation_type AS ENUM ('active', 'decline', 'ended');

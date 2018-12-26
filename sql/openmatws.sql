CREATE TABLE openmatdata(
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    days_hours CHAR(50) NOT NULL,
    street TEXT NOT NULL,
    phone CHAR(50) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);

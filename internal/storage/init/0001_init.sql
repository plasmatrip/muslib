CREATE USER muslib WITH ENCRYPTED PASSWORD 'password';

CREATE DATABASE muslib OWNER 'muslib';

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO muslib;

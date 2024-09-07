
-- Create table tasks if not exists
CREATE TABLE IF NOT ESIXT tasks (
  id SERIAL PRIMARY KEY,
  type INT NOT NULL,
  value INT NOT NULL,
  state VARCHAR(50),
  creation_time TIMESTAMP,
  last_update_time TIMESTAMP
);

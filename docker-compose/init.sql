CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid()
, name STRING UNIQUE NOT NULL
, age INT
, married BOOLEAN
, location JSONB
, phone_numbers STRING[]
, creation_time TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO users (name, age, married
, location, phone_numbers) VALUES
('Tom', '31', TRUE, '{"state": "CA", "city": "LA"}', ARRAY['000-0000', '001-1234'])
, ('Nathan', '25', TRUE, '{"state": "UT", "city": "Salt Lake City"}', ARRAY['987-6666'])
, ('Mary', NULL, FALSE, '{"state": "AZ", "city": "Phoenix", "street": "Main"}', NULL)
ON CONFLICT (name) DO NOTHING
;

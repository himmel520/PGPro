CREATE TABLE IF NOT EXISTS commands (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(64) NOT NULL,
    description TEXT,
    script TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS commands_info (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    commands_id UUID REFERENCES commands(id),
    start_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP DEFAULT NULL,
    CONSTRAINT valid_end_time CHECK (end_time >= start_time),
    exitcode INTEGER DEFAULT -1,
    output text DEFAULT ''
);
CREATE TABLE inputs (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    format TEXT NOT NULL
);

CREATE TABLE audio_profiles (
    id SERIAL PRIMARY KEY,
    input_id UUID NOT NULL,
    bitrate INT,
    codec TEXT NOT NULL,
    FOREIGN KEY (input_id) REFERENCES inputs(id)
);

CREATE TABLE video_profiles (
    id SERIAL PRIMARY KEY,
    input_id UUID NOT NULL,
    codec TEXT NOT NULL,
    bitrate INT,
    max_key_interval INT,
    framerate INT,
    width INT,
    height INT,
    FOREIGN KEY (input_id) REFERENCES inputs(id)
);

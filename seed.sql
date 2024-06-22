-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create the videos table
CREATE TABLE videos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create the video_tags table as a join table between users and videos
CREATE TABLE video_tags (
    user_id INT NOT NULL,
    video_id INT NOT NULL,
    PRIMARY KEY (user_id, video_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (video_id) REFERENCES videos(id)
);

-- Insert sample data into users table
INSERT INTO users (name) VALUES 
('Alice'),
('Bob'),
('Charlie');

-- Insert sample data into videos table
INSERT INTO videos (name) VALUES 
('Video 1'),
('Video 2'),
('Video 3');

-- Insert sample data into video_tags table
INSERT INTO video_tags (user_id, video_id) VALUES 
(1, 1),
(1, 2),
(2, 3),
(3, 1),
(3, 3);

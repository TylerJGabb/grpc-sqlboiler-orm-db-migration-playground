-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- Ensure that all timestamp operations are in UTC
SET TIME ZONE 'UTC';

CREATE TABLE change_requests (
    id SERIAL PRIMARY KEY,
    github_pr_id varchar(255),
    github_pr_url VARCHAR(255),
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    docs TEXT
);

CREATE TABLE status_events (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    change_request_id INTEGER,
    FOREIGN KEY (change_request_id) REFERENCES change_requests(id)
);

INSERT INTO change_requests (github_pr_id, github_pr_url, created_by, docs)
VALUES
(101, 'https://github.com/example/repo/pull/101', 'alice', 'Initial setup of the project.'),
(102, 'https://github.com/example/repo/pull/102', 'bob', 'Added new features.'),
(103, 'https://github.com/example/repo/pull/103', 'charlie', 'Fixed bugs and performance issues.');

INSERT INTO status_events (status, timestamp, change_request_id)
VALUES
('pending', '2024-06-01 10:00:00+00', 1),
('open', '2024-06-02 15:00:00+00', 1),
('rebased', '2024-06-03 09:30:00+00', 1),
('confict', '2024-06-04 16:45:00+00', 1),
('pending', '2024-06-05 11:00:00+00', 3),
('open', '2024-06-06 14:00:00+00', 3),
('closed', '2024-06-07 14:00:00+00', 3);

-- +migrate Down
-- SQL in section 'Down' is executed when this migration is rolled back

DROP TABLE IF EXISTS status_events;
DROP TABLE IF EXISTS change_requests;

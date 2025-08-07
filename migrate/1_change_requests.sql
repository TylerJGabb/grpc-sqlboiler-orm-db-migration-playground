-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

-- Ensure that all timestamp operations are in UTC
SET TIME ZONE 'UTC';


CREATE TABLE change_requests (
    id SERIAL PRIMARY KEY,
    github_branch_name VARCHAR(255),
    github_pr_id varchar(255),
    github_pr_url VARCHAR(255),
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    closed BOOLEAN DEFAULT FALSE,
    type VARCHAR(255) NOT NULL,
    docs TEXT
);

INSERT INTO change_requests (github_pr_id, github_pr_url, created_by, docs, type)
VALUES
('101', 'https://github.com/example/repo/pull/101', 'alice', 'Initial setup of the project.', 'tmt-project'),
('102', 'https://github.com/example/repo/pull/102', 'bob', 'Added new features.', 'tmt-project'),
('103', 'https://github.com/example/repo/pull/103', 'charlie', 'Fixed bugs and performance issues.', 'tmt-project');

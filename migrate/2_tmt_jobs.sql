-- +migrate Up

CREATE TYPE job_status AS ENUM (
    'pending',
    'failed',
    'completed'
);

CREATE TABLE tmt_jobs(
  id SERIAL PRIMARY KEY,
  change_request_id INT NOT NULL,
  project_name VARCHAR(255) NOT NULL,
  orchestration_repository VARCHAR(255) NOT NULL,
  application VARCHAR(255) NOT NULL,
  dv01_domain VARCHAR(255) NOT NULL,
  user_email VARCHAR(255) NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,

  status job_status NOT NULL DEFAULT 'pending',
  status_message TEXT,

  FOREIGN KEY (change_request_id) REFERENCES change_requests(id)
);

INSERT INTO tmt_jobs (
  change_request_id,
  project_name,
  orchestration_repository,
  application,
  dv01_domain,user_email,
  status,
  status_message,
  completed_at
) VALUES
(1, 'project1', 'repo1', 'app1', 'domain1', 'alice', 'pending', null, null),
(2, 'project2', 'repo2', 'app2', 'domain2', 'bob', 'failed', 'Failed to create PR for TMT', '2021-01-01 00:00:00'),
(3, 'project3', 'repo3', 'app3', 'domain3', 'charlie', 'completed', 'Successfully created PR for TMT', '2021-01-02 00:00:00');

-- +migration Down
DROP TABLE IF EXISTS tmt_request;

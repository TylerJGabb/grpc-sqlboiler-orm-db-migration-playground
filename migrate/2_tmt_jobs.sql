-- +migrate Up

CREATE TABLE tmt_jobs(
  id SERIAL PRIMARY KEY,
  change_request_id INT NOT NULL,
  project_name VARCHAR(255) NOT NULL,
  orchestration_repository VARCHAR(255) NOT NULL,
  application VARCHAR(255) NOT NULL,
  tenant_domain VARCHAR(255) NOT NULL,
  user_email VARCHAR(255) NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,

  status VARCHAR(255) NOT NULL,
  status_message TEXT,

  FOREIGN KEY (change_request_id) REFERENCES change_requests(id)
);

INSERT INTO tmt_jobs (
  change_request_id,
  project_name,
  orchestration_repository,
  application,
  tenant_domain,user_email,
  status,
  status_message,
  completed_at
) VALUES
(1, 'project1', 'repo1', 'app1', 'domain1', 'alice', 'PENDING', null, null),
(2, 'project2', 'repo2', 'app2', 'domain2', 'bob', 'FAILED', 'Failed to create PR for TMT', '2021-01-01 00:00:00'),
(3, 'project3', 'repo3', 'app3', 'domain3', 'charlie', 'COMPLETED', 'Successfully created PR for TMT', '2021-01-02 00:00:00');


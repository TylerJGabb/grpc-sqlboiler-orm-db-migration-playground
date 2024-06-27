-- +migrate Up
CREATE TABLE rebase_jobs(
  id SERIAL PRIMARY KEY,
  change_request_id INT NOT NULL,

  created_at TIMESTAMP DEFAULT NOT NULL CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,

  status job_status NOT NULL DEFAULT 'pending',
  status_message TEXT,

  FOREIGN KEY (change_request_id) REFERENCES change_requests(id)
);

INSERT INTO rebase_jobs (
  change_request_id,
  status,
  status_message
) VALUES
(1, 'pending', null),
(2, 'failed', 'Failed to rebase becasue of blah blah.... need manual intervention'),
(3, 'completed', 'Successfully rebased off of commit abcdefg');

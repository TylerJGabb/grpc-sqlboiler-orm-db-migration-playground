-- +migrate Up
CREATE TABLE rebase_jobs(
  id SERIAL PRIMARY KEY,
  change_request_id INT NOT NULL,

  created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP,

  status VARCHAR(255) NOT NULL,
  status_message TEXT,

  FOREIGN KEY (change_request_id) REFERENCES change_requests(id)
);

INSERT INTO rebase_jobs (
  change_request_id,
  status,
  status_message
) VALUES
(1, 'PENDING', null),
(2, 'FAILED', 'Failed to rebase becasue of blah blah.... need manual intervention'),
(3, 'COMPLETED', 'Successfully rebased off of commit abcdefg');

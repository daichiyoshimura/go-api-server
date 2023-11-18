CREATE TABLE db.accounts (
  id bigint PRIMARY KEY
  name varchar(255) NOT NULL
  created_at time NOT NULL DEFAULT CURRENT_TIMESTAMP()
  updated_at time NOT NULL DEFAULT CURRENT_TIMESTAMP()
  deleted_at time
);

--bun:split

/*For Unit Test*/
INSERT INTO db.account (:name) VALUES ("JohnSmith");
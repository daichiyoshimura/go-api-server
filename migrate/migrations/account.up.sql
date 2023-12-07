CREATE TABLE db.accounts (
  id binary(16) PRIMARY KEY
  name varchar(255) NOT NULL
  created_at time NOT NULL DEFAULT CURRENT_TIMESTAMP()
  updated_at time NOT NULL DEFAULT CURRENT_TIMESTAMP()
  deleted_at time
);

--bun:split

/*For Unit Test*/
INSERT INTO db.account (:id, :name) VALUES (UUID_TO_BIN("dummy",1), "JohnSmith");
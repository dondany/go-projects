CREATE TABLE lists (
   id SERIAL PRIMARY KEY,
   name TEXT
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    name TEXT,
    completed BOOLEAN,
    list_id INT NOT NULL,
    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE
);

INSERT INTO lists (name) VALUES ('My Todo List');

INSERT INTO todos (name, completed, list_id) VALUES ('Wake up', true, 1);
INSERT INTO todos (name, completed, list_id) VALUES ('Add', false, 1);
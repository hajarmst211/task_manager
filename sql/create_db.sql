-- create_db.sql

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE status_types AS ENUM ("not started", "in progress", "done")
CREATE TYPE priorities AS ENUM ("high", "medium", "low")
CREATE TYPES task_types AS ENUM ("to do list", "text", "date")


CREATE TABLE IF NOT EXISTS blocks(
    block_id UUID UUID DEFAULT gen_random_uuid(),
    parent_id UUID REFERENCES block(block_id) ON DELETE CASCADE,
    tasks TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks(
    task_id UUID DEFAULT gen_random_uuid(),
    block_id UUID REFERENCES blocks(block_id),
    title TEXT NOT NULL,
    task_type task_types,
    task_properties REFERENCES properties(propriety_id),
    properties JSONB DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS properties(
    propriety_id UUID DEFAULT gen_random_uuid(),
    task_id UUID REFERENCES task(task_id),
    is_done BOOLEAN,
    status  status_types,
    deadline DATE
);

CREATE TABLE IF NOT EXISTS tags(
    tag_id UUID DEFAULT gen_random_uuid(),
    label TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS task_tags(
    tag_id UUID REFERENCES tags(tag_id),
    task_id UUID REFERENCES tasks(task_id),
    PRIMARY KEY (task_id, tag_id)
);



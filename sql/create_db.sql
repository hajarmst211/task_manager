-- create_db.sql

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE status_types AS ENUM ('not started', 'in progress', 'done');
CREATE TYPE priorities AS ENUM ('high', 'medium', 'low');
CREATE TYPE task_types AS ENUM ('to do list', 'text', 'date');


CREATE TABLE IF NOT EXISTS blocks(
    block_id UUID DEFAULT gen_random_uuid(),
    parent_id UUID REFERENCES block(block_id) ON DELETE CASCADE,
    tasks TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS tasks(
    task_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    block_id UUID REFERENCES blocks(block_id),
    title TEXT NOT NULL,
    task_type task_types NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    status  status_types DEFAULT 'not started',
    properties JSONB DEFAULT '{}'::jsonb
);


CREATE TABLE IF NOT EXISTS tags(
    tag_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    label TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS task_tags(
    tag_id UUID REFERENCES tags(tag_id) ON DELETE CASCADE,
    task_id UUID REFERENCES tasks(task_id)ON DELETE CASCADE,
    PRIMARY KEY (task_id, tag_id)
);



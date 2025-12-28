-- 1. Create Tags

INSERT INTO tags (label) VALUES 
('Personal project'), 
('School work'), 
('Sports'), 
('Others');


WITH main_block AS (
    INSERT INTO blocks(block_theme) VALUES('2026 goals')
    RETURNING block_id
)
INSERT INTO blocks (parent_id, block_theme) (
    SELECT block_id, 'Professional goals' FROM main_block
); -- this creates another block inside the main block dedicated to professional goals


INSERT INTO blocks(parent_id, block_theme) VALUES(
    (SELECT block_id FROM blocks WHERE block_theme='2026 goals' LIMIT 1),
    'Personal goals'
);


INSERT INTO tasks(block_id, title, task_type, properties) VALUES(
    (SELECT block_id FROM blocks WHERE block_theme = '2026 goals' LIMIT 1),
    'Defeat your social anxiety',
    'to do list',
    '{"checklist": [
                    "Relax while talking", 
                    "Dont judge what you say", 
                    "Think before talking"
                    ], 
        "People to help": "Anouss"
    }'::jsonb
);
 

# Database logic:
- Blocks (1:N) Blocks: Self-referencing relationship for hierarchy.

- Blocks (1:N) Tasks: One block contains many tasks.

- Tasks (M:N) Tags: Handled via the task_tags table.

-- Drop tasks table and its indexes
DROP INDEX IF EXISTS idx_tasks_status;
DROP INDEX IF EXISTS idx_tasks_type;
DROP INDEX IF EXISTS idx_tasks_campaign_id;
DROP TABLE IF EXISTS tasks;

-- Drop campaigns table and its indexes
DROP INDEX IF EXISTS idx_campaigns_mode_status;
DROP INDEX IF EXISTS idx_campaigns_status;
DROP INDEX IF EXISTS idx_campaigns_start_time;
DROP TABLE IF EXISTS campaigns;

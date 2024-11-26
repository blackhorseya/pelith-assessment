-- Drop indexes for the rewards table
DROP INDEX IF EXISTS idx_rewards_user_address;
DROP INDEX IF EXISTS idx_rewards_campaign_id;
DROP INDEX IF EXISTS idx_rewards_redeemed_at;
DROP INDEX IF EXISTS idx_rewards_points;

-- Drop the rewards table
DROP TABLE IF EXISTS rewards;

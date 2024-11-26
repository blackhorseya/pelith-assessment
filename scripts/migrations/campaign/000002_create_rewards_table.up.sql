-- Create rewards table
CREATE TABLE rewards
(
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),                   -- Reward ID
    user_address  VARCHAR(42) NOT NULL,                                         -- Associated user address (e.g., wallet address)
    campaign_id   UUID NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,    -- Associated campaign ID
    points        BIGINT NOT NULL CHECK (points >= 0),                          -- Points used to redeem the reward (must be non-negative)
    redeemed_at   TIMESTAMPTZ,                                                  -- Redemption timestamp (nullable, for pending rewards)
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),                           -- Created timestamp
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()                            -- Updated timestamp
);

-- Optional: Add indexes for query optimization
CREATE INDEX idx_rewards_user_address ON rewards (user_address);
CREATE INDEX idx_rewards_campaign_id ON rewards (campaign_id);
CREATE INDEX idx_rewards_redeemed_at ON rewards (redeemed_at);
CREATE INDEX idx_rewards_points ON rewards (points);

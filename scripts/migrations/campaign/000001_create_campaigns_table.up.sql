-- Create campaigns table
CREATE TABLE campaigns
(
    id          UUID PRIMARY KEY      DEFAULT gen_random_uuid(), -- Campaign ID
    name        VARCHAR(255) NOT NULL UNIQUE,                    -- Campaign name
    description TEXT,                                            -- Campaign description
    start_time  TIMESTAMPTZ  NOT NULL,                           -- Campaign start time
    end_time    TIMESTAMPTZ,                                     -- Campaign end time (optional)
    mode        SMALLINT     NOT NULL DEFAULT 0,                 -- Campaign mode (0: Unspecified, 1: Real-Time, 2: Backtest)
    status      SMALLINT     NOT NULL DEFAULT 0,                 -- Campaign status (0: Unspecified, 1: Pending, 2: Active, 3: Completed)
    pool_id     VARCHAR(42)  NOT NULL,                           -- Pool ID
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),             -- Created timestamp
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()              -- Updated timestamp
);

-- Create tasks table
CREATE TABLE tasks
(
    id          UUID PRIMARY KEY      DEFAULT gen_random_uuid(),                   -- Task ID
    campaign_id UUID         NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE, -- Associated campaign
    name        VARCHAR(255) NOT NULL,                                             -- Task name
    description TEXT,                                                              -- Task description
    type        SMALLINT     NOT NULL DEFAULT 0,                                   -- Task type (0: Unspecified, 1: Onboarding, 2: Share Pool)
    criteria    JSONB        NOT NULL DEFAULT '{}'::JSONB,                         -- Task criteria (stored as JSON)
    status      SMALLINT     NOT NULL DEFAULT 0,                                   -- Task status (0: Unspecified, 1: Active, 2: Inactive)
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),                               -- Created timestamp
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()                                -- Updated timestamp
);

-- Indexes for campaigns
CREATE INDEX idx_campaigns_start_time ON campaigns (start_time); -- Optimize queries by start_time
CREATE INDEX idx_campaigns_status ON campaigns (status); -- Optimize queries by status
CREATE INDEX idx_campaigns_mode_status ON campaigns (mode, status);
-- Composite index for mode and status

-- Indexes for tasks
CREATE INDEX idx_tasks_campaign_id ON tasks (campaign_id); -- Optimize queries by campaign_id
CREATE INDEX idx_tasks_type ON tasks (type); -- Optimize queries by type
CREATE INDEX idx_tasks_status ON tasks (status); -- Optimize queries by status

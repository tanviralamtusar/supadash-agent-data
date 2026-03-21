CREATE TABLE IF NOT EXISTS public.refresh_tokens
(
    id            serial      not null,
    account_id    int         not null,
    token         text        not null unique,
    expires_at    timestamptz not null,
    revoked       boolean     not null default false,
    
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),

    primary key (id),
    CONSTRAINT fk_refresh_token_account FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_token ON public.refresh_tokens(token);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_account_id ON public.refresh_tokens(account_id);

CREATE TABLE IF NOT EXISTS public.audit_logs
(
    id               serial      not null,
    target_project   text,       -- optional: project_ref if Action is project-specific
    actor_id         int         not null, -- account_id of the user performing the action
    action           text        not null, -- e.g. "PROJECT_PAUSED", "ENV_VAR_UPDATED", "MEMBER_INVITED"
    ip_address       text,       -- optional: IP address of the user
    user_agent       text,       -- optional: Browser/Client
    details          jsonb,      -- additional arbitrary payload
    
    created_at       timestamptz not null default now(),

    primary key (id),
    CONSTRAINT fk_audit_log_account FOREIGN KEY (actor_id) REFERENCES accounts (id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_audit_logs_project ON public.audit_logs(target_project);
CREATE INDEX IF NOT EXISTS idx_audit_logs_actor ON public.audit_logs(actor_id);

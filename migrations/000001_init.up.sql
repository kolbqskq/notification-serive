CREATE TABLE notifications (
    ID UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    event_type VARCHAR(30) NOT NULL,
    message VARCHAR(1000) NOT NULL,
    source_service VARCHAR(30),
    status VARCHAR(30) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    sent_at TIMESTAMPTZ,
    error_message VARCHAR(100)
);
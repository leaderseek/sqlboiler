CREATE SCHEMA leaderseek

CREATE TABLE player (
    id UUID PRIMARY KEY,
    email_address TEXT UNIQUE NOT NULL,
    forenames TEXT,
    surnames TEXT,
    mobile_telephone_address TEXT
)

CREATE TABLE team (
    id UUID PRIMARY KEY,
    display_name TEXT UNIQUE NOT NULL
)

CREATE TABLE team_member (
    id UUID PRIMARY KEY REFERENCES player,
    team_id UUID NOT NULL REFERENCES team
)

CREATE TABLE leader_supporter (
    id UUID PRIMARY KEY REFERENCES player,
    candidate_id UUID NOT NULL REFERENCES team_member
)

CREATE VIEW leader_supporter_count AS (
    SELECT
    ls.candidate_id,
    COUNT(ls.id) AS supporter_count
    FROM leader_supporter ls
    GROUP BY ls.candidate_id
);

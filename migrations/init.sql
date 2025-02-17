CREATE TABLE orders (
    id Uuid not null,
    created_at Datetime not null,
    deleted_at Datetime null,
    PRIMARY KEY(id)
)
CREATE TABLE referrals (
    id Uuid not null,
    created_at Datetime not null,
    deleted_at Datetime null,
    PRIMARY KEY(id)
)
CREATE TABLE referral_tests (
    referral_id Uuid not null,
    test_id int not null,
    PRIMARY KEY(id, test_id)
)
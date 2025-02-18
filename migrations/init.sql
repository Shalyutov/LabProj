CREATE TABLE orders (
    id Uuid not null,
    created_at Datetime not null,
    deleted_at Datetime null,
    PRIMARY KEY(id)
);
CREATE TABLE referrals (
    id Uuid not null,
    patient_id Uuid null,
    order_id Uuid null,
    issued_at Datetime not null,
    deleted_at Datetime null,
    send_at Datetime null,
    height float null,
    weight float null,
    tick_bite bool null,
    hiv_status int null,
    pregnancy_week int null,
    PRIMARY KEY(id)
);
CREATE TABLE referral_tests (
    referral_id Uuid not null,
    test_id int not null,
    PRIMARY KEY(referral_id, test_id)
);
CREATE TABLE samples (
    id Uuid not null,
    referral_id Uuid not null,
    issued_at Datetime not null,
    is_valid bool null,
    case_id int not null,
    PRIMARY KEY(id)
);
CREATE TABLE patients (
     id Uuid not null,
     surname Utf8 null,
     name Utf8 null,
     lastname Utf8 null,
     gender Utf8 null,
     email Utf8 null,
     representative Utf8 null,
     document Uint64 null,
     phone Uint64 null,
     birth_year int null,
     birth_month int null,
     birth_day int null,
     PRIMARY KEY(id)
);
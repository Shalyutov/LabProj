-- noinspection SqlNoDataSourceInspectionForFile
--CORE
--preanalytic
CREATE TABLE orders (
    id Uuid not null,
    created_at Datetime not null,
    paid_at Datetime null,
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
    accepted_at Datetime null,
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
     birth_date Date32 null,
     PRIMARY KEY(id)
);
--laboratory
CREATE TABLE work_units (
    id Uuid not null,
    referral_id Uuid not null,
    test_id int not NULL,
    equipment_id int not null,
    queued_at Datetime null,
    processed_at Datetime null,
    unit_result int null,
    PRIMARY KEY(id)
);
CREATE TABLE results (
    unit_id Uuid not null,
    indicator_id int not null,
    string_value Utf8 null,
    binary_value bool null,
    integer_value float null,
    issued_at Datetime null,
    confirmed_at Datetime null,
    is_valid bool null,
    PRIMARY KEY(unit_id, indicator_id)
);
--IAM
CREATE TABLE users (
    username Utf8 not null,
    full_name Utf8 null,
    is_blocked bool null,
    PRIMARY KEY (username)
);
CREATE TABLE user_scopes (
    username Utf8 not null,
    scope Utf8 not null,
    PRIMARY KEY (username, scope)
);
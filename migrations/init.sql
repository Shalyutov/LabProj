-- noinspection SqlNoDataSourceInspectionForFile
--CORE
--preanalytic
create table orders (
   id         uuid not null,
   created_at datetime not null,
   paid_at    datetime null,
   deleted_at datetime null,
   primary key ( id )
);
create table referrals (
   id             uuid not null,
   patient_id     uuid null,
   order_id       uuid null,
   issued_at      datetime not null,
   deleted_at     datetime null,
   send_at        datetime null,
   accepted_at    datetime null,
   height         float null,
   weight         float null,
   tick_bite      bool null,
   hiv_status     int null,
   pregnancy_week int null,
   primary key ( id )
);
create table referral_tests (
   referral_id uuid not null,
   test_id     int not null,
   primary key ( referral_id,
                 test_id )
);
create table samples (
   id          uuid not null,
   referral_id uuid not null,
   issued_at   datetime not null,
   is_valid    bool null,
   case_id     int not null,
   primary key ( id )
);
create table patients (
   id             uuid not null,
   surname        utf8 null,
   name           utf8 null,
   lastname       utf8 null,
   gender         utf8 null,
   email          utf8 null,
   representative utf8 null,
   document       uint64 null,
   phone          uint64 null,
   birth_date     date32 null,
   primary key ( id )
);
--laboratory
create table work_units (
   id           uuid not null,
   referral_id  uuid not null,
   test_id      int not null,
   equipment_id int not null,
   queued_at    datetime null,
   processed_at datetime null,
   unit_result  int null,
   primary key ( id )
);
create table results (
   unit_id       uuid not null,
   indicator_id  int not null,
   string_value  utf8 null,
   binary_value  bool null,
   integer_value float null,
   issued_at     datetime null,
   confirmed_at  datetime null,
   is_valid      bool null,
   primary key ( unit_id,
                 indicator_id )
);
--IAM
create table users (
   username   utf8 not null,
   full_name  utf8 null,
   is_blocked bool null,
   primary key ( username )
);
create table user_scopes (
   username utf8 not null,
   scope    utf8 not null,
   primary key ( username,
                 scope )
);
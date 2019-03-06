select * from reported_records
select count(distinct id) as active from reported_records
select count(distinct id) as deleted from reported_records where deleted_at is not null

select * from reported_records where extract(MONTH from date) = '02' AND extract(YEAR from date) = '2019' AND consultant = 'Stanislav Valasek'
select count(*) from reported_records where extract(MONTH from date) = '03' AND extract(YEAR from date) = '2019' AND consultant = 'Stanislav Valasek'
select count(*) from reported_records where date(date) in ('2019-02-25', '2019-02-26', '2019-02-27', '2019-02-28')
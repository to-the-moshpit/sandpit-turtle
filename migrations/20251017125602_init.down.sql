alter table event_artist drop constraint fk_event_artist_artist;
alter table event_artist drop constraint fk_event_artist_event;

drop table if exists event_artist;
drop table if exists artist;

alter table event drop constraint fk_event_venue;
alter table event drop constraint fk_event_hall;

drop table if exists event;

alter table hall drop constraint fk_hall_venue;

drop table if exists hall;

drop table if exists venue;
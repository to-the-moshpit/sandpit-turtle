begin;

  create table venue
  (
    id uuid primary key default gen_random_uuid(),
    created timestamp default current_timestamp not null,
    modified timestamp default current_timestamp not null,
    "name" varchar not null,
    address varchar not null
  );

  create table hall
  (
    id uuid primary key default gen_random_uuid(),
    created timestamp default current_timestamp not null,
    modified timestamp default current_timestamp not null,
    "name" varchar not null,
    venue_id uuid not null,
    constraint fk_hall_venue
      foreign key(venue_id) references venue(id)
      on delete cascade
  );

  create table event
  (
    id uuid primary key default gen_random_uuid(),
    created timestamp default current_timestamp not null,
    modified timestamp default current_timestamp not null,
    "start" timestamp not null,
    "end" timestamp not null,
    "name" varchar not null,
    "url" varchar not null,
    venue_id uuid not null,
    hall_id uuid not null,
    constraint fk_event_venue
      foreign key(venue_id) references venue(id)
      on delete cascade,
    constraint fk_event_hall
      foreign key(hall_id) references hall(id)
      on delete cascade
  );

  create table artist
  (
    id uuid primary key default gen_random_uuid(),
    created timestamp default current_timestamp not null,
    modified timestamp default current_timestamp not null,
    "name" varchar not null
  );

  create table event_artist
  (
    id uuid primary key default gen_random_uuid(),
    created timestamp default current_timestamp not null,
    modified timestamp default current_timestamp not null,
    artist_id uuid not null,
    event_id uuid not null,
    lined_up int not null,
    constraint fk_event_artist_artist
      foreign key(artist_id) references artist(id)
      on delete cascade,
    constraint fk_event_artist_event
      foreign key(event_id) references event(id)
      on delete cascade
  );

commit;
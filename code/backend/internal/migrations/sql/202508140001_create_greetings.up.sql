create table if not exists greetings (
  id smallint primary key,
  text text not null check (length(text) > 0),
  updated_at timestamptz not null default now(),
  constraint greetings_single_row check (id = 1)
);

insert into greetings (id, text) values (1, 'Hello Word')
on conflict (id) do nothing;

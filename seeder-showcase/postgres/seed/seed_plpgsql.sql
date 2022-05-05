do language plpgsql
$$
    declare
        person_1_id uuid = gen_random_uuid();
        person_2_id uuid = gen_random_uuid();
        person_3_id uuid = gen_random_uuid();

    begin
        insert into app.person (id, name, surname)
        values (person_1_id, 'Frederick', 'Simpson'),
               (person_2_id, 'James', 'Anderson'),
               (person_3_id, 'Timothy', 'Pennington')
        on conflict (id) do nothing;
    end;
$$;

do language plpgsql
$$
    declare
        person_4_id uuid = gen_random_uuid();
        person_5_id uuid = gen_random_uuid();
        person_6_id uuid = gen_random_uuid();

    begin
        insert into app.person (id, name, surname)
        values (person_4_id, 'Dorothy', 'Martinez'),
               (person_5_id, 'Pamela', 'Lawson'),
               (person_6_id, 'Barbara', 'Roth')
        on conflict (id) do nothing;
    end;
$$;

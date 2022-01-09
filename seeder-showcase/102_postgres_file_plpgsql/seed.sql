do language plpgsql
$$
    declare
        person_1_id uuid = gen_random_uuid();
        person_2_id uuid = gen_random_uuid();
        person_3_id uuid = gen_random_uuid();

    begin
        insert into person (id, name, surname)
        values
            (person_1_id, 'John', 'Doe'),
            (person_2_id, 'Kelvin', 'Houston'),
            (person_3_id, 'Brett', 'Vaught');
    end;
$$;

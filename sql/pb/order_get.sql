CREATE OR REPLACE FUNCTION order_get(_uidr character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_o orders;
	_response JSONB;
BEGIN
	SELECT *
	FROM orders
	WHERE uid_order = _uidr
	INTO _o;

	IF _o.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid не существует';
	END IF;

	SELECT
		COALESCE(og.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT o.id, o.uid_order, o.created, o.status, od.bin_brand, od.bin_type, od.bin_bank, od.bin_country, od.name, od.mobile, od.address, od.card_number, od.card_holder_name, od.exp_month, od.exp_year, od.security_code, od.user_agent, od.ip_address, od.current_url, od.lang, od.operating_system, od.browser
			FROM orders o, order_data od
			WHERE o.uid_order = _uidr
			AND od.uid_order = _uidr
		) ag
	) og
	INTO _response;

	RETURN _response;
END;
$function$
CREATE OR REPLACE FUNCTION room_leave(_uidr character varying, _uidu character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
	_u users;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = _uidr
	INTO _r;

	SELECT *
	FROM users
	WHERE uid = _uidu
	INTO _u;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid не существует';
	END IF;

	IF _r.member_count > 1 THEN
		UPDATE rooms SET member_count = member_count-1 WHERE uid_room = _uidr;
	ELSE
		UPDATE rooms SET member_count=0, is_works=FALSE WHERE uid_room = _uidr;
	END IF;

	IF _u.id ISNULL THEN
		UPDATE orders SET status = 'Chat closed' WHERE uid_order = _uidr;
	END IF;
END;
$function$
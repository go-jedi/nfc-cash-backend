CREATE OR REPLACE FUNCTION room_join(_uidr character varying, _uidu character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = _uidr
	INTO _r;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid не существует';
	END IF;

	UPDATE rooms SET member_count = member_count+1, is_works = TRUE WHERE uid_room = _uidr;
	INSERT INTO users_room(uid_user, entry_room) VALUES(_uidu, _uidr);
END;
$function$
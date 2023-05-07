CREATE OR REPLACE FUNCTION refresh_token_add(_id INTEGER, _rtkn character varying, _expt TIMESTAMP)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_rt refresh_tokens;
BEGIN
	SELECT *
	FROM refresh_tokens
	WHERE user_id = _id
	INTO _rt;

	IF _rt.id ISNULL THEN
		INSERT INTO refresh_tokens(user_id, token, expires_at) VALUES(_id, _rtkn, _expt);
	ELSE
		UPDATE refresh_tokens SET token = _rtkn, expires_at = _expt WHERE user_id = _id;
	END IF;
END;
$function$ 
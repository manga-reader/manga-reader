# Redis data structure

`<user_id>:favorite`: `{<comic_id1>, <comic_id2>....}`:
ranked by the latest updated one to the least updated one

`<user_id>:<website>:<comic_id>`: `<vol>:<page>`

`<comic_id>`: `{"id": <comic_id>, "name": <comic_name>, "latest_volume": <latest_volume>, "updated_at": <timestamp>}`

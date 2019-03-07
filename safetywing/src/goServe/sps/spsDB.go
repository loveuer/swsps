package sps

func getUserLastSearchHistory(uid string) ([]string, error) {
	var lastHiss []string
	rows, err := conn.Query("select content from sw_search_his where uid=$1 order by time desc limit 5", uid)
	if err != nil {
		return lastHiss, err
	}
	for rows.Next() {
		var onehis string
		rows.Scan(&onehis)
		lastHiss = append(lastHiss, onehis)
	}
	return lastHiss, nil
}

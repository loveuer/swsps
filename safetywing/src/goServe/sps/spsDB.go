package sps

import "log"

func getSpByID(spid string) (spsModel, error) {
	var sp spsModel
	row := conn.QueryRow("select * from swsps where id=$1", spid)
	err = row.Scan(&sp.id, &sp.name, &sp.pn, &sp.sn, &sp.oldSim, &sp.newSim, &sp.position, &sp.status, &sp.amount, &sp.durables, &sp.comment, &sp.img1, &sp.img2, &sp.wanted)
	if err != nil {
		log.Println(" ** scan one sp error")
		return sp, err
	}

	rows, err := conn.Query("select tag from swsps_tag where spid=$1", spid)
	if err != nil {
		log.Println("  * query sps tags error")
		return sp, err
	}
	for rows.Next() {
		var tag string
		rows.Scan(&tag)
		sp.tags = append(sp.tags, tag)
	}
	return sp, nil
}

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

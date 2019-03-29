package sps

import (
	"fmt"
	"log"
)

// 获取单个 备件 的具体信息
func getSpByID(spid string) (SpsModel, error) {
	var sp SpsModel
	row := conn.QueryRow("select id,name,pn,sn,orgsim,nowsim,pos,status,amount,sp_limit,is_consumable,comment,img1,img2 from sw_sps where id=$1", spid)
	err = row.Scan(&sp.ID, &sp.Name, &sp.PN, &sp.SN, &sp.OrgSim, &sp.NowSim, &sp.Position, &sp.Status, &sp.Amount, &sp.Limit, &sp.IsConsumable, &sp.Comment, &sp.Img1, &sp.Img2)
	if err != nil {
		log.Println(" ** scan one sp error")
		log.Println(" ** " + err.Error())
		return sp, err
	}

	return sp, nil
}

// 用户 搜索历史 (not done)
func getUserLastSearchHistory(uid string) ([]string, error) {
	var lastHiss []string
	rows, err := conn.Query("select content from sw_search_his where uid=$1 order by time desc limit 10", uid)
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

// 搜索 备件
func getSpsBySearchKey(key string) (*[]SpsModel, error) {

	var spsList []SpsModel

	sentence := fmt.Sprintf("select id,name,pn,sn,orgsim,nowsim,pos,status,amount,sp_limit,is_consumable,comment,img1,img2 from sw_sps where name ilike '%%%s%%' or pn ilike '%%%s%%' or sn ilike '%%%s%%' or comment ilike '%%%s%%'", key, key, key, key)

	rows, err := conn.Query(sentence)
	if err != nil {
		log.Println(" ** get sps by search key(name) error")
		log.Println(err)
	} else {
		for rows.Next() {
			var onesps SpsModel
			rows.Scan(&onesps.ID, &onesps.Name, &onesps.PN, &onesps.SN, &onesps.OrgSim, &onesps.NowSim, &onesps.Position, &onesps.Status, &onesps.Amount, &onesps.Limit, &onesps.IsConsumable, &onesps.Comment, &onesps.Img1, &onesps.Img2)
			spsList = append(spsList, onesps)
		}
	}

	return &spsList, nil
}

// 历史界面获取历史,不分哪个sp, limit 20
func getHisbyStart(start string) *[]SpHistory {
	var hisval []SpHistory
	sentence := fmt.Sprintf("select spid,time,last_sim,next_sim,last_pos,next_pos,last_status,next_status,amount_chg,auth,comment,short from sw_sphis order by id desc offset %s limit 20", start)

	rows, err := conn.Query(sentence)
	if err != nil {
		log.Printf(" ** sp db operate err(getHisbyStart): %v\n", err)
		return &hisval
	}

	for rows.Next() {
		var onehis SpHistory
		rows.Scan(&onehis.Spid, &onehis.Time, &onehis.LastSim, &onehis.NextSim, &onehis.LastPos, &onehis.NextPos, &onehis.LastStatus, &onehis.NextStatus, &onehis.AmountChg, &onehis.Auth, &onehis.Comment, &onehis.Short)
		onehis.SD = onehis.Time.Format("2006-01-02")
		onehis.ST = onehis.Time.Format("15:04:05")
		hisval = append(hisval, onehis)
	}

	return &hisval
}

// sp one 界面获取历史, 用spid获取, limit 20
func getOneSpsHis(spid, start string) *[]SpHistory {
	var hisval []SpHistory
	sentence := fmt.Sprintf("select spid,time,last_sim,next_sim,last_pos,next_pos,last_status,next_status,amount_chg,auth,comment,short from sw_sphis where spid=%s order by id desc offset %s limit 20", spid, start)
	rows, err := conn.Query(sentence)
	if err != nil {
		log.Printf(" ** sp db operate err(getOneSpsHis): %v\n", err)
		return &hisval
	}

	for rows.Next() {
		var onehis SpHistory
		rows.Scan(&onehis.Spid, &onehis.Time, &onehis.LastSim, &onehis.NextSim, &onehis.LastPos, &onehis.NextPos, &onehis.LastStatus, &onehis.NextStatus, &onehis.AmountChg, &onehis.Auth, &onehis.Comment, &onehis.Short)
		onehis.SD = onehis.Time.Format("2006-01-02")
		onehis.ST = onehis.Time.Format("15:04:05")
		hisval = append(hisval, onehis)
	}

	return &hisval
}

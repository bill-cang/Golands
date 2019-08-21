package main

import (
	"fmt"
	"strconv"
)

func main() {

	mmp := make(map[string]string, 0)
	mmp["00"] = "21"
	mmp["01"] = "18"
	mmp["02"] = "21"
	mmp["03"] = "19"
	mmp["04"] = "16"
	mmp["05"] = "17"
	mmp["06"] = "16"
	mmp["07"] = "21"
	mmp["08"] = "20"
	mmp["09"] = "25"
	mmp["10"] = "14"
	mmp["11"] = "13"
	mmp["12"] = "11"
	mmp["13"] = "10"
	mmp["14"] = "14"
	mmp["15"] = "19"
	mmp["16"] = "12"
	mmp["17"] = "22"
	mmp["18"] = "29"
	mmp["19"] = "18"
	mmp["20"] = "16"
	mmp["21"] = "18"
	mmp["22"] = "16"
	mmp["23"] = "15"
	mmp["24"] = "19"
	mmp["25"] = "18"
	mmp["26"] = "17"
	mmp["27"] = "21"
	mmp["28"] = "23"
	mmp["29"] = "16"
	mmp["30"] = "20"
	mmp["31"] = "23"
	mmp["32"] = "22"
	mmp["33"] = "15"
	mmp["34"] = "23"
	mmp["35"] = "18"
	mmp["36"] = "13"
	mmp["37"] = "17"
	mmp["38"] = "18"
	mmp["39"] = "16"
	mmp["40"] = "19"
	mmp["41"] = "16"
	mmp["42"] = "12"
	mmp["43"] = "19"
	mmp["44"] = "20"
	mmp["45"] = "17"
	mmp["46"] = "16"
	mmp["47"] = "21"
	mmp["48"] = "18"
	mmp["49"] = "15"
	mmp["50"] = "18"
	mmp["51"] = "20"
	mmp["52"] = "20"
	mmp["53"] = "19"
	mmp["54"] = "22"
	mmp["55"] = "21"
	mmp["56"] = "18"
	mmp["57"] = "19"
	mmp["58"] = "19"
	mmp["59"] = "17"
	mmp["60"] = "16"
	mmp["61"] = "16"
	mmp["62"] = "19"
	mmp["63"] = "16"
	mmp["64"] = "19"
	mmp["65"] = "23"
	mmp["66"] = "22"
	mmp["67"] = "21"
	mmp["68"] = "17"
	mmp["69"] = "19"
	mmp["70"] = "18"
	mmp["71"] = "21"
	mmp["72"] = "12"
	mmp["73"] = "15"
	mmp["74"] = "16"
	mmp["75"] = "20"
	mmp["76"] = "17"
	mmp["77"] = "21"
	mmp["78"] = "19"
	mmp["79"] = "19"
	mmp["80"] = "18"
	mmp["81"] = "25"
	mmp["82"] = "17"
	mmp["83"] = "15"
	mmp["84"] = "21"
	mmp["85"] = "21"
	mmp["86"] = "24"
	mmp["87"] = "17"
	mmp["88"] = "15"
	mmp["89"] = "17"
	mmp["90"] = "22"
	mmp["91"] = "18"
	mmp["92"] = "22"
	mmp["93"] = "17"
	mmp["94"] = "20"
	mmp["95"] = "21"
	mmp["96"] = "22"
	mmp["97"] = "24"
	mmp["98"] = "18"
	mmp["99"] = "17"

	resMap := make(map[string]float64)
	su, _ := strconv.ParseFloat("1843", 64)
	for k, v := range mmp {
		f, _ := strconv.ParseFloat(v, 64)
		resMap[k] = (f / su) * 100
	}

	var key = "00"
	for i := 0; i < 100; i++ {
		istr := strconv.Itoa(i)
		if i < 10 {
			key = "0" + istr
		} else {
			key = istr
		}
		f := resMap[key]
		fmt.Println(key, "：", f)
	}

}

/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111689	,600058099	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111697	,600058094	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111704	,600058091	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111706	,600058090	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111713	,600058089	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111722	,600058088	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111728	,600058087	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111732	,600058086	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111738	,600058085	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111744	,600058083	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111748	,600058065	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111757	,600058064	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111768	,600058063	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111777	,600058041	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111784	,600057927	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111789	,600057921	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111792	,600057920	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111798	,600056344	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111811	,600055676	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111821	,600055670	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111824	,600055637	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111831	,600010721	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111840	,34125		,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111844	,300023936	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111852	,300023064	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111856	,300022974	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111867	,300017728	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111873	,300017276	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111878	,300010567	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111884	,20686		,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111890	,900052346	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111895	,900052345	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111900	,900052327	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111911	,900052326	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111918	,900052325	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111923	,900052191	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111934	,900052190	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111938	,900052189	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111949	,900052188	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111955	,900052187	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111960	,900052185	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111965	,900052184	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111970	,900052133	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111980	,900050597	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111985	,900049942	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (111990	,900049934	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112004	,900049933	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112011	,900049906	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112016	,600058138	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112018	,600058137	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112025	,600058136	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112034	,600058135	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112038	,600058134	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112042	,600058133	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112050	,600058132	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112057	,600058131	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112062	,600058130	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112069	,600058124	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112081	,600058122	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112086	,600058116	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112096	,900052402	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112100	,900052401	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112103	,900052400	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112110	,900052399	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112112	,900052397	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112118	,900052396	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112127	,900052395	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112131	,900052394	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112137	,900052393	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112141	,900052387	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112146	,900052385	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112155	,900052384	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112164	,900052383	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112173	,900052378	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112183	,900052362	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112186	,900052359	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112191	,900052357	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112198	,900052354	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112202	,900052352	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112206	,900052351	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112218	,900052350	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112230	,900052349	,0	,'pxg'	,1566151200	,0	,1)*/
/*INSERT INTO ixo_tickets ( code, uid, is_lock, token, create_at, type, flag) VALUES (112232	,900052348	,0	,'pxg'	,1566151200	,0	,1)*/

26798
300018079
300023176
32280
600056221
900049422
900049425
900049426
900049827
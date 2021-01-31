// Package functions implements functions for returning URL and BODY content as struced data
package _functions

import (
		
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	
	//"encoding/base64"
	
	//"image"

	"encoding/json"
	_struct "fbrest/FBxRESTCore/struct"		
	_sessions "fbrest/FBxRESTCore/sessions"
	_httpstuff "fbrest/FBxRESTCore/httpstuff"
	_dbscheme "fbrest/FBxRESTCore/dbscheme"
	//"os"
	"net/http"
	//"net/url"
	"html/template"	
	"path"
	"fbrest/FBxRESTCore/config"
	_ "image/png"
)

//Returns respone for HTML site of FBRest usage
func ResponseHelpHTML(w http.ResponseWriter, code int) {
	
	const funcstr = "func ResponseHelpHTML"
    log.Debug(funcstr)
	profile := _struct.Profile{Appname:config.AppName,  Version:config.Version,Copyright:config.Copyright, Key:"-MNhE7Yf50sz6U9Hgqae", Duration:_sessions.MaxDuration}
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	if err := tmpl.Execute(w, profile); err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ResponseHelpDesignHTML(w http.ResponseWriter, code int) {
	
	/*
	reader, err := os.Open("templates/selfhtml.png")
	if err != nil {
	     log.Fatal(err)
	}
	m, _, err := image.Decode(reader)
	defer reader.Close()
*/
    const funcstr = "func ResponseHelpDesignHTML"
    log.Debug(funcstr)
	profile := _struct.Profile{Appname:config.AppName,  Version:config.Version,Copyright:config.Copyright, Key:"-MNhE7Yf50sz6U9Hgqae", Duration:_sessions.MaxDuration}
	fp := path.Join("templates", "design.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	if err := tmpl.Execute(w, profile); err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func ResponseHelpCommandsHTML(w http.ResponseWriter, code int) {
	
	profile := _struct.Profile{Appname:config.AppName,  Version:config.Version,Copyright:config.Copyright, Key:"-MNhE7Yf50sz6U9Hgqae", Duration:_sessions.MaxDuration}
	fp := path.Join("templates", "commands.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	if err := tmpl.Execute(w, profile); err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ResponseInfoBusyText(w http.ResponseWriter, code int) {
	
	profile := _struct.Profile{Appname:config.AppName,  Version:config.Version,Copyright:config.Copyright, Key:"-MNhE7Yf50sz6U9Hgqae", Duration:_sessions.MaxDuration}
	fp := path.Join("templates", "busy.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	if err := tmpl.Execute(w, profile); err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func OutParameters(entitiesData _struct.SQLAttributes) {
	const funcstr = "func OutParameters"
	log.Debug(funcstr)
	
	log.WithFields(log.Fields{"Given command     ": entitiesData.Cmd,}).Info("")
	log.WithFields(log.Fields{"Given sepfill     ": entitiesData.Sepfill,}).Info("")
	log.WithFields(log.Fields{"Given info        ": entitiesData.Info,}).Info("")
}

func OutTableParameters(entitiesData _struct.GetTABLEAttributes) {
	const funcstr = "func OutTableParameters"
	log.Debug(funcstr)
	
	log.WithFields(log.Fields{"Given fields      ": entitiesData.Fields,}).Info("")
	log.WithFields(log.Fields{"Given order by    ": entitiesData.OrderBy,}).Info("")
	log.WithFields(log.Fields{"Given group by    ": entitiesData.GroupBy,}).Info("")
	log.WithFields(log.Fields{"Given sepfill     ": entitiesData.Sepfill,}).Info("")
	log.WithFields(log.Fields{"Given info        ": entitiesData.Info,}).Info("")	
}

func GetSQLParamsFromBODY(r *http.Request , entitiesData *_struct.SQLAttributes) (bool) {
	const funcstr = "func GetSQLParamsFromBODY"
	log.Debug(funcstr) 
	var xdata _struct.GetUrlSQLAttributes
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error(funcstr)
		return false
	}
	entitiesData.Cmd = ReplaceAritmetics(xdata.Cmd)	
	entitiesData.Info = xdata.Info
	return true
}

func GetTableParamsFromBODY(r *http.Request , entitiesData *_struct.GetTABLEAttributes) (ok bool) {
	const funcstr = "func GetTableParamsFromBODY"
	log.Debug(funcstr) 
	var xdata _struct.GetUrlTABLEAttributes
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error(funcstr)
		return false
	}
	entitiesData.Fields = strings.Join(xdata.Fields,",")
	entitiesData.Filter = xdata.Filter
	entitiesData.GroupBy = strings.Join(xdata.GroupBy,",")
	entitiesData.OrderBy = strings.Join(xdata.OrderBy,",")
	entitiesData.Skip = xdata.Skip
	entitiesData.First = xdata.First	
	return true
}

func GetFieldParamsFromBODY(r *http.Request , entitiesData *_struct.GetTABLEAttributes) (ok bool) {
	const funcstr = "func GetFieldParamsFromBODY"
	log.Debug(funcstr) 
	var xdata _struct.GetUrlTABLEAttributes
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error(funcstr)
		return false
	}
	entitiesData.Fields = strings.Join(xdata.Fields,",")
	entitiesData.Filter = xdata.Filter
	entitiesData.GroupBy = strings.Join(xdata.GroupBy,",")
	entitiesData.OrderBy = strings.Join(xdata.OrderBy,",")
	entitiesData.Skip = xdata.Skip
	entitiesData.First = xdata.First	
	return true
}

func GetSQLParamsFromURL(r *http.Request , entitiesData *_struct.SQLAttributes) (ok bool) {
	const funcstr = "func GetSQLParamsFromURL"
	var u = r.URL
	
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)
			
		if(len(par) > 0) {
			xdata := &_struct.GetUrlSQLAttributes{}

			err := json.Unmarshal([]byte(par), &xdata)

			if(err != nil) {
				return false
			}

			log.Debug(xdata)
			entitiesData.Cmd = xdata.Cmd
			entitiesData.Info = xdata.Info			
			return true
		}
	}
	var okret bool
	urlparams, ok := r.URL.Query()["cmd"]
	infoparams, okinfo := r.URL.Query()["info"]
	log.WithFields(log.Fields{"URL params length": len(urlparams),	}).Debug(funcstr)

    if ok && len(urlparams[0]) > 0 {
		
		var cmd string =  string(urlparams[0])
		if(strings.HasPrefix(cmd,"'")||strings.HasPrefix(cmd,"-")) {
				
			entitiesData.Sepfill = cmd[:1]
            cmd = cmd[1:]
			cmd = strings.ReplaceAll(cmd, entitiesData.Sepfill, " ")	
			log.WithFields(log.Fields{"Sepfill": entitiesData.Sepfill,	}).Debug(funcstr+"->set key")				
		}
			    
		log.WithFields(log.Fields{"Cmd": cmd,	}).Debug(funcstr+"->set key")
		entitiesData.Cmd = cmd
		okret = true;
	}
	if okinfo && len(infoparams[0]) > 0 {
				
		var info string =  string(infoparams[0])				
		log.WithFields(log.Fields{"Info": info,	}).Debug(funcstr+"->set key")
		entitiesData.Info = info
		okret = true
	}

	return okret
}

func GetSessionParamsFromBODY(r *http.Request , entitiesData *_dbscheme.DatabaseAttributes) (bool) {
	const funcstr = "func GetSessionParamsFromBODY"
	log.Debug(funcstr) 
	var xdata _dbscheme.DatabaseAttributes
	
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error(funcstr)
		return false
	}
	
	entitiesData.Database = xdata.Database
	entitiesData.Location = xdata.Location
	entitiesData.Password = xdata.Password
	entitiesData.User     = xdata.User
	entitiesData.Port     = xdata.Port
	
	return true
}

func GetSessionSchemeParamsFromBODY(r *http.Request , entitiesData *_dbscheme.GetUrlSessionSchemeAttributes) (bool) {
	const funcstr = "func GetSessionSchemeParamsFromBODY"
	log.Debug(funcstr) 
	var xdata _dbscheme.GetUrlSessionSchemeAttributes
	
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error(funcstr)
		return false
	}
	
	entitiesData.Password = xdata.Password
	entitiesData.User     = xdata.User
	entitiesData.DBScheme = xdata.DBScheme
	
	return true
}

func GetSessionParamsFromURL(r *http.Request , entitiesData *_dbscheme.DatabaseAttributes) (bool) {
	
	const funcstr = "func GetSessionParamsFromURL"

	var u = r.URL
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)		
	//	 par = "{\"location\":\"localhost\",\"database\":\"D:/Data/Dokuments/DOKUMENTS30.FDB\",\"port\":3050,\"password\":\"su\",\"user\":\"superuser\"}"	
		if(len(par) > 0) {
			xdata := &_dbscheme.DatabaseAttributes{}

			err := json.Unmarshal([]byte(par), &xdata)

			if(err != nil) {
				return false
			}

			log.Info(xdata)
			entitiesData.Database = xdata.Database
			entitiesData.Port = xdata.Port
			entitiesData.Password = xdata.Password
			entitiesData.User = xdata.User		
			entitiesData.Location = xdata.Location	
			return true
		}
	}

	var okret bool
	databaseparams, databaseok := u.Query()["Database"]
	locationparams, locationok := u.Query()["Location"]
	portparams, portok := u.Query()["Port"]
	userparams, userok := u.Query()["User"]
	passwordparams, passwordok := u.Query()["Password"]
	
	if databaseok && len(databaseparams[0]) > 0 {
		log.WithFields(log.Fields{"URL params length": len(databaseparams),	}).Debug(funcstr)
		urlparam := databaseparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Database = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Database = urlparam
		}
		okret = true
	}

	if locationok && len(locationparams[0]) > 0 {
		log.WithFields(log.Fields{"URL location length": len(locationparams),	}).Debug(funcstr)
		urlparam := locationparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Location = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Location = urlparam
		}
		okret = true
	}

	if portok && len(portparams[0]) > 0 {
		log.WithFields(log.Fields{"URL port length": len(portparams),	}).Debug(funcstr)
		urlparam := portparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Port,_ = strconv.Atoi(urlparam[1:len(urlparam)-1])
		} else {
			entitiesData.Port,_ = strconv.Atoi(urlparam)
		}
		okret = true
	}

	if userok && len(userparams[0]) > 0 {
		log.WithFields(log.Fields{"URL user length": len(userparams),	}).Debug(funcstr)
		urlparam := userparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.User = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.User = urlparam
		}
		okret = true
	}

	if passwordok && len(passwordparams[0]) > 0 {
		log.WithFields(log.Fields{"URL password length": len(passwordparams),	}).Debug(funcstr)
		urlparam := passwordparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Password = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Password = urlparam
		}
		okret = true
	}

	return okret
}

func GetSessionSchemeParamsFromURL(r *http.Request , entitiesData *_dbscheme.GetUrlSessionSchemeAttributes) (bool) {
	
	const funcstr = "func GetSessionParamsFromURL"

	var u = r.URL
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)		
	//old	 par = "{\"location\":\"localhost\",\"database\":\"D:/Data/Dokuments/DOKUMENTS30.FDB\",\"port\":3050,\"password\":\"su\",\"user\":\"superuser\"}"	
	//old	 par = "{\"dbscheme\":\"health_ffm1\",\"password\":\"su\",\"user\":\"superuser\"}"	
		if(len(par) > 0) {
			xdata := &_dbscheme.GetUrlSessionSchemeAttributes{}

			err := json.Unmarshal([]byte(par), &xdata)

			if(err != nil) {
				return false
			}

			log.Info(xdata)
			entitiesData.Password = xdata.Password
			entitiesData.User     = xdata.User
			entitiesData.DBScheme = xdata.DBScheme		
			return true
		}
	}

	var okret bool
	dbschemeparams, dbschemeok := u.Query()["DBScheme"]
	
	userparams, userok := u.Query()["User"]
	passwordparams, passwordok := u.Query()["Password"]
	
	if dbschemeok && len(dbschemeparams[0]) > 0 {
		log.WithFields(log.Fields{"URL params length": len(dbschemeparams),	}).Debug(funcstr)
		urlparam := dbschemeparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.DBScheme = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.DBScheme = urlparam
		}
		okret = true
	}

	if userok && len(userparams[0]) > 0 {
		log.WithFields(log.Fields{"URL user length": len(userparams),	}).Debug(funcstr)
		urlparam := userparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.User = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.User = urlparam
		}
		okret = true
	}

	if passwordok && len(passwordparams[0]) > 0 {
		log.WithFields(log.Fields{"URL password length": len(passwordparams),	}).Debug(funcstr)
		urlparam := passwordparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Password = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Password = urlparam
		}
		okret = true
	}

	return okret
}

func GetFIELDPayloadFromString2(params string ,  entitiesData *_struct.FIELDVALUEAttributes) {
	
	// payload=(id:1, username: 'admin', email: 'email@example.org')

	const funcstr = "func GetFIELDPayloadFromString"
	var psplit  = "&"
	var csplit  = "="
	var par = strings.Split(params,psplit)

	log.WithFields(log.Fields{"URL params length": len(par),	}).Debug(funcstr)
	
    if len(par) > 0 {
		for _, pars := range par {
			params = _httpstuff.UnEscape(pars)
			keyval :=  strings.SplitN(params,csplit,2)
						
			log.WithFields(log.Fields{"Key": string(keyval[0]),	}).Debug(funcstr+"->found key")
			log.WithFields(log.Fields{"Val": string(keyval[1]),	}).Debug(funcstr+"->found value")
			
			if(strings.EqualFold(string(keyval[0]), string("FIELDS"))) {								
				log.WithFields(log.Fields{"Fields": string(keyval[1]),	}).Debug(funcstr+"->set Fields")
				
			}			
		}		
	} 
	return
}

func GetSQLParamsFromString2(params string ,entitiesData *_struct.SQLAttributes) {
	
	const funcstr = "func GetSQLParamsFromString"
	var csplit  = "="
	var par = params

	log.WithFields(log.Fields{"URL params length": len(par),	}).Debug(funcstr)
    params = _httpstuff.UnEscape(par)
	log.WithFields(log.Fields{"SQL": params,	}).Debug(funcstr+"->set key")

	
	keyval :=  strings.SplitN(params,csplit,2)
	
	log.WithFields(log.Fields{"Key": string(keyval[0]),	}).Debug(funcstr+"->found key")
	log.WithFields(log.Fields{"Val": string(keyval[1]),	}).Debug(funcstr+"->found value")
	
	if(strings.EqualFold(string(keyval[0]), string("CMD"))) {
		var cmd string =  string(keyval[1])
		if(strings.HasPrefix(cmd,"'")||strings.HasPrefix(cmd,"-")) {
		
			entitiesData.Sepfill = cmd[:1]
			cmd = cmd[1:]
			cmd = strings.ReplaceAll(cmd, entitiesData.Sepfill, " ")	
			log.WithFields(log.Fields{"Sepfill": entitiesData.Sepfill,	}).Debug(funcstr+"->set key")				
		}
		
		log.WithFields(log.Fields{"Cmd": cmd,	}).Debug(funcstr+"->set key")
		entitiesData.Cmd = cmd
	}
	
	if(strings.EqualFold(string(keyval[0]), string("INFO"))) {								
		log.WithFields(log.Fields{"Info": string(keyval[1]),	}).Debug(funcstr+"->set key")
		entitiesData.Info = string(keyval[1])
	}							
	return
}

//Returns the last-nLeft slice from URL
//e.g. when nLeft == 0 returns the last slice
func GetRightPathSliceFromURL(r *http.Request, nLeft int) (key string) {
	
	urlstr := string(r.URL.String())
	var keyval =  strings.SplitN(urlstr,"?",2)
	
	urlstr = keyval[0]
	t2 :=  strings.Split(urlstr,"/")
	key = t2[len(t2)-1-nLeft]
	log.WithFields(log.Fields{_sessions.SessionTokenStr: key,	}).Debug("func GetPathSliceFromURL")	
	return key
}

func GetLeftPathSliceFromURL(r *http.Request, nLeft int) (key string) {
	const funcstr = "func GetLeftPathSliceFromURL"
	urlstr := string(r.URL.String())
	var keyval =  strings.SplitN(urlstr,"?",2)
	
	urlstr = keyval[0]

	t2 :=  strings.Split(urlstr,"/")
	key = t2[nLeft+1]
	log.WithFields(log.Fields{_sessions.SessionTokenStr: key,	}).Debug(funcstr)	
	return key
}

func GetTableParamsFromURL(r *http.Request , entitiesData *_struct.GetTABLEAttributes) (bool) {
	
	//  http://localhost:4488/{{.Key}}/rest/get/TSTANDORT?fjson={"table": "TSTANDORT","fields": ["ID","BEZ","GUELTIG"],"filter":"ID=1 AND BEZ like 'x%'","order": ["BEZ ASC","ID DESC"],"groupby": ["ID","BEZ"],"first": 0}
	
	const funcstr = "func GetTableParamsFromURL"

	var u = r.URL
	
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)
			

		if(len(par) > 0) {
			xdata := &_struct.GetUrlTABLEAttributes{}

			err := json.Unmarshal([]byte(par), &xdata)

			if(err != nil) {
				return false
			}

			log.Info(xdata)
			entitiesData.Fields = strings.Join(xdata.Fields,",")
			entitiesData.Filter = xdata.Filter
			entitiesData.GroupBy = strings.Join(xdata.GroupBy,",")
			entitiesData.OrderBy = strings.Join(xdata.OrderBy,",")
			entitiesData.Skip = xdata.Skip
			entitiesData.First = xdata.First			
			return true
		}
	}
	
	var s = _httpstuff.UnEscape(u.RawQuery)

	var pars = strings.Split(s,"&")
	var okret bool
	for _, par := range pars {
		if(strings.HasPrefix(par,_struct.Fields+"=")){
			var par = par[len(_struct.Fields)+1:]
			if(par[:1] == "(") {
				entitiesData.Fields = par[1:len(par)-1]
			} else {
				entitiesData.Fields = par
			}
			if(len(entitiesData.Fields) < 1) {
				entitiesData.Fields = "*"
			}
			okret = true
		} else if(strings.HasPrefix(par,_struct.Filter+"=")){
			var par = par[len(_struct.Filter)+1:]
			if(par[:1] == "(") {
				entitiesData.Filter = par[1:len(par)-1]
			} else {
				entitiesData.Filter = par
			}
			entitiesData.Filter = ReplaceAritmetics(entitiesData.Filter)
			okret = true
		} else if(strings.HasPrefix(par,_struct.Order+"=")){
			var par = par[len(_struct.Order)+1:]
			if(par[:1] == "(") {
				entitiesData.OrderBy = par[1:len(par)-1]
			} else {
				entitiesData.OrderBy = par
			}
			okret = true
		} else if(strings.HasPrefix(par,_struct.Group+"=")){
			var par = par[len(_struct.Group)+1:]
			if(par[:1] == "(") {
				entitiesData.GroupBy = par[1:len(par)-1]
			} else {
				entitiesData.GroupBy = par
			}
			okret = true
		
		} else if(strings.HasPrefix(par,_struct.Info+"=")){
			var par = par[len(_struct.Info)+1:]
			if(par[:1] == "(") {
				entitiesData.Info = par[1:len(par)-1]
			} else {
				entitiesData.Info = par
			}
			okret = true
		
		} else if(strings.HasPrefix(par,_struct.Limit+"=")){
			var par = par[len(_struct.Limit)+1:]
			var limit string
			if(par[:1] == "(") {
				limit = par[1:len(par)-1]
			} else {
				limit = par
			}
			var lm = strings.Split(limit,",")
			if(len(lm) == 1) {
				entitiesData.First,_ = strconv.Atoi(lm[0])
				entitiesData.Skip = 0
			} else if (len(lm) == 2) {
				entitiesData.First,_ = strconv.Atoi(lm[0])
				entitiesData.Skip,_ = strconv.Atoi(lm[1])		
			}
			okret = true
		}
	}

/*
	u, _ = url.Parse(u.Path+"?"+s)

	log.Debug(s)
	var okret bool
	fieldparams, okfields := u.Query()[_struct.Fields]
	orderparams, okorder := u.Query()[_struct.Order]
	filterparams, okfilter := u.Query()[_struct.Filter]
	groupparams, okgroup := u.Query()[_struct.Group]
	infoparams, okinfo := u.Query()[_struct.Info]
	limitparams, oklimit := u.Query()[_struct.Limit]
	
	log.WithFields(log.Fields{"fieldparams": fieldparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"filterparams": filterparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"orderparams": orderparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"groupparams": groupparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"infoparams": infoparams,}).Debug(funcstr)

		
    if okfields && len(fieldparams[0]) > 0 {
		
		urlparam := fieldparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Fields = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Fields = urlparam
		}
		if(len(entitiesData.Fields) < 1) {
			entitiesData.Fields = "*"
		}
		okret = true
	} 

	if okfilter && len(filterparams[0]) > 0 {
		
		urlparam := filterparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Filter = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Filter = urlparam
		}
		entitiesData.Filter = ReplaceAritmetics(entitiesData.Filter)
		okret = true
	} 

	if okorder && len(orderparams[0]) > 0 {
		
		urlparam := orderparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.OrderBy = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.OrderBy = urlparam
		}
		okret = true
	} 

	if okgroup && len(groupparams[0]) > 0 {
		
		urlparam := groupparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.GroupBy = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.GroupBy = urlparam
		}
		okret = true
	}

	entitiesData.First = -1
	entitiesData.Skip  = -1
	if oklimit && len(limitparams[0]) > 0 {
		
		urlparam := limitparams[0]
		var limit string
		if(urlparam[:1] == "(") {
			limit = urlparam[1:len(urlparam)-1]
		} else {
			limit = urlparam
		}
		var lm = strings.Split(limit,",")
		if(len(lm) == 1) {
			entitiesData.First,_ = strconv.Atoi(lm[0])
			entitiesData.Skip = 0
		} else if (len(lm) == 2) {
			entitiesData.First,_ = strconv.Atoi(lm[0])
			entitiesData.Skip,_ = strconv.Atoi(lm[1])		
		}
		okret = true
	}

	if okinfo && len(infoparams[0]) > 0 {
		
		urlparam := infoparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Info = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Info = urlparam
		}
		okret = true
	}
    */
	return okret
}

func GetFieldParamsFromURL(r *http.Request , entitiesData *_struct.GetTABLEAttributes) (bool) {
	
	//  http://localhost:4488/{{.Key}}/rest/get/TSTANDORT?fjson={"table": "TSTANDORT","fields": ["ID","BEZ","GUELTIG"],"filter":"ID=1 AND BEZ like 'x%'","order": ["BEZ ASC","ID DESC"],"groupby": ["ID","BEZ"],"first": 0}
	
	const funcstr = "func GetFieldParamsFromURL"

	var u = r.URL
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)
			

		if(len(par) > 0) {
			xdata := &_struct.GetUrlTABLEAttributes{}

			err := json.Unmarshal([]byte(par), &xdata)

			if(err != nil) {
				return false
			}

			log.Info(xdata)
			entitiesData.Fields = strings.Join(xdata.Fields,",")
			entitiesData.Filter = xdata.Filter
			entitiesData.GroupBy = strings.Join(xdata.GroupBy,",")
			entitiesData.OrderBy = strings.Join(xdata.OrderBy,",")
			entitiesData.Skip = xdata.Skip
			entitiesData.First = xdata.First			
			return true
		}
	}

	var s = _httpstuff.UnEscape(u.RawQuery)

	var pars = strings.Split(s,"&")
	var okret bool
	for _, par := range pars {
		if(strings.HasPrefix(par,_struct.Filter+"=")){
			var par = par[len(_struct.Filter)+1:]
			if(par[:1] == "(") {
				entitiesData.Filter = par[1:len(par)-1]
			} else {
				entitiesData.Filter = par
			}
			entitiesData.Filter = ReplaceAritmetics(entitiesData.Filter)
			okret = true
		} else if(strings.HasPrefix(par,_struct.Order+"=")){
			var par = par[len(_struct.Order)+1:]
			if(par[:1] == "(") {
				entitiesData.OrderBy = par[1:len(par)-1]
			} else {
				entitiesData.OrderBy = par
			}
			okret = true
		} else if(strings.HasPrefix(par,_struct.Group+"=")){
			var par = par[len(_struct.Group)+1:]
			if(par[:1] == "(") {
				entitiesData.GroupBy = par[1:len(par)-1]
			} else {
				entitiesData.GroupBy = par
			}
			okret = true
		
		} else if(strings.HasPrefix(par,_struct.Info+"=")){
			var par = par[len(_struct.Info)+1:]
			if(par[:1] == "(") {
				entitiesData.Info = par[1:len(par)-1]
			} else {
				entitiesData.Info = par
			}
			okret = true
		
		} else if(strings.HasPrefix(par,_struct.Limit+"=")){
			var par = par[len(_struct.Limit)+1:]
			var limit string
			if(par[:1] == "(") {
				limit = par[1:len(par)-1]
			} else {
				limit = par
			}
			var lm = strings.Split(limit,",")
			if(len(lm) == 1) {
				entitiesData.First,_ = strconv.Atoi(lm[0])
				entitiesData.Skip = 0
			} else if (len(lm) == 2) {
				entitiesData.First,_ = strconv.Atoi(lm[0])
				entitiesData.Skip,_ = strconv.Atoi(lm[1])		
			}
			okret = true
		}
	}

	/*
	orderparams, okorder := u.Query()[_struct.Order]
	filterparams, okfilter := u.Query()[_struct.Filter]
	groupparams, okgroup := u.Query()[_struct.Group]
	infoparams, okinfo := u.Query()[_struct.Info]
	limitparams, oklimit := u.Query()[_struct.Limit]
	
	log.WithFields(log.Fields{"filterparams": filterparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"orderparams": orderparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"groupparams": groupparams,}).Debug(funcstr)
	log.WithFields(log.Fields{"infoparams": infoparams,}).Debug(funcstr)

	if okfilter && len(filterparams[0]) > 0 {
		urlparam := filterparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Filter = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Filter = urlparam
		}
		okret = true
	} 

	if okorder && len(orderparams[0]) > 0 {
		urlparam := orderparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.OrderBy = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.OrderBy = urlparam
		}
		okret = true
	} 

	if okgroup && len(groupparams[0]) > 0 {	
		urlparam := groupparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.GroupBy = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.GroupBy = urlparam
		}
		okret = true
	}

	entitiesData.First = -1
	entitiesData.Skip  = -1
	if oklimit && len(limitparams[0]) > 0 {
		
		urlparam := limitparams[0]
		var limit string
		if(urlparam[:1] == "(") {
			limit = urlparam[1:len(urlparam)-1]
		} else {
			limit = urlparam
		}
		var lm = strings.Split(limit,",")
		if(len(lm) == 1) {
			entitiesData.First,_ = strconv.Atoi(lm[0])
			entitiesData.Skip = 0
		} else if (len(lm) == 2) {
			entitiesData.First,_ = strconv.Atoi(lm[0])
			entitiesData.Skip,_ = strconv.Atoi(lm[1])		
		}
		okret = true
	}

	if okinfo && len(infoparams[0]) > 0 {
		
		urlparam := infoparams[0]
		if(urlparam[:1] == "(") {
			entitiesData.Info = urlparam[1:len(urlparam)-1]
		} else {
			entitiesData.Info = urlparam
		}
		okret = true
	}
    */
	return okret
}

func GetFIELDPayloadFromBODY(r *http.Request , entitiesData *_struct.FIELDVALUEAttributes) (bool) {
	log.Debug("func GetFIELDPayloadFromBODY") 
	var xdata _struct.GetUrlPayloadAttributes
	//body, err2 := ioutil.ReadAll(r.Body)
    //var str string = string(body)
	//log.Info(str)
	err2 := json.NewDecoder(r.Body).Decode(&xdata)
	if err2 != nil {			
		log.WithFields(log.Fields{"Decode params to JSON": err2.Error(),	}).Error("func GetFIELDPayloadFromBODY")
		return false
	}
	
	entitiesData.FieldValue = xdata.Payload
	xdata.Filter = ReplaceAritmetics(xdata.Filter)
	
	entitiesData.Filter = xdata.Filter
	return true
}

func GetFIELDPayloadFromURL(r *http.Request , entitiesData *_struct.FIELDVALUEAttributes) (bool){
	//  http://localhost:4488/{{.Key}}/rest/put/TSTANDORT?payload=(bez='Nürnberg2')&filter=(bez='Nürnberg1')  
	//  http://localhost:4488/{{.Key}}/rest/put/TSTANDORT?ftext="payload=(bez='Nürnberg2')&filter=(bez='Nürnberg1')"
	// http://localhost:4488/{{.Key}}/rest/put/TSTANDORT?fjson={"payload":["ID='123'","BEZ='test'","GUELTIG=1"], "filter": "ID=1 AND BEZ like 'x%'"}

	const funcstr = "func GetFIELDPayloadFromURL"
	var u = r.URL
	log.Debug(funcstr)
	if(strings.HasPrefix(u.RawQuery,_struct.FormatJson)) {			
		var par = u.RawQuery[len(_struct.FormatJson)+1:]
		par = _httpstuff.UnEscape(par)
		if(len(par) > 0) {
			xdata := &_struct.GetUrlPayloadAttributes{}
			err := json.Unmarshal([]byte(par), &xdata)
			if(err != nil) {
				return false
			}
			log.Info("xdata")
			log.Info(xdata)
			for _, vals := range xdata.Payload {
				entitiesData.FieldValue = append(entitiesData.FieldValue,vals)
			}
			entitiesData.Filter = xdata.Filter
			return true
		}
	}
	

	var s = _httpstuff.UnEscape(u.RawQuery)

	var pars = strings.Split(s,"&")
	var okret bool
	log.Debug(funcstr+"->pars")
	log.Debug(pars)
	for _, par := range pars {
		if(strings.HasPrefix(par,_struct.Filter+"=")){
			var par = par[len(_struct.Filter)+1:]
			if(par[:1] == "(") {
				entitiesData.Filter = par[1:len(par)-1]
			} else {
				entitiesData.Filter = par
			}
			entitiesData.Filter = ReplaceAritmetics(entitiesData.Filter)
			okret = true
		} else if(strings.HasPrefix(par,_struct.Payload+"=")){
			var pars = par[len(_struct.Payload)+1:]		
			var st string = pars[:1]
			log.Debug(funcstr+"->st:"+st)
			if(st == "(") {
				pars = pars[1:]
			}	
			log.Debug(funcstr+"->pars:"+pars)
			st = pars[len(pars)-1:]	
			log.Debug(funcstr+"->st:"+st)
			if(st == ")") {
				pars = pars[:len(pars)-1]
			}
			log.Debug(funcstr+"->pars:"+pars)
			//keyval :=  strings.SplitN(pars,",",-1)
			keyval :=  SplitPars(pars,",")
	
			for _, pars := range keyval {
				entitiesData.FieldValue = append(entitiesData.FieldValue,pars)	
			}
			
			okret = true
		
		}
	}

    /*
	payloadparams, okpayload := r.URL.Query()[_struct.Payload]
	filterparams, okfilter := r.URL.Query()[_struct.Filter]

	log.WithFields(log.Fields{"payloadparams": payloadparams,}).Debug(funcstr)
			
    if okpayload && len(payloadparams[0]) > 0 {
		var pars = payloadparams[0]
		log.Info(pars)
		
		var st = pars[:1]
		if(st == "(") {
			pars = pars[1:]
		}	

		st = pars[len(pars)-1:]	
		if(st == ")") {
			pars = pars[:len(pars)-1]
		}

		keyval :=  strings.SplitN(pars,",",2)

		for _, pars := range keyval {
			entitiesData.FieldValue = append(entitiesData.FieldValue,pars)	
		}
		ok = true
	} 

	if okfilter && len(filterparams[0]) > 0 {	
		entitiesData.Filter = filterparams[0]
		ok = true
	}
	*/
	return okret
}
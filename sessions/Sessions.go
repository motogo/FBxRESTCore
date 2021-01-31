package _sessions

import (
	log "github.com/sirupsen/logrus"
	_httpstuff "fbrest/FBxRESTCore/httpstuff"
	_permissions "fbrest/FBxRESTCore/permissions"
	"sync"
	"time"
	"strconv"
	"net/http"
	"encoding/xml"
	"io/ioutil"
	_struct "fbrest/FBxRESTCore/struct"		
	_apperrors "fbrest/FBxRESTCore/apperrors"
)

const MaxDuration = 30*60*1E9  //ns
const SessionTokenStr = "session token"

type Item struct {
	Token string  `json:"Token"` 
	Value string  `json:"Value"` 
	Permission _permissions.PermissionType `json:"Permission"`
	Start time.Time  `json:"Time"` 
	Duration time.Duration  `json:"Duration"` 
	Valid bool `json:"Valid"` 
}

type Sessions struct {
    Session []Item `xml:"Session"`
}

type repository struct {
	items map[string]Item
	mu    sync.RWMutex
}



func (r *repository) Add(token string, permission _permissions.PermissionType, conn string) (ky Item) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var data Item
	data.Token = token
	data.Value = conn
	data.Start = time.Now()
	data.Duration = MaxDuration
	data.Permission = permission
	r.items[token] = data
	log.WithFields(log.Fields{SessionTokenStr: token,	}).Debug("func Add "+SessionTokenStr)	
	return data
}

func (r *repository) AddStruct(data Item) (ky Item) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[data.Token] = data
	log.WithFields(log.Fields{SessionTokenStr: data.Token,	}).Debug("func Add "+SessionTokenStr)	
	return data
}

func (r *repository) Delete(token string) {
	r.mu.Lock()
	defer r.mu.Unlock()	
	delete(r.items, token)
	log.WithFields(log.Fields{SessionTokenStr: token,	}).Debug("func Delete "+SessionTokenStr)	
}

func (r *repository) Get(token string) (Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[token]
	if !ok {
		var err error = _apperrors.ErrPermissionItemNotFound
		log.WithFields(log.Fields{SessionTokenStr+" not found": token,	}).Debug("func Get "+SessionTokenStr)	
		return item,err
	}
	log.WithFields(log.Fields{SessionTokenStr+" found": token,	}).Debug("func Get "+SessionTokenStr)	
	return item, nil
}

func GetTokenDataFromRepository(token string) (kval Item) {
	const funcStr = "func Sessions.GetTokenDataFromRepository"
	log.WithFields(log.Fields{token: token,	}).Debug(funcStr)	
	var rep = Repository() 
	var result,_ = rep.Get(token)	 
	result.Token = token
	return result 
}

func TokenValid(response http.ResponseWriter, key string) (kv Item) {
	const funcStr = "func Sessions.TokenValid"
	log.WithFields(log.Fields{key: key,	}).Debug(funcStr)	

	var Response _struct.ResponseData
	kv  = GetTokenDataFromRepository(key)	
	if(len(kv.Value) < 1) {
		Response.Status = http.StatusForbidden
		Response.Message = "No valid database connection found by "+SessionTokenStr+" "+kv.Token
		Response.Data = _apperrors.DataNil
		_httpstuff.RestponWithJson(response, http.StatusInternalServerError, Response)
		kv.Valid = false
		return kv
	}
	var duration time.Duration = kv.Duration
	var end = time.Now();
	difference := end.Sub(kv.Start)
	if(difference > duration) {
		//Zeit für Key abgelaufen
		Response.Status  = http.StatusForbidden
		Response.Message = SessionTokenStr+" "+kv.Token + " has expired after "+ strconv.Itoa(MaxDuration/1E9) +" seconds"
		Response.Data    = _apperrors.DataNil
		log.WithFields(log.Fields{funcStr+"->Session expired:": Response.Message,	}).Debug("func KeyValid")
		var rep = Repository() 
	    rep.Delete(key)
		_httpstuff.RestponWithJson(response, http.StatusInternalServerError, Response)
		kv.Valid = false

		return kv
	}
	log.WithFields(log.Fields{funcStr+"->Session vaild->Remaining "+SessionTokenStr+" duration (s)": (duration-difference),	}).Debug("func KeyValid")
	kv.Valid = true
	return kv
}

func WritePermanantSessions(pfile string) {
	var dataarr Sessions
	var data Item

	data.Token = "ffm1health" 
	data.Value = "SYSDBA:masterkey@localhost:3050/D:/Data/LAR/HEALTHDATAS304.FDB"
	data.Permission = _permissions.All
	data.Start = time.Now()
	data.Duration  = 100000 
	data.Valid = true
	dataarr.Session = append(dataarr.Session,data)

	data.Token = "bln11health" 
	data.Value = "SYSDBA:masterkey@localhost:3050/D:/Data/LAR/HEALTHDATAS304.FDB"
	data.Permission = _permissions.All
	data.Start  = time.Now()
	data.Duration  = 100000 
	data.Valid = true
	dataarr.Session = append(dataarr.Session,data)
	file, _ :=  xml.MarshalIndent(dataarr, "", " ")
 
	_ = ioutil.WriteFile(pfile, file, 0644)
	
}

func ReadPermanantSessions(pfile string) {
	const funcstr = "func ReadPermanantSessions"
	log.Debug(funcstr);
	dataarr, err := ioutil.ReadFile(pfile)
    if err != nil {		
		log.WithFields(log.Fields{"File reading error": err,	}).Error(funcstr)	
        return
    }
	//data := &[]Item{}
	data := &Sessions{}
	xml.Unmarshal(dataarr,&data)
	var rep = Repository()
	for _, pars := range data.Session {
		pars.Duration = pars.Duration * 1E9
		log.Debug(" ")	
		log.WithFields(log.Fields{"Read Token     :": pars.Token,}).Debug(funcstr)	
		log.WithFields(log.Fields{"Read Connection:": pars.Value,}).Debug(funcstr)	
		log.WithFields(log.Fields{"Read Start     :": pars.Start,}).Debug(funcstr)	
		log.WithFields(log.Fields{"Read Duration  :": pars.Duration,}).Debug(funcstr)	
		
		rep.AddStruct(pars)
	}
}


	
var instance *repository

	
func Repository() *repository {
	if instance == nil {
		instance = &repository {
			items: make(map[string]Item),
		}
	}
	return instance
}




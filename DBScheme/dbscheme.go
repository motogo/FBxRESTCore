package _dbscheme

import (  
	"encoding/xml"
	"io/ioutil"
	"sync"
	"strconv"
	_apperrors "fbrest/FBxRESTCore/apperrors"
	
	log "github.com/sirupsen/logrus"
)

type DatabaseAttributes struct {
	Name string `json:"name"`
	Key string `json:"key"`
	Location string  `json:"location"`
	Database string  `json:"database"`
	Port int  `json:"port"`	
	Password string `json:"password"`
	User string `json:"user"`	
   }


func WriteDefaultDBScheme(pfile string) {
	var dataarr [] DatabaseAttributes
	var data DatabaseAttributes
	data.Name = "health bln11"
	data.Key = "health_bln11"
	data.Password = "masterkey"	
	data.User     = "SYSDBA"
	data.Location = "192.168.12.99"
	data.Database = "D:/Data/LAR/HEALTHDATAS304.FDB"
	data.Port = 3050

	dataarr = append(dataarr,data)

	data.Name = "health ffm1"
	data.Key  = "health_ffm1"
	data.Password = "masterkey"	
	data.User     = "SYSDBA"
	data.Location = "192.168.11.98"
	data.Database = "D:/Data/LAR/HEALTHDATAS304.FDB"
	data.Port = 3050

	dataarr = append(dataarr,data)

	file, _ := xml.MarshalIndent(dataarr, "", " ")
 
	_ = ioutil.WriteFile(pfile, file, 0644)
	
}

type GetUrlSessionSchemeAttributes struct {	
	User string `json:"user"` 
	Password string `json:"password"` 
	DBScheme string `json:"dbscheme"` 
   }

const DBSchemeKeyStr = "dbschema key"

type repository struct {
	datas map[string]DatabaseAttributes
	mu    sync.RWMutex
}


type DBSchemes struct {
    DatabaseAttributes []DatabaseAttributes `xml:"DatabaseAttributes"`
}

func (r *repository) Add(name string, key string,user string,password string, database string, location string, port int) (ky DatabaseAttributes) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var data DatabaseAttributes
	data.Key  = key
	data.Name  = name
	data.User  = user
	data.Password  = password	
	data.Database = database
	data.Location = location
	data.Port = port

	r.datas[key] = data
	log.WithFields(log.Fields{"Key": key,	}).Debug("func Add ")	
	return data
}

func (r *repository) Delete(token string) {
	r.mu.Lock()
	defer r.mu.Unlock()	
	delete(r.datas, token)
	log.WithFields(log.Fields{DBSchemeKeyStr: token,	}).Debug("func Delete "+DBSchemeKeyStr)	
}

func (r *repository) Get(token string) (item DatabaseAttributes, err error ) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.datas[token]
	if !ok {
		err = _apperrors.ErrPermissionKeyNotFound
		log.WithFields(log.Fields{DBSchemeKeyStr+" not found": token,	}).Debug("func Get "+DBSchemeKeyStr)	
		return item, err
	}
	log.WithFields(log.Fields{DBSchemeKeyStr+" found": token,	}).Debug("func Get "+DBSchemeKeyStr)	
	return item, nil
}


func (r *repository) Exists(token string) (bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.datas[token]
	if !ok {
		return false
	}
	return true
}

var (
	pr *repository
)

func Repository() *repository {
	if pr == nil {
		pr = &repository {
			datas: make(map[string]DatabaseAttributes),
		}
	}
	return pr
}

func GetSchemeFromRepository(dbscheme string) (perm DatabaseAttributes, err error ) {
	
	log.WithFields(log.Fields{DBSchemeKeyStr: dbscheme,	}).Debug("func GetSchemeFromRepository")	
	var rep = Repository() 
	var result,err1 = rep.Get(dbscheme)	 
	if(err1 != nil) {
		return result,err1
	}
	return result,err
}

func ReadDBScheme(pfile string) {
	const funcstr = "func ReadDBScheme"
	log.Debug("funcstr")
	log.Info(log.GetLevel())
	data, err := ioutil.ReadFile(pfile)
    if err != nil {		
		log.WithFields(log.Fields{"File reading error": err,	}).Error("funcstr")	
        return
    }
	xdata := &DBSchemes{}
	xml.Unmarshal(data,&xdata)
	var rep = Repository() 
	for _, xd := range xdata.DatabaseAttributes {
		if(len(xd.Key) > 0) {				
			if(!rep.Exists(xd.Key)) {
				var itm = rep.Add(xd.Name, xd.Key, xd.User, xd.Password, xd.Database, xd.Location, xd.Port)		
				log.WithFields(log.Fields{"Added dbscheme": "Name:"+itm.Name+" Key:"+itm.Key+" User:"+itm.User+" Port:"+strconv.Itoa(itm.Port) +" Database:"+itm.Database+" Location:"+itm.Location,	}).Debug("funcstr")	
			}
		}
	}  
}
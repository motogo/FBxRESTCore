package _tests

import (

		_struct "fbrest/FBxRESTCore/struct"
		//_permissions "fbrest/FBxRESTCore/permissions"
		//_sessions "fbrest/FBxRESTCore/sessions"
		_dbscheme "fbrest/FBxRESTCore/dbscheme"
		"encoding/json"
		//"encoding/xml"
		"io/ioutil"
		//"time"
		log "github.com/sirupsen/logrus"
)

func WriteUrlTABLEAttributesJson(pfile string) {
	var data _struct.GetUrlTABLEAttributes
	data.Table = "TSTANDORT"	
	data.Fields = append(data.Fields,"ID")
	data.Fields = append(data.Fields,"BEZ")
	data.Fields = append(data.Fields,"GUELTIG")
	
	data.Filter = "ID=1 AND BEZ like 'x%'"
	data.OrderBy = append(data.OrderBy,"BEZ,ASC")
	data.OrderBy = append(data.OrderBy,"ID,DESC")
	data.GroupBy = append(data.GroupBy,"ID")
	data.GroupBy = append(data.GroupBy,"BEZ")

	file, _ := json.MarshalIndent(data, "", " ") 
	_ = ioutil.WriteFile(pfile, file, 0644)

}

func WriteGetUrlPayloadAttributesJson(pfile string) {
	var data _struct.GetUrlPayloadAttributes
	var va _struct.FieldValueAttributes
	va.Field = "ID"
	va.Value = "'123'"
	
	data.Payload = append(data.Payload,va.Field+"="+va.Value)
	va.Field = "BEZ"
	va.Value = "'test'"
	data.Payload = append(data.Payload,va.Field+"="+va.Value)
	va.Field = "GUELTIG"
	va.Value = "1"
	data.Payload = append(data.Payload,va.Field+"="+va.Value)		
	data.Filter = "ID=1 AND BEZ like 'x%'"
	
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile(pfile, file, 0644)
	
}




func ReadUrlTABLEAttributesJson(pfile string) {
	data, err := ioutil.ReadFile(pfile)
    if err != nil {		
		log.WithFields(log.Fields{"File reading error": err,	}).Error("func ReadUrlTABLEAttributesJson")	
        return
    }
	xdata := &_struct.GetUrlPayloadAttributes{}
	json.Unmarshal(data,&xdata)
	log.Info(data)
}

func WriteUrlSessionAttributesJson(pfile string) {
	var data _dbscheme.DatabaseAttributes
	
	data.Password = "su"	
	data.User     = "superuser"
	data.Location = "localhost"
	data.Database = "D:/Data/Dokuments/DOKUMENTS30.FDB"
	
	file, _ := json.MarshalIndent(data, "", " ")
 
	_ = ioutil.WriteFile(pfile, file, 0644)
	
}






func ReadUrlSessionAttributesJson(pfile string) {
	data, err := ioutil.ReadFile(pfile)
    if err != nil {		
		log.WithFields(log.Fields{"File reading error": err,	}).Error("func ReadUrlTABLEAttributesJson")	
        return
    }
	xdata := &_dbscheme.DatabaseAttributes{}
	json.Unmarshal(data,&xdata)
	log.Info(data)
}







func Dummy() {

}

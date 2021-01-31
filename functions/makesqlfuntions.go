// Package makesqlfunctions
package _functions

import (
	"strconv"
	"strings"	
	log "github.com/sirupsen/logrus"
	_struct "fbrest/FBxRESTCore/struct"	
//	bguuid "github.com/kjk/betterguid"	
	guuid "github.com/google/uuid"
)

// Makes SQL Field-Value string from given field and value
// e.g. 
// Field: BEZ, Value:'123' -> return BEZ = '123'
// Field: BEZ, Value:'1*3' -> return BEZ like '1%3'
func MakeFieldValue(field string, value string) string {
	var filter string
	if(strings.HasPrefix(value,"'")) {
		if(value == "'*'") {
			filter = ""
		} else if (strings.Contains(value,"*")) {
			value = strings.Replace(value,"*","%",-1)		
			filter = field+" like ("+value+")"
		} else {
			filter = field+" = "+value	
		}
	} else {
		if(HasOperator(value)){
			filter = field+" "+value
		} else {
			filter = field+" = "+value	
		}
	}
	return filter
}

// Makes select SQL string for database from struct GetTABLEAttributes
func MakeSelectSQL(entitiesData _struct.GetTABLEAttributes) (cmd string) {
	var sql string
	if(len(entitiesData.Fields) < 1) {
		entitiesData.Fields = "*"
	}

	var limitstr string
	if(entitiesData.First > 0) {
		limitstr = " FIRST "+strconv.Itoa(entitiesData.First)	
	}

	if(entitiesData.Skip > 0) {
		limitstr = limitstr + " SKIP "+strconv.Itoa(entitiesData.Skip)	
	}

	sql = "SELECT"+limitstr+" "+entitiesData.Fields+" FROM " + entitiesData.Table
	if(len(entitiesData.Filter) > 0) {
		sql = sql + " WHERE " + entitiesData.Filter
	}
	if(len(entitiesData.GroupBy) > 0) {
		sql = sql + " GROUP BY " + entitiesData.GroupBy
	}
	if(len(entitiesData.OrderBy) > 0) {
		sql = sql + " ORDER BY " + entitiesData.OrderBy
	}
	cmd = sql
	return cmd
}



// Makes update table SQL string for database from struct FIELDVALUEAttributes
// insert into Scholars (firstname, lastname, address, phone, email)  values ('Henry', 'Higgins', '27A Wimpole Street', '3231212', null)
func MakeInsertTableSQL(entitiesData _struct.FIELDVALUEAttributes) (cmd string) {
	var cmdHead  = "INSERT INTO  " + entitiesData.Table 
	var cmdcol = ""
	var cmdval = ""
	var csplit  = "="
	log.Debug("fieldvalue")
	log.Debug(entitiesData.FieldValue)
	for _, fv := range entitiesData.FieldValue {
		log.Debug("fv")
		log.Debug(fv)
		//var par = strings.Split(fv,csplit)
		var par = SplitPars(fv,csplit)
		log.Debug(par)
		if(len(par) > 1){
			if(len(cmdcol)>0) { 				    
				cmdcol = cmdcol + " , "
			}
			if(len(cmdval)>0) { 				    
				cmdval = cmdval + " , "
			}
			cmdcol = cmdcol + par[0]
			cmdval = cmdval + par[1]
	    }
	}

	var id = guuid.New().String()	
	cmd = cmdHead + "(ID," + cmdcol +") VALUES ('"+id+"',"+ cmdval+ ")"
	return
}

// Makes update table SQL string for database from struct FIELDVALUEAttributes
func MakeUpdateTableSQL(entitiesData _struct.FIELDVALUEAttributes) (cmd string) {
	var cmdHead  = "UPDATE " + entitiesData.Table + " SET " 
	for _, fv := range entitiesData.FieldValue {
		if(len(cmd)>0) { 
			cmd = cmd + " , "
		}
		cmd = cmd + fv
	}
	cmd = cmdHead + cmd + " WHERE " + entitiesData.Filter
	return
}

// Makes drop table SQL string for database from struct MakeDeleteTableSQL
func MakeDeleteTableSQL(entitiesData _struct.FIELDVALUEAttributes) (cmd string) {
	cmd = "DROP TABLE " + entitiesData.Table
	return
}

// Makes drop field SQL string for database table from struct MakeDeleteTableSQL
func MakeDeleteTableFieldSQL(entitiesData _struct.FIELDVALUEAttributes) (cmd string) {
	cmd = "ALTER TABLE " + entitiesData.Table + " DROP " + entitiesData.FieldValue[0]
	return
}
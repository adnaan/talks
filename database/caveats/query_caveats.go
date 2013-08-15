// +build ignore
package main


//won't execute

//SINGLEINIT OMIT
err := db.QueryRow("select id, name, age from person where id = ?", 1).Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		log.Fatal(err)
	}
//SINGLEEND OMIT

//NULLINIT OMIT
for rows.Next() {
    var s sql.NullString
    err := rows.Scan(&s)
    if err!=nil{
    	//do something
    }
    if s.Valid {
       // use s.String
    } else {
       // NULL value
    }
}
//NULLEND OMIT

//COLINIT OMIT
cols, err := rows.Columns() 
if err!=nil{
    	//do something
    }            
vals := make([]sql.RawBytes, len(cols)) 
ints := make([]interface{}, len(cols))  
for i := range ints {
    vals[i] = &ints[i] 
}
for rows.Next() {
    err := rows.Scan(vals...)
  	if err!=nil{
  		//do something
  	}
}
//COLEND OMIT

//RESINIT OMIT
_, err := db.Exec("DELETE FROM person")  // Ok
_, err := db.Query("DELETE FROM person") // Not Ok, returns sql.Result which hangs on until gargbage collected
//RESEND OMIT

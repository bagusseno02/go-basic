/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package database

var connection string

func init(){
	connection = "POSTGRESQL"
}

func GetDatabase() string  {
	return connection
}


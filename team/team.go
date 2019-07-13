package team

import (
	"fmt"
	"github.com/Ankr-network/dccn-common/db-servicee"
	"os"
	"strings"
)

var (
	db  commondbservice.DBService
	err error
)

func init() {
	fmt.Println("team package initialization")
	ankr_micro.LoadConfigFromEnv()
	if db, err = dbservice.New(); err != nil {
		log.Printf("team package initialization error %s \n", err.Error())
	}
	defer db.Close()
}

func CheckUserAccessPrivilege(path string, uid string) (bool bool) { // read and write
    record, _ := db.GetRole(uid)


}

func CheckUserAccessReadPrivilege(path string, uid string)bool{
     return true
}

func CheckUserAccessWritePrivilege(path string, uid string)bool{
      return true
}

func GetTeamUser(uid string) string { // return team uid

}

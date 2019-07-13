package commondbservice

import (
	"errors"
	micro2 "github.com/Ankr-network/dccn-common/ankr-micro"
	"github.com/Ankr-network/dccn-common/protos"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)





type DBService interface {
	GetTeamUsers(team_id string)(*[]*UserRole, error)
	InsertRole(role UserRole)(error)
	UpdateRole(role UserRole)(error)
	GetRole(id string)(*UserRole, error)
	CreateTeam(record TeamRecord)(error)
    GetTeam(team_id string)(*TeamRecord, error)
	Close()
}


// UserDB implements DBService
type DB struct {
	role *mgo.Collection
	team *mgo.Collection
	user *mgo.Collection
}

// New returns DBService.
func New() (*DB, error) {
	config := micro2.GetConfig()
	config.DatabaseName = "dccn"
	roleCollection := micro2.GetCollection("role")
	userCollection := micro2.GetCollection("user")
	teamCollection := micro2.GetCollection("team")
	return &DB{
		role: roleCollection,
		team: teamCollection,
		user: userCollection,
	}, nil
}

func (p *DB) Close() {
	//p.Close()
}



type UserRole struct {
	ID               string             `bson:"uid"`
	Email            string             `bson:"email"`
	Name             string             `bson:"name"`
	Role             string             `bson:"role"`
	TeamID           string             `bson:"team_id"`
	Privilege        string             `bson:"privilege"`
}



type TeamRecord struct {
	Name             string             `bson:"name"`
	TeamID           string             `bson:"team_id"`
	Uid              string             `bson:"uid"`
}




func(p *DB) GetTeamUsers(team_id string)(*[]*UserRole, error){
	var users []*UserRole
	if err := p.role.Find(bson.M{"team_id": team_id}).All(&users); err != nil {
		return nil, err
	}

	return &users, nil
}

func(p *DB) InsertRole(role UserRole)(error){
	err := p.role.Insert(role)
	if err != nil {
		return errors.New(ankr_default.DbError+err.Error())
	}
	return err
}


func(p *DB) UpdateRole(role UserRole)(error){
	err := p.role.Update(bson.M{"uid": role.ID},
	bson.M{
		"$set": bson.M{
		"role":       role.Role,
		"privilege":        role.Privilege}})

	log.Printf("UpdateRole %s \n", role.Privilege)
	if err != nil {
		return errors.New(ankr_default.DbError+err.Error())
	}
	return err
}

func(p *DB) GetRole(userid string)(*UserRole, error){
	var role UserRole
	err := p.role.Find(bson.M{"uid": userid}).One(&role)
	if err != nil {
		return nil, errors.New(ankr_default.DbError+err.Error())
	}
	return &role, err
}

func(p *DB) CreateTeam(record TeamRecord)(error){
	err := p.team.Insert(record)
	if err != nil {
		return errors.New(ankr_default.DbError+err.Error())
	}
	return err
}

func(p *DB)GetTeam(team_id string)(*TeamRecord, error){
	var record TeamRecord
	err := p.role.Find(bson.M{"team_id": team_id}).One(&record)
	if err != nil {
		return nil, errors.New(ankr_default.DbError+err.Error())
	}
	return &record, err
}

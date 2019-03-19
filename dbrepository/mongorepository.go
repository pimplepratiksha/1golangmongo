package dbrepository

import (
	domain "assignment1/domain"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	"encoding/json"
	"bufio"
	"os"
	"fmt"
	"strings"
)

type MongoRepository struct {
	mongoSession *mgo.Session
	db           string
}

var collectionName = "restaurant"

//*****************************************************************************************************

//NewMongoRepository create new repository
func NewMongoRepository(mongoSession *mgo.Session, db string) *MongoRepository {
	return &MongoRepository{
		mongoSession: mongoSession,
		db:           db,
	}
}

func (r *MongoRepository) Insert(filename string) (int,error){
	fname,_:=os.Open(filename)
	defer fname.Close()
	fp:=bufio.NewScanner(fname)
	line:=""
	var final domain.Restaurant
	rcnt:=0
	for fp.Scan(){
			rcnt+=1
			line = fp.Text()
			json.Unmarshal([]byte(line),&final)
			final.DBID=domain.NewID()
			p,_:=r.Store(&final)
			fmt.Println(p)
	}
	return rcnt,nil
}

//Find a Restaurant(reader)
func (r *MongoRepository) Get(id domain.ID) (*domain.Restaurant, error) {
	result := domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"_id": id}).One(&result)
	switch err {
		case nil:
			return &result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//get all restaurants(reader)
func (r *MongoRepository) GetAll() ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{}).All(&result)
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Find a Restaurant By Name(reader)
func (r *MongoRepository) FindByName(name string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	err := coll.Find(bson.M{"name":name}).All(&result) 	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Store a Restaurant record(writer)
func (r *MongoRepository) Store(b *domain.Restaurant) (domain.ID, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	if domain.ID(0) == b.DBID {
		b.DBID = domain.NewID()
	}
	_, err := coll.UpsertId(b.DBID, b)

	if err != nil {
		return domain.ID(0), err
	}
	return b.DBID, nil
}

//Delete a Restaurant record(writer)
func (r *MongoRepository) Delete(id domain.ID)(error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Remove(bson.M{"_id": id})
	switch err {
		case nil:
			return nil
		case mgo.ErrNotFound:
			return domain.ErrNotFound
		default:
			return err
	}
}

//Find a Restaurant By Type Of Food(filter)
func (r *MongoRepository) FindByTypeOfFood(foodtype string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"type_of_food": foodtype}).All(&result) 	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}

//Find a Restaurant By Type Of Post Code(filter)
func (r *MongoRepository) FindByTypeOfPostCode(postcode string) ([]*domain.Restaurant, error) {
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	err := coll.Find(bson.M{"postcode": postcode}).All(&result) 	
	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}
//Count number of Restaurant By Type Of Food(filter)
func (r *MongoRepository) CountByTypeOfFood(foodtype string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	fcnt,err := coll.Find(bson.M{"type_of_food": foodtype}).Count() 	
	switch err {
		case nil:
			return fcnt, nil
		case mgo.ErrNotFound:
			return 0, domain.ErrNotFound
		default:
			return 1,err
	}
}

//Count number of Restaurant By Type Of Post Code(filter)
func (r *MongoRepository) CountByTypeOfPostCode(postcode string) (int, error) {
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	fcnt,err := coll.Find(bson.M{"postcode": postcode}).Count() 	
	switch err {
		case nil:
			return fcnt, nil
		case mgo.ErrNotFound:
			return 0, domain.ErrNotFound
		default:
			return 1,err
	}
}

//Search on Restaurant (filter)
func (r *MongoRepository)  Search(query string) ([]*domain.Restaurant, error){
	result := []*domain.Restaurant{}
	session := r.mongoSession.Clone()
	defer session.Close()
	coll := session.DB(r.db).C(collectionName)
	
	arr:=strings.Split(query,"=")
	key:=arr[0]
	value:=arr[1]
			
	err := coll.Find(bson.M{key:value}).All(&result) 	
	switch err {
		case nil:
			return result, nil
		case mgo.ErrNotFound:
			return nil, domain.ErrNotFound
		default:
			return nil, err
	}
}


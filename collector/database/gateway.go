package database

import (
	"github.com/gesiel/go-collect/collector/access"
	"github.com/gesiel/go-collect/collector/subscriber"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type SubscriberGatewayMongo struct {
	session *mgo.Session
}

type AccessGatewayMongo struct {
	session *mgo.Session
}

type accessDTO struct {
	ClientId   string
	Path       string
	Date       time.Time
	Identified bool
}

func (this *SubscriberGatewayMongo) Save(subscriber *subscriber.Subscriber) error {
	session := this.session.Copy()
	defer session.Close()

	db := session.DB("collector")

	_, err := db.C("access_data").UpdateAll(bson.M{
		"clientId": subscriber.ClientId,
	}, bson.M{
		"$set": bson.M{
			"name":       subscriber.Name,
			"email":      subscriber.Email,
			"identified": true,
		},
	})

	if err != nil {
		return err
	}

	err = db.C("subscribers").Insert(bson.M{
		"clientId": subscriber.ClientId,
		"name":     subscriber.Name,
		"email":    subscriber.Email,
	})

	return err
}

func (*SubscriberGatewayMongo) All() ([]*subscriber.SubscribersAccessData, error) {
	return []*subscriber.SubscribersAccessData{}, nil
}

func (this *AccessGatewayMongo) Save(access *access.Access) error {
	session := this.session.Copy()
	defer session.Close()

	collection := session.DB("collector").C("access_data")

	exists := map[string]*interface{}{}
	err := collection.Find(bson.M{"clientId": access.ClientId}).One(&exists)

	if err == nil {
		id := bson.NewObjectId()
		err = collection.Insert(bson.M{
			"_id":        id,
			"clientId":   access.ClientId,
			"path":       access.Path,
			"date":       access.Date,
			"identified": exists["identified"],
			"name":       exists["name"],
			"email":      exists["email"],
		})
		if err != nil {
			return err
		}
		access.Id = id.Hex()
	} else if err == mgo.ErrNotFound {
		id := bson.NewObjectId()
		err = collection.Insert(bson.M{
			"_id":        id,
			"clientId":   access.ClientId,
			"path":       access.Path,
			"date":       access.Date,
			"identified": false,
		})
		if err != nil {
			return err
		}
		access.Id = id.Hex()
	} else {
		return err
	}

	return nil
}

func NewSubscriberGateway(session *mgo.Session) subscriber.Gateway {
	gateway := new(SubscriberGatewayMongo)
	gateway.session = session
	return gateway
}

func NewAccessGateway(session *mgo.Session) access.Gateway {
	gateway := new(AccessGatewayMongo)
	gateway.session = session
	return gateway
}

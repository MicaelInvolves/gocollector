package database

import (
	"time"

	"github.com/gesiel/gocollector/access"
	"github.com/gesiel/gocollector/subscriber"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubscriberGatewayMongo struct {
	session *mgo.Session
}

type AccessGatewayMongo struct {
	session *mgo.Session
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

type dto struct {
	Id         string    `bson:"_id,omitempty"`
	ClientId   string    `bson:"clientId"`
	Path       string    `bson:"path"`
	Date       time.Time `bson:"date"`
	Identified bool      `bson:"identified"`
	Name       string    `bson:"name"`
	Email      string    `bson:"email"`
}

func (this *SubscriberGatewayMongo) All() ([]*subscriber.SubscribersAccessData, error) {
	session := this.session.Copy()
	defer session.Close()

	iter := session.DB("collector").C("access_data").Find(bson.M{
		"identified": true,
	}).Iter()

	resultMap := map[string]*subscriber.SubscribersAccessData{}
	row := dto{}
	for iter.Next(&row) {
		clientId := row.ClientId
		data := resultMap[clientId]
		if data == nil {
			data = &subscriber.SubscribersAccessData{
				Subscriber: &subscriber.Subscriber{
					ClientId: row.ClientId,
					Email:    row.Email,
					Name:     row.Name,
				},
				AccessCount: 0,
				AccessPaths: []string{},
			}
			resultMap[clientId] = data
		}

		data.AccessCount++
		if !contains(data.AccessPaths, row.Path) {
			data.AccessPaths = append(data.AccessPaths, row.Path)
		}
	}

	if iter.Err() != nil {
		return nil, iter.Err()
	}

	var result []*subscriber.SubscribersAccessData
	for _, value := range resultMap {
		result = append(result, value)
	}
	return result, nil
}

func contains(s []string, v string) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
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

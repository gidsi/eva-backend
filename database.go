package main

import (
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

func writeSpaceData(data SpaceData) {
	if(data.Space != "") {
		session, err := mgo.Dial(config.MongoDbServer)
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)

		c := session.DB(config.MongoDbDatabase).C("spacedata")
		_, err = c.Upsert(bson.M{ "space": data.Space }, data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func writeSpaceurl(spaceUrl SpaceUrl) {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spaceurl")
	count, _ := c.Find(bson.M{ "url": spaceUrl.Url }).Count()
	if(count == 0) {
		c.Insert(spaceUrl);
	}
}

func writeCalendar(calendar Calendar) {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("calendar")
	c.Upsert(bson.M{ "space": calendar.Space }, calendar)
}

func updateSpaceurl(spaceUrl SpaceUrl) {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spaceurl")
	c.Update(bson.M{ "url": spaceUrl.Url }, spaceUrl);
}

func readSpacedata() []SpaceData {
        session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spacedata")
        result := []SpaceData{}
        c.Find(bson.M{}).Iter().All(&result)

	return result
}

func readSpaceurl() []SpaceUrl {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("spaceurl")
	result := []SpaceUrl{}
	c.Find(bson.M{}).Iter().All(&result)

	return result
}

func readCalendar() []Calendar {
	session, err := mgo.Dial(config.MongoDbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB(config.MongoDbDatabase).C("calendar")
	result := []Calendar{}
	c.Find(bson.M{}).Iter().All(&result)

	return result
}
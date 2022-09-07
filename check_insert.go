package main

import (
	"context"
	"fmt"
	"github.com/netwatcherio/netwatcher-agent/agent_models"
	"github.com/netwatcherio/netwatcher-control/control_models"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func insertIcmpData(agent *control_models.Agent, data []agent_models.IcmpTarget, timestamp time.Time, c *mongo.Database) (bool, error) {
	var icmpData = control_models.IcmpData{
		ID:        primitive.NewObjectID(),
		Agent:     agent.ID,
		Data:      data,
		Timestamp: timestamp,
	}

	mar, err := bson.Marshal(icmpData)
	if err != nil {
		log.Errorf("1 %s", err)
		return false, err
	}
	var b *bson.D
	err = bson.Unmarshal(mar, &b)
	if err != nil {
		log.Errorf("2 %s", err)
		return false, err
	}
	result, err := c.Collection("icmp_data").InsertOne(context.TODO(), b)
	if err != nil {
		log.Errorf("3 %s", err)
		return false, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return true, nil
}

func insertMtrData(agent *control_models.Agent, data []agent_models.MtrTarget, timestamp time.Time, c *mongo.Database) (bool, error) {
	var icmpData = control_models.MtrData{
		ID:        primitive.NewObjectID(),
		Agent:     agent.ID,
		Data:      data,
		Timestamp: timestamp,
	}

	mar, err := bson.Marshal(icmpData)
	if err != nil {
		log.Errorf("1 %s", err)
		return false, err
	}
	var b *bson.D
	err = bson.Unmarshal(mar, &b)
	if err != nil {
		log.Errorf("2 %s", err)
		return false, err
	}
	result, err := c.Collection("mtr_data").InsertOne(context.TODO(), b)
	if err != nil {
		log.Errorf("3 %s", err)
		return false, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return true, nil
}

func insertNetworkInfo(agent *control_models.Agent, data agent_models.NetworkInfo, timestamp time.Time, c *mongo.Database) (bool, error) {
	var networkData = control_models.NetworkData{
		ID:        primitive.NewObjectID(),
		Agent:     agent.ID,
		Data:      data,
		Timestamp: timestamp,
	}

	mar, err := bson.Marshal(networkData)
	if err != nil {
		log.Errorf("1 %s", err)
		return false, err
	}
	var b *bson.D
	err = bson.Unmarshal(mar, &b)
	if err != nil {
		log.Errorf("2 %s", err)
		return false, err
	}
	result, err := c.Collection("network_data").InsertOne(context.TODO(), b)
	if err != nil {
		log.Errorf("3 %s", err)
		return false, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return true, nil
}

func insertSpeedTestData(agent *control_models.Agent, data agent_models.SpeedTestInfo, timestamp time.Time, c *mongo.Database) (bool, error) {
	var networkData = control_models.SpeedTestData{
		ID:        primitive.NewObjectID(),
		Agent:     agent.ID,
		Data:      data,
		Timestamp: timestamp,
	}

	mar, err := bson.Marshal(networkData)
	if err != nil {
		log.Errorf("1 %s", err)
		return false, err
	}
	var b *bson.D
	err = bson.Unmarshal(mar, &b)
	if err != nil {
		log.Errorf("2 %s", err)
		return false, err
	}
	result, err := c.Collection("speedtest_data").InsertOne(context.TODO(), b)
	if err != nil {
		log.Errorf("3 %s", err)
		return false, err
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return true, nil
}
package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/netwatcherio/netwatcher-agent/api"
	"github.com/netwatcherio/netwatcher-agent/checks"
	_ "github.com/netwatcherio/netwatcher-agent/checks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"netwatcher-control/handler"
)

func (r *Router) getConfiguration() {
	r.App.Post("/api/v2/config/", func(c *fiber.Ctx) error {
		c.Accepts("Application/json") // "Application/json"
		respB := api.Data{}

		var dataRequest api.Data

		err := json.Unmarshal(c.Body(), dataRequest)
		if err != nil {
			respB.Error = "500"
		}

		var agentSearch handler.Agent
		if dataRequest.ID != "" && dataRequest.PIN != "" {
			agentSearch.Pin = dataRequest.PIN
			hexId, err := primitive.ObjectIDFromHex(dataRequest.ID)
			if err != nil {
				respB.Error = "500"
			}
			agentSearch.ID = hexId
		} else if dataRequest.ID == "" && dataRequest.PIN != "" {
			agentSearch.Pin = dataRequest.PIN
			agentSearch.Initialized = false
		} else {
			respB.Error = "500"
		}

		if (agentSearch.ID != primitive.ObjectID{0} || agentSearch.Initialized == false) && (agentSearch.Pin != "") {
			err := agentSearch.Verify(r.DB)
			if err != nil {
				respB.Error = "500"
			}

			respB.ID = agentSearch.ID.Hex()
			respB.PIN = agentSearch.Pin
			respB.Checks = []checks.CheckData{}
			//todo add checks to be processed

			var agentCheck handler.AgentCheck

			all, err := agentCheck.GetAll(r.DB)
			if err != nil {
				return err
			}

			for _, ac := range all {
				modifiedData := checks.CheckData{
					Type:     string(ac.Type),
					Target:   ac.Target,
					ID:       ac.ID.Hex(),
					Duration: ac.Duration,
					Count:    ac.Count,
					Server:   ac.Server,
				}

				respB.Checks = append(respB.Checks, modifiedData)
			}
		}

		jRespB, err := json.Marshal(respB)
		if err != nil {
			return c.SendString("Something went wrong...") // => ✋
		}
		return c.SendString(string(jRespB)) // => ✋ good
	})
}

/*app.Post("/v1/agent/update/icmp", func(c *fiber.Ctx) error {
	var str = apiUpdateIcmp(c, db)
	if str != "" {
		return c.SendString(str)
	}

	return c.SendString("Something went wrong...") // => ✋
})*/

/*func apiUpdateIcmp(c *fiber.Ctx, db *mongo.Database) string {
	c.Accepts("Application/json") // "Application/json"
	respB := agent_models.ApiResponse{}
	respB.Response = 200

	var data agent_models.ApiPushData
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		log.Errorf("2 %s", err)
		respB.Response = 500
	}

	agentId, hash, b, err := verifyAgentHash(data.Pin, data.Hash, db)
	if err != nil {
		log.Errorf("0 %s", err)
		respB.Response = 401
	}

	if hash != "" && b {
		icmpD, _ := json.Marshal(data.Data)
		if err != nil {
			log.Errorf("1 %s", err)
			respB.Response = 500
		}

		var data2 []agent_models.IcmpTarget
		err := json.Unmarshal(icmpD, &data2)
		if err != nil {
			log.Errorf("2 %s", err)
			respB.Response = 500
		}

		var agent, _ = getAgent(agentId, db)
		if err != nil {
			log.Errorf("5 %s", err)
			respB.Response = 500
		}

		_, err = insertIcmpData(agent, data2, data.Timestamp, db)
		if err != nil {
			respB.Response = 500
		}
	}

	jRespB, err := json.Marshal(respB)
	if err != nil {
		log.Errorf("3 Unable to marshal API response.")
	} else {
		return string(jRespB) // => ✋ good
	}

	return ""
}*/

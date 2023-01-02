package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"netwatcher-control/handler"
	"netwatcher-control/models"
)

app.Get("/agent/:agent?", func(c *fiber.Ctx) error {
	b, _ := handler.ValidateSession(c, session, db)
	if !b {
		return c.Redirect("/auth/login")
	}

	user, err := handler.GetUserFromSession(c, session, db)
	if err != nil {
		return c.Redirect("/auth")
	}

	user.Password = ""

	if c.Params("agent") == "" {
		return c.RedirectBack("/home")
	}
	objId, err := primitive.ObjectIDFromHex(c.Params("agent"))
	if err != nil {
		return c.RedirectBack("/home")
	}

	agent := models.Agent{ID: objId}
	err = agent.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/agents")
	}

	site := models.Site{ID: objId}
	err = site.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/home")
	}

	marshal, err := json.Marshal(agent)
	if err != nil {
		log.Errorf("13 %s", err)
	}

	getAgentStats, err := handler.GetAgentStats(agent, db)
	if err != nil {
		return err
	}

	// TODO process if they are logged in or not, otherwise send them to registration/login
	return c.Render("agent", fiber.Map{
		"title":        agent.Name,
		"siteSelected": true,
		"siteName":     site.Name,
		"siteId":       site.ID.Hex(),
		"agents":       html.UnescapeString(string(marshal)),
		/*"mtr":              html.UnescapeString(string(marshalMtr)),
		"speed":            html.UnescapeString(string(marshalSpeed)),*/
		/*"online":           stats.Online,*/
		/*"agentStats":       html.UnescapeString(string(getAgentStats)),*/
		"publicAddress":    getAgentStats.NetInfo.PublicAddress,
		"localAddress":     getAgentStats.NetInfo.LocalAddress,
		"defaultGateway":   getAgentStats.NetInfo.DefaultGateway,
		"internetProvider": getAgentStats.NetInfo.InternetProvider,
		"uploadSpeed":      math.Round(getAgentStats.SpeedTestInfo.ULSpeed),
		"downloadSpeed":    math.Round(getAgentStats.SpeedTestInfo.DLSpeed),
		/*"speedtestPending": agent.AgentConfig.SpeedTestPending,*/
		"agentId":   agent.ID.Hex(),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		//"icmpMetrics":  html.UnescapeString(string(j)),
	},
		"layouts/main")
})
app.Get("/agents/:siteid?", func(c *fiber.Ctx) error {
	b, _ := handler.ValidateSession(c, session, db)
	if !b {
		return c.Redirect("/auth/login")
	}

	user, err := handler.GetUserFromSession(c, session, db)
	if err != nil {
		return c.Redirect("/auth")
	}

	user.Password = ""

	if c.Params("siteid") == "" {
		return c.Redirect("/home")
	}
	objId, err := primitive.ObjectIDFromHex(c.Params("siteid"))
	if err != nil {
		return c.Redirect("/home")
	}

	site := models.Site{ID: objId}
	err = site.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/home")
	}

	/*var agentStatList models.AgentStatsList

	stats, err := getAgentStatsForSite(objId, db)
	if err != nil {
		//todo handle error
		//return err
	}
	agentStatList.List = stats

	var hasAgents = true
	if len(agentStatList.List) == 0 {
		hasAgents = false
	}*/

	/*doc, err := json.Marshal(agentStatList)
	if err != nil {
		log.Errorf("1 %s", err)
	}*/

	/*agents, err := getAgents(objId, db)
	if err != nil {
		// todo handle error
		//return err
	}

	doc, err := json.Marshal(agents)
	if err != nil {
		log.Errorf("1 %s", err)
	}

	var hasAgentsBool = true
	if len(agents) == 0 {
		hasAgentsBool = false
		log.Warnf("%s", "site does NOT have agents")
	}*/

	// Render index within layouts/main
	// TODO process if they are logged in or not, otherwise send them to registration/login
	//log.Errorf("%s", string(doc))
	return c.Render("agents", fiber.Map{
		"title":        "agents",
		"siteSelected": true,
		"siteId":       site.ID.Hex(),
		"siteName":     site.Name,
		"firstName":    user.FirstName,
		"lastName":     user.LastName,
		"email":        user.Email},
		/*"agents":       html.UnescapeString(string(doc)),
		"hasAgents":    hasAgents},*/
		"layouts/main")
})
app.Get("/agent/new/:siteid?", func(c *fiber.Ctx) error {
	// Render index within layouts/main
	b, _ := handler.ValidateSession(c, session, db)
	if !b {
		return c.Redirect("/auth/login")
	}

	user, err := handler.GetUserFromSession(c, session, db)
	if err != nil {
		return c.Redirect("/auth")
	}

	user.Password = ""

	if c.Params("siteid") == "" {
		return c.Redirect("/home")
	}
	objId, err := primitive.ObjectIDFromHex(c.Params("siteid"))
	if err != nil {
		return c.Redirect("/home")
	}

	site := models.Site{ID: objId}
	err = site.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/home")
	}

	// TODO process if they are logged in or not, otherwise send them to registration/login
	return c.Render("agent_new", fiber.Map{
		"title":        "new agent",
		"firstName":    user.FirstName,
		"lastName":     user.LastName,
		"email":        user.Email,
		"siteId":       site.ID.Hex(),
		"siteName":     site.Name,
		"siteSelected": true,
	},
		"layouts/main")
})
/*app.Post("/agent/new/:siteid?", func(c *fiber.Ctx) error {
	c.Accepts("application/x-www-form-urlencoded") // "Application/json"

	// todo recevied body is in url format, need to convert to new struct??
	//

	// Render index within layouts/main
	b, _ := handler.ValidateSession(c, session, db)
	if !b {
		return c.Redirect("/auth/login")
	}

	user, err := handler.GetUserFromSession(c, session, db)
	if err != nil {
		return c.Redirect("/auth")
	}

	user.Password = ""

	if c.Params("siteid") == "" {
		return c.Redirect("/agents")
	}
	objId, err := primitive.ObjectIDFromHex(c.Params("siteid"))
	if err != nil {
		return c.Redirect("/agents")
	}

	site := models.Site{ID: objId}
	err = site.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/home")
	}

	cAgent := new(models.CreateAgent)
	if err := c.BodyParser(cAgent); err != nil {
		log.Warnf("4 %s", err)
		return err
	}

	icmpTargets := strings.Split(cAgent.IcmpTargets, ",")
	mtrTargets := strings.Split(cAgent.MtrTargets, ",")

	agentId, err := CreateAgent(cAgent.Name, icmpTargets, mtrTargets, site.ID, db)
	if err != nil {
		//todo handle error??
		return c.Redirect("/agents")
	}

	// todo handle error/success and return to home
	return c.Redirect("/agent/install/" + agentId.String())
})*/
app.Get("/agent/install/:agentid", func(c *fiber.Ctx) error {
	b, _ := handler.ValidateSession(c, session, db)
	if !b {
		return c.Redirect("/auth/login")
	}

	user, err := handler.GetUserFromSession(c, session, db)
	if err != nil {
		return c.Redirect("/auth")
	}

	user.Password = ""

	if c.Params("agentid") == "" {
		return c.Redirect("/home")
	}
	objId, err := primitive.ObjectIDFromHex(c.Params("agentid"))
	if err != nil {
		return c.Redirect("/home")
	}

	agent := models.Agent{ID: objId}
	err = agent.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/agents")
	}

	site := models.Site{ID: objId}
	err = site.Get(db)
	if err != nil {
		log.Error(err)
		return c.Redirect("/home")
	}

	//todo handle if already installed

	return c.Render("agent_install", fiber.Map{
		"title":        "agent install",
		"siteSelected": true,
		"siteId":       agent.Site.Hex(),
		"siteName":     site.Name,
		"firstName":    user.FirstName,
		"lastName":     user.LastName,
		"email":        user.Email,
		"agentPin":     agent.Pin,
	},
		"layouts/main")
})

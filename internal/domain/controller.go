package domain

type Controller struct {
	inputService  InputService
	checkService  CheckService
	outputService OutputService
}

func NewController(inputService InputService, checkService CheckService, outputService OutputService) *Controller {
	return &Controller{inputService: inputService, checkService: checkService, outputService: outputService}
}

func (c *Controller) InitChecks() error {
	sitesToCheck, err := c.inputService.GetSitesToCheck()
	if err != nil {
		return err
	}
	checkedSites, err := c.checkService.CheckSites(sitesToCheck)
	if err != nil {
		return err
	}
	err = c.outputService.SendToOutput(checkedSites)
	return nil
}

package output_service

import "health_check/internal/domain"

type MultiOutput struct {
	outputs []domain.OutputService
}

func NewComplexOutput(outputs []domain.OutputService) *MultiOutput {
	return &MultiOutput{outputs: outputs}
}

func (c MultiOutput) SendToOutput(report []*domain.SiteChecked) error {
	for _, output := range c.outputs {
		err := output.SendToOutput(report)
		if err != nil {
			return err
		}
	}
	return nil
}

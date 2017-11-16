package controllers

import (
	"github.com/goadesign/goa"
	"github.com/linh-snoopy/go-examples/goatest/gen/app"
)

// OperandsController implements the Operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a Operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Sum runs the sum action.
func (c *OperandsController) Sum(ctx *app.SumOperandsContext) error {
	// OperandsController_Sum: start_implement

	// Put your logic here
	v := ctx.Left + ctx.Right

	// OperandsController_Sum: end_implement
	res := &app.MyResult{&v}
	return ctx.OK(res)
}

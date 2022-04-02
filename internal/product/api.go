package product

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/globant/crud_project/internal/errors"
	"github.com/globant/crud_project/pkg/log"
	"net/http"
	"strconv"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/products/<id>", res.get)
	r.Post("/products", res.create)
	r.Put("/products/<id>", res.update)
	r.Delete("/products/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	
	product, err := r.service.Get(c.Request.Context(), uint32(id))
	if err != nil {
		return err
	}

	return c.Write(product)
}

func (r resource) create(c *routing.Context) error {
	var input CreateProductRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	product, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(product, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	var input UpdateProductRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	product, err := r.service.Update(c.Request.Context(), uint32(id), input)
	if err != nil {
		return err
	}

	return c.Write(product)
}

func (r resource) delete(c *routing.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	product, err := r.service.Delete(c.Request.Context(), uint32(id))
	if err != nil {
		return err
	}

	return c.Write(product)
}

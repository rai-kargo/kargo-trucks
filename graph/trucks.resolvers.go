package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rai-kargo/kargo-trucks/graph/generated"
	"github.com/rai-kargo/kargo-trucks/graph/model"
)

func (r *mutationResolver) SaveTruck(ctx context.Context, id *string, plateNo string) (*model.Truck, error) {
	truck := &model.Truck{
		ID:      fmt.Sprintf("TRUCK-%d", len(r.Trucks)+1),
		PlateNo: plateNo,
	}
	r.Trucks = append(r.Trucks, truck)
	return truck, nil
}

func (r *mutationResolver) SaveShipment(ctx context.Context, id *string, name string, origin string, destination string, deliveryDate string, truckID *string) (*model.Shipment, error) {
	shipment := &model.Shipment{
		ID:           fmt.Sprintf("SHIPMENT-%d", len(r.Shipments)+1),
		Name:         name,
		Origin:       origin,
		Destination:  destination,
		DeliveryDate: deliveryDate,
		Truck:        &model.Truck{ID: *truckID},
	}
	r.Shipments = append(r.Shipments, shipment)
	return shipment, nil
}

func (r *queryResolver) PaginatedTrucks(ctx context.Context) ([]*model.Truck, error) {
	return r.Trucks, nil
}

func (r *queryResolver) PaginatedShipments(ctx context.Context, id *string, name *string, origin *string, destination *string, deliveryDate *string, page int, first int) ([]*model.Shipment, error) {
	return r.Shipments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) PaginatedShipmentss(ctx context.Context) ([]*model.Shipment, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) ID(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) PlateNo(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Page(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) First(ctx context.Context) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

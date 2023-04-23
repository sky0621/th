package controller

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/sky0621/familiagildo/adapter/controller/custommodel"
	"github.com/sky0621/familiagildo/domain/vo"
)

// RequestCreateGuildByGuest is the resolver for the requestCreateGuildByGuest field.
func (r *mutationResolver) RequestCreateGuildByGuest(ctx context.Context, input RequestCreateGuildInput) (*GuestToken, error) {
	acceptedNumber, err := r.Guild.RequestCreateGuildByGuest(ctx, vo.ToGuildName(input.GuildName), vo.ToOwnerMail(input.OwnerMail))
	if err != nil {
		log.Err(err).Send()
		AddGraphQLError(ctx, err)
		return nil, err
	}
	return &GuestToken{
		AcceptedNumber: acceptedNumber,
	}, err
}

// CreateOwnerByGuest is the resolver for the createOwnerByGuest field.
func (r *mutationResolver) CreateOwnerByGuest(ctx context.Context, input CreateOwnerByGuestInput) (*custommodel.Void, error) {
	panic(fmt.Errorf("not implemented: CreateOwnerByGuest - createOwnerByGuest"))
}

// CreateParticipantByGuest is the resolver for the createParticipantByGuest field.
func (r *mutationResolver) CreateParticipantByGuest(ctx context.Context, input CreateParticipantByGuestInput) (*custommodel.Void, error) {
	panic(fmt.Errorf("not implemented: CreateParticipantByGuest - createParticipantByGuest"))
}

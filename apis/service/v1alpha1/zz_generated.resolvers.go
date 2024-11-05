// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/macpaw/provider-opsgenie/apis/team/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TeamIDRef,
		Selector:     mg.Spec.ForProvider.TeamIDSelector,
		To: reference.To{
			List:    &v1alpha1.TeamList{},
			Managed: &v1alpha1.Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TeamID")
	}
	mg.Spec.ForProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TeamIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TeamID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TeamIDRef,
		Selector:     mg.Spec.InitProvider.TeamIDSelector,
		To: reference.To{
			List:    &v1alpha1.TeamList{},
			Managed: &v1alpha1.Team{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TeamID")
	}
	mg.Spec.InitProvider.TeamID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TeamIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ServiceIncidentRule.
func (mg *ServiceIncidentRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceIDRef,
		Selector:     mg.Spec.ForProvider.ServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceID")
	}
	mg.Spec.ForProvider.ServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ServiceIDRef,
		Selector:     mg.Spec.InitProvider.ServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceID")
	}
	mg.Spec.InitProvider.ServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceIDRef = rsp.ResolvedReference

	return nil
}

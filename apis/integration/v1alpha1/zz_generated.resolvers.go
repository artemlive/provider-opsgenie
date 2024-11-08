// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha11 "github.com/macpaw/provider-opsgenie/apis/api/v1alpha1"
	v1alpha1 "github.com/macpaw/provider-opsgenie/apis/team/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this IntegrationAction.
func (mg *IntegrationAction) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Create); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Create[i3].Responders); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Create[i3].Responders[i4].ID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Create[i3].Responders[i4].IDRef,
				Selector:     mg.Spec.ForProvider.Create[i3].Responders[i4].IDSelector,
				To: reference.To{
					List:    &v1alpha1.TeamList{},
					Managed: &v1alpha1.Team{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Create[i3].Responders[i4].ID")
			}
			mg.Spec.ForProvider.Create[i3].Responders[i4].ID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Create[i3].Responders[i4].IDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.IntegrationIDRef,
		Selector:     mg.Spec.ForProvider.IntegrationIDSelector,
		To: reference.To{
			List:    &v1alpha11.IntegrationList{},
			Managed: &v1alpha11.Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IntegrationID")
	}
	mg.Spec.ForProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IntegrationIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Create); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Create[i3].Responders); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Create[i3].Responders[i4].ID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.Create[i3].Responders[i4].IDRef,
				Selector:     mg.Spec.InitProvider.Create[i3].Responders[i4].IDSelector,
				To: reference.To{
					List:    &v1alpha1.TeamList{},
					Managed: &v1alpha1.Team{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Create[i3].Responders[i4].ID")
			}
			mg.Spec.InitProvider.Create[i3].Responders[i4].ID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Create[i3].Responders[i4].IDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IntegrationID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.IntegrationIDRef,
		Selector:     mg.Spec.InitProvider.IntegrationIDSelector,
		To: reference.To{
			List:    &v1alpha11.IntegrationList{},
			Managed: &v1alpha11.Integration{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IntegrationID")
	}
	mg.Spec.InitProvider.IntegrationID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IntegrationIDRef = rsp.ResolvedReference

	return nil
}

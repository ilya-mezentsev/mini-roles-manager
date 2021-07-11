package mock

import (
	sharedMock "mini-roles-backend/source/domains/shared/mock"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
	"mini-roles-backend/source/domains/shared/services/hash"
)

const count = 10

func MakeValidAppData() sharedModels.AppData {
	var resources []sharedModels.Resource
	role := sharedModels.Role{
		Id: sharedMock.ExistsRoleId,
	}

	for i := 0; i < 10; i++ {
		permitCreateId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("permit-create"))
		denyCreateId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("deny-create"))
		permitReadId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("permit-read"))
		denyReadId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("deny-read"))
		permitUpdateId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("permit-update"))
		denyUpdateId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("deny-update"))
		permitDeleteId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("permit-delete"))
		denyDeleteId := sharedModels.PermissionId(hash.Md5WithTimeAsKey("deny-delete"))

		role.Permissions = append(
			role.Permissions,
			denyCreateId,
			permitReadId,
			permitUpdateId,
			denyDeleteId,
		)

		resources = append(resources, sharedModels.Resource{
			Id: sharedModels.ResourceId(hash.Md5WithTimeAsKey(string(sharedMock.ExistsResourceId))),
			Permissions: []sharedModels.Permission{
				{
					Id:        permitCreateId,
					Operation: sharedResource.CreateOperation,
					Effect:    sharedResource.PermitEffect,
				},
				{
					Id:        denyCreateId,
					Operation: sharedResource.CreateOperation,
					Effect:    sharedResource.DenyEffect,
				},

				{
					Id:        permitReadId,
					Operation: sharedResource.ReadOperation,
					Effect:    sharedResource.PermitEffect,
				},
				{
					Id:        denyReadId,
					Operation: sharedResource.ReadOperation,
					Effect:    sharedResource.DenyEffect,
				},

				{
					Id:        permitUpdateId,
					Operation: sharedResource.UpdateOperation,
					Effect:    sharedResource.PermitEffect,
				},
				{
					Id:        denyUpdateId,
					Operation: sharedResource.UpdateOperation,
					Effect:    sharedResource.DenyEffect,
				},

				{
					Id:        permitDeleteId,
					Operation: sharedResource.DeleteOperation,
					Effect:    sharedResource.PermitEffect,
				},
				{
					Id:        denyDeleteId,
					Operation: sharedResource.DeleteOperation,
					Effect:    sharedResource.DenyEffect,
				},
			},
		})
	}

	return sharedModels.AppData{
		Resources: resources,
		Roles:     []sharedModels.Role{role},
	}
}

func MakeLoadsValidAppData() sharedModels.AppData {
	var res sharedModels.AppData
	for i := 0; i < count; i++ {
		r := MakeValidAppData()
		res.Resources = append(res.Resources, r.Resources...)
		res.Roles = append(res.Roles, r.Roles...)
	}

	return res
}

package validation

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	sharedModels "mini-roles-backend/source/domains/shared/models"
	sharedResource "mini-roles-backend/source/domains/shared/resource"
)

type CheckFn func(appData sharedModels.AppData) []string

var resourcesCheckFns = []CheckFn{
	resourcesIdsAreUnique,
	allLinksAreExist,
	permissionsIdsAreUnique,
	allOperationsAreExist,
	rolesIdsAreUnique,
	extendsIdsAreExist,
	permissionsAreExist,
	permissionsDoNotConflict,
	defaultRolesVersionIdIsPresent,
}

func Validate(appDataBytes []byte) (sharedModels.AppData, []string) {
	var appData sharedModels.AppData
	err := json.Unmarshal(appDataBytes, &appData)
	if err != nil {
		log.Errorf("Unable to unmarshal json data: %v", err)

		return appData, []string{
			"Unable to unmarshal json data",
		}
	}

	return appData, checkFields(appData)
}

func checkFields(appData sharedModels.AppData) []string {
	var (
		errorMessages    []string
		receivedMessages int
		messagesCount    int
	)
	errorsChan := make(chan []string)

	for _, resourcesCheckFn := range resourcesCheckFns {
		messagesCount++
		go func(resourcesCheckFn CheckFn) {
			errorsChan <- resourcesCheckFn(appData)
		}(resourcesCheckFn)
	}

	for {
		select {
		case errorMessage := <-errorsChan:
			receivedMessages++
			if errorMessage != nil {
				errorMessages = append(errorMessages, errorMessage...)
			}

		default:
			if receivedMessages >= messagesCount {
				return errorMessages
			}
		}
	}
}

func resourcesIdsAreUnique(appData sharedModels.AppData) []string {
	var messages []string
	resourcesIds := make(map[sharedModels.ResourceId]uint)
	for _, resource := range appData.Resources {
		_, hasResourceId := resourcesIds[resource.Id]
		if hasResourceId {
			resourcesIds[resource.Id] += 1
		} else {
			resourcesIds[resource.Id] = 1
		}
	}

	for resourceId, foundCount := range resourcesIds {
		if foundCount > 1 {
			messages = append(messages, fmt.Sprintf("Resource id %s found %d times", resourceId, foundCount))
		}
	}

	return messages
}

func allLinksAreExist(appData sharedModels.AppData) []string {
	var messages []string
	resourcesIds := make(map[sharedModels.ResourceId]struct{})
	for _, resource := range appData.Resources {
		resourcesIds[resource.Id] = struct{}{}
	}

	for _, resource := range appData.Resources {
		for _, resourceLink := range resource.LinksTo {
			_, resourceExists := resourcesIds[resourceLink]
			if !resourceExists {
				messages = append(messages, fmt.Sprintf(
					"Resource with id %s has not exists link with id: %s",
					resource.Id,
					resourceLink,
				))
			}
		}
	}

	return messages
}

func permissionsIdsAreUnique(appData sharedModels.AppData) []string {
	var messages []string
	permissionsIds := make(map[sharedModels.PermissionId]uint)
	for _, resource := range appData.Resources {
		for _, permission := range resource.Permissions {
			_, hasPermissionId := permissionsIds[permission.Id]
			if hasPermissionId {
				permissionsIds[permission.Id] += 1
			} else {
				permissionsIds[permission.Id] = 1
			}
		}
	}

	for _, resource := range appData.Resources {
		for _, permission := range resource.Permissions {
			if permissionsIds[permission.Id] > 1 {
				messages = append(messages, fmt.Sprintf(
					"Resource with id %s has not unique permission id - %s",
					resource.Id,
					permission.Id,
				))
				break
			}
		}
	}

	return messages
}

// allOperationsAreExist is about create|read|update|delete with permit|deny combination
func allOperationsAreExist(appData sharedModels.AppData) []string {
	/*
		structure like
			resource_id:
				create:
					permit
					deny
				read:
					permit
					deny
				...
	*/
	operations := make(map[sharedModels.ResourceId]map[string]map[string]bool)
	var messages []string

	for _, resource := range appData.Resources {
		for _, permission := range resource.Permissions {
			if _, found := operations[resource.Id]; !found {
				operations[resource.Id] = make(map[string]map[string]bool)
			}

			if _, found := operations[resource.Id][permission.Operation]; !found {
				operations[resource.Id][permission.Operation] = make(map[string]bool)
			}

			operations[resource.Id][permission.Operation][permission.Effect] = true
		}
	}

	for resourceId, operationsToEffect := range operations {
		for _, operation := range []string{
			sharedResource.CreateOperation,
			sharedResource.ReadOperation,
			sharedResource.UpdateOperation,
			sharedResource.DeleteOperation,
		} {
			operationEffects := operationsToEffect[operation]
			if operationEffects == nil {
				messages = append(messages, fmt.Sprintf(
					"Resource with id %s has no %s operation",
					resourceId,
					operation,
				))
				continue
			}

			for _, effect := range []string{
				sharedResource.PermitEffect,
				sharedResource.DenyEffect,
			} {
				effectExists := operationEffects[effect]
				if !effectExists {
					messages = append(messages, fmt.Sprintf(
						"Resource with id %s has no %s effect for %s operation",
						resourceId,
						effect,
						operation,
					))
				}
			}
		}
	}

	return messages
}

func rolesIdsAreUnique(appData sharedModels.AppData) []string {
	var messages []string
	rolesIds := make(map[sharedModels.RolesVersionId]map[sharedModels.RoleId]uint)
	for _, role := range appData.Roles {
		if _, hasRolesVersionId := rolesIds[role.VersionId]; !hasRolesVersionId {
			rolesIds[role.VersionId] = map[sharedModels.RoleId]uint{}
		}

		_, hasRoleId := rolesIds[role.VersionId][role.Id]
		if hasRoleId {
			rolesIds[role.VersionId][role.Id] += 1
		} else {
			rolesIds[role.VersionId][role.Id] = 1
		}
	}

	for rolesVersionId, roles := range rolesIds {
		for roleId, foundCount := range roles {
			if foundCount > 1 {
				messages = append(
					messages,
					fmt.Sprintf(
						"Role id %s found %d times for version %s",
						roleId,
						foundCount,
						rolesVersionId,
					),
				)
			}
		}
	}

	return messages
}

func extendsIdsAreExist(appData sharedModels.AppData) []string {
	var messages []string
	rolesIds := make(map[sharedModels.RolesVersionId]map[sharedModels.RoleId]struct{})
	for _, role := range appData.Roles {
		if _, hasRolesVersionId := rolesIds[role.VersionId]; !hasRolesVersionId {
			rolesIds[role.VersionId] = map[sharedModels.RoleId]struct{}{}
		}

		rolesIds[role.VersionId][role.Id] = struct{}{}
	}

	for _, role := range appData.Roles {
		for _, extendsId := range role.Extends {
			_, extendsIdFound := rolesIds[role.VersionId][extendsId]
			if !extendsIdFound {
				messages = append(messages, fmt.Sprintf(
					"Role with id %s has not exists extends id: %s, for version %s",
					role.Id,
					extendsId,
					role.VersionId,
				))
			}
		}
	}

	return messages
}

func permissionsAreExist(appData sharedModels.AppData) []string {
	var messages []string
	permissionsIds := make(map[sharedModels.PermissionId]struct{})
	for _, resource := range appData.Resources {
		for _, permission := range resource.Permissions {
			permissionsIds[permission.Id] = struct{}{}
		}
	}

	for _, role := range appData.Roles {
		for _, permissionId := range role.Permissions {
			_, permissionExists := permissionsIds[permissionId]
			if !permissionExists {
				messages = append(messages, fmt.Sprintf(
					"Role with id %s has not exists permission id: %s",
					role.Id,
					permissionId,
				))
			}
		}
	}

	return messages
}

func permissionsDoNotConflict(appData sharedModels.AppData) []string {
	var messages []string
	/*
		structure like
			roles_version_1:
				role_1:
					resource_1:
						create: 1
						read: 2
						...

		if we get 2 it means that role contains conflict permission
	*/
	rolesResourcesOperationsEffects := make(map[sharedModels.RolesVersionId]map[sharedModels.RoleId]map[sharedModels.ResourceId]map[string]uint)
	permissions := make(map[sharedModels.PermissionId]sharedModels.Permission)
	for _, resource := range appData.Resources {
		for _, permission := range resource.Permissions {
			permissions[permission.Id] = sharedModels.Permission{
				Id:        permission.Id,
				Operation: permission.Operation,
				Effect:    permission.Effect,
				Resource:  resource,
			}
		}
	}

	for _, role := range appData.Roles {
		for _, permissionId := range role.Permissions {
			permission, found := permissions[permissionId]
			if !found {
				// here we cannot be sure that permission with particular id is existing
				continue
			}

			if _, found = rolesResourcesOperationsEffects[role.VersionId]; !found {
				rolesResourcesOperationsEffects[role.VersionId] = make(map[sharedModels.RoleId]map[sharedModels.ResourceId]map[string]uint)
			}

			if _, found = rolesResourcesOperationsEffects[role.VersionId][role.Id]; !found {
				rolesResourcesOperationsEffects[role.VersionId][role.Id] = make(map[sharedModels.ResourceId]map[string]uint)
			}

			if _, found = rolesResourcesOperationsEffects[role.VersionId][role.Id][permission.Resource.Id]; !found {
				rolesResourcesOperationsEffects[role.VersionId][role.Id][permission.Resource.Id] = make(map[string]uint)
			}

			if _, found = rolesResourcesOperationsEffects[role.VersionId][role.Id][permission.Resource.Id][permission.Operation]; !found {
				rolesResourcesOperationsEffects[role.VersionId][role.Id][permission.Resource.Id][permission.Operation] = 1
			} else {
				rolesResourcesOperationsEffects[role.VersionId][role.Id][permission.Resource.Id][permission.Operation] += 1
			}
		}
	}

	for rolesVersionId, roles := range rolesResourcesOperationsEffects {
		for roleId, resourceIdToOperationEffectCount := range roles {
			for resourceId, operationToEffectCount := range resourceIdToOperationEffectCount {
				for _, operation := range []string{
					sharedResource.CreateOperation,
					sharedResource.ReadOperation,
					sharedResource.UpdateOperation,
					sharedResource.DeleteOperation,
				} {
					count := operationToEffectCount[operation]
					if count > 1 {
						messages = append(messages, fmt.Sprintf(
							"Role with id %s has conflict permissions for resource %s on %s operation for version %s",
							roleId,
							resourceId,
							operation,
							rolesVersionId,
						))
					}
				}
			}
		}
	}

	return messages
}

func defaultRolesVersionIdIsPresent(appData sharedModels.AppData) []string {
	var messages []string
	if appData.DefaultRolesVersionId == "" {
		messages = append(messages, "Default roles version id is not present")
	}

	return messages
}

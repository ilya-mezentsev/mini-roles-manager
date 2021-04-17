# mini-roles-manager
## [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control) for your systems

### Why?
This project can help you to manage access to your resources for different roles of your systems users. <br/>
All you need is Sign-Up, create resources and roles and adapt your system to this project API

### Deployment
Deployment instructions is available [here](./deployment.md)

### API usage
After you created account, needed resources and roles, added permissions for your roles you can use API for checking access for users. <br/>
The simplest example of API call is cURL:
```bash
$ curl -X POST localhost:8000/api/permissions \ 
  -H "X-RM-Auth-Token: YOUR_API_TOKEN" \ 
  -d '{"roleId": "role-1", "resourceId": "resource-1", "operation": "create"}'
```

#### Request body:
All parameters are required
* roleId - id of user role
* resourceId - id of resource
* operation - operation that user is going to perform under resource (create, read, update or delete)

#### Responses:
Each response has next fields:
* status (string) - "ok" or "error", status of response
* data (any) - response data

Examples <br/>
HTTP 200 Ok
```json5
{
  "status": "ok",
  "data": {
    "effect": "permit"
  }
}
```
```json5
{
  "status": "ok",
  "data": {
    "effect": "deny"
  }
}
```

400 Bad Request (passed invalid operation)
```json5
{
  "data": {
    "code": "validation-error",
    "description": "Key: 'PermissionAccess.Operation' Error:Field validation for 'Operation' failed on the 'oneof' tag"
  },
  "status":"error"
}
```

Unauthorized (No token in headers): <br/>
401 Unauthorized
```json5
{
  "data": {
    "code": "missed-token-in-headers",
    "description": "No auth token in headers"
  },
  "status":"error"
}
```

Forbidden (Provided token does not exists): <br/>
403 Forbidden
```json5
{
  "data":{
    "code":"no-account-by-token",
    "description":"Unable to find account by provided token"
  },
  "status":"error"
}
```

Server Error (Something horrible happened): <br/>
500 Internal Server Error
```json5
{
  "data":{
    "code":"unknown-error",
    "description":"Unknown error happened"
  },
  "status":"error"
}
```

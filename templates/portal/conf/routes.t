# routes
module:testrunner

# static
GET     /favicon.ico                            404
GET     /public/*filepath                       Static.Serve("public")
GET     /upload/*filepath                       Static.Serve("upload")

# portal
GET		/										PortalController.Index

# auth
GET     /auth.logout                            AuthController.Logout
GET     /auth.login                             AuthController.Login
POST    /auth.login                             AuthController.LoginPost
GET     /auth.password                          AuthController.Password
POST    /auth.password                          AuthController.PasswordPost

# widget
GET     /widget.index                           WidgetController.Index
GET     /widget.upload                          WidgetController.Upload
POST    /widget.upload                          WidgetController.UploadPost
GET     /widget.editor                          WidgetController.Editor
POST    /widget.editor                          WidgetController.EditorPost

[[range .tables]]
# [[.Name | module]]
GET     /[[.Name | module]].index                    [[.Name | class]]Controller.Index     
POST    /[[.Name | module]].query                    [[.Name | class]]Controller.Query
POST    /[[.Name | module]].finder.index             [[.Name | class]]Controller.FinderIndex
POST    /[[.Name | module]].finder.query             [[.Name | class]]Controller.FinderQuery         
GET     /[[.Name | module]].detail/:id               [[.Name | class]]Controller.Detail      
GET     /[[.Name | module]].create                   [[.Name | class]]Controller.Create     
POST    /[[.Name | module]].create                   [[.Name | class]]Controller.CreatePost
GET     /[[.Name | module]].update/:id               [[.Name | class]]Controller.Update
POST    /[[.Name | module]].update                   [[.Name | class]]Controller.UpdatePost
GET     /[[.Name | module]].remove/:id               [[.Name | class]]Controller.Remove     
POST    /[[.Name | module]].remove                   [[.Name | class]]Controller.RemovePost     
[[end]] 
 
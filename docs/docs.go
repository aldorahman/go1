// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `
swagger: '2.0'
info:
  version: '1.0'
  title: go1
  contact: {}
host: golang-bn.herokuapp.com
basePath: /api
securityDefinitions: {}
schemes:
  - https
consumes:
  - application/json
produces:
  - application/json
paths:
  /auth/login:
    post:
      summary: login
      tags:
        - auth
      operationId: login
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/loginrequest'
      responses:
        '200':
          description: ''
          headers: {}
  /auth/register:
    post:
      summary: register
      tags:
        - auth
      operationId: register
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/registerrequest'
      responses:
        '200':
          description: ''
          headers: {}
  /user/profile:
    get:
      summary: detail user
      tags:
        - user
      operationId: detailuser
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    put:
      summary: edit user
      tags:
        - user
      operationId: edituser
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/edituserrequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
  /todo:
    get:
      summary: list todo
      tags:
        - todo
      operationId: listtodo
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    post:
      summary: create todo
      tags:
        - todo
      operationId: createtodo
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/createtodorequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
  /todo/{id}:
    get:
      summary: detail todo
      tags:
        - todo
      operationId: detailtodo
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    put:
      summary: edit todo
      tags:
        - todo
      operationId: edittodo
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/edittodorequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    delete:
      summary: delete todo
      tags:
        - todo
      operationId: deletetodo
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/deletetodorequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
  /color:
    get:
      summary: list color
      tags:
        - color
      operationId: listcolor
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    post:
      summary: create color
      tags:
        - color
      operationId: createcolor
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/createcolorrequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
  /color/{id}:
    get:
      summary: detail color
      tags:
        - color
      operationId: detailcolor
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    put:
      summary: edit color
      tags:
        - color
      operationId: editcolor
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
        - name: Body
          in: body
          required: true
          description: ''
          schema:
            $ref: '#/definitions/editcolorrequest'
      responses:
        '200':
          description: ''
          headers: {}
      security: []
    delete:
      summary: delete color
      tags:
        - color
      operationId: deletecolor
      deprecated: false
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          required: true
          type: integer
          format: int32
          description: ''
        - name: Authorization
          in: header
          required: true
          type: string
          description: ''
      responses:
        '200':
          description: ''
          headers: {}
      security: []
definitions:
  loginrequest:
    title: loginrequest
    example:
      email: admin@gmail.com
      password: admin
    type: object
    properties:
      email:
        type: string
      password:
        type: string
    required:
      - email
      - password
  registerrequest:
    title: registerrequest
    example:
      name: admin
      email: admin@gmail.com
      password: admin
    type: object
    properties:
      name:
        type: string
      email:
        type: string
      password:
        type: string
    required:
      - name
      - email
      - password
  edituserrequest:
    title: edituserrequest
    example:
      name: admin
      email: admin@gmail.com
      password: admin
    type: object
    properties:
      name:
        type: string
      email:
        type: string
      password:
        type: string
    required:
      - name
      - email
      - password
  createtodorequest:
    title: createtodorequest
    example:
      title: macam2
      isi: pemangsa tempat
      reminder: '2022-06-04T06:18:30Z'
      colorId: 6
    type: object
    properties:
      title:
        type: string
      isi:
        type: string
      reminder:
        type: string
      colorId:
        type: integer
        format: int32
    required:
      - title
      - isi
      - reminder
      - colorId
  edittodorequest:
    title: edittodorequest
    example:
      title: macam2
      isi: pemangsa tempat
      reminder:
        Time: '2022-06-04T06:18:30Z'
        Valid: true
      colorId: 2
    type: object
    properties:
      title:
        type: string
      isi:
        type: string
      reminder:
        $ref: '#/definitions/Reminder'
      colorId:
        type: integer
        format: int32
    required:
      - title
      - isi
      - reminder
      - colorId
  Reminder:
    title: Reminder
    example:
      Time: '2022-06-04T06:18:30Z'
      Valid: true
    type: object
    properties:
      Time:
        type: string
      Valid:
        type: boolean
    required:
      - Time
      - Valid
  deletetodorequest:
    title: deletetodorequest
    example:
      title: macam2
      isi: pemangsa tempat
      colorId: 6
    type: object
    properties:
      title:
        type: string
      isi:
        type: string
      colorId:
        type: integer
        format: int32
    required:
      - title
      - isi
      - colorId
  createcolorrequest:
    title: createcolorrequest
    example:
      color_type: merah
      color_name: merah gelap
    type: object
    properties:
      color_type:
        type: string
      color_name:
        type: string
    required:
      - color_type
      - color_name
  editcolorrequest:
    title: editcolorrequest
    example:
      color_type: merah
      color_name: merah gelap
    type: object
    properties:
      color_type:
        type: string
      color_name:
        type: string
    required:
      - color_type
      - color_name
tags:
  - name: auth
  - name: user
  - name: todo
  - name: color
`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "golang-bn.herokuapp.com",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Swagger  demo service API",
	Description: "This is demo server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}

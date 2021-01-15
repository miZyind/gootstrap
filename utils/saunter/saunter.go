package saunter

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"net/http"
	"strings"

	spec "github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

func parseAttribute(value string) string {
	return strings.Split(value, " ")[0]
}

func parseInfo(context []string) *spec.Info {
	scheme := &spec.Info{}

	for _, value := range context {
		attribute := parseAttribute(value)
		term := strings.TrimSpace(value[len(attribute):])
		switch attribute {
		case "@title":
			scheme.Title = term
		case "@version":
			scheme.Version = term
		case "@description":
			scheme.Description = term
		}
	}

	return scheme
}

func parseSecurity(context []string) *spec.SecuritySchemeRef {
	scheme := &spec.SecuritySchemeRef{
		Value: &spec.SecurityScheme{},
	}

	for _, raw := range context {
		attribute := parseAttribute(raw)
		value := strings.TrimSpace(raw[len(attribute):])
		switch attribute {
		case "@name":
			scheme.Value.Name = value
		case "@type":
			scheme.Value.Type = value
		case "@description":
			scheme.Value.Description = value
		case "@scheme":
			scheme.Value.Scheme = value
		}
	}

	return scheme
}

func generateSpec(routes gin.RoutesInfo) spec.Swagger {
	swagger := spec.Swagger{OpenAPI: "3.0.0"}

	fileSet := token.NewFileSet()

	// Parse general
	fileTree, err := parser.ParseFile(fileSet, "./routers/v1/main.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("cannot parse source files %s: %s", "main.go", err)
	}

	securitySchemes := map[string]*spec.SecuritySchemeRef{}

	for _, comment := range fileTree.Comments {
		context := strings.Split(comment.Text(), "\n")

		for i, value := range context {
			switch parseAttribute(value) {
			case "@Info":
				swagger.Info = parseInfo(context[i+1:])
			case "@Security":
				security := parseSecurity(context[i+1:])
				securitySchemes[security.Value.Name] = security
			}

		}
	}

	if len(securitySchemes) > 0 {
		swagger.Components.SecuritySchemes = securitySchemes
	}

	// Parse API
	apiTree, err := parser.ParseFile(fileSet, "./routers/v1/todo.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("cannot parse source files %s: %s", "todo.go", err)
	}

	for _, route := range routes {
		if strings.Contains(route.Path, "todo") {
			for _, scope := range apiTree.Scope.Objects {
				if scope.Kind == ast.Fun && scope.Decl.(*ast.FuncDecl).Doc != nil {
					operation := &spec.Operation{
						Parameters: spec.NewParameters(),
						Security:   spec.NewSecurityRequirements(),
					}
					for _, raw := range strings.Split(scope.Decl.(*ast.FuncDecl).Doc.Text(), "\n") {
						attribute := parseAttribute(raw)
						value := strings.TrimSpace(raw[len(attribute):])
						switch attribute {
						case "@security":
							operation.Security.With(spec.NewSecurityRequirement().Authenticate(value))
						case "@summary":
							operation.Summary = value
						case "@200":
							operation.AddResponse(200, spec.NewResponse().WithDescription(value).WithJSONSchema(&spec.Schema{}))
						case "@400":
							operation.AddResponse(400, spec.NewResponse().WithDescription(value).WithJSONSchema(&spec.Schema{}))
						case "@500":
							operation.AddResponse(500, spec.NewResponse().WithDescription(value).WithJSONSchema(&spec.Schema{}))
						}
					}
					swagger.AddOperation(route.Path, route.Method, operation)
				}
			}
		}
	}

	return swagger
}

// Handler wraps `http.Handler` into `gin.HandlerFunc`
func Handler(routes gin.RoutesInfo) gin.HandlerFunc {
	tmpl, err := template.New("index").Parse(IndexTemplate)
	if err != nil {
		panic(err)
	}

	swagger := generateSpec(routes)

	return func(c *gin.Context) {
		if err := tmpl.Execute(c.Writer, struct{ Spec spec.Swagger }{swagger}); err != nil {
			panic(err)
		}
	}
}

// Static creates saunter static file system
func Static() http.FileSystem {
	static, err := fs.New()
	if err != nil {
		panic(err)
	}

	return static
}

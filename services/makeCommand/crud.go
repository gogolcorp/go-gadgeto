package makeCommand

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/edwinvautier/go-cli/helpers"
	"github.com/edwinvautier/go-cli/services/filesystem"
)

// AddControllersToRouter creates the routes inside the router file for the 5 controllers created by the crud command
func AddControllersToRouter(modelNamePascalCase string) error {
	workdir := filesystem.GetWorkdirOrDie()
	routerFile, err := ioutil.ReadFile(workdir + "/api/routes/routes.go")
	if err != nil {
		return err
	}
	routerContent := string(routerFile)

	var finalRouterLines []string
	routerLines := strings.Split(routerContent, "\n")
	// Create string
	linesToAppend := []string{
		"\t\tapi.POST(\"/" + helpers.LowerCase(modelNamePascalCase) + "\", controllers.Create" + modelNamePascalCase + ")",
		"\t\tapi.GET(\"/" + helpers.LowerCase(modelNamePascalCase) + "/:id\", controllers.Get" + modelNamePascalCase + "ByID)",
		"\t\tapi.GET(\"/" + helpers.LowerCase(modelNamePascalCase) + "\", controllers.GetAll" + modelNamePascalCase + ")",
		"\t\tapi.PUT(\"/" + helpers.LowerCase(modelNamePascalCase) + "/:id\", controllers.Update" + modelNamePascalCase + ")",
		"\t\tapi.DELETE(\"/" + helpers.LowerCase(modelNamePascalCase) + "/:id\", controllers.Delete" + modelNamePascalCase + "ByID)",
	}
	finalRouterLines = routerLines[:len(routerLines)-2]
	finalRouterLines = append(finalRouterLines, linesToAppend...)
	finalRouterLines = append(finalRouterLines, "}")
	file, err := os.OpenFile(workdir+"/api/routes/routes.go", os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(finalRouterLines, "\n"))

	return err
}

package pkg

import (
	"os"

	"text/template"
	"k8s.io/helm/pkg/chartutil"
	"path/filepath"
	"fmt"
	"os/exec"
)

func Build(config map[string]string) (error){
	err := ParseFile(config["path"], config, "values.yaml")
	if err != nil {
		return err
	}

	err = ParseFile(config["path"], config, "Chart.yaml")
	if err != nil {
		return err
	}

	name, err := Package(config["path"], config["destination"])

	cmd := exec.Command("curl", "-v","-T", name, "-X", "PUT", config["repository"])
	err = cmd.Run()

	if err != nil {
		return err
	}

	os.Remove(name)

	return err
}

func ParseFile(path string, config map[string]string, fileName string) error{
	file := filepath.Join(path, fileName)
	t, err := template.ParseFiles(file)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	err = t.Execute(f, config)
	if err != nil {
		return err
	}
	f.Close()

	return nil
}


func Package(path string, destination string) (string, error){
	path, err := filepath.Abs(path)

	ch, err := chartutil.LoadDir(path)
	if err != nil {
		return "",err
	}

	if filepath.Base(path) != ch.Metadata.Name {
		return "", fmt.Errorf("directory name (%s) and Chart.yaml name (%s) must match", filepath.Base(path), ch.Metadata.Name)
	}

	_, err = chartutil.LoadRequirements(ch)

	var dest string
	if destination == "." {
		// Save to the current working directory.
		dest, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		// Otherwise save to set destination
		dest = destination
	}

	name, err := chartutil.Save(ch, dest)

	return name, err
}
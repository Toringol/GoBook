/*
* Создайте инструмент, который сообщает о множестве всех
* пакетов в рабочей области, которые транзитивно зависят
* от пакетов, указанных аргументами командной строки.
* Указание: вам нужно будет выполнить go list дважды:
* один раз - для исходных пакетов и один раз - для
* всех пакетов. Вы можете проанализировать вывод в
* формате JSON с помощью пакета encoding/json (раздел 4.5).
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	ImportPath string
	Deps       []string
}

func main() {
	dependees, err := Dependees(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	for _, dependee := range dependees {
		fmt.Println(dependee.ImportPath)
	}
}

func List(template ...string) ([]Package, error) {
	cmd := exec.Command("go", append([]string{"list", "-json"}, template...)...)
	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var packages []Package
	dec := json.NewDecoder(bytes.NewReader(b))
	for {
		var val Package
		err := dec.Decode(&val)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		packages = append(packages, val)
	}
	return packages, nil
}

func Dependees(template ...string) ([]Package, error) {
	packages, err := List(template...)
	if err != nil {
		return nil, err
	}
	all, err := List("...")
	if err != nil {
		return nil, err
	}

	var dependees []Package
	for _, dependee := range all {
		depends := false
	loopDependency:
		for _, dependency := range dependee.Deps {
			for _, pack := range packages {
				if dependency == pack.ImportPath {
					depends = true
					break loopDependency
				}
			}
		}
		if depends {
			dependees = append(dependees, dependee)
		}
	}
	return dependees, nil
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	dgi "github.com/sataga/go-kubernetes-sample/domain/getimage"
	duc "github.com/sataga/go-kubernetes-sample/domain/updateconfigmap"
	ikb "github.com/sataga/go-kubernetes-sample/infra/kubernetes"
)

var (
	updateConfigMapFlag = flag.NewFlagSet("update-cm", flag.ExitOnError)
	target              = updateConfigMapFlag.String("target", "default", "update ConfigMap target (e.g. prometheus-cm)")
	dryrun              = updateConfigMapFlag.Bool("dry-run", false, "This option to just show what would be done")
)

func printDefaultsAll() {
	fmt.Println("usage: go-kubernetes-sample [global options] subcommand [subcommand options]")
	flag.PrintDefaults()
	fmt.Println("\nsubcommands:")
	fmt.Println("update-cm:    update user k8s configmaps")
	updateConfigMapFlag.PrintDefaults()
}

func main() {
	flag.Parse()
	subCommandArgs := os.Args[1+flag.NFlag():]
	if len(subCommandArgs) == 0 {
		printDefaultsAll()
		log.Fatalln("specify subcommand")
	}
	switch subCommand := subCommandArgs[0]; subCommand {
	case "update-cm":
		if err := updateConfigMapFlag.Parse(subCommandArgs[1:]); err != nil {
			log.Fatalf("parsing updating configmap flag: %s", err)
		}
		kcl, err := ikb.NewKubernetesClient()
		if err != nil {
			log.Fatalf("Initialize Kubernetes client: %s", err)
		}
		cmu := duc.NewConfigMapUpdater(kcl)

		if err = cmu.UpdateConfigMap(*target, *dryrun); err != nil {
			log.Fatalf("UpdateConfigMap %s: %s", *target, err)
		}
		fmt.Println("Finished!")
	case "get-image":
		if err := updateConfigMapFlag.Parse(subCommandArgs[1:]); err != nil {
			log.Fatalf("parsing updating configmap flag: %s", err)
		}
		kcl, err := ikb.NewKubernetesClient()
		if err != nil {
			log.Fatalf("Initialize Kubernetes client: %s", err)
		}
		ig := dgi.NewImageGetter(kcl)

		if err = ig.GetImage(*target, *dryrun); err != nil {
			log.Fatalf("UpdateConfigMap %s: %s", *target, err)
		}
		fmt.Println("Finished!")
	}

}

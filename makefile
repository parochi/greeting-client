# help messages
help:
	@echo 'Usage'
	@echo 'gen-client : Generate the client'

    
gen-client:
	@echo "Generating Client"
	GO111MODULE=off
	cd /Users/raz/workspace/go/src/github.com/parochi/greeting-client
	/Users/raz/workspace/go/src/k8s.io/code-generator/generate-groups.sh all "github.com/parochi/greeting-client/pkg/client" "github.com/parochi/greeting-client/pkg/apis" "parochi":"v1"

<!--
Thank you for your pull request. Please provide a description and check the items in the checklist.

Bug fixes and new features should include tests.
-->

### Change summary
<!-- Please provide what changed from an end users perspective. -->

### Technical details
<!-- Please provide a description what was changed here. -->

##### Checklist
<!-- Remove items that do not apply. For completed items, change [ ] to [x]. -->

- [ ] `go test -tags="unit,psql" ./...` passes
- [ ] `go test -tags=integration,psql -parallel=1 -p=1 ./... -args $PWD/config_psql.yaml` passes (needs running psql instance see [README](./README.md#prepare-dev-database))
- [ ] `go test -tags=integration,mysql -parallel=1 -p=1 ./... -args $PWD/config_mysql.yaml` passes (needs running mysql instance see [README](./README.md#prepare-dev-database))
- [ ] `golangci-lint run --build-tags psql,unit,integration` passes
- [ ] swagger go code comments have been adopted accordingly
- [ ] documentation(/docs) is adopted accordingly

<!-- _NOTE: these things are not required to open a PR and can be done afterwards / while the PR is open._ -->


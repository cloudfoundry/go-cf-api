# Architectural Decision Log

This log lists the architectural decisions for cloudgontroller.

<!-- adrlog -- Regenerate the content by using "adr-log -i" and inserting into here. You can install it via "npm install -g adr-log" -->

* [ADR-0000](0000-use-markdown-architectural-decision-records.md) - Use Markdown Architectural Decision Records
* [ADR-0001](0001-use-sql-boiler-and-build-tags.md) - Use [sqlboiler](https://github.com/volatiletech/sqlboiler) together with build tags to support multiple databases
* [ADR-0002](0002-use-zap-logging-framework.md) - Use [Zap](https://github.com/uber-go/zap) logging as a performant and customizable structured logging framework
* [ADR-0003](0003-split-traffic-between-both-implementations.md) - Deploy [HAProxy](https://www.haproxy.org/) to route specific endpoints/methods to the new implementation

<!-- adrlogstop -->

For new ADRs, please use [template.md](https://github.com/adr/madr/blob/2.1.2/template/template.md) as basis.
More information on MADR is available at <https://adr.github.io/madr/>.
General information about architectural decision records is available at <https://adr.github.io/>.

<!-- markdownlint-disable-file MD013 -->
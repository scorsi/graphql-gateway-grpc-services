local container = import "./container.jsonnet";
local deployment = import "./deployment.jsonnet";
local service = import "./service.jsonnet";

{
   Simple:: function(name, environment, port, replicas=3) [
      service.Simple(name, port),
      deployment.Simple(name,
          container.Simple(
              name,
              "eu.gcr.io/monorepo-base/%s:%s" % [name, environment],
              port),
          replicas=replicas)
  ]
}

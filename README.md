# go-opensearch
Extension for go-elasticsearch library supports Opensearch/Opendistro for Elasticsearch APIs.
Based on official **go-elasticsearch** package - https://github.com/elastic/go-elasticsearch.
---

# Info
Base - **_github.com/elastic/go-elasticsearch/v7 v7.13.1_**

This package wraps **esapi.Client** stuct from **go-elasticsearch** with **API** struct from this package
into single **opensearch.Client** stuct. That means what you can still use requests provided by **go-elasticsearch**
as before, but also you can use specific APIs requests for **Opensearch\Opendistro for Elasticsearch** and their plugins
(like **Security plugin**) via **opensearch.Client.OSAPI** field.

Also package provides replacement for **NewClient()** and **NewDefaultClient()** methods which return **opensearch.Client**.
Default client uses **https** connection instead of **http**(**go-elasticsearh**), ignores certificate verification by default and
has **admin/admin** credentials.


# Usage
## Import to your project
**go.mod**
```bigquery
module my-cool-module

require github.com/thethir13en/go-opensearch v0.1.0
```

## Examples

### Get role from Opensearch security plugin(same for other objects)
```bigquery
package main

import (
    osclient "github.com/thethir13en/go-opensearch"
    "log"
    "context"
)

func main() {
    // Create new Opensearch Client
    //
    es, _ := osclient.NewDefaultClient()
    
    // Prepare request
    //
    req := osclient.GetRoleRequest{
        Name:     "all_access",
        Pretty:   true,
    }

    // Send request
    //
    res, _ := req.Do(context.TODO(),es)
    defer res.Body.Close()

    // Log response
    //
    log.Println(res.String())
}
```
Output:
```
2021/07/09 14:04:23 [200 OK] {
  "kibana_user" : {
    "reserved" : true,
    "hidden" : false,
    "description" : "Provide the minimum permissions for a kibana user",
    "cluster_permissions" : [
      "cluster_composite_ops"
    ],
    "index_permissions" : [
      {
        "index_patterns" : [
          ".kibana",
          ".kibana-6",
          ".kibana_*",
          ".opensearch_dashboards",
          ".opensearch_dashboards-6",
          ".opensearch_dashboards_*"
        ],
        "fls" : [ ],
        "masked_fields" : [ ],
        "allowed_actions" : [
          "read",
          "delete",
          "manage",
          "index"
        ]
      },
      {
        "index_patterns" : [
          ".tasks",
          ".management-beats",
          "*:.tasks",
          "*:.management-beats"
        ],
        "fls" : [ ],
        "masked_fields" : [ ],
        "allowed_actions" : [
          "indices_all"
        ]
      }
    ],
    "tenant_permissions" : [ ],
    "static" : true
  }
}
```
### Delete role
```bigquery
package main

import (
    osclient "github.com/thethir13en/go-opensearch"
    "log"
    "context"
)

func main() {
    // Create new Opensearch Client
    //
    es, _ := osclient.NewDefaultClient()
    
    // Prepare request
    //
    req := osclient.DeleteRoleRequest{
        Name:     "custom_role2",
        Pretty:   true,
    }

    // Send request
    //
    res, _ := req.Do(context.TODO(),es)
    defer res.Body.Close()

    // Log response
    //
    log.Println(res.String())
}
```
Output:
```bigquery
2021/07/09 14:08:56 [200 OK] {
  "status" : "OK",
  "message" : "'custom_role' deleted."
}
```
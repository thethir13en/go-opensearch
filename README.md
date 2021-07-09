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


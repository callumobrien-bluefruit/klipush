# `klipush`

A tool for refreshing klipfolio data sources.

## Usage

```
$ klipush -id source-id <data.json
```

where `source-id` is the ID of the data source to update.

Expects a file `secrets.json` in the current directory with the format
```
{
  "api-key": "<your-klipfolio-api-key>"
}
```

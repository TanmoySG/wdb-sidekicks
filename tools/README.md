# wdb Tools

Various tools (mostly commandline) to do various operations with/on WDB Data. Can be used from wdb-server.

## List of Tools

- [roles_hidden_field_update](./roles_hidden_field_update) : a tool to migrate persisted roles data to new model that has a new field `hidden`. 
- [migrate_databases_primary_key](./migrate_databases_primary_key/) : a tool to migrate existing records in Collections from existing non-primary-key-ed model to new primary-key-ed model and assings recordId as existing id.

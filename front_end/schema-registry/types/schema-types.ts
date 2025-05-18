export type Schema = {
  id: string;
  name: string;
  schema: string;
};

export type FetchSchemasResponse = {
  statusCode: number;
  schemas: Schema[];
};

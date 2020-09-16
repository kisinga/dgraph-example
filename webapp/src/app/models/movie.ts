import { Person } from "./person";

export interface Movie {
  name: string;
  starring: string;
  genre: string;
  initial_relrease_date: Date;
  performance: Performance;
}

interface Performance {
  actor: Person;
}

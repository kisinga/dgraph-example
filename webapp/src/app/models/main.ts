export interface Film {
  uid: string;
  name: string;
  genre: Array<string>;
  release_date: Date;
  actors: Performance;
}

export interface Flattened {
  name: string;
  uid: string;
}

export interface Actor extends Flattened {
  films_acted: Array<Film>;
}
